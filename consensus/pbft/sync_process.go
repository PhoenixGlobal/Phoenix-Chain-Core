package pbft

import (
	"container/list"
	"fmt"
	"sort"
	"time"

	"Phoenix-Chain-Core/consensus/pbft/state"

	"Phoenix-Chain-Core/consensus/pbft/network"
	"Phoenix-Chain-Core/consensus/pbft/protocols"
	ctypes "Phoenix-Chain-Core/consensus/pbft/types"
	"Phoenix-Chain-Core/consensus/pbft/utils"
	"Phoenix-Chain-Core/ethereum/core/types"
	"Phoenix-Chain-Core/libs/common"
)

var syncPrepareVotesInterval = 3 * time.Second
var syncPreCommitsInterval = 3 * time.Second

// Get the block from the specified connection, get the block into the fetcher, and execute the block PBFT update state machine
func (pbft *Pbft) fetchBlock(id string, hash common.Hash, number uint64, qc *ctypes.QuorumCert) {
	if pbft.fetcher.Len() != 0 {
		pbft.log.Trace("Had fetching block")
		return
	}
	highestQC := pbft.state.HighestPreCommitQCBlock()
	if highestQC.NumberU64()+2 < number {
		pbft.log.Debug(fmt.Sprintf("Local state too low, local.highestQC:%s,%d, remote.msg:%s,%d", highestQC.Hash().String(), highestQC.NumberU64(), hash.String(), number))
		return
	}

	baseBlockHash, baseBlockNumber := common.Hash{}, uint64(0)
	var parentBlock *types.Block

	if highestQC.NumberU64() < number {
		parentBlock = highestQC
		baseBlockHash, baseBlockNumber = parentBlock.Hash(), parentBlock.NumberU64()
		//baseBlockHash, baseBlockNumber = hash, number
	}  else {
		pbft.log.Trace("No suitable block need to request")
		return
	}

	if qc != nil {
		if qc.Epoch==pbft.state.Epoch(){
			if err := pbft.verifyPrepareQC(number, hash, qc,0); err != nil {
				pbft.log.Warn(fmt.Sprintf("Verify prepare qc failed, error:%s", err.Error()))
				return
			}
		}else {
			pbft.log.Warn("fetchBlock: epoch of qc is higher then local", "qc epoch",qc.Epoch, "local epoch",pbft.state.Epoch())
		}
	}

	match := func(msg ctypes.Message) bool {
		_, ok := msg.(*protocols.QCBlockList)
		return ok
	}

	executor := func(msg ctypes.Message) {
		defer func() {
			pbft.log.Debug("Close fetching")
			//utils.SetFalse(&pbft.fetching)
		}()
		pbft.log.Debug("Receive QCBlockList", "msg", msg.String())
		if blockList, ok := msg.(*protocols.QCBlockList); ok {
			// Execution block
			//fmt.Println("555555555555,executor receive msg: ",msg,blockList.Blocks[0].Number(),blockList.QC)
			for _, block := range blockList.Blocks {
				if block.ParentHash() != parentBlock.Hash() {
					//fmt.Println("6666666,","Response block's is error", "blockHash", block.Hash(), "blockNumber", block.NumberU64(), "parentHash", parentBlock.Hash(), "parentNumber", parentBlock.NumberU64())
					pbft.log.Debug("Response block's is error",
						"blockHash", block.Hash(), "blockNumber", block.NumberU64(),
						"parentHash", parentBlock.Hash(), "parentNumber", parentBlock.NumberU64())
					return
				}
				//fmt.Println("66666111111")
				if err := pbft.blockCacheWriter.Execute(block, parentBlock); err != nil {
					pbft.log.Error("Execute block failed", "hash", block.Hash(), "number", block.NumberU64(), "error", err)
					return
				}
				parentBlock = block
			}

			// Update the results to the PBFT state machine
			pbft.asyncCallCh <- func() {
				for i, block := range blockList.Blocks {
					if err := pbft.verifyPrepareQC(block.NumberU64(), block.Hash(), blockList.QC[i],0); err != nil {
						pbft.log.Error("Verify block prepare qc failed", "hash", block.Hash(), "number", block.NumberU64(), "error", err)
						return
					}
				}
				if err := pbft.OnInsertQCBlock(blockList.Blocks, blockList.QC); err != nil {
					pbft.log.Error("Insert block failed", "error", err)
				}
			}
			if blockList.ForkedBlocks == nil || len(blockList.ForkedBlocks) == 0 {
				pbft.log.Trace("No forked block need to handle")
				return
			}
			// Remove local forks that already exist.
			filteredForkedBlocks := make([]*types.Block, 0)
			filteredForkedQCs := make([]*ctypes.QuorumCert, 0)
			//localForkedBlocks, _ := pbft.blockTree.FindForkedBlocksAndQCs(parentBlock.Hash(), parentBlock.NumberU64())
			localForkedBlocks, _ := pbft.blockTree.FindBlocksAndQCs(parentBlock.NumberU64())

			if localForkedBlocks != nil && len(localForkedBlocks) > 0 {
				pbft.log.Debug("LocalForkedBlocks", "number", localForkedBlocks[0].NumberU64(), "hash", localForkedBlocks[0].Hash().TerminalString())
			}

			for i, forkedBlock := range blockList.ForkedBlocks {
				for _, localForkedBlock := range localForkedBlocks {
					if forkedBlock.NumberU64() == localForkedBlock.NumberU64() && forkedBlock.Hash() != localForkedBlock.Hash() {
						filteredForkedBlocks = append(filteredForkedBlocks, forkedBlock)
						filteredForkedQCs = append(filteredForkedQCs, blockList.ForkedQC[i])
						break
					}
				}
			}
			if filteredForkedBlocks != nil && len(filteredForkedBlocks) > 0 {
				pbft.log.Debug("FilteredForkedBlocks", "number", filteredForkedBlocks[0].NumberU64(), "hash", filteredForkedBlocks[0].Hash().TerminalString())
			}

			// Execution forked block.
			//var forkedParentBlock *types.Block
			for _, forkedBlock := range filteredForkedBlocks {
				if forkedBlock.NumberU64() != parentBlock.NumberU64() {
					pbft.log.Error("Invalid forked block", "lastParentNumber", parentBlock.NumberU64(), "forkedBlockNumber", forkedBlock.NumberU64())
					break
				}
				//for _, block := range blockList.Blocks {
				//	if block.Hash() == forkedBlock.ParentHash() && block.NumberU64() == forkedBlock.NumberU64()-1 {
				//		forkedParentBlock = block
				//		break
				//	}
				//}
				//if forkedParentBlock != nil {
				//	break
				//}
			}

			// Verify forked block and execute.
			for _, forkedBlock := range filteredForkedBlocks {
				parentBlock := pbft.blockTree.FindBlockByHash(forkedBlock.ParentHash())
				if parentBlock == nil {
					pbft.log.Debug("Response forked block's is error",
						"blockHash", forkedBlock.Hash(), "blockNumber", forkedBlock.NumberU64())
					return
				}
				//if forkedParentBlock == nil || forkedBlock.ParentHash() != forkedParentBlock.Hash() {
				//	pbft.log.Debug("Response forked block's is error",
				//		"blockHash", forkedBlock.Hash(), "blockNumber", forkedBlock.NumberU64(),
				//		"parentHash", parentBlock.Hash(), "parentNumber", parentBlock.NumberU64())
				//	return
				//}

				if err := pbft.blockCacheWriter.Execute(forkedBlock, parentBlock); err != nil {
					pbft.log.Error("Execute forked block failed", "hash", forkedBlock.Hash(), "number", forkedBlock.NumberU64(), "error", err)
					return
				}
			}

			//pbft.asyncCallCh <- func() {
			//	for i, forkedBlock := range filteredForkedBlocks {
			//		if err := pbft.verifyPrepareQC(forkedBlock.NumberU64(), forkedBlock.Hash(), blockList.ForkedQC[i],0); err != nil {
			//			pbft.log.Error("Verify forked block prepare qc failed", "hash", forkedBlock.Hash(), "number", forkedBlock.NumberU64(), "error", err)
			//			return
			//		}
			//	}
			//	if err := pbft.OnInsertQCBlock(filteredForkedBlocks, filteredForkedQCs); err != nil {
			//		pbft.log.Error("Insert forked block failed", "error", err)
			//	}
			//}
		}
	}

	expire := func() {
		pbft.log.Debug("Fetch timeout, close fetching", "targetId", id, "baseBlockHash", baseBlockHash, "baseBlockNumber", baseBlockNumber)
		utils.SetFalse(&pbft.fetching)
	}

	pbft.log.Debug("Start fetching", "id", id, "basBlockHash", baseBlockHash, "baseBlockNumber", baseBlockNumber)

	//utils.SetTrue(&pbft.fetching)
	pbft.fetcher.AddTask(id, match, executor, expire)
	msg:=&protocols.GetQCBlockList{BlockHash: baseBlockHash, BlockNumber: baseBlockNumber,Epoch: pbft.state.Epoch()}
	pbft.network.Send(id, msg)
	//fmt.Println("44444444444444,id is ",id,"number is ",number)
	//fmt.Println("send protocols.GetQCBlockList:",msg.String(),msg)
}

// Obtain blocks that are not in the local according to the proposed block
func (pbft *Pbft) prepareBlockFetchRules(id string, pb *protocols.PrepareBlock) {
	//fmt.Println(22222,"pbft.state.HighestPreCommitQCBlock().NumberU64():",pbft.state.HighestPreCommitQCBlock().NumberU64())
	if pb.Block.NumberU64() >= pbft.state.BlockNumber() {
		baseBlockNumber:=pbft.state.BlockNumber()
		b := pbft.state.ViewBlockByIndex(baseBlockNumber)
		if b == nil {
			pbft.SyncPrepareBlock(id, pbft.state.Epoch(), baseBlockNumber, 0,pbft.state.ViewNumber())
		}

		//for i := uint32(0); i <= pb.BlockIndex; i++ {
		//	b := pbft.state.ViewBlockByIndex(i)
		//	if b == nil {
		//		pbft.SyncPrepareBlock(id, pbft.state.Epoch(), pbft.state.BlockNumber(), i)
		//	}
		//}
	}
}

// Get votes and blocks that are not available locally based on the height of the vote
func (pbft *Pbft) prepareVoteFetchRules(id string, vote *protocols.PrepareVote) {
	if vote.BlockNumber == pbft.state.BlockNumber() && vote.ViewNumber == pbft.state.ViewNumber() {
		baseBlockNumber:=pbft.state.BlockNumber()
		b, qc := pbft.state.ViewBlockAndQC(baseBlockNumber)
		if b == nil {
			pbft.SyncPrepareBlock(id, pbft.state.Epoch(), baseBlockNumber, 0,pbft.state.ViewNumber())
		} else if qc == nil {
			pbft.SyncBlockQuorumCert(id, b.NumberU64(), b.Hash(), 0)
		}
	}
}

// Get votes and blocks that are not available locally based on the height of the vote
func (pbft *Pbft) preCommitFetchRules(id string, vote *protocols.PreCommit) {
	if vote.BlockNumber == pbft.state.BlockNumber() && vote.ViewNumber == pbft.state.ViewNumber() {
		baseBlockNumber:=pbft.state.BlockNumber()
		b, qc := pbft.state.ViewBlockAndPreCommitQC(baseBlockNumber)
		if b == nil {
			pbft.SyncPrepareBlock(id, pbft.state.Epoch(), baseBlockNumber, 0,pbft.state.ViewNumber())
		} else if qc == nil {
			pbft.SyncBlockPreCommitQuorumCert(id, b.NumberU64(), b.Hash(), 0)
		}
	}
}

// OnGetPrepareBlock handles the  message type of GetPrepareBlockMsg.
func (pbft *Pbft) OnGetPrepareBlock(id string, msg *protocols.GetPrepareBlock) error {
	if msg.Epoch == pbft.state.Epoch() && msg.BlockNumber == pbft.state.BlockNumber()&& msg.ViewNumber == pbft.state.ViewNumber() {
		prepareBlock := pbft.state.PrepareBlockByIndex(msg.BlockNumber)
		if prepareBlock != nil {
			pbft.log.Debug("Send PrepareBlock", "peer", id, "prepareBlock", prepareBlock.String())
			pbft.network.Send(id, prepareBlock)
		}

		_, qc := pbft.state.ViewBlockAndQC(msg.BlockNumber)
		if qc != nil {
			pbft.log.Debug("Send BlockQuorumCert on GetPrepareBlock", "peer", id, "qc", qc.String())
			pbft.network.Send(id, &protocols.BlockQuorumCert{BlockQC: qc})
		}
	}
	return nil
}

// OnGetBlockQuorumCert handles the message type of GetBlockQuorumCertMsg.
func (pbft *Pbft) OnGetBlockQuorumCert(id string, msg *protocols.GetBlockQuorumCert) error {
	_, qc := pbft.state.ViewBlockAndQC(msg.BlockNumber)
	if qc != nil {
		pbft.log.Debug("Send BlockQuorumCert on GetBlockQuorumCert", "peer", id, "qc", qc.String())
		pbft.network.Send(id, &protocols.BlockQuorumCert{BlockQC: qc})
	}
	//_, qc := pbft.blockTree.FindBlockAndQC(msg.BlockHash, msg.BlockNumber)
	//if qc != nil {
	//	qc.Step=ctypes.RoundStepPrepareVote
	//	pbft.network.Send(id, &protocols.BlockQuorumCert{BlockQC: qc})
	//}
	return nil
}

// OnBlockQuorumCert handles the message type of BlockQuorumCertMsg.
func (pbft *Pbft) OnBlockQuorumCert(id string, msg *protocols.BlockQuorumCert) error {
	pbft.log.Debug("Receive BlockQuorumCert", "peer", id, "msg", msg.String())
	if msg.BlockQC.Epoch != pbft.state.Epoch() || (msg.BlockQC.BlockNumber != pbft.state.BlockNumber()&&pbft.state.BlockNumber()!=0) || msg.BlockQC.ViewNumber != pbft.state.ViewNumber(){
		pbft.log.Trace("Receive BlockQuorumCert response failed", "local.epoch", pbft.state.Epoch(), "local.blockNumber", pbft.state.BlockNumber(),"local.viewNumber", pbft.state.ViewNumber(),  "msg", msg.String())
		return fmt.Errorf("msg is not match current state")
	}

	if _, qc := pbft.blockTree.FindBlockAndQC(msg.BlockQC.BlockHash, msg.BlockQC.BlockNumber); qc != nil {
		pbft.log.Trace("Block has exist", "msg", msg.String())
		return fmt.Errorf("block already exists")
	}

	if pbft.state.HadPrepareVoteQC(msg.BlockQC.BlockHash, msg.BlockQC.BlockNumber,msg.BlockQC.ViewNumber){
		pbft.log.Trace("BlockQC has exist", "msg", msg.String())
		return fmt.Errorf("BlockQC already exists")
	}

	// If blockQC comes the block must exist
	block := pbft.state.ViewBlockByIndex(msg.BlockQC.BlockNumber)
	if block == nil {
		pbft.log.Debug("Block not exist", "msg", msg.String())
		return fmt.Errorf("block not exist")
	}
	if err := pbft.verifyPrepareQC(block.NumberU64(), block.Hash(), msg.BlockQC,ctypes.RoundStepPrepareVote); err != nil {
		pbft.log.Error("Failed to verify prepareQC", "err", err.Error())
		return &authFailedError{err}
	}

	pbft.TrySetHighestQCBlock(block)
	//pbft.trySendPrepareVote()
	pbft.state.SetPrepareVoteQC(msg.BlockQC)
	pbft.state.AddQC(msg.BlockQC)

	pbft.genPreCommit(msg.BlockQC)

	pbft.log.Info("Receive new blockQuorumCert, ", "view", pbft.state.ViewString(), "blockQuorumCert", msg.BlockQC.String())
	//pbft.insertPrepareQC(msg.BlockQC)
	return nil
}

// OnGetBlockQuorumCert handles the message type of GetBlockQuorumCertMsg.
func (pbft *Pbft) OnGetBlockPreCommitQuorumCert(id string, msg *protocols.GetBlockPreCommitQuorumCert) error {
	_, qc := pbft.state.ViewBlockAndPreCommitQC(msg.BlockNumber)
	if qc != nil {
		pbft.log.Debug("Send BlockQuorumCert on GetBlockPreCommitQuorumCert", "peer", id, "qc", qc.String())
		pbft.network.Send(id, &protocols.BlockPreCommitQuorumCert{BlockQC: qc})
	}

	//_, qc := pbft.blockTree.FindBlockAndQC(msg.BlockHash, msg.BlockNumber)
	//if qc != nil {
	//	pbft.network.Send(id, &protocols.BlockPreCommitQuorumCert{BlockQC: qc})
	//}
	return nil
}

// OnBlockQuorumCert handles the message type of BlockQuorumCertMsg.
func (pbft *Pbft) OnBlockPreCommitQuorumCert(id string, msg *protocols.BlockPreCommitQuorumCert) error {
	pbft.log.Debug("Receive BlockPreCommitQuorumCert", "peer", id, "msg", msg.String())
	if msg.BlockQC.Epoch != pbft.state.Epoch() || (msg.BlockQC.BlockNumber != pbft.state.BlockNumber()&&pbft.state.BlockNumber()!=0) || msg.BlockQC.ViewNumber != pbft.state.ViewNumber(){
		pbft.log.Trace("Receive BlockPreCommitQuorumCert response failed", "local.epoch", pbft.state.Epoch(), "local.BlockNumber", pbft.state.BlockNumber(),"local.ViewNumber", pbft.state.ViewNumber(),  "msg", msg.String())
		return fmt.Errorf("msg is not match current state")
	}

	if _, qc := pbft.blockTree.FindBlockAndQC(msg.BlockQC.BlockHash, msg.BlockQC.BlockNumber); qc != nil {
		pbft.log.Trace("Block has exist", "msg", msg.String())
		return fmt.Errorf("block already exists")
	}

	if pbft.state.HadPreCommitQC(msg.BlockQC.BlockHash, msg.BlockQC.BlockNumber,msg.BlockQC.ViewNumber){
		pbft.log.Trace("BlockPreCommitQC has exist", "msg", msg.String())
		return fmt.Errorf("BlockPreCommitQC already exists")
	}

	// If blockQC comes the block must exist
	block := pbft.state.ViewBlockByIndex(msg.BlockQC.BlockNumber)
	if block == nil {
		pbft.log.Debug("Block not exist", "msg", msg.String())
		return fmt.Errorf("block not exist")
	}
	if err := pbft.verifyPrepareQC(block.NumberU64(), block.Hash(), msg.BlockQC,ctypes.RoundStepPreCommit); err != nil {
		pbft.log.Error("Failed to verify preCommitQC", "err", err.Error())
		return &authFailedError{err}
	}

	pbft.log.Info("Receive blockPreCommitQuorumCert, try insert preCommitQC", "view", pbft.state.ViewString(), "blockPreCommitQuorumCert", msg.BlockQC.String())
	pbft.insertPreCommitQC(msg.BlockQC)
	//pbft.trySendPreCommit()
	return nil
}

// OnGetQCBlockList handles the message type of GetQCBlockListMsg.
func (pbft *Pbft) OnGetQCBlockList(id string, msg *protocols.GetQCBlockList) error {
	highestQC := pbft.state.HighestPreCommitQCBlock()

	if highestQC.NumberU64() <= msg.BlockNumber {
		pbft.log.Debug(fmt.Sprintf("Receive GetQCBlockList failed, local.highestQC:%s,%d, msg:%s", highestQC.Hash().TerminalString(), highestQC.NumberU64(), msg.String()))
		return fmt.Errorf("peer state too high")
	}

	if highestQC.NumberU64() > msg.BlockNumber+2{
		pbft.log.Debug(fmt.Sprintf("Receive GetQCBlockList failed, local.highestQC:%s,%d, msg:%s", highestQC.Hash().TerminalString(), highestQC.NumberU64(), msg.String()))
		return fmt.Errorf("peer state too low")
	}

	//if highestQC.Hash() == msg.BlockHash && highestQC.NumberU64() == msg.BlockNumber {
	//	pbft.log.Debug(fmt.Sprintf("Receive GetQCBlockList failed, local.highestQC:%s,%d, msg:%s", highestQC.Hash().TerminalString(), highestQC.NumberU64(), msg.String()))
	//	return fmt.Errorf("peer state too low")
	//}

	//lock := pbft.state.HighestLockBlock()
	//commit := pbft.state.HighestCommitBlock()

	qcs := make([]*ctypes.QuorumCert, 0)
	blocks := make([]*types.Block, 0)

	//if commit.ParentHash() == msg.BlockHash {
	//	block, qc := pbft.blockTree.FindBlockAndQC(commit.Hash(), commit.NumberU64())
	//	qcs = append(qcs, qc)
	//	blocks = append(blocks, block)
	//}
	//
	//if lock.ParentHash() == msg.BlockHash || commit.ParentHash() == msg.BlockHash {
	//	block, qc := pbft.blockTree.FindBlockAndQC(lock.Hash(), lock.NumberU64())
	//	qcs = append(qcs, qc)
	//	blocks = append(blocks, block)
	//}

	for blockNumber:=msg.BlockNumber+1;blockNumber<=highestQC.NumberU64();blockNumber++{
		block,qc:=pbft.state.ViewBlockAndPreCommitQC(blockNumber)
		if block!=nil&&qc!=nil{
			if qc.Epoch>msg.Epoch{
				break
			}
			qcs = append(qcs, qc)
			blocks = append(blocks, block)
		}
	}

	//if highestQC.ParentHash() == msg.BlockHash {
	//	block, qc := pbft.blockTree.FindBlockAndQC(highestQC.Hash(), highestQC.NumberU64())
	//	qcs = append(qcs, qc)
	//	blocks = append(blocks, block)
	//}

	// If the height of the QC exists in the blocktree,
	// collecting forked blocks.
	forkedQCs := make([]*ctypes.QuorumCert, 0)
	forkedBlocks := make([]*types.Block, 0)
	if highestQC.ParentHash() == msg.BlockHash {
		bs, qcs := pbft.blockTree.FindForkedBlocksAndQCs(highestQC.Hash(), highestQC.NumberU64())
		if bs != nil && qcs != nil && len(bs) != 0 && len(qcs) != 0 {
			pbft.log.Debug("Find forked block and return", "forkedBlockLen", len(bs), "forkedQcLen", len(qcs))
			forkedBlocks = append(forkedBlocks, bs...)
			forkedQCs = append(forkedQCs, qcs...)
		}
	}

	if len(qcs) != 0 {
		pbft.network.Send(id, &protocols.QCBlockList{QC: qcs, Blocks: blocks, ForkedBlocks: forkedBlocks, ForkedQC: forkedQCs})
		pbft.log.Debug("Send QCBlockList", "len", len(qcs))
	}
	return nil
}

// OnGetPrepareVote is responsible for processing the business logic
// of the GetPrepareVote message. It will synchronously return a
// PrepareVotes message to the sender.
func (pbft *Pbft) OnGetPrepareVote(id string, msg *protocols.GetPrepareVote) error {
	pbft.log.Debug("Received message on OnGetPrepareVote", "from", id, "msgHash", msg.MsgHash(), "message", msg.String())

	if msg.Epoch== pbft.state.Epoch() && msg.BlockNumber == pbft.state.BlockNumber()&& msg.ViewNumber == pbft.state.ViewNumber(){
		// If the block has already QC, that response QC instead of votes.
		// Avoid the sender spent a lot of time to verifies PrepareVote msg.
		_, qc := pbft.state.ViewBlockAndQC(msg.BlockNumber)
		if qc != nil {
			pbft.network.Send(id, &protocols.BlockQuorumCert{BlockQC: qc})
			pbft.log.Debug("OnGetPrepareVote Send BlockQuorumCert", "peer", id, "qc", qc.String())
			return nil
		}

		prepareVoteMap := pbft.state.AllPrepareVoteByNumber(msg.BlockNumber)
		// Defining an array for receiving PrepareVote.
		votes := make([]*protocols.PrepareVote, 0, len(prepareVoteMap))
		if prepareVoteMap != nil {
			threshold := pbft.threshold(pbft.currentValidatorLen())
			remain := threshold - (pbft.currentValidatorLen() - int(msg.UnKnownSet.Size()))
			for k, v := range prepareVoteMap {
				if msg.UnKnownSet.GetIndex(k) {
					votes = append(votes, v)
				}

				// Limit response votes
				if remain > 0 && len(votes) >= remain {
					break
				}
			}
		}
		if len(votes) > 0 {
			pbft.network.Send(id, &protocols.PrepareVotes{Epoch: msg.Epoch, ViewNumber: msg.ViewNumber, BlockIndex: msg.BlockIndex, Votes: votes})
			pbft.log.Debug("Send PrepareVotes", "peer", id, "epoch", msg.Epoch, "viewNumber", msg.ViewNumber, "blockIndex", msg.BlockIndex)
		}
	}
	return nil
}

// OnGetPreCommit is responsible for processing the business logic
// of the GetPrepareVote message. It will synchronously return a
// PrepareVotes message to the sender.
func (pbft *Pbft) OnGetPreCommit(id string, msg *protocols.GetPreCommit) error {
	pbft.log.Debug("Received message on OnGetPreCommit", "from", id, "msgHash", msg.MsgHash(), "message", msg.String())
	if msg.Epoch== pbft.state.Epoch() && msg.BlockNumber == pbft.state.BlockNumber()&& msg.ViewNumber == pbft.state.ViewNumber() {
		// If the block has already QC, that response QC instead of votes.
		// Avoid the sender spent a lot of time to verifies PrepareVote msg.
		_, qc := pbft.state.ViewBlockAndPreCommitQC(msg.BlockNumber)
		if qc != nil {
			pbft.network.Send(id, &protocols.BlockPreCommitQuorumCert{BlockQC: qc})
			pbft.log.Debug("OnGetPreCommit Send BlockPreCommitQuorumCert", "peer", id, "qc", qc.String())
			return nil
		}

		prepareVoteMap := pbft.state.AllPreCommitByNumber(msg.BlockNumber)
		// Defining an array for receiving PrepareVote.
		votes := make([]*protocols.PreCommit, 0, len(prepareVoteMap))
		if prepareVoteMap != nil {
			threshold := pbft.threshold(pbft.currentValidatorLen())
			remain := threshold - (pbft.currentValidatorLen() - int(msg.UnKnownSet.Size()))
			for k, v := range prepareVoteMap {
				if msg.UnKnownSet.GetIndex(k) {
					votes = append(votes, v)
				}

				// Limit response votes
				if remain > 0 && len(votes) >= remain {
					break
				}
			}
		}
		if len(votes) > 0 {
			pbft.network.Send(id, &protocols.PreCommits{Epoch: msg.Epoch, ViewNumber: msg.ViewNumber, BlockIndex: msg.BlockIndex, Votes: votes})
			pbft.log.Debug("Send PreCommits", "peer", id, "epoch", msg.Epoch, "viewNumber", msg.ViewNumber, "blockIndex", msg.BlockIndex)
		}
	}
	return nil
}

// OnPrepareVotes handling response from GetPrepareVote response.
func (pbft *Pbft) OnPrepareVotes(id string, msg *protocols.PrepareVotes) error {
	pbft.log.Debug("Received message on OnPrepareVotes", "from", id, "msgHash", msg.MsgHash(), "message", msg.String())
	for _, vote := range msg.Votes {
		//_, qc := pbft.blockTree.FindBlockAndQC(vote.BlockHash, vote.BlockNumber)
		_, qc := pbft.state.ViewBlockAndQC(vote.BlockNumber)
		if qc == nil && !pbft.network.ContainsHistoryMessageHash(vote.MsgHash()) {
			if err := pbft.OnPrepareVote(id, vote); err != nil {
				if e, ok := err.(HandleError); ok && e.AuthFailed() {
					pbft.log.Error("OnPrepareVotes failed", "peer", id, "err", err)
				}
				return err
			}
		}
	}
	return nil
}

// OnPrepareVotes handling response from GetPrepareVote response.
func (pbft *Pbft) OnPreCommits(id string, msg *protocols.PreCommits) error {
	pbft.log.Debug("Received message on OnPreCommits", "from", id, "msgHash", msg.MsgHash(), "message", msg.String())
	for _, vote := range msg.Votes {
		//_, qc := pbft.blockTree.FindBlockAndQC(vote.BlockHash, vote.BlockNumber)
		_, qc := pbft.state.ViewBlockAndPreCommitQC(vote.BlockNumber)
		if qc == nil && !pbft.network.ContainsHistoryMessageHash(vote.MsgHash()) {
			if err := pbft.OnPreCommit(id, vote); err != nil {
				if e, ok := err.(HandleError); ok && e.AuthFailed() {
					pbft.log.Error("OnPreCommits failed", "peer", id, "err", err)
				}
				return err
			}
		}
	}
	return nil
}

// OnGetLatestStatus hands GetLatestStatus messages.
//
// main logic:
// 1.Compare the blockNumber of the sending node with the local node,
// and if the blockNumber of local node is larger then reply LatestStatus message,
// the message contains the status information of the local node.
func (pbft *Pbft) OnGetLatestStatus(id string, msg *protocols.GetLatestStatus) error {
	pbft.log.Debug("Received message on OnGetLatestStatus", "from", id, "logicType", msg.LogicType, "msgHash", msg.MsgHash(), "message", msg.String())
	if msg.BlockNumber != 0 && msg.QuorumCert == nil || msg.LBlockNumber != 0 && msg.LQuorumCert == nil {
		pbft.log.Error("Invalid getLatestStatus,lack corresponding quorumCert", "getLatestStatus", msg.String())
		return nil
	}
	// Define a function that performs the send action.
	launcher := func(bType uint64, targetId string, blockNumber uint64, blockHash common.Hash, qc *ctypes.QuorumCert) error {
		err := pbft.network.PeerSetting(targetId, bType, blockNumber)
		if err != nil {
			pbft.log.Error("GetPeer failed", "err", err, "peerId", targetId)
			return err
		}
		// Synchronize block data with fetchBlock.
		pbft.fetchBlock(targetId, blockHash, blockNumber, qc)
		return nil
	}
	//
	if msg.LogicType == network.TypeForPreCommitQCBn {
		localQCNum, localQCHash := pbft.state.HighestPreCommitQCBlock().NumberU64(), pbft.state.HighestPreCommitQCBlock().Hash()
		localLockNum, localLockHash := pbft.state.HighestLockBlock().NumberU64(), pbft.state.HighestLockBlock().Hash()
		if localQCNum == msg.BlockNumber && localQCHash == msg.BlockHash {
			pbft.log.Debug("Local qcBn is equal the sender's qcBn", "remoteBn", msg.BlockNumber, "localBn", localQCNum, "remoteHash", msg.BlockHash, "localHash", localQCHash)
			if forkedHash, forkedNum, forked := pbft.blockTree.IsForked(localQCHash, localQCNum); forked {
				pbft.log.Debug("Local highest QC forked", "forkedQCHash", forkedHash, "forkedQCNumber", forkedNum, "localQCHash", localQCHash, "localQCNumber", localQCNum)
				_, qc := pbft.blockTree.FindBlockAndQC(forkedHash, forkedNum)
				_, lockQC := pbft.blockTree.FindBlockAndQC(localLockHash, localLockNum)
				pbft.network.Send(id, &protocols.LatestStatus{
					BlockNumber:  forkedNum,
					BlockHash:    forkedHash,
					QuorumCert:   qc,
					LBlockNumber: localLockNum,
					LBlockHash:   localLockHash,
					LQuorumCert:  lockQC,
					LogicType:    msg.LogicType})
			}
			return nil
		}
		if localQCNum < msg.BlockNumber || (localQCNum == msg.BlockNumber && localQCHash != msg.BlockHash) {
			pbft.log.Debug("Local qcBn is less than the sender's qcBn", "remoteBn", msg.BlockNumber, "localBn", localQCNum)
			//if msg.LBlockNumber == localQCNum && msg.LBlockHash != localQCHash {
			//	return launcher(msg.LogicType, id, msg.LBlockNumber, msg.LBlockHash, msg.LQuorumCert)
			//}
			return launcher(msg.LogicType, id, msg.BlockNumber, msg.BlockHash, msg.QuorumCert)
		}
		// must carry block qc
		_, qc := pbft.blockTree.FindBlockAndQC(localQCHash, localQCNum)
		_, lockQC := pbft.blockTree.FindBlockAndQC(localLockHash, localLockNum)
		pbft.log.Debug("Local qcBn is larger than the sender's qcBn", "remoteBn", msg.BlockNumber, "localBn", localQCNum)
		pbft.network.Send(id, &protocols.LatestStatus{
			BlockNumber:  localQCNum,
			BlockHash:    localQCHash,
			QuorumCert:   qc,
			LBlockNumber: localLockNum,
			LBlockHash:   localLockHash,
			LQuorumCert:  lockQC,
			LogicType:    msg.LogicType,
		})
	}
	return nil
}

// OnLatestStatus is used to process LatestStatus messages that received from peer.
func (pbft *Pbft) OnLatestStatus(id string, msg *protocols.LatestStatus) error {
	pbft.log.Debug("Received message on OnLatestStatus", "from", id, "msgHash", msg.MsgHash(), "message", msg.String())
	if msg.BlockNumber != 0 && msg.QuorumCert == nil || msg.LBlockNumber != 0 && msg.LQuorumCert == nil {
		pbft.log.Error("Invalid LatestStatus,lack corresponding quorumCert", "latestStatus", msg.String())
		return nil
	}
	switch msg.LogicType {
	case network.TypeForPreCommitQCBn:
		localQCBn, localQCHash := pbft.state.HighestPreCommitQCBlock().NumberU64(), pbft.state.HighestPreCommitQCBlock().Hash()
		if localQCBn < msg.BlockNumber || (localQCBn == msg.BlockNumber && localQCHash != msg.BlockHash) {
			err := pbft.network.PeerSetting(id, msg.LogicType, msg.BlockNumber)
			if err != nil {
				pbft.log.Error("PeerSetting failed", "err", err)
				return err
			}
			pbft.log.Debug("LocalQCBn is lower than sender's", "localBn", localQCBn, "remoteBn", msg.BlockNumber)
			//if localQCBn == msg.LBlockNumber && localQCHash != msg.LBlockHash {
			//	pbft.log.Debug("OnLatestStatus ~ fetchBlock by LBlockHash and LBlockNumber")
			//	pbft.fetchBlock(id, msg.LBlockHash, msg.LBlockNumber, msg.LQuorumCert)
			//} else {
			//	pbft.log.Debug("OnLatestStatus ~ fetchBlock by QCBlockHash and QCBlockNumber")
			//	pbft.fetchBlock(id, msg.BlockHash, msg.BlockNumber, msg.QuorumCert)
			//}
			pbft.log.Debug("OnLatestStatus ~ fetchBlock by QCBlockHash and QCBlockNumber")
			pbft.fetchBlock(id, msg.BlockHash, msg.BlockNumber, msg.QuorumCert)
		}
	}
	return nil
}

// OnPrepareBlockHash responsible for handling PrepareBlockHash message.
//
// Note: After receiving the PrepareBlockHash message, it is determined whether the
// block information exists locally. If not, send a network request to get
// the block data.
func (pbft *Pbft) OnPrepareBlockHash(id string, msg *protocols.PrepareBlockHash) error {
	pbft.log.Debug("Received message OnPrepareBlockHash", "from", id, "msgHash", msg.MsgHash(), "message", msg.String())
	if msg.Epoch == pbft.state.Epoch() && msg.BlockNumber == pbft.state.BlockNumber()&& msg.ViewNumber == pbft.state.ViewNumber() {
		block := pbft.state.ViewBlockByIndex(msg.BlockNumber)
		if block == nil {
			pbft.network.RemoveMessageHash(id, msg.MsgHash())
			pbft.SyncPrepareBlock(id, msg.Epoch, msg.BlockNumber, msg.BlockIndex,msg.ViewNumber)
		}
	}
	return nil
}

// OnGetViewChange responds to nodes that require viewChange.
//
// The Epoch and viewNumber of viewChange must be consistent
// with the state of the current node.
func (pbft *Pbft) OnGetViewChange(id string, msg *protocols.GetViewChange) error {
	pbft.log.Debug("Received message on OnGetViewChange", "from", id, "msgHash", msg.MsgHash(), "message", msg.String(), "local", pbft.state.ViewString())

	localEpoch, localViewNumber,localBlockNumber := pbft.state.Epoch(), pbft.state.ViewNumber(),pbft.state.BlockNumber()

	isLocalView := func() bool {
		return msg.Epoch == localEpoch && msg.ViewNumber == localViewNumber && msg.BlockNumber == localBlockNumber
	}

	isLastView := func() bool {
		return (msg.Epoch == localEpoch && msg.ViewNumber+1 == localViewNumber&& msg.BlockNumber == localBlockNumber) || (msg.Epoch == localEpoch && localViewNumber == state.DefaultViewNumber&& msg.BlockNumber+1 == localBlockNumber)
	}

	isPreviousView := func() bool {
		return (msg.Epoch == localEpoch && msg.ViewNumber+1 < localViewNumber&& msg.BlockNumber == localBlockNumber)||(msg.Epoch == localEpoch && msg.BlockNumber+1<localBlockNumber)
	}

	if isLocalView() {
		viewChanges := pbft.state.AllViewChange()

		vcs := &protocols.ViewChanges{}
		for k, v := range viewChanges {
			if msg.ViewChangeBits.GetIndex(k) {
				vcs.VCs = append(vcs.VCs, v)
			}
		}
		pbft.log.Debug("Send ViewChanges", "peer", id, "len", len(vcs.VCs))
		if len(vcs.VCs) != 0 {
			pbft.network.Send(id, vcs)
		}
		return nil
	}
	// Return view QC in the case of less than 1.
	if isLastView() {
		lastViewChangeQC := pbft.state.LastViewChangeQC()
		if lastViewChangeQC == nil {
			pbft.log.Info("Not found lastViewChangeQC")
			return nil
		}
		err := lastViewChangeQC.EqualAll(msg.Epoch, msg.ViewNumber,msg.BlockNumber-1)
		if err != nil {
			pbft.log.Error("Last view change is not equal msg.BlockNumber", "err", err)
			return err
		}
		viewChangeQuorumCert := &protocols.ViewChangeQuorumCert{
			ViewChangeQC: lastViewChangeQC,
		}
		pbft.log.Debug("IsLastView，Send last viewChange quorumCert", "viewChangeQuorumCert", viewChangeQuorumCert.String())
		pbft.network.Send(id, viewChangeQuorumCert)
		return nil
	}
	// get previous viewChangeQC from wal db
	if isPreviousView() {
		if qc, err := pbft.bridge.GetViewChangeQC(msg.Epoch, msg.BlockNumber-1,msg.ViewNumber); err == nil && qc != nil {
			// also inform the local highest view
			highestqc, _ := pbft.bridge.GetViewChangeQC(localEpoch, localBlockNumber-1,localViewNumber-1)
			viewChangeQuorumCert := &protocols.ViewChangeQuorumCert{
				ViewChangeQC: qc,
			}
			if highestqc != nil {
				viewChangeQuorumCert.HighestViewChangeQC = highestqc
			}
			pbft.log.Debug("Send previous viewChange quorumCert", "viewChangeQuorumCert", viewChangeQuorumCert.String())
			pbft.network.Send(id, viewChangeQuorumCert)
			return nil
		}
	}
	pbft.log.Debug("request is not match local view", "local state",pbft.state.ViewString(),"msg", msg.String())
	return fmt.Errorf("request is not match local view, local:%s,msg:%s", pbft.state.ViewString(), msg.String())
}

// OnViewChangeQuorumCert handles the message type of ViewChangeQuorumCertMsg.
func (pbft *Pbft) OnViewChangeQuorumCert(id string, msg *protocols.ViewChangeQuorumCert) error {
	pbft.log.Debug("Received message on OnViewChangeQuorumCert", "from", id, "msgHash", msg.MsgHash(), "message", msg.String())
	viewChangeQC := msg.ViewChangeQC
	epoch, viewNumber, _, _, _, blockNumber := viewChangeQC.MaxBlock()
	if pbft.state.Epoch() == epoch && pbft.state.ViewNumber() == viewNumber && pbft.state.BlockNumber()==blockNumber+1 {
		if err := pbft.verifyViewChangeQC(msg.ViewChangeQC); err == nil {
			pbft.log.Info("Receive viewChangeQuorumCert, try change view by viewChangeQC", "view", pbft.state.ViewString(), "viewChangeQC", viewChangeQC.String())
			pbft.tryChangeViewByViewChange(msg.ViewChangeQC)
		} else {
			pbft.log.Debug("Verify ViewChangeQC failed", "err", err)
			return &authFailedError{err}
		}
	}
	// if the other party's view is still higher than the local one, continue to synchronize the view
	pbft.trySyncViewChangeQuorumCert(id, msg)
	return nil
}

// Synchronize view one by one according to the highest view notified by the other party
func (pbft *Pbft) trySyncViewChangeQuorumCert(id string, msg *protocols.ViewChangeQuorumCert) {
	highestViewChangeQC := msg.HighestViewChangeQC
	if highestViewChangeQC == nil {
		return
	}
	epoch, viewNumber, _, _, _, blockNumber := highestViewChangeQC.MaxBlock()
	if pbft.state.Epoch() != epoch {
		return
	}
	if pbft.state.ViewNumber() == viewNumber&&pbft.state.BlockNumber()==blockNumber+1{
		if err := pbft.verifyViewChangeQC(highestViewChangeQC); err == nil {
			pbft.log.Debug("The highest view is equal to local, change view by highestViewChangeQC directly", "localView", pbft.state.ViewString(), "futureView", highestViewChangeQC.String())
			pbft.tryChangeViewByViewChange(highestViewChangeQC)
		}
		return
	}
	if pbft.state.ViewNumber() < viewNumber && pbft.state.BlockNumber()==blockNumber+1{
		// if local view lags, synchronize view one by one
		if err := pbft.verifyViewChangeQC(highestViewChangeQC); err == nil {
			pbft.log.Debug("Receive future viewChange quorumCert, sync viewChangeQC with fast mode", "localView", pbft.state.ViewString(), "futureView", highestViewChangeQC.String())
			// request viewChangeQC for the current view
			pbft.network.Send(id, &protocols.GetViewChange{
				Epoch:          pbft.state.Epoch(),
				BlockNumber:    pbft.state.BlockNumber(),
				ViewNumber:     pbft.state.ViewNumber(),
				ViewChangeBits: utils.NewBitArray(uint32(pbft.currentValidatorLen())),
			})
		}
	}
}

// OnViewChanges handles the message type of ViewChangesMsg.
func (pbft *Pbft) OnViewChanges(id string, msg *protocols.ViewChanges) error {
	pbft.log.Debug("Received message on OnViewChanges", "from", id, "msgHash", msg.MsgHash(), "message", msg.String())
	for _, v := range msg.VCs {
		if !pbft.network.ContainsHistoryMessageHash(v.MsgHash()) {
			if err := pbft.OnViewChange(id, v); err != nil {
				if e, ok := err.(HandleError); ok && e.AuthFailed() {
					pbft.log.Error("OnViewChanges failed", "peer", id, "err", err)
				}
				return err
			}
		}
	}
	return nil
}

// MissingViewChangeNodes returns the node ID of the missing vote.
//
// Notes:
// Use the channel to complete serial execution to prevent concurrency.
func (pbft *Pbft) MissingViewChangeNodes() (v *protocols.GetViewChange, err error) {
	result := make(chan struct{})

	pbft.asyncCallCh <- func() {
		defer func() { result <- struct{}{} }()
		allViewChange := pbft.state.AllViewChange()

		length := pbft.currentValidatorLen()
		vbits := utils.NewBitArray(uint32(length))

		// enough qc or did not reach deadline
		if len(allViewChange) >= pbft.threshold(length) || !pbft.state.IsDeadline() {
			v, err = nil, fmt.Errorf("no need sync viewchange")
			return
		}
		for i := uint32(0); i < vbits.Size(); i++ {
			if _, ok := allViewChange[i]; !ok {
				vbits.SetIndex(i, true)
			}
		}

		v, err = &protocols.GetViewChange{
			Epoch:          pbft.state.Epoch(),
			ViewNumber:     pbft.state.ViewNumber(),
			BlockNumber:    pbft.state.BlockNumber(),
			ViewChangeBits: vbits,
		}, nil
	}
	<-result
	return
}

// MissingPrepareVote returns missing vote.
func (pbft *Pbft) MissingPrepareVote() (v *protocols.GetPrepareVote, err error) {
	result := make(chan struct{})

	pbft.asyncCallCh <- func() {
		defer func() { result <- struct{}{} }()

		//begin := pbft.state.MaxQCNumber() + 1
		index := pbft.state.BlockNumber()
		len := pbft.currentValidatorLen()
		pbft.log.Debug("MissingPrepareVote", "epoch", pbft.state.Epoch(), "viewNumber", pbft.state.ViewNumber(), "BlockNumber", index, "validatorLen", len)

		block := pbft.state.HighestQCBlock()
		blockTime := common.MillisToTime(int64(block.Time()))

		//for index := begin; index <= end; index++ {
		size := pbft.state.PrepareVoteLenByNumber(index)
		pbft.log.Debug("The length of prepare vote", "number", index, "size", size)

		// We need sync prepare votes when a long time not arrived QC.
		if size < pbft.threshold(len) && time.Since(blockTime) >= syncPrepareVotesInterval { // need sync prepare votes
			knownVotes := pbft.state.AllPrepareVoteByNumber(index)
			unKnownSet := utils.NewBitArray(uint32(len))
			for i := uint32(0); i < unKnownSet.Size(); i++ {
				if vote := pbft.csPool.GetPrepareVote(pbft.state.Epoch(), index, 0, i); vote != nil {
					cMsg := vote.Msg.(ctypes.ConsensusMsg)
					if cMsg.ViewNum()==pbft.state.ViewNumber(){
						go pbft.ReceiveMessage(vote)
						continue
					}
				}
				if _, ok := knownVotes[i]; !ok {
					unKnownSet.SetIndex(i, true)
				}
			}

			v, err = &protocols.GetPrepareVote{
				Epoch:      pbft.state.Epoch(),
				ViewNumber: pbft.state.ViewNumber(),
				BlockNumber: index,
				BlockIndex: 0,
				UnKnownSet: unKnownSet,
			}, nil
			//break
		}
		//}
		if v == nil {
			err = fmt.Errorf("not need sync prepare vote")
		}
	}
	<-result
	return
}

// MissingPreCommit returns missing preCommit.
func (pbft *Pbft) MissingPreCommit() (v *protocols.GetPreCommit, err error) {
	result := make(chan struct{})

	pbft.asyncCallCh <- func() {
		defer func() { result <- struct{}{} }()

		//begin := pbft.state.MaxPreCommitQCNumber() + 1
		index := pbft.state.BlockNumber()
		len := pbft.currentValidatorLen()
		pbft.log.Debug("MissingPreCommit", "epoch", pbft.state.Epoch(), "viewNumber", pbft.state.ViewNumber(), "BlockNumber", index,"validatorLen", len)

		block := pbft.state.HighestPreCommitQCBlock()
		blockTime := common.MillisToTime(int64(block.Time()))

		//for index := begin; index <= end; index++ {
		size := pbft.state.PreCommitLenByNumber(index)
		pbft.log.Debug("The length of PreCommit", "number", index, "size", size)

		// We need sync prepare votes when a long time not arrived QC.
		if size < pbft.threshold(len) && time.Since(blockTime) >= syncPreCommitsInterval { // need sync prepare votes
			knownVotes := pbft.state.AllPreCommitByNumber(index)
			unKnownSet := utils.NewBitArray(uint32(len))
			for i := uint32(0); i < unKnownSet.Size(); i++ {
				if vote := pbft.csPool.GetPreCommit(pbft.state.Epoch(), index, 0, i); vote != nil {
					cMsg := vote.Msg.(ctypes.ConsensusMsg)
					if cMsg.ViewNum()==pbft.state.ViewNumber(){
						go pbft.ReceiveMessage(vote)
						continue
					}
				}
				if _, ok := knownVotes[i]; !ok {
					unKnownSet.SetIndex(i, true)
				}
			}

			v, err = &protocols.GetPreCommit{
				Epoch:      pbft.state.Epoch(),
				ViewNumber: pbft.state.ViewNumber(),
				BlockNumber: index,
				BlockIndex: 0,
				UnKnownSet: unKnownSet,
			}, nil
			//break
		}
		//}
		if v == nil {
			err = fmt.Errorf("not need sync preCommit")
		}
	}
	<-result
	return
}

// LatestStatus returns latest status.
func (pbft *Pbft) LatestStatus() (v *protocols.GetLatestStatus) {
	result := make(chan struct{})

	pbft.asyncCallCh <- func() {
		defer func() { result <- struct{}{} }()

		qcBn, qcHash := pbft.HighestPreCommitQCBlockBn()
		_, qc := pbft.blockTree.FindBlockAndQC(qcHash, qcBn)

		lockBn, lockHash := pbft.HighestLockBlockBn()
		_, lockQC := pbft.blockTree.FindBlockAndQC(lockHash, lockBn)

		v = &protocols.GetLatestStatus{
			BlockNumber:  qcBn,
			BlockHash:    qcHash,
			QuorumCert:   qc,
			LBlockNumber: lockBn,
			LBlockHash:   lockHash,
			LQuorumCert:  lockQC,
		}
	}
	<-result
	return
}

// OnPong is used to receive the average delay time.
func (pbft *Pbft) OnPong(nodeID string, netLatency int64) error {
	pbft.log.Trace("OnPong", "nodeID", nodeID, "netLatency", netLatency)
	pbft.netLatencyLock.Lock()
	defer pbft.netLatencyLock.Unlock()
	latencyList, exist := pbft.netLatencyMap[nodeID]
	if !exist {
		pbft.netLatencyMap[nodeID] = list.New()
		pbft.netLatencyMap[nodeID].PushBack(netLatency)
	} else {
		if latencyList.Len() > 5 {
			e := latencyList.Front()
			pbft.netLatencyMap[nodeID].Remove(e)
		}
		pbft.netLatencyMap[nodeID].PushBack(netLatency)
	}
	return nil
}

// BlockExists is used to query whether the specified block exists in this node.
func (pbft *Pbft) BlockExists(blockNumber uint64, blockHash common.Hash) error {
	result := make(chan error, 1)
	pbft.asyncCallCh <- func() {
		if (blockHash == common.Hash{}) {
			result <- fmt.Errorf("invalid blockHash")
			return
		}
		block := pbft.blockTree.FindBlockByHash(blockHash)
		if block = pbft.blockChain.GetBlock(blockHash, blockNumber); block == nil {
			result <- fmt.Errorf("not found block by hash:%s, number:%d", blockHash.TerminalString(), blockNumber)
			return
		}
		if block.Hash() != blockHash || blockNumber != block.NumberU64() {
			result <- fmt.Errorf("not match from block, hash:%s, number:%d, queriedHash:%s, queriedNumber:%d",
				blockHash.TerminalString(), blockNumber,
				block.Hash().TerminalString(), block.NumberU64())
			return
		}
		result <- nil
	}
	return <-result
}

// AvgLatency returns the average delay time of the specified node.
//
// The average is the average delay between the current
// node and all consensus nodes.
// Return value unit: milliseconds.
func (pbft *Pbft) AvgLatency() time.Duration {
	pbft.netLatencyLock.Lock()
	defer pbft.netLatencyLock.Unlock()
	// The intersection of peerSets and consensusNodes.
	target, err := pbft.network.AliveConsensusNodeIDs()
	if err != nil {
		return time.Duration(0)
	}
	var (
		avgSum     int64
		result     int64
		validCount int64
	)
	// Take 2/3 nodes from the target.
	var pair utils.KeyValuePairList
	for _, v := range target {
		if latencyList, exist := pbft.netLatencyMap[v]; exist {
			avg := calAverage(latencyList)
			pair.Push(utils.KeyValuePair{Key: v, Value: avg})
		}
	}
	sort.Sort(pair)
	if pair.Len() == 0 {
		return time.Duration(0)
	}
	validCount = int64(pair.Len() * 2 / 3)
	if validCount == 0 {
		validCount = 1
	}
	for _, v := range pair[:validCount] {
		avgSum += v.Value
	}

	result = avgSum / validCount
	pbft.log.Debug("Get avg latency", "avg", result)
	return time.Duration(result) * time.Millisecond
}

// DefaultAvgLatency returns the avg latency of default.
func (pbft *Pbft) DefaultAvgLatency() time.Duration {
	return time.Duration(protocols.DefaultAvgLatency) * time.Millisecond
}

func calAverage(latencyList *list.List) int64 {
	var (
		sum    int64
		counts int64
	)
	for e := latencyList.Front(); e != nil; e = e.Next() {
		if latency, ok := e.Value.(int64); ok {
			counts++
			sum += latency
		}
	}
	if counts > 0 {
		return sum / counts
	}
	return 0
}

func (pbft *Pbft) SyncPrepareBlock(id string, epoch uint64, blockNumber uint64, blockIndex uint32,viewNumber uint64) {
	if msg := pbft.csPool.GetPrepareBlock(epoch, blockNumber, blockIndex); msg != nil {
		go pbft.ReceiveMessage(msg)
	}
	if pbft.syncingCache.AddOrReplace(blockNumber) {
		msg := &protocols.GetPrepareBlock{Epoch: epoch, BlockNumber: blockNumber, BlockIndex: blockIndex,ViewNumber: viewNumber}
		if id == "" {
			pbft.network.PartBroadcast(msg)
			pbft.log.Debug("Send GetPrepareBlock by part broadcast", "msg", msg.String())
		} else {
			pbft.network.Send(id, msg)
			pbft.log.Debug("Send GetPrepareBlock", "peer", id, "msg", msg.String())
		}
	}
}

func (pbft *Pbft) SyncBlockQuorumCert(id string, blockNumber uint64, blockHash common.Hash, blockIndex uint32) {
	if msg := pbft.csPool.GetPrepareQC(pbft.state.Epoch(), blockNumber, blockIndex); msg != nil {
		go pbft.ReceiveMessage(msg)
	}
	if pbft.syncingCache.AddOrReplace(blockHash) {
		msg := &protocols.GetBlockQuorumCert{BlockHash: blockHash, BlockNumber: blockNumber}
		pbft.network.Send(id, msg)
		pbft.log.Debug("Send GetBlockQuorumCert", "peer", id, "msg", msg.String())
	}

}

func (pbft *Pbft) SyncBlockPreCommitQuorumCert(id string, blockNumber uint64, blockHash common.Hash, blockIndex uint32) {
	if msg := pbft.csPool.GetPreCommitQC(pbft.state.Epoch(), blockNumber, blockIndex); msg != nil {
		go pbft.ReceiveMessage(msg)
	}
	if pbft.syncingCache.AddOrReplace(blockHash) {
		msg := &protocols.GetBlockPreCommitQuorumCert{BlockHash: blockHash, BlockNumber: blockNumber}
		pbft.network.Send(id, msg)
		pbft.log.Debug("Send GetBlockPreCommitQuorumCert", "peer", id, "msg", msg.String())
	}

}