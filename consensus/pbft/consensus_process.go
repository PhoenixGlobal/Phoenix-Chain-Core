package pbft

import (
	"Phoenix-Chain-Core/ethereum/core/types/pbfttypes"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"Phoenix-Chain-Core/libs/crypto/bls"

	"Phoenix-Chain-Core/consensus/pbft/utils"

	"Phoenix-Chain-Core/consensus/pbft/executor"

	"Phoenix-Chain-Core/libs/log"

	"Phoenix-Chain-Core/consensus/pbft/protocols"
	"Phoenix-Chain-Core/consensus/pbft/state"
	ctypes "Phoenix-Chain-Core/consensus/pbft/types"
	"Phoenix-Chain-Core/ethereum/core/types"
	"Phoenix-Chain-Core/libs/common"
)

// OnPrepareBlock performs security rule verification，store in blockTree,
// Whether to start synchronization
func (pbft *Pbft) OnPrepareBlock(id string, msg *protocols.PrepareBlock) error {
	pbft.log.Debug("Receive PrepareBlock", "id", id, "msg", msg.String())
	if err := pbft.VerifyHeader(nil, msg.Block.Header(), false); err != nil {
		pbft.log.Error("Verify header fail", "number", msg.Block.Number(), "hash", msg.Block.Hash(), "err", err)
		return err
	}
	if err := pbft.safetyRules.PrepareBlockRules(msg); err != nil {
		blockCheckFailureMeter.Mark(1)

		if err.Common() {
			pbft.log.Debug("Prepare block rules fail", "number", msg.Block.Number(), "hash", msg.Block.Hash(), "err", err)
			return err
		}
		// verify consensus signature
		if err := pbft.verifyConsensusSign(msg); err != nil {
			signatureCheckFailureMeter.Mark(1)
			return err
		}

		if err.Fetch() {
			if pbft.isProposer(msg.Epoch, msg.BlockNum(), msg.ProposalIndex) {
				pbft.log.Info("Epoch or viewNumber higher than local, try to fetch block", "fetchHash", msg.Block.ParentHash(), "fetchNumber", msg.Block.NumberU64()-1)
				pbft.fetchBlock(id, msg.Block.ParentHash(), msg.Block.NumberU64()-1, nil)
			}
			return err
		}
		if err.FetchPrepare() {
			if pbft.isProposer(msg.Epoch, msg.BlockNum(), msg.ProposalIndex) {
				pbft.log.Debug("err.FetchPrepare(), prepareBlockFetchRules", "prepare", msg.String())
				pbft.prepareBlockFetchRules(id, msg)
			}
			return err
		}
		if err.NewView() {
			var block *types.Block
			var qc *ctypes.QuorumCert
			if e := pbft.checkViewChangeQC(msg); e != nil {
				return e
			}
			if msg.ViewChangeQC != nil {
				_, _, _, _, hash, number := msg.ViewChangeQC.MaxBlock()
				block, qc = pbft.blockTree.FindBlockAndQC(hash, number)
			} else {
				block, qc = pbft.blockTree.FindBlockAndQC(msg.Block.ParentHash(), msg.Block.NumberU64()-1)
			}
			pbft.log.Debug("Receive new view's block, change view", "newEpoch", msg.Epoch, "newView", msg.ViewNumber)
			pbft.changeView(msg.Epoch, msg.ViewNumber, block, qc, msg.ViewChangeQC)
		}
	}

	if _, err := pbft.verifyConsensusMsg(msg); err != nil {
		pbft.log.Error("Failed to verify prepareBlock", "prepare", msg.String(), "error", err.Error())
		signatureCheckFailureMeter.Mark(1)
		return err
	}
	// The new block is notified by the PrepareBlockHash to the nodes in the network.
	pbft.state.AddPrepareBlock(msg)
	pbft.log.Debug("Receive new prepareBlock", "msgHash", msg.MsgHash(), "prepare", msg.String())
	pbft.state.UpdateStep(ctypes.RoundStepPrepareBlock)
	pbft.exePrepareBlock(msg)
	return nil
}

// OnPrepareVote perform security rule verification，store in blockTree,
// Whether to start synchronization
func (pbft *Pbft) OnPrepareVote(id string, msg *protocols.PrepareVote) error {
	pbft.log.Debug("Receive PrepareVote", "id", id, "msg", msg.String())
	if err := pbft.safetyRules.PrepareVoteRules(msg); err != nil {
		if err.Common() {
			pbft.log.Debug("Preparevote rules fail", "number", msg.BlockHash, "hash", msg.BlockHash, "err", err)
			return err
		}

		// verify consensus signature
		if pbft.verifyConsensusSign(msg) != nil {
			signatureCheckFailureMeter.Mark(1)
			return err
		}

		if err.Fetch() {
			if msg.ParentQC != nil {
				pbft.log.Info("Epoch or viewNumber higher than local, try to fetch block", "fetchHash", msg.ParentQC.BlockHash, "fetchNumber", msg.ParentQC.BlockNumber)
				pbft.fetchBlock(id, msg.ParentQC.BlockHash, msg.ParentQC.BlockNumber, msg.ParentQC)
			}
		} else if err.FetchPrepare() {
			pbft.log.Debug("err.FetchPrepare(), prepareVoteFetchRules", "prepareVote", msg.String())
			pbft.prepareVoteFetchRules(id, msg)
		}
		return err
	}

	var node *pbfttypes.ValidateNode
	var err error
	if node, err = pbft.verifyConsensusMsg(msg); err != nil {
		pbft.log.Error("Failed to verify prepareVote", "vote", msg.String(), "error", err.Error())
		return err
	}

	pbft.state.AddPrepareVote(uint32(node.Index), msg)
	pbft.log.Debug("Receive new prepareVote", "msgHash", msg.MsgHash(), "vote", msg.String(), "votes", pbft.state.PrepareVoteLenByNumber(msg.BlockNumber))

	pbft.state.UpdateStep(ctypes.RoundStepPrepareVote)
	//pbft.insertPrepareQC(msg.ParentQC)
	pbft.findPreVoteQC(msg.BlockNumber)
	return nil
}

// OnPrepareVote perform security rule verification，store in blockTree,
// Whether to start synchronization
func (pbft *Pbft) OnPreCommit(id string, msg *protocols.PreCommit) error {
	pbft.log.Debug("Receive PreCommit", "id", id, "msg", msg.String())
	if err := pbft.safetyRules.PreCommitRules(msg); err != nil {
		if err.Common() {
			pbft.log.Debug("PreCommit rules fail", "number", msg.BlockHash, "hash", msg.BlockHash, "err", err)
			return err
		}

		// verify consensus signature
		if pbft.verifyConsensusSign(msg) != nil {
			signatureCheckFailureMeter.Mark(1)
			return err
		}

		if err.Fetch() {
			if msg.ParentQC != nil {
				pbft.log.Info("Epoch or viewNumber higher than local, try to fetch block", "fetchHash", msg.ParentQC.BlockHash, "fetchNumber", msg.ParentQC.BlockNumber)
				pbft.fetchBlock(id, msg.ParentQC.BlockHash, msg.ParentQC.BlockNumber, msg.ParentQC)
			}
		} else if err.FetchPrepare() {
			pbft.log.Debug("err.FetchPrepare(), preCommitFetchRules", "preCommit", msg.String())
			pbft.preCommitFetchRules(id, msg)
		}
		return err
	}

	var node *pbfttypes.ValidateNode
	var err error
	if node, err = pbft.verifyConsensusMsg(msg); err != nil {
		pbft.log.Error("Failed to verify PreCommit", "vote", msg.String(), "error", err.Error())
		return err
	}

	pbft.state.AddPreCommit(uint32(node.Index), msg)
	pbft.log.Debug("Receive new PreCommit", "msgHash", msg.MsgHash(), "vote", msg.String(), "votes", pbft.state.PreCommitLenByNumber(msg.BlockNumber))

	pbft.state.UpdateStep(ctypes.RoundStepPreCommit)
	//pbft.insertPrepareQC(msg.ParentQC)
	pbft.findPreCommitQC(msg)
	return nil
}

// OnViewChange performs security rule verification, view switching.
func (pbft *Pbft) OnViewChange(id string, msg *protocols.ViewChange) error {
	pbft.log.Debug("Receive ViewChange", "id", id, "msg", msg.String())
	if err := pbft.safetyRules.ViewChangeRules(msg); err != nil {
		if err.Fetch() {
			if msg.PrepareQC != nil {
				pbft.log.Info("Epoch or viewNumber higher than local, try to fetch block", "fetchHash", msg.BlockHash, "fetchNumber", msg.BlockNumber)
				pbft.fetchBlock(id, msg.BlockHash, msg.BlockNumber, msg.PrepareQC)
			}
		}
		return err
	}

	var node *pbfttypes.ValidateNode
	var err error

	if node, err = pbft.verifyConsensusMsg(msg); err != nil {
		pbft.log.Error("Failed to verify viewChange", "viewChange", msg.String(), "error", err.Error())
		return err
	}

	pbft.state.AddViewChange(uint32(node.Index), msg)
	pbft.log.Debug("Receive new viewChange", "msgHash", msg.MsgHash(), "viewChange", msg.String(), "total", pbft.state.ViewChangeLen())
	//pbft.state.UpdateStep(ctypes.RoundStepNewRound)
	// It is possible to achieve viewchangeQC every time you add viewchange
	pbft.tryChangeView()
	return nil
}

// OnViewTimeout performs timeout logic for view.
func (pbft *Pbft) OnViewTimeout() {
	pbft.log.Info("Current view timeout", "view", pbft.state.ViewString())
	node, err := pbft.isCurrentValidator()
	if err != nil {
		pbft.log.Info("ViewTimeout local node is not validator")
		return
	}

	if pbft.state.ViewChangeByIndex(node.Index) != nil {
		pbft.log.Debug("Had send viewchange, don't send again")
		return
	}

	hash, number := pbft.state.HighestPreCommitQCBlock().Hash(), pbft.state.HighestPreCommitQCBlock().NumberU64()
	_, qc := pbft.blockTree.FindBlockAndQC(hash, number)

	viewChange := &protocols.ViewChange{
		Epoch:          pbft.state.Epoch(),
		ViewNumber:     pbft.state.ViewNumber(),
		BlockHash:      hash,
		BlockNumber:    number,
		ValidatorIndex: uint32(node.Index),
		PrepareQC:      qc,
	}

	if err := pbft.signMsgByBls(viewChange); err != nil {
		pbft.log.Error("Sign ViewChange failed", "err", err)
		return
	}

	// write sendViewChange info to wal
	if !pbft.isLoading() {
		pbft.bridge.SendViewChange(viewChange)
	}

	pbft.state.AddViewChange(uint32(node.Index), viewChange)
	pbft.network.Broadcast(viewChange)
	pbft.log.Info("Local add viewChange", "index", node.Index, "viewChange", viewChange.String(), "total", pbft.state.ViewChangeLen())

	pbft.tryChangeView()
}

// OnInsertQCBlock performs security rule verification, view switching.
func (pbft *Pbft) OnInsertQCBlock(blocks []*types.Block, qcs []*ctypes.QuorumCert) error {
	if len(blocks) != len(qcs) {
		return fmt.Errorf("block qc is inconsistent")
	}
	//todo insert tree, update view
	for i := 0; i < len(blocks); i++ {
		block, qc := blocks[i], qcs[i]
		//todo verify qc

		if err := pbft.safetyRules.QCBlockRules(block, qc); err != nil {
			if err.NewView() {
				pbft.log.Info("The block to be written belongs to the next view, need change view", "view", pbft.state.ViewString(), "qcBlock", qc.String())
				pbft.changeView(qc.Epoch, state.DefaultViewNumber, block, qc, nil)
			} else {
				return err
			}
		}
		pbft.insertQCBlock(block, qc)
		pbft.log.Info("Insert QC block success", "qcBlock", qc.String())
		pbft.tryAddStateBlockNumber(qc)
		pbft.tryChangeView()
	}

	//pbft.findExecutableBlock()
	return nil
}

// Update blockTree, try commit new block
func (pbft *Pbft) insertQCBlock(block *types.Block, qc *ctypes.QuorumCert) {
	pbft.log.Info("Insert preCommit QC block", "qc", qc.String())
	if block.NumberU64() <= pbft.state.HighestLockBlock().NumberU64() || pbft.HasBlock(block.Hash(), block.NumberU64()) {
		pbft.log.Debug("The inserted block has exists in chain",
			"number", block.Number(), "hash", block.Hash(),
			"CommitNumber", pbft.state.HighestCommitBlock().Number(),
			"CommitHash", pbft.state.HighestCommitBlock().Hash())
		return
	}
	if pbft.state.Epoch() == qc.Epoch{
		if pbft.state.ViewBlockByIndex(qc.BlockNumber) == nil {
			pbft.state.AddQCBlock(block, qc)
			pbft.state.AddQC(qc)
			pbft.state.AddPreCommitQC(qc)
		} else {
			pbft.state.AddQC(qc)
			pbft.state.AddPreCommitQC(qc)
		}
	}

	lock, commit := pbft.blockTree.InsertQCBlock(block, qc)
	pbft.TrySetHighestQCBlock(block)
	pbft.TrySetHighestPreCommitQCBlock(block)
	pbft.TrySetHighestLockBlock(lock)
	isOwn := func() bool {
		node, err := pbft.isCurrentValidator()
		if err != nil {
			return false
		}
		proposer := pbft.currentProposer()
		// The current node is the proposer and the block is generated by itself
		if node.Index == proposer.Index && pbft.state.Epoch() == qc.Epoch {
			return true
		}
		return false
	}()
	if !isOwn {
		pbft.txPool.Reset(block)
	}
	pbft.tryCommitNewBlock(lock, commit, block)
	//pbft.trySendPrepareVote()
	//pbft.tryChangeView()
	if pbft.insertBlockQCHook != nil {
		// test hook
		pbft.insertBlockQCHook(block, qc)
	}
}

func (pbft *Pbft) TrySetHighestQCBlock(block *types.Block) {
	h := pbft.state.HighestQCBlock()
	if h.NumberU64()<=block.NumberU64(){
		pbft.state.SetHighestQCBlock(block)
	}
	return
}

func (pbft *Pbft) TrySetHighestPreCommitQCBlock(block *types.Block) {
	_, qc := pbft.blockTree.FindBlockAndQC(block.Hash(), block.NumberU64())
	h := pbft.state.HighestPreCommitQCBlock()
	_, hqc := pbft.blockTree.FindBlockAndQC(h.Hash(), h.NumberU64())
	if hqc == nil || qc.HigherQuorumCert(hqc.BlockNumber, hqc.Epoch, hqc.ViewNumber) {
		pbft.state.SetHighestPreCommitQCBlock(block)
	}
}

func (pbft *Pbft) TrySetHighestLockBlock(block *types.Block) {
	_, qc := pbft.blockTree.FindBlockAndQC(block.Hash(), block.NumberU64())
	h := pbft.state.HighestLockBlock()
	_, hqc := pbft.blockTree.FindBlockAndQC(h.Hash(), h.NumberU64())
	if hqc == nil || qc.HigherQuorumCert(hqc.BlockNumber, hqc.Epoch, hqc.ViewNumber) {
		pbft.state.SetHighestLockBlock(block)
	}
}

func (pbft *Pbft) insertPreCommitQC(qc *ctypes.QuorumCert) {
	if qc != nil {
		block := pbft.state.ViewBlockByIndex(qc.BlockNumber)

		linked := func(blockNumber uint64) bool {
			if block != nil {
				parent, _ := pbft.blockTree.FindBlockAndQC(block.ParentHash(), block.NumberU64()-1)
				return parent != nil && pbft.state.HighestPreCommitQCBlock().NumberU64()<= blockNumber
			}
			return false
		}
		hasExecuted := func() bool {
			if pbft.validatorPool.IsValidator(qc.Epoch, pbft.config.Option.NodeID) {
				return pbft.state.HadSendPreVoteByQC(qc) && linked(qc.BlockNumber)
			} else if pbft.validatorPool.IsCandidateNode(pbft.config.Option.NodeID) {
				maxExecutedBlockNumber:=pbft.state.MaxExecutedBlockNumber()
				//blockIndex, finish := pbft.state.Executing()
				return (qc.BlockNumber <= maxExecutedBlockNumber) && linked(qc.BlockNumber)
			}
			return false
		}
		if block != nil && hasExecuted() {
			pbft.state.SetPreCommitQC(qc)
			pbft.state.AddPreCommitQC(qc)
			pbft.insertQCBlock(block, qc)
			pbft.log.Debug("insertPreCommitQC success", "view", pbft.state.ViewString(), "hasExecuted",hasExecuted(),"linked",linked(qc.BlockNumber))
			pbft.tryAddStateBlockNumber(qc)
			pbft.tryChangeView()
		}else {
			pbft.log.Debug("insertPreCommitQC fail", "view", pbft.state.ViewString(), "hasExecuted",hasExecuted(),"linked",linked(qc.BlockNumber))
		}
	}
}

// Asynchronous execution block callback function
func (pbft *Pbft) onAsyncExecuteStatus(s *executor.BlockExecuteStatus) {
	pbft.log.Debug("Async Execute Block", "hash", s.Hash, "number", s.Number)
	if s.Err != nil {
		pbft.log.Error("Execute block failed", "err", s.Err, "hash", s.Hash, "number", s.Number)
		return
	}
	//index, finish := pbft.state.Executing()
	//if index==math.MaxUint32{
	//	pbft.log.Error("view.executing.BlockIndex==math.MaxUint32")
	//	return
	//}
	//if !finish {
	var index uint32
	index=0
	block := pbft.state.ViewBlockByIndex(s.Number)
	if block != nil {
		if block.Hash() == s.Hash {
			pbft.state.SetExecuting(index, true)
			if pbft.executeFinishHook != nil {
				pbft.executeFinishHook(index)
			}
			pbft.state.SetMaxExecutedBlockNumber(block.NumberU64())
			_, err := pbft.validatorPool.GetValidatorByNodeID(pbft.state.Epoch(), pbft.config.Option.NodeID)
			if err != nil {
				pbft.log.Debug("Current node is not validator,no need to sign block", "err", err, "hash", s.Hash, "number", s.Number)
				return
			}
			if err := pbft.signBlock(block.Hash(), block.NumberU64(), index); err != nil {
				pbft.log.Error("Sign block failed", "err", err, "hash", s.Hash, "number", s.Number)
				return
			}

			pbft.log.Debug("Sign block", "hash", s.Hash, "number", s.Number)
			//if msg := pbft.csPool.GetPrepareQC(pbft.state.Epoch(), pbft.state.ViewNumber(), index); msg != nil {
			//	go pbft.ReceiveMessage(msg)
			//}
			pbft.findPreVoteQC(s.Number)
		}
	}
	//}
	//pbft.findExecutableBlock()
}

// Sign the block that has been executed
// Every time try to trigger a send PrepareVote
func (pbft *Pbft) signBlock(hash common.Hash, number uint64, index uint32) error {
	// todo sign vote
	// parentQC added when sending
	// Determine if the current consensus node is
	node, err := pbft.validatorPool.GetValidatorByNodeID(pbft.state.Epoch(), pbft.config.Option.NodeID)
	if err != nil {
		return err
	}
	prepareVote := &protocols.PrepareVote{
		Epoch:          pbft.state.Epoch(),
		ViewNumber:     pbft.state.ViewNumber(),
		BlockHash:      hash,
		BlockNumber:    number,
		BlockIndex:     index,
		ValidatorIndex: uint32(node.Index),
	}

	if err := pbft.signMsgByBls(prepareVote); err != nil {
		return err
	}
	//pbft.state.PendingPrepareVote().Push(prepareVote)
	//pbft.state.PrepareVote=prepareVote
	// Record the number of participating consensus
	consensusCounter.Inc(1)

	pbft.trySendPrepareVote(prepareVote)
	return nil
}

// Every time try to trigger a send PreCommit
func (pbft *Pbft) genPreCommit(qc *ctypes.QuorumCert) error {
	if qc.BlockNumber!=pbft.state.BlockNumber(){
		pbft.log.Debug("qc.BlockNumber!=pbft.state.BlockNumber()","qc.BlockNumber", qc.BlockNumber, "pbft.state.BlockNumber()", pbft.state.BlockNumber())
		return errors.New("blockNumber is not equal to pbft.state.BlockNumber()")
	}

	//pv:=pbft.state.PrepareVote
	//if pv.BlockNumber!=qc.BlockNumber{
	//	return errors.New("qc.BlockNumber is not equal to pbft.state.PrepareVote.BlockNumber")
	//}

	node, err := pbft.validatorPool.GetValidatorByNodeID(pbft.state.Epoch(), pbft.config.Option.NodeID)
	if err != nil {
		return err
	}
	preCommit := &protocols.PreCommit{
		Epoch:          qc.Epoch,
		ViewNumber:     qc.ViewNumber,
		BlockHash:      qc.BlockHash,
		BlockNumber:    qc.BlockNumber,
		BlockIndex:     qc.BlockIndex,
		ValidatorIndex: uint32(node.Index),
		//ParentQC:       qc.ParentQC,
		//Signature:      pv.Signature,
	}

	block := pbft.state.ViewBlockByIndex(qc.BlockNumber)
	if b, qc2 := pbft.blockTree.FindBlockAndQC(block.ParentHash(), block.NumberU64()-1); b != nil || block.NumberU64() == 0 {
		preCommit.ParentQC=qc2
		if err := pbft.signMsgByBls(preCommit); err != nil {
			return err
		}
		//pbft.state.PreCommit=preCommit
		pbft.trySendPreCommit(preCommit)
		return nil
	}
    err=errors.New("genPreCommit find ParentQC fail")
	pbft.log.Debug("genPreCommit find ParentQC fail", "qc", qc.String())
	return err
}

// Send a signature,
// obtain a signature from the pending queue,
// determine whether the parent block has reached QC,
// and send a signature if it is reached, otherwise exit the sending logic.
func (pbft *Pbft) trySendPrepareVote(msg *protocols.PrepareVote) {
	// Check timeout
	if pbft.state.IsDeadline() {
		pbft.log.Debug("Current view had timeout, Refuse to send prepareVotes")
		return
	}

	//pending := pbft.state.PendingPrepareVote()
	//hadSend := pbft.state.HadSendPrepareVote()
	//
	//for !pending.Empty() {
	//p := pbft.state.PrepareVote
	p:=msg
	if err := pbft.voteRules.AllowVote(p); err != nil {
		pbft.log.Debug("Not allow send vote", "err", err, "msg", p.String())
		return
	}

	block := pbft.state.ViewBlockByIndex(p.BlockNumber)
	// The executed block has a signature.
	// Only when the view is switched, the block is cleared but the vote is also cleared.
	// If there is no block, the consensus process is abnormal and should not run.
	if block == nil {
		pbft.log.Crit("Try send PrepareVote failed", "err", "vote corresponding block not found", "view", pbft.state.ViewString(), p.String())
	}
	if b, qc := pbft.blockTree.FindBlockAndQC(block.ParentHash(), block.NumberU64()-1); b != nil || block.NumberU64() == 0 {
		p.ParentQC = qc
		//hadSend.Push(p)
		pbft.tryAddHadSendVote(p.BlockNumber,p)
		//Determine if the current consensus node is
		node, _ := pbft.validatorPool.GetValidatorByNodeID(pbft.state.Epoch(), pbft.config.Option.NodeID)
		pbft.log.Info("Add local prepareVote", "vote", p.String())
		pbft.state.AddPrepareVote(uint32(node.Index), p)
		//pending.Pop()

		// write sendPrepareVote info to wal
		if !pbft.isLoading() {
			pbft.bridge.SendPrepareVote(block, p)
		}

		pbft.network.Broadcast(p)
	}
	return
	//}
}

// Send a signature,
// obtain a signature from the pending queue,
// determine whether the parent block has reached QC,
// and send a signature if it is reached, otherwise exit the sending logic.
func (pbft *Pbft) trySendPreCommit(msg *protocols.PreCommit) {
	// Check timeout
	if pbft.state.IsDeadline() {
		pbft.log.Debug("Current view had timeout, Refuse to send preCommits")
		return
	}

	//pending := pbft.state.PendingPrepareVote()
	//hadSend := pbft.state.HadSendPrepareVote()
	//
	//for !pending.Empty() {
	//p := pbft.state.PreCommit
	p:=msg
	if err := pbft.voteRules.AllowPreCommit(p); err != nil {
		pbft.log.Debug("Not allow send preCommit", "err", err, "msg", p.String())
		return
	}

	block := pbft.state.ViewBlockByIndex(p.BlockNumber)
	// The executed block has a signature.
	// Only when the view is switched, the block is cleared but the vote is also cleared.
	// If there is no block, the consensus process is abnormal and should not run.
	if block == nil {
		pbft.log.Crit("Try send PreCommit failed", "err", "vote corresponding block not found", "view", pbft.state.ViewString(), p.String())
	}
	if b, qc := pbft.blockTree.FindBlockAndQC(block.ParentHash(), block.NumberU64()-1); b != nil || block.NumberU64() == 0 {
		p.ParentQC = qc
		//hadSend.Push(p)
		//pbft.state.HadSendPreCommit=p
		pbft.tryAddHadSendPreCommit(p.BlockNumber,p)
		//Determine if the current consensus node is
		node, _ := pbft.validatorPool.GetValidatorByNodeID(pbft.state.Epoch(), pbft.config.Option.NodeID)
		pbft.log.Info("Add local preCommit", "vote", p.String())
		pbft.state.AddPreCommit(uint32(node.Index), p)
		//pending.Pop()

		// write sendPrepareVote info to wal
		if !pbft.isLoading() {
			pbft.bridge.SendPreCommit(block, p)
		}
		pbft.network.Broadcast(p)
	}
	return
	//}
}

func (pbft *Pbft) exePrepareBlock(msg *protocols.PrepareBlock) {
	blockIndex:=msg.BlockIndex
	block := pbft.state.ViewBlockByIndex(msg.BlockNum())
	if block==nil{
		pbft.log.Error(fmt.Sprintf("exePrepareBlock ViewBlockByIndex failed, blockIndex:[%d],msg:[%v]", blockIndex,msg))
		return
	}
	parent, _ := pbft.blockTree.FindBlockAndQC(block.ParentHash(), block.NumberU64()-1)
	if parent == nil {
		pbft.log.Error(fmt.Sprintf("exePrepareBlock find block's parent failed :[%d,%d,%s]", blockIndex, block.NumberU64(), block.Hash().String()))
		return
	}

	pbft.log.Debug("exePrepareBlock find Block", "hash", block.Hash(), "number", block.NumberU64())
	if err := pbft.asyncExecutor.Execute(block, parent); err != nil {
		pbft.log.Error("Async Execute block failed", "error", err)
	}
	pbft.state.SetExecuting(blockIndex, false)
}

func (pbft *Pbft) findPreVoteQC(blockNumber uint64) {
	//index:= pbft.state.PrepareVote.BlockIndex
	//if index==math.MaxUint32{
	//	fmt.Println("pbft.state.PrepareVote.BlockIndex==math.MaxUint32")
	//	pbft.log.Error("pbft.state.PrepareVote.BlockIndex==math.MaxUint32")
	//	return
	//}
	size := pbft.state.PrepareVoteLenByNumber(blockNumber)

	prepareQC := func() bool {
		return size >= pbft.threshold(pbft.currentValidatorLen())
	}

	if prepareQC() {
		block := pbft.state.ViewBlockByIndex(blockNumber)
		qc := pbft.generatePrepareQC(pbft.state.AllPrepareVoteByNumber(blockNumber))
		if qc != nil {
			pbft.log.Info("New qc block have been created", "qc", qc.String())
			//pbft.insertQCBlock(block, qc)
			pbft.TrySetHighestQCBlock(block)
			pbft.network.Broadcast(&protocols.BlockQuorumCert{BlockQC: qc})
			// metrics
			blockQCCollectedGauage.Update(int64(block.Time()))
			//pbft.trySendPrepareVote()
			pbft.state.SetPrepareVoteQC(qc)
			pbft.state.AddQC(qc)
			pbft.genPreCommit(qc)
		}
	}
	//pbft.tryChangeView()
}

// Each time a new vote is triggered, a new QC Block will be triggered, and a new one can be found by the commit block.
func (pbft *Pbft) findPreCommitQC(msg *protocols.PreCommit) {
	//index:= pbft.state.PreCommit.BlockIndex
	//if index==math.MaxUint32{
	//	pbft.log.Error("pbft.state.PreCommit.BlockIndex==math.MaxUint32")
	//	return
	//}

	size := pbft.state.PreCommitLenByNumber(msg.BlockNumber)

	prepareQC := func() bool {
		return size >= pbft.threshold(pbft.currentValidatorLen())
	}

	if prepareQC() {
		block := pbft.state.ViewBlockByIndex(msg.BlockNumber)
		qc := pbft.generatePreCommitQC(pbft.state.AllPreCommitByNumber(msg.BlockNumber))
		if qc != nil {
			pbft.state.SetPreCommitQC(qc)
			pbft.state.AddPreCommitQC(qc)
			pbft.log.Info("New PreCommit qc block have been created", "qc", qc.String())
			pbft.insertQCBlock(block, qc)
			pbft.network.Broadcast(&protocols.BlockPreCommitQuorumCert{BlockQC: qc})
			// metrics
			blockQCCollectedGauage.Update(int64(block.Time()))
			//pbft.trySendPreCommit()
			pbft.tryAddStateBlockNumber(qc)
			pbft.tryChangeView()
		}
	}
}

func (pbft *Pbft) initStateBlockNumberByQC(qc *ctypes.QuorumCert) {
	if qc==nil{
		return
	}
	pbft.state.AddBlockNumber(qc)
	return
}

func (pbft *Pbft) initStateBlockNumberByBlockNumber(blockNumber uint64) {
	pbft.state.SetBlockNumber(blockNumber)
	return
}

func (pbft *Pbft) tryAddStateBlockNumber(qc *ctypes.QuorumCert) {
	if qc==nil{
		return
	}
	if qc.BlockNumber==0{
		pbft.state.AddBlockNumber(qc)
		return
	}
	if qc.BlockNumber==pbft.state.BlockNumber(){
		pbft.state.AddBlockNumber(qc)
	}
	return
}

func (pbft *Pbft) tryAddHadSendVote(number uint64,vote *protocols.PrepareVote) {
	hadSend:=pbft.state.HadSendVoteNumber()
	hadSendVote:=pbft.state.HadSendVoteLast()
	if number>hadSend||(number==hadSend&&(vote.Epoch>hadSendVote.Epoch||vote.Epoch==hadSendVote.Epoch&&vote.ViewNumber>hadSendVote.ViewNumber)){
		pbft.state.AddHadSendVoteNumber(number,vote)
	}
	return
}

func (pbft *Pbft) tryAddHadSendPreCommit(number uint64,vote *protocols.PreCommit) {
	hadSend:=pbft.state.HadSendPreCommitNumber()
	hadSendPreCommit:=pbft.state.HadSendPreCommitLast()
	if number>hadSend||(number==hadSend&&(vote.Epoch>hadSendPreCommit.Epoch||vote.Epoch==hadSendPreCommit.Epoch&&vote.ViewNumber>hadSendPreCommit.ViewNumber)){
		pbft.state.AddHadSendPreCommitNumber(number,vote)
	}
	return
}

// Try commit a new block
func (pbft *Pbft) tryCommitNewBlock(lock *types.Block, commit *types.Block, qc *types.Block) {
	if lock == nil || commit == nil {
		pbft.log.Warn("Try commit failed", "hadLock", lock != nil, "hadCommit", commit != nil)
		return
	}
	if commit.NumberU64()+1 != qc.NumberU64() {
		pbft.log.Warn("Try commit failed,the requirement of commit block is not achieved", "commit", commit.NumberU64(), "lock", lock.NumberU64(), "qc", qc.NumberU64())
		return
	}
	//highestqc := pbft.state.HighestQCBlock()
	highestqc := qc
	commitBlock, commitQC := pbft.blockTree.FindBlockAndQC(commit.Hash(), commit.NumberU64())

	_, oldCommit := pbft.state.HighestLockBlock(), pbft.state.HighestCommitBlock()
	// Incremental commit block
	if oldCommit.NumberU64()+1 == commit.NumberU64() {
		pbft.commitBlock(commit, commitQC, lock, highestqc)
		pbft.state.SetHighestCommitBlock(commit)
		pbft.blockTree.PruneBlock(commit.Hash(), commit.NumberU64(), nil)
		pbft.blockTree.NewRoot(commit)
		// metrics
		blockNumberGauage.Update(int64(commit.NumberU64()))
		highestQCNumberGauage.Update(int64(highestqc.NumberU64()))
		highestLockedNumberGauage.Update(int64(lock.NumberU64()))
		highestCommitNumberGauage.Update(int64(commit.NumberU64()))
		blockConfirmedMeter.Mark(1)
	} else if oldCommit.NumberU64() == commit.NumberU64() && oldCommit.NumberU64() > 0 {
		pbft.log.Info("Fork block", "number", highestqc.NumberU64(), "hash", highestqc.Hash())
		lockBlock, lockQC := pbft.blockTree.FindBlockAndQC(lock.Hash(), lock.NumberU64())
		qcBlock, qcQC := pbft.blockTree.FindBlockAndQC(highestqc.Hash(), highestqc.NumberU64())

		qcState := &protocols.State{Block: qcBlock, QuorumCert: qcQC}
		lockState := &protocols.State{Block: lockBlock, QuorumCert: lockQC}
		commitState := &protocols.State{Block: commitBlock, QuorumCert: commitQC}
		pbft.bridge.UpdateChainState(qcState, lockState, commitState)
	}
}

// Asynchronous processing of errors generated by the submission block
func (pbft *Pbft) OnCommitError(err error) {
	// FIXME Do you want to panic and stop the consensus?
	pbft.log.Error("Commit block error", "err", err)
}

// According to the current view QC situation, try to switch view
func (pbft *Pbft) tryChangeView() {
	// Had receive all qcs of current view
	block, qc := pbft.blockTree.FindBlockAndQC(pbft.state.HighestPreCommitQCBlock().Hash(), pbft.state.HighestPreCommitQCBlock().NumberU64())

	//increasing := func() uint64 {
	//	return pbft.state.ViewNumber() + 1
	//}

	shouldSwitch := func() bool {
		return pbft.validatorPool.ShouldSwitch(block.NumberU64())
	}()

	enough := func() bool {
		return pbft.state.MaxQCIndex()+1 == pbft.config.Sys.Amount ||
			(qc != nil && qc.Epoch == pbft.state.Epoch() && shouldSwitch)
	}()

	if shouldSwitch {
		if err := pbft.validatorPool.Update(block.NumberU64(), pbft.state.Epoch()+1, pbft.eventMux); err == nil {
			pbft.log.Info("Update validator success", "number", block.NumberU64())
		}
	}

	if enough {
		// If current has produce enough blocks, then change view immediately.
		// Otherwise, waiting for view's timeout.
		if pbft.validatorPool.EqualSwitchPoint(block.NumberU64()) {
			pbft.log.Info("BlockNumber is equal to switchPoint, change epoch", "blockNumber", block.NumberU64(), "view", pbft.state.ViewString())
			pbft.changeView(pbft.state.Epoch()+1, state.DefaultViewNumber, block, qc, nil)
		} else {
			pbft.log.Info("Produce enough blocks, change view", "view", pbft.state.ViewString())
			pbft.changeView(pbft.state.Epoch(), state.DefaultViewNumber, block, qc, nil)
		}
		return
	}

	viewChangeQC := func() bool {
		if pbft.state.ViewChangeLen() >= pbft.threshold(pbft.currentValidatorLen()) {
			return true
		}
		pbft.log.Debug("Try change view failed, had receive viewchange", "len", pbft.state.ViewChangeLen(), "view", pbft.state.ViewString())
		return false
	}

	if viewChangeQC() {
		viewChangeQC := pbft.generateViewChangeQC(pbft.state.AllViewChange())
		pbft.log.Info("Receive enough viewchange, try change view by viewChangeQC", "view", pbft.state.ViewString(), "viewChangeQC", viewChangeQC.String())
		pbft.tryChangeViewByViewChange(viewChangeQC)
	}
}

func (pbft *Pbft) richViewChangeQC(viewChangeQC *ctypes.ViewChangeQC) {
	node, err := pbft.isCurrentValidator()
	if err != nil {
		pbft.log.Info("Local node is not validator")
		return
	}
	hadSend := pbft.state.ViewChangeByIndex(uint32(node.Index))
	if hadSend != nil && !viewChangeQC.ExistViewChange(hadSend.Epoch, hadSend.ViewNumber, hadSend.BlockHash) {
		cert, err := pbft.generateViewChangeQuorumCert(hadSend)
		if err != nil {
			pbft.log.Error("Generate viewChangeQuorumCert error", "err", err)
			return
		}
		pbft.log.Info("Already send viewChange, append viewChangeQuorumCert to ViewChangeQC", "cert", cert.String())
		viewChangeQC.AppendQuorumCert(cert)
	}

	_, qc := pbft.blockTree.FindBlockAndQC(pbft.state.HighestQCBlock().Hash(), pbft.state.HighestQCBlock().NumberU64())
	_, _, blockEpoch, blockView, _, number := viewChangeQC.MaxBlock()

	if qc == nil{
		return
	}
	if qc.HigherQuorumCert(number, blockEpoch, blockView) {
		if hadSend == nil {
			v, err := pbft.generateViewChange(qc)
			if err != nil {
				pbft.log.Error("Generate viewChange error", "err", err)
				return
			}
			cert, err := pbft.generateViewChangeQuorumCert(v)
			if err != nil {
				pbft.log.Error("Generate viewChangeQuorumCert error", "err", err)
				return
			}
			pbft.log.Info("Not send viewChange, append viewChangeQuorumCert to ViewChangeQC", "cert", cert.String())
			viewChangeQC.AppendQuorumCert(cert)
		}
	}
}

func (pbft *Pbft) tryChangeViewByViewChange(viewChangeQC *ctypes.ViewChangeQC) {
	increasing := func() uint64 {
		return pbft.state.ViewNumber() + 1
	}

	block, qc := pbft.blockTree.FindBlockAndQC(pbft.state.HighestPreCommitQCBlock().Hash(), pbft.state.HighestPreCommitQCBlock().NumberU64())
	if block.NumberU64() != 0 {
		pbft.richViewChangeQC(viewChangeQC)
		_, _, blockEpoch, _, _, number := viewChangeQC.MaxBlock()
		if pbft.state.HighestPreCommitQCBlock().NumberU64()<number || (pbft.state.HighestPreCommitQCBlock().NumberU64()==number&&pbft.state.Epoch()<blockEpoch){
			// fixme
			pbft.log.Warn("Local node is behind other validators", "blockState", pbft.state.HighestBlockString(), "viewChangeQC", viewChangeQC.String())
			return
		}
		//if block == nil || qc == nil {
		//	// fixme get qc block
		//	pbft.log.Warn("Local node is behind other validators", "blockState", pbft.state.HighestBlockString(), "viewChangeQC", viewChangeQC.String())
		//	return
		//}
	}
	_, _, blockEpoch, _, _, number := viewChangeQC.MaxBlock()

	if pbft.validatorPool.EqualSwitchPoint(number) && blockEpoch == pbft.state.Epoch() {
		// Validator already switch, new epoch
		pbft.log.Info("BlockNumber is equal to switchPoint, change epoch", "blockNumber", number, "view", pbft.state.ViewString())
		pbft.changeView(pbft.state.Epoch()+1, state.DefaultViewNumber, block, qc, viewChangeQC)
	} else {
		pbft.changeView(pbft.state.Epoch(), increasing(), block, qc, viewChangeQC)
	}
}

func (pbft *Pbft) generateViewChangeQuorumCert(v *protocols.ViewChange) (*ctypes.ViewChangeQuorumCert, error) {
	node, err := pbft.isCurrentValidator()
	if err != nil {
		return nil, errors.Wrap(err, "local node is not validator")
	}
	total := uint32(pbft.validatorPool.Len(pbft.state.Epoch()))
	var aggSig bls.Sign
	if err := aggSig.Deserialize(v.Sign()); err != nil {
		return nil, err
	}

	blockEpoch, blockView := uint64(0), uint64(0)
	if v.PrepareQC != nil {
		blockEpoch, blockView = v.PrepareQC.Epoch, v.PrepareQC.ViewNumber
	}
	cert := &ctypes.ViewChangeQuorumCert{
		Epoch:           v.Epoch,
		ViewNumber:      v.ViewNumber,
		BlockHash:       v.BlockHash,
		BlockNumber:     v.BlockNumber,
		BlockEpoch:      blockEpoch,
		BlockViewNumber: blockView,
		ValidatorSet:    utils.NewBitArray(total),
	}
	cert.Signature.SetBytes(aggSig.Serialize())
	cert.ValidatorSet.SetIndex(node.Index, true)
	return cert, nil
}

func (pbft *Pbft) generateViewChange(qc *ctypes.QuorumCert) (*protocols.ViewChange, error) {
	node, err := pbft.isCurrentValidator()
	if err != nil {
		return nil, errors.Wrap(err, "local node is not validator")
	}
	v := &protocols.ViewChange{
		Epoch:          pbft.state.Epoch(),
		ViewNumber:     pbft.state.ViewNumber(),
		BlockHash:      qc.BlockHash,
		BlockNumber:    qc.BlockNumber,
		ValidatorIndex: uint32(node.Index),
		PrepareQC:      qc,
	}
	if err := pbft.signMsgByBls(v); err != nil {
		return nil, errors.Wrap(err, "Sign ViewChange failed")
	}

	return v, nil
}

// change view
func (pbft *Pbft) changeView(epoch, viewNumber uint64, block *types.Block, qc *ctypes.QuorumCert, viewChangeQC *ctypes.ViewChangeQC) {
	interval := func() uint64 {
		return 1
	}
	// last epoch and last viewNumber
	// when pbft is started or fast synchronization ends, the preEpoch, preViewNumber defaults to 0, 0
	// but pbft is now in the loading state and lastViewChangeQC is nil, does not save the lastViewChangeQC
	preEpoch, preViewNumber := pbft.state.Epoch(), pbft.state.ViewNumber()
	// syncingCache is belong to last view request, clear all sync cache
	pbft.syncingCache.Purge()
	pbft.csPool.Purge(epoch, block.NumberU64())

	pbft.state.ResetView(epoch, viewNumber,pbft.state.BlockNumber())
	pbft.state.SetViewTimer(interval())
	pbft.state.SetLastViewChangeQC(viewChangeQC)

	// metrics.
	viewNumberGauage.Update(int64(viewNumber))
	epochNumberGauage.Update(int64(epoch))
	viewChangedTimer.UpdateSince(time.Unix(int64(block.Time()), 0))

	// write confirmed viewChange info to wal
	if !pbft.isLoading() {
		pbft.bridge.ConfirmViewChange(epoch, viewNumber, block, qc, viewChangeQC, preEpoch, preViewNumber)
	}
	pbft.clearInvalidBlocks(block)
	pbft.evPool.Clear(epoch, block.NumberU64())
	// view change maybe lags behind the other nodes,active sync prepare block
	pbft.SyncPrepareBlock("", epoch, block.NumberU64(), 0)
	pbft.log = log.New("epoch", pbft.state.Epoch(), "view", pbft.state.ViewNumber(),"blockNumber",pbft.state.BlockNumber())
	pbft.log.Info("Success to change view, current view deadline", "deadline", pbft.state.Deadline())
}

// Clean up invalid blocks in the previous view
func (pbft *Pbft) clearInvalidBlocks(newBlock *types.Block) {
	var rollback []*types.Block
	newHead := newBlock.Header()
	for _, p := range pbft.state.HadSendPrepareVote().Peek() {
		if p.BlockNumber > newBlock.NumberU64() {
			block := pbft.state.ViewBlockByIndex(p.BlockNumber)
			rollback = append(rollback, block)
		}
	}
	for _, p := range pbft.state.PendingPrepareVote().Peek() {
		if p.BlockNumber > newBlock.NumberU64() {
			block := pbft.state.ViewBlockByIndex(p.BlockNumber)
			rollback = append(rollback, block)
		}
	}
	pbft.blockCacheWriter.ClearCache(pbft.state.HighestCommitBlock())

	//todo proposer is myself
	pbft.txPool.ForkedReset(newHead, rollback)
}
