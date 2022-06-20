package pbft

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"Phoenix-Chain-Core/consensus/pbft/state"

	"Phoenix-Chain-Core/libs/common/math"
	"Phoenix-Chain-Core/libs/log"

	"Phoenix-Chain-Core/ethereum/node"

	ctypes "Phoenix-Chain-Core/consensus/pbft/types"
	"Phoenix-Chain-Core/consensus/pbft/wal"

	"Phoenix-Chain-Core/consensus/pbft/protocols"
	"Phoenix-Chain-Core/ethereum/core/types"
)

var (
	errNonContiguous = errors.New("non contiguous chain block state")
)

var (
	viewChangeQCPrefix = []byte("qc") // viewChangeQCPrefix + epoch (uint64 big endian) + viewNumber (uint64 big endian) -> viewChangeQC
	epochPrefix        = []byte("e")
)

// Bridge encapsulates functions required to update consensus state and consensus msg.
// As a bridge layer for pbft and wal.
type Bridge interface {
	UpdateChainState(qcState, lockState, commitState *protocols.State)
	ConfirmViewChange(epoch, viewNumber uint64, block *types.Block, qc *ctypes.QuorumCert, viewChangeQC *ctypes.ViewChangeQC, preEpoch, preViewNumber uint64)
	SendViewChange(view *protocols.ViewChange)
	SendPrepareBlock(pb *protocols.PrepareBlock)
	SendPrepareVote(block *types.Block, vote *protocols.PrepareVote)
	SendPreCommit(block *types.Block, vote *protocols.PreCommit)
	GetViewChangeQC(epoch uint64, blockNumber uint64,viewNumber uint64) (*ctypes.ViewChangeQC, error)

	Close()
}

// emptyBridge is a empty implementation for Bridge
type emptyBridge struct {
}

func (b *emptyBridge) UpdateChainState(qcState, lockState, commitState *protocols.State) {
}

func (b *emptyBridge) ConfirmViewChange(epoch, viewNumber uint64, block *types.Block, qc *ctypes.QuorumCert, viewChangeQC *ctypes.ViewChangeQC, preEpoch, preViewNumber uint64) {
}

func (b *emptyBridge) SendViewChange(view *protocols.ViewChange) {
}

func (b *emptyBridge) SendPrepareBlock(pb *protocols.PrepareBlock) {
}

func (b *emptyBridge) SendPrepareVote(block *types.Block, vote *protocols.PrepareVote) {
}

func (b *emptyBridge) SendPreCommit(block *types.Block, vote *protocols.PreCommit) {
}

func (b *emptyBridge) GetViewChangeQC(epoch uint64, blockNumber uint64,viewNumber uint64) (*ctypes.ViewChangeQC, error) {
	return nil, nil
}

func (b *emptyBridge) Close() {

}

// baseBridge is a default implementation for Bridge
type baseBridge struct {
	pbft *Pbft
}

// NewBridge creates a new Bridge to update consensus state and consensus msg.
func NewBridge(ctx *node.ServiceContext, pbft *Pbft) (Bridge, error) {
	if ctx == nil {
		return &emptyBridge{}, nil
	}
	baseBridge := &baseBridge{
		pbft: pbft,
	}
	return baseBridge, nil
}

// UpdateChainState tries to update consensus state to wal
// If the write fails, the process will stop
// lockChainState or commitChainState may be nil, if it is nil, we only append qc to the qcChain array
func (b *baseBridge) UpdateChainState(qcState, lockState, commitState *protocols.State) {
	tStart := time.Now()
	var chainState *protocols.ChainState
	var err error
	// load current consensus state
	b.pbft.wal.LoadChainState(func(cs *protocols.ChainState) error {
		chainState = cs
		return nil
	})

	if chainState == nil {
		log.Trace("ChainState is empty,may be the first time to update chainState")
		if err = b.newChainState(commitState, lockState, qcState); err != nil {
			panic(fmt.Sprintf("update chain state error: %s", err.Error()))
		}
		return
	}
	if !chainState.ValidChainState() {
		panic(fmt.Sprintf("invalid chain state from wal"))
	}
	walCommitNumber := chainState.Commit.QuorumCert.BlockNumber
	commitNumber := commitState.QuorumCert.BlockNumber
	if walCommitNumber != commitNumber && walCommitNumber+1 != commitNumber {
		log.Warn("The chainState of wal and updating are discontinuous,ignore this updating", "walCommit", chainState.Commit.String(), "updateCommit", commitState.String())
		return
	}

	if chainState.Commit.EqualState(commitState) && chainState.Lock.EqualState(lockState) && chainState.QC[0].QuorumCert.BlockNumber == qcState.QuorumCert.BlockNumber {
		err = b.addQCState(qcState, chainState)
	} else {
		err = b.newChainState(commitState, lockState, qcState)
	}

	if err != nil {
		panic(fmt.Sprintf("update chain state error: %s", err.Error()))
	}
	log.Info("Success to update chainState", "commitState", commitState.String(), "lockState", lockState.String(), "qcState", qcState.String(), "elapsed", time.Since(tStart))
}

// newChainState tries to update consensus state to wal
// Need to do continuous block check before writing.
func (b *baseBridge) newChainState(commit *protocols.State, lock *protocols.State, qc *protocols.State) error {
	log.Debug("New chainState", "commitState", commit.String(), "lockState", lock.String(), "qcState", qc.String())
	if !commit.ValidState() || !lock.ValidState() || !qc.ValidState() {
		return errNonContiguous
	}
	// check continuous block chain
	//if !contiguousChainBlock(commit.Block, lock.Block) || !contiguousChainBlock(lock.Block, qc.Block) {
	if !contiguousChainBlock(commit.Block, lock.Block) {
		return errNonContiguous
	}
	chainState := &protocols.ChainState{
		Commit: commit,
		Lock:   lock,
		QC:     []*protocols.State{qc},
	}
	return b.pbft.wal.UpdateChainState(chainState)
}

// addQCState tries to add consensus qc state to wal
// Need to do continuous block check before writing.
func (b *baseBridge) addQCState(qc *protocols.State, chainState *protocols.ChainState) error {
	log.Debug("Add qcState", "qcState", qc.String())
	//lock := chainState.Lock
	//// check continuous block chain
	//if !contiguousChainBlock(lock.Block, qc.Block) {
	//	return errNonContiguous
	//}
	chainState.QC = append(chainState.QC, qc)
	return b.pbft.wal.UpdateChainState(chainState)
}

// ConfirmViewChange tries to update ConfirmedViewChange consensus msg to wal.
// at the same time we will record the current fileID and fileSequence.
// the next time the phoenixchain node restart, we will recovery the msg from this check point.
func (b *baseBridge) ConfirmViewChange(epoch, viewNumber uint64, block *types.Block, qc *ctypes.QuorumCert, viewChangeQC *ctypes.ViewChangeQC, preEpoch, preViewNumber uint64) {
	tStart := time.Now()
	// save the identity location of the wal message in the file system
	meta := &wal.ViewChangeMessage{
		Epoch:      epoch,
		ViewNumber: viewNumber,
		BlockNumber: block.NumberU64(),
	}
	if err := b.pbft.wal.UpdateViewChange(meta); err != nil {
		panic(fmt.Sprintf("update viewChange meta error, err:%s", err.Error()))
	}
	// save ConfirmedViewChange message, the viewChangeQC is last viewChangeQC
	vc := &protocols.ConfirmedViewChange{
		Epoch:        epoch,
		ViewNumber:   viewNumber,
		Block:        block,
		QC:           qc,
		ViewChangeQC: viewChangeQC,
	}
	if err := b.pbft.wal.WriteSync(vc); err != nil {
		panic(fmt.Sprintf("write confirmed viewChange error, err:%s", err.Error()))
	}
	// save last viewChangeQC, for viewChangeQC synchronization
	if viewChangeQC != nil {
		b.pbft.wal.UpdateViewChangeQC(preEpoch, block.NumberU64(),preViewNumber, viewChangeQC)
	}
	log.Debug("Success to confirm viewChange", "confirmedViewChange", vc.String(), "elapsed", time.Since(tStart))
}

// SendViewChange tries to update SendViewChange consensus msg to wal.
func (b *baseBridge) SendViewChange(view *protocols.ViewChange) {
	tStart := time.Now()
	s := &protocols.SendViewChange{
		ViewChange: view,
	}
	if err := b.pbft.wal.WriteSync(s); err != nil {
		panic(fmt.Sprintf("write send viewChange error, err:%s", err.Error()))
	}
	log.Debug("Success to send viewChange", "view", view.String(), "elapsed", time.Since(tStart))
}

// SendPrepareBlock tries to update SendPrepareBlock consensus msg to wal.
func (b *baseBridge) SendPrepareBlock(pb *protocols.PrepareBlock) {
	tStart := time.Now()
	s := &protocols.SendPrepareBlock{
		Prepare: pb,
	}
	if err := b.pbft.wal.WriteSync(s); err != nil {
		panic(fmt.Sprintf("write send prepareBlock error, err:%s", err.Error()))
	}
	log.Debug("Success to send prepareBlock", "prepareBlock", pb.String(), "elapsed", time.Since(tStart))
}

// SendPrepareVote tries to update SendPrepareVote consensus msg to wal.
func (b *baseBridge) SendPrepareVote(block *types.Block, vote *protocols.PrepareVote) {
	tStart := time.Now()
	s := &protocols.SendPrepareVote{
		Block: block,
		Vote:  vote,
	}
	if err := b.pbft.wal.WriteSync(s); err != nil {
		panic(fmt.Sprintf("write send prepareVote error, err:%s", err.Error()))
	}
	log.Debug("Success to send prepareVote", "prepareVote", vote.String(), "elapsed", time.Since(tStart))
}

// SendPrepareVote tries to update SendPrepareVote consensus msg to wal.
func (b *baseBridge) SendPreCommit(block *types.Block, vote *protocols.PreCommit) {
	tStart := time.Now()
	s := &protocols.SendPreCommit{
		Block: block,
		Vote:  vote,
	}
	if err := b.pbft.wal.WriteSync(s); err != nil {
		panic(fmt.Sprintf("write send PreCommit error, err:%s", err.Error()))
	}
	log.Debug("Success to send PreCommit", "PreCommit", vote.String(), "elapsed", time.Since(tStart))
}

func (b *baseBridge) GetViewChangeQC(epoch uint64, blockNumber uint64,viewNumber uint64) (*ctypes.ViewChangeQC, error) {
	return b.pbft.wal.GetViewChangeQC(epoch, blockNumber,viewNumber)
}

func (b *baseBridge) Close() {
	b.pbft.wal.Close()
}

// recoveryChainState tries to recovery consensus chainState from wal when the phoenixchain node restart.
// need to do some necessary checks based on the latest blockchain block.
// execute commit/lock/qcs block and load the corresponding state to pbft consensus.
func (pbft *Pbft) recoveryChainState(chainState *protocols.ChainState) error {
	pbft.log.Info("Recover chain state from wal", "chainState", chainState.String())
	commit, lock, qcs := chainState.Commit, chainState.Lock, chainState.QC
	// The highest block that has been written to disk

	//rootBlock := pbft.blockChain.GetBlock(pbft.blockChain.CurrentHeader().Hash(), pbft.blockChain.CurrentHeader().Number.Uint64())
	rootBlock := pbft.blockChain.CurrentBlock()
	//currentBlock := pbft.blockChain.CurrentBlock()

	pbft.log.Debug("begin recovery","local CurrentBlock.blockNumber",rootBlock.NumberU64())

	isCurrent := rootBlock.NumberU64() == commit.Block.NumberU64() && rootBlock.Hash() == commit.Block.Hash()
	isParent := contiguousChainBlock(rootBlock, commit.Block)

	currentIsCommit:=false

	if !isCurrent && !isParent {
		return fmt.Errorf("recovery chain state errror,non contiguous chain block state, curNum:%d, curHash:%s, commitNum:%d, commitHash:%s", rootBlock.NumberU64(), rootBlock.Hash().String(), commit.Block.NumberU64(), commit.Block.Hash().String())
	}
	if isParent {
		// recovery commit state
		if err := pbft.recoveryCommitState(commit, rootBlock); err != nil {
			return err
		}
		pbft.log.Debug("recoveryCommitState success","commit",commit.String(),"now local CurrentBlock.blockNumber",pbft.blockChain.CurrentBlock().NumberU64())
		currentIsCommit=false
	}else {
		currentIsCommit=true
	}

	// recovery lock state
	rootBlock = commit.Block
	isCurrent = rootBlock.NumberU64() == lock.Block.NumberU64() && rootBlock.Hash() == lock.Block.Hash()
	isParent = contiguousChainBlock(rootBlock, lock.Block)
	if !isCurrent && !isParent {
		return fmt.Errorf("recoveryLockState recovery chain state error,non contiguous chain block state, curNum:%d, curHash:%s, lockNum:%d, lockHash:%s", rootBlock.NumberU64(), rootBlock.Hash().String(), lock.Block.NumberU64(), lock.Block.Hash().String())
	}
	if isParent {
		if err := pbft.recoveryLockState(lock, commit.Block); err != nil {
			return err
		}
		pbft.log.Debug("recoveryLockState success","lock",lock.String(),"now local CurrentBlock.blockNumber",pbft.blockChain.CurrentBlock().NumberU64())
	} else {
		if !currentIsCommit{
			pbft.state.SetHighestLockBlock(lock.Block)
			//if err := pbft.recoveryLockState(lock, currentBlock); err != nil {
			//	return err
			//}
		}
	}

	// recovery qc state
	rootBlock = lock.Block
	isCurrent = rootBlock.NumberU64() == qcs[0].Block.NumberU64() && rootBlock.Hash() == qcs[0].Block.Hash()
	isParent = contiguousChainBlock(rootBlock, qcs[0].Block)
	if !isCurrent && !isParent {
		return fmt.Errorf("recoveryQCState recovery chain state error,non contiguous chain block state, curNum:%d, curHash:%s, commitNum:%d, commitHash:%s", rootBlock.NumberU64(), rootBlock.Hash().String(), qcs[0].Block.NumberU64(), qcs[0].Block.Hash().String())
	}
	if isParent {
		if err := pbft.recoveryQCState(qcs, lock.Block); err != nil {
			return err
		}
		pbft.log.Debug("recoveryQCState success","qcs",qcs[0].String(),"now local CurrentBlock.blockNumber",pbft.blockChain.CurrentBlock().NumberU64())
	} else {
		if !currentIsCommit{
			pbft.recoveryHighestQCBlock(qcs)
			//if err := pbft.recoveryQCState(qcs, currentBlock); err != nil {
			//	return err
			//}
		}
	}
	return nil
}

func (pbft *Pbft) recoveryCommitState(commit *protocols.State, parent *types.Block) error {
	log.Info("Recover commit state", "commitNumber", commit.Block.NumberU64(), "commitHash", commit.Block.Hash(), "parentNumber", parent.NumberU64(), "parentHash", parent.Hash())
	// recovery commit state
	if err := pbft.executeBlock(commit.Block, parent, math.MaxUint32); err != nil {
		return err
	}
	// write commit block to chain
	extra, err := ctypes.EncodeExtra(byte(pbftVersion), commit.QuorumCert)
	if err != nil {
		return err
	}
	commit.Block.SetExtraData(extra)
	if err := pbft.blockCacheWriter.WriteBlock(commit.Block); err != nil {
		return err
	}
	if err := pbft.validatorPool.Commit(commit.Block); err != nil {
		return err
	}
	pbft.recoveryChainStateProcess(protocols.CommitState, commit)
	pbft.blockTree.NewRoot(commit.Block)
	return nil
}

func (pbft *Pbft) recoveryLockState(lock *protocols.State, parent *types.Block) error {
	// recovery lock state
	if err := pbft.executeBlock(lock.Block, parent, math.MaxUint32); err != nil {
		return err
	}
	pbft.recoveryChainStateProcess(protocols.LockState, lock)
	return nil
}

func (pbft *Pbft) recoveryQCState(qcs []*protocols.State, parent *types.Block) error {
	// recovery qc state
	for _, qc := range qcs {
		if err := pbft.executeBlock(qc.Block, parent, math.MaxUint32); err != nil {
			return err
		}
		pbft.recoveryChainStateProcess(protocols.QCState, qc)
	}
	return nil
}

func (pbft *Pbft) recoveryHighestQCBlock(qcs []*protocols.State) error {
	// recovery qc state
	for _, qc := range qcs {
		pbft.TrySetHighestQCBlock(qc.Block)
		pbft.TrySetHighestPreCommitQCBlock(qc.Block)
	}
	return nil
}

//func (pbft *Pbft) recoveryPreCommitQCState(qcs []*protocols.State, parent *types.Block) error {
//	// recovery qc state
//	for _, qc := range qcs {
//		if err := pbft.executeBlock(qc.Block, parent, math.MaxUint32); err != nil {
//			return err
//		}
//		pbft.recoveryChainStateProcess(protocols.PreCommitQCState, qc)
//	}
//	return nil
//}

// recoveryChainStateProcess tries to recovery the corresponding state to pbft consensus.
func (pbft *Pbft) recoveryChainStateProcess(stateType uint16, s *protocols.State) {
	pbft.trySwitchValidator(s.Block.NumberU64())
	pbft.tryWalChangeView(s.QuorumCert.Epoch, s.QuorumCert.ViewNumber, s.Block, s.QuorumCert, nil)
	pbft.state.AddQCBlock(s.Block, s.QuorumCert)
	pbft.state.AddQC(s.QuorumCert)
	pbft.state.AddPreCommitQC(s.QuorumCert)
	pbft.blockTree.InsertQCBlock(s.Block, s.QuorumCert)
	pbft.tryAddStateBlockNumber(s.QuorumCert)
	pbft.state.SetExecuting(s.QuorumCert.BlockIndex, true)
	pbft.state.SetMaxExecutedBlockNumber(s.QuorumCert.BlockNumber)

	switch stateType {
	case protocols.CommitState:
		pbft.TrySetHighestPreCommitQCBlock(s.Block)
		pbft.state.SetHighestCommitBlock(s.Block)
	case protocols.LockState:
		pbft.state.SetHighestLockBlock(s.Block)
		pbft.TrySetHighestQCBlock(s.Block)
		pbft.TrySetHighestPreCommitQCBlock(s.Block)
	//case protocols.QCState:
	//	pbft.TrySetHighestQCBlock(s.Block)
	case protocols.QCState:
		pbft.TrySetHighestQCBlock(s.Block)
		pbft.TrySetHighestPreCommitQCBlock(s.Block)
	}

	// The state may have reached the automatic switch point, then advance to the next view
	if pbft.validatorPool.EqualSwitchPoint(s.Block.NumberU64()) {
		pbft.log.Info("QCBlock is equal to switchPoint, change epoch", "state", s.String(), "view", pbft.state.ViewString())
		pbft.tryWalChangeView(pbft.state.Epoch()+1, state.DefaultViewNumber, s.Block, s.QuorumCert, nil)
		return
	}
	if s.QuorumCert.BlockIndex+1 == pbft.config.Sys.Amount {
		pbft.log.Info("QCBlock is the last index on the view, change view", "state", s.String(), "view", pbft.state.ViewString())
		pbft.tryWalChangeView(pbft.state.Epoch(), pbft.state.ViewNumber()+1, s.Block, s.QuorumCert, nil)
		return
	}
}

// trySwitch tries to switch next validator.
func (pbft *Pbft) trySwitchValidator(blockNumber uint64) {
	if pbft.validatorPool.ShouldSwitch(blockNumber) {
		if err := pbft.validatorPool.Update(blockNumber, pbft.state.Epoch()+1, pbft.eventMux); err != nil {
			pbft.log.Debug("Update validator error", "err", err.Error())
		}
	}
}

// tryWalChangeView tries to change view.
func (pbft *Pbft) tryWalChangeView(epoch, viewNumber uint64, block *types.Block, qc *ctypes.QuorumCert, viewChangeQC *ctypes.ViewChangeQC) {
	if epoch > pbft.state.Epoch()|| epoch == pbft.state.Epoch() && block.NumberU64()+1 > pbft.state.BlockNumber()|| epoch == pbft.state.Epoch() && block.NumberU64()+1 == pbft.state.BlockNumber()&& viewNumber > pbft.state.ViewNumber() {
		pbft.changeView(epoch, viewNumber, block, qc, viewChangeQC)
	}
}

// tryWalChangeView tries to change view.
func (pbft *Pbft) walChangeView(epoch, viewNumber uint64, block *types.Block, qc *ctypes.QuorumCert, viewChangeQC *ctypes.ViewChangeQC) {
	if viewChangeQC!=nil{
		if epoch > pbft.state.Epoch()|| epoch == pbft.state.Epoch() && block.NumberU64()+1 > pbft.state.BlockNumber()||epoch == pbft.state.Epoch() && block.NumberU64()+1 == pbft.state.BlockNumber()&& viewNumber > pbft.state.ViewNumber() {
			pbft.changeView(epoch, viewNumber, block, qc, viewChangeQC)
		}
		return
	}
	if epoch > pbft.state.Epoch()|| epoch == pbft.state.Epoch() && block.NumberU64()+1 >= pbft.state.BlockNumber()&& viewNumber >= pbft.state.ViewNumber() {
		pbft.changeView(epoch, viewNumber, block, qc, viewChangeQC)
	}
}

// recoveryMsg tries to recovery consensus msg from wal when the phoenixchain node restart.
func (pbft *Pbft) recoveryMsg(msg interface{}) error {
	pbft.log.Info("Recover journal message from wal", "msgType", reflect.TypeOf(msg))

	switch m := msg.(type) {
	case *protocols.ConfirmedViewChange:
		pbft.log.Debug("Load journal message from wal", "msgType", reflect.TypeOf(msg), "confirmedViewChange", m.String())
		pbft.walChangeView(m.Epoch, m.ViewNumber, m.Block, m.QC, m.ViewChangeQC)

	case *protocols.SendViewChange:
		pbft.log.Debug("Load journal message from wal", "msgType", reflect.TypeOf(msg), "sendViewChange", m.String())

		should, err := pbft.shouldViewChangeRecovery(m)
		if err != nil {
			return err
		}
		if should {
			node, err := pbft.validatorPool.GetValidatorByNodeID(m.ViewChange.Epoch, pbft.config.Option.NodeID)
			if err != nil {
				return err
			}
			pbft.state.AddViewChange(uint32(node.Index), m.ViewChange)
		}

	case *protocols.SendPrepareBlock:
		pbft.log.Debug("Load journal message from wal", "msgType", reflect.TypeOf(msg), "sendPrepareBlock", m.String())

		should, err := pbft.shouldRecovery(m)
		if err != nil {
			return err
		}
		if should {
			// execute block
			block := m.Prepare.Block
			if pbft.state.ViewBlockByIndex(m.Prepare.BlockNum()) == nil {
				if err := pbft.executeBlock(block, nil, m.Prepare.BlockIndex); err != nil {
					return err
				}
				pbft.state.SetExecuting(m.Prepare.BlockIndex, true)
				pbft.state.SetMaxExecutedBlockNumber(m.Prepare.BlockNum())
			}
			pbft.signMsgByBls(m.Prepare)
			pbft.state.AddPrepareBlock(m.Prepare)
		}

	case *protocols.SendPrepareVote:
		pbft.log.Debug("Load journal message from wal", "msgType", reflect.TypeOf(msg), "sendPrepareVote", m.String())

		should, err := pbft.shouldRecovery(m)
		if err != nil {
			return err
		}
		if should {
			// execute block
			block := m.Block
			if pbft.state.ViewBlockByIndex(m.Vote.BlockNumber) == nil {
				if err := pbft.executeBlock(block, nil, m.Vote.BlockIndex); err != nil {
					return err
				}
				pbft.state.SetExecuting(m.Vote.BlockIndex, true)
				pbft.state.SetMaxExecutedBlockNumber(m.Vote.BlockNumber)
				pbft.state.AddPrepareBlock(&protocols.PrepareBlock{
					Epoch:      m.Vote.Epoch,
					ViewNumber: m.Vote.ViewNumber,
					Block:      block,
					BlockIndex: m.Vote.BlockIndex,
				})
			}

			//pbft.state.HadSendPrepareVote().Push(m.Vote)
			pbft.tryAddHadSendVote(m.Vote.BlockNumber-1,m.Vote)
			node, _ := pbft.validatorPool.GetValidatorByNodeID(m.Vote.Epoch, pbft.config.Option.NodeID)
			pbft.state.AddPrepareVote(uint32(node.Index), m.Vote)
		}
	case *protocols.SendPreCommit:
		pbft.log.Debug("Load journal message from wal", "msgType", reflect.TypeOf(msg), "sendPrepareVote", m.String())

		should, err := pbft.shouldRecovery(m)
		if err != nil {
			return err
		}
		if should {
			// execute block
			block := m.Block
			if pbft.state.ViewBlockByIndex(m.Vote.BlockNumber) == nil {
				if err := pbft.executeBlock(block, nil, m.Vote.BlockIndex); err != nil {
					return err
				}
				pbft.state.SetExecuting(m.Vote.BlockIndex, true)
				pbft.state.SetMaxExecutedBlockNumber(m.Vote.BlockNumber)
				pbft.state.AddPrepareBlock(&protocols.PrepareBlock{
					Epoch:      m.Vote.Epoch,
					ViewNumber: m.Vote.ViewNumber,
					Block:      block,
					BlockIndex: m.Vote.BlockIndex,
				})
			}

			//pbft.state.HadSendPreCommit=m.Vote
			pbft.tryAddHadSendPreCommit(m.Vote.BlockNumber-1,m.Vote)
			node, _ := pbft.validatorPool.GetValidatorByNodeID(m.Vote.Epoch, pbft.config.Option.NodeID)
			pbft.state.AddPreCommit(uint32(node.Index), m.Vote)
		}
	}
	return nil
}

// contiguousChainBlock check if the two incoming blocks are continuous.
func contiguousChainBlock(p *types.Block, s *types.Block) bool {
	contiguous := p.NumberU64()+1 == s.NumberU64() && p.Hash() == s.ParentHash()
	if !contiguous {
		log.Info("Non contiguous block", "sNumber", s.NumberU64(), "sParentHash", s.ParentHash(), "pNumber", p.NumberU64(), "pHash", p.Hash())
	}
	return contiguous
}

// executeBlock call blockCacheWriter to execute block.
func (pbft *Pbft) executeBlock(block *types.Block, parent *types.Block, index uint32) error {
	if parent == nil {
		if parent, _ = pbft.blockTree.FindBlockAndQC(block.ParentHash(), block.NumberU64()-1); parent == nil {
			if parent = pbft.state.ViewBlockByIndex(block.NumberU64()-1); parent == nil {
				return fmt.Errorf("find executable block's parent failed, blockNum:%d, blockHash:%s", block.NumberU64(), block.Hash().String())
			}
		}
	}
	if err := pbft.blockCacheWriter.Execute(block, parent); err != nil {
		return fmt.Errorf("execute block failed, blockNum:%d, blockHash:%s, parentNum:%d, parentHash:%s, err:%s", block.NumberU64(), block.Hash().String(), parent.NumberU64(), parent.Hash().String(), err.Error())
	}
	return nil
}

// shouldRecovery check if the consensus msg needs to be recovery.
// if the msg does not belong to the current view or the msg number is smaller than the qc number discard it.
func (pbft *Pbft) shouldRecovery(msg protocols.WalMsg) (bool, error) {
	if pbft.higherViewState(msg) {
		return false, fmt.Errorf("higher view state, curEpoch:%d, curBlockNum:%d, curViewNum:%d, msgEpoch:%d,msgBlockNum:%d, msgViewNum:%d", pbft.state.Epoch(),pbft.state.BlockNumber(), pbft.state.ViewNumber(), msg.Epoch(), msg.BlockNumber(),msg.ViewNumber())
	}
	if pbft.lowerViewState(msg) {
		// The state may have reached the automatic switch point, so advance to the next view
		return false, nil
	}
	// equalViewState
	highestQCBlockBn, _ := pbft.HighestPreCommitQCBlockBn()
	return msg.BlockNumber() > highestQCBlockBn+1, nil
}

// shouldRecovery check if the consensus msg needs to be recovery.
// if the msg does not belong to the current view or the msg number is smaller than the qc number discard it.
func (pbft *Pbft) shouldViewChangeRecovery(msg protocols.WalMsg) (bool, error) {
	if msg.Epoch() > pbft.state.Epoch() || msg.Epoch() == pbft.state.Epoch() && msg.BlockNumber()+1 > pbft.state.BlockNumber()|| msg.Epoch() == pbft.state.Epoch() && msg.BlockNumber()+1 == pbft.state.BlockNumber()&& msg.ViewNumber() > pbft.state.ViewNumber() {
		return false, fmt.Errorf("higher view state, curEpoch:%d, curBlockNum:%d, curViewNum:%d, msgEpoch:%d,msgBlockNum:%d, msgViewNum:%d", pbft.state.Epoch(),pbft.state.BlockNumber(), pbft.state.ViewNumber(), msg.Epoch(), msg.BlockNumber(),msg.ViewNumber())
	}
	if msg.Epoch() < pbft.state.Epoch() || msg.Epoch() == pbft.state.Epoch() && msg.BlockNumber()+1 < pbft.state.BlockNumber()|| msg.Epoch() == pbft.state.Epoch() && msg.BlockNumber()+1 == pbft.state.BlockNumber()&& msg.ViewNumber() < pbft.state.ViewNumber(){
		// The state may have reached the automatic switch point, so advance to the next view
		return false, nil
	}
	return true, nil
}

// equalViewState check if the msg view is equal with current.
func (pbft *Pbft) equalViewState(msg protocols.WalMsg) bool {
	return msg.Epoch() == pbft.state.Epoch() && msg.BlockNumber() == pbft.state.BlockNumber()&& msg.ViewNumber() == pbft.state.ViewNumber()
}

// lowerViewState check if the msg Epoch is lower than current.
func (pbft *Pbft) lowerViewState(msg protocols.WalMsg) bool {
	return msg.Epoch() < pbft.state.Epoch() || msg.Epoch() == pbft.state.Epoch() && msg.BlockNumber() < pbft.state.BlockNumber()|| msg.Epoch() == pbft.state.Epoch() && msg.BlockNumber() == pbft.state.BlockNumber()&& msg.ViewNumber() < pbft.state.ViewNumber()
}

// higherViewState check if the msg Epoch is higher than current.
func (pbft *Pbft) higherViewState(msg protocols.WalMsg) bool {
	return msg.Epoch() > pbft.state.Epoch() || msg.Epoch() == pbft.state.Epoch() && msg.BlockNumber() > pbft.state.BlockNumber()|| msg.Epoch() == pbft.state.Epoch() && msg.BlockNumber() == pbft.state.BlockNumber()&& msg.ViewNumber() > pbft.state.ViewNumber()
}
