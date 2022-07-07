package rules

import (
	"fmt"
	"time"

	"Phoenix-Chain-Core/consensus/pbft/validator"
	"Phoenix-Chain-Core/libs/common"

	"Phoenix-Chain-Core/consensus/pbft/protocols"
	"Phoenix-Chain-Core/consensus/pbft/state"
	ctypes "Phoenix-Chain-Core/consensus/pbft/types"
	"Phoenix-Chain-Core/ethereum/core/types"
)

const (
	riseTimeLimit = 1000 * time.Millisecond
)

type SafetyError interface {
	error
	Common() bool
	Fetch() bool   //Is the error need fetch
	NewView() bool //need change view
	FetchPrepare() bool
}

type safetyError struct {
	text         string
	common       bool
	fetch        bool
	newView      bool
	fetchPrepare bool
}

func (s safetyError) Error() string {
	return s.text
}

func (s safetyError) Common() bool {
	return s.common
}

func (s safetyError) Fetch() bool {
	return s.fetch
}

func (s safetyError) NewView() bool {
	return s.newView
}

func (s safetyError) FetchPrepare() bool {
	return s.fetchPrepare
}

func newCommonError(text string) SafetyError {
	return &safetyError{
		text:         text,
		common:       true,
		fetch:        false,
		newView:      false,
		fetchPrepare: false,
	}
}

func newFetchError(text string) SafetyError {
	return &safetyError{
		text:         text,
		common:       false,
		fetch:        true,
		newView:      false,
		fetchPrepare: false,
	}
}

func newViewError(text string) SafetyError {
	return &safetyError{
		text:         text,
		common:       false,
		fetch:        false,
		newView:      true,
		fetchPrepare: false,
	}
}

func newFetchPrepareError(text string) SafetyError {
	return &safetyError{
		text:         text,
		common:       false,
		fetch:        false,
		newView:      false,
		fetchPrepare: true,
	}
}

func newError(text string) SafetyError {
	return &safetyError{
		text:         text,
		common:       false,
		fetch:        false,
		newView:      false,
		fetchPrepare: false,
	}
}

type SafetyRules interface {
	// Security rules for proposed blocks
	PrepareBlockRules(block *protocols.PrepareBlock) SafetyError

	// Security rules for proposed votes
	PrepareVoteRules(vote *protocols.PrepareVote) SafetyError

	// Security rules for precommits
	PreCommitRules(vote *protocols.PreCommit) SafetyError

	// Security rules for viewChange
	ViewChangeRules(vote *protocols.ViewChange) SafetyError

	// Security rules for qcblock
	QCBlockRules(block *types.Block, qc *ctypes.QuorumCert) SafetyError
}

type baseSafetyRules struct {
	viewState     *state.ViewState
	blockTree     *ctypes.BlockTree
	config        *ctypes.Config
	validatorPool *validator.ValidatorPool
}

// PrepareBlock rules
// 1.Less than local viewNumber drop
// 2.Synchronization greater than local viewNumber
// 3.Lost more than the time window
func (r *baseSafetyRules) PrepareBlockRules(block *protocols.PrepareBlock) SafetyError {
	isCommitChild := func() bool {
		//return block.BlockNum() == r.viewState.HighestQCBlock().NumberU64()+1 &&
		//	block.Block.ParentHash() == r.viewState.HighestQCBlock().Hash() &&
		//	r.blockTree.FindBlockByHash(block.Block.ParentHash()) != nil
		parent := r.blockTree.FindBlockByHash(block.Block.ParentHash())
		if parent == nil {
			return false
		}
		return parent.NumberU64() == r.viewState.HighestPreCommitQCBlock().NumberU64()
	}

	isLockChild := func() bool {
		return block.BlockNum() == r.viewState.HighestLockBlock().NumberU64()+1 &&
			block.Block.ParentHash() == r.viewState.HighestLockBlock().Hash() &&
			r.blockTree.FindBlockByHash(block.Block.ParentHash()) != nil
	}

	acceptViewChangeQC := func() bool {
		if block.ViewChangeQC == nil {
			return r.config.Sys.Amount == r.viewState.MaxQCIndex()+1 || r.validatorPool.EqualSwitchPoint(block.Block.NumberU64()-1)
		}
		_, _, _, _, hash, number := block.ViewChangeQC.MaxBlock()
		return number+1 == block.Block.NumberU64() && block.Block.ParentHash() == hash
	}

	isFirstBlock := func() bool {
		return block.BlockIndex == 0
	}

	doubtDuplicate := func() bool {
		//min:=r.viewState.MinViewBlockNumber()
		//for i := min; i <= r.viewState.MaxViewBlockNumber(); i++ {
		//	local := r.viewState.ViewBlockByIndex(i)
		//	if local != nil && local.NumberU64() == block.BlockNum() && local.Hash() != block.Block.Hash() {
		//		return true
		//	}
		//}
		local := r.viewState.ViewBlockByIndex(block.BlockNum())
		if local != nil && local.NumberU64() == block.BlockNum() && local.Hash() != block.Block.Hash() {
			return true
		}
		return false
	}

	acceptBlockTime := func() SafetyError {
		parentBlock := r.blockTree.FindBlockByHash(block.Block.ParentHash())
		//if parentBlock == nil && !isFirstBlock() {
		//	parentBlock = r.viewState.ViewBlockByIndex(block.BlockIndex - 1)
		//}
		if parentBlock == nil {
			return newCommonError(fmt.Sprintf("parentBlock does not exist(parentHash:%s, parentNum:%d, blockIndex:%d)", block.Block.ParentHash().String(), block.BlockNum()-1, block.BlockIndex))
		}

		// parent block time must before than prepareBlock time
		if parentBlock.Time() > block.Block.Time() {
			return newCommonError(fmt.Sprintf("prepareBlock time is before parent(parentHash:%s, parentNum:%d, parentTime:%d, blockHash:%s, blockNum:%d, blockTime:%d)",
				parentBlock.Hash().String(), parentBlock.NumberU64(), parentBlock.Time(), block.Block.Hash().String(), block.BlockNum(), block.Block.Time()))
		}

		// prepareBlock time cannot exceed system time by 1000 ms(default)
		riseTime := common.Millis(time.Now().Add(riseTimeLimit))
		if block.Block.Time() > uint64(riseTime) {
			return newCommonError(fmt.Sprintf("prepareBlock time is advance(blockHash:%s, blockNum:%d, blockTime:%d, sysTime:%d)",
				block.Block.Hash().String(), block.BlockNum(), block.Block.Time(), riseTime))
		}
		return nil
	}

	// if local epoch and viewNumber is the same with msg
	// Note:
	// 1. block index is greater than or equal to the Amount value, discard the msg.
	// 2. the index block exist, discard the msg.
	// 3. the first index block of this view, and is a subblock of the local highestQC or highestLock, accept the msg.
	// 4. the previous index block does not exist, discard the msg.
	// 5. block index is continuous, but number or hash is not, discard the msg.
	acceptIndexBlock := func() SafetyError {
		if block.BlockIndex >= r.config.Sys.Amount {
			return newCommonError(fmt.Sprintf("blockIndex higher than amount(index:%d, amount:%d)", block.BlockIndex, r.config.Sys.Amount))
		}
		if doubtDuplicate() {
			return nil
		}
		current := r.viewState.ViewBlockByIndex(block.BlockNum())
		if current != nil {
			return newCommonError(fmt.Sprintf("blockNum already exists(blockNum:%d)", block.BlockNum()))
		}
		if isFirstBlock() {
			if !isCommitChild() && !isLockChild() {
				return newCommonError(fmt.Sprintf("the first index block is not contiguous by local highestQC or highestLock"))
			}
			return acceptBlockTime()
		}
		// If block index is greater than 0, query the parent block from the viewBlocks
		//pre := r.viewState.ViewBlockByIndex(block.BlockIndex - 1)
		//if pre == nil {
		//	return newFetchPrepareError(fmt.Sprintf("previous index block not exists,discard msg(index:%d)", block.BlockIndex-1))
		//}
		//if pre.NumberU64()+1 != block.BlockNum() || pre.Hash() != block.Block.ParentHash() {
		//	return newCommonError(fmt.Sprintf("non contiguous index block(preIndex:%d,preNum:%d,preHash:%s,curIndex:%d,curNum:%d,curParentHash:%s)",
		//		block.BlockIndex-1, pre.NumberU64(), pre.Hash().String(), block.BlockIndex, block.BlockNum(), block.Block.ParentHash().String()))
		//}
		// Verify the prepareBlock time
		if err := acceptBlockTime(); err != nil {
			return err
		}
		return nil
	}

	changeEpochBlockRules := func(block *protocols.PrepareBlock) SafetyError {
		if r.viewState.Epoch() > block.Epoch {
			return newCommonError(fmt.Sprintf("epoch too low(local:%d, msg:%d)", r.viewState.Epoch(), block.Epoch))
		}

		if isFirstBlock() && acceptViewChangeQC() && isCommitChild() {
			return newViewError("new epoch, need change view")
		}

		b, _ := r.blockTree.FindBlockAndQC(block.Block.ParentHash(), block.BlockNum()-1)
		if b == nil {
			return newCommonError(fmt.Sprintf("epoch higher than local, but not find parent block(local:%d, msg:%d)", r.viewState.Epoch(), block.Epoch))
		}

		return newFetchError(fmt.Sprintf("epoch higher than local(local:%d, msg:%d)", r.viewState.Epoch(), block.Epoch))
	}

	if r.viewState.Epoch() != block.Epoch {
		return changeEpochBlockRules(block)
	}

	if r.viewState.BlockNumber() > block.BlockNum() {
		return newCommonError(fmt.Sprintf("BlockNumber too low(local:%d, msg:%d)", r.viewState.BlockNumber(), block.BlockNum()))
	}
	if r.viewState.BlockNumber() < block.BlockNum() {
		if block.BlockNum()>1{
			b, _ := r.blockTree.FindBlockAndQC(block.Block.ParentHash(), block.BlockNum()-1)
			if b == nil {
				//if block.BlockNum()-r.viewState.BlockNumber()==1{
				//	return newCommonError(fmt.Sprintf("BlockNumber higher than local, but not find parent block(local:%d, msg:%d)", r.viewState.BlockNumber(), block.BlockNum()))
				//}else {
				//	return newFetchError(fmt.Sprintf("BlockNumber higher than local(local:%d, msg:%d)", r.viewState.BlockNumber(), block.BlockNum()))
				//}
				return newFetchError(fmt.Sprintf("BlockNumber higher than local(local:%d, msg:%d)", r.viewState.BlockNumber(), block.BlockNum()))
			}
			//return newFetchError(fmt.Sprintf("BlockNumber higher than local(local:%d, msg:%d)", r.viewState.BlockNumber(), block.BlockNum()))
		}
	}

	if r.viewState.ViewNumber() > block.ViewNumber {
		return newCommonError(fmt.Sprintf("viewNumber too low(local:%d, msg:%d)", r.viewState.ViewNumber(), block.ViewNumber))
	}

	if r.viewState.ViewNumber() < block.ViewNumber {
		if (isCommitChild()|| isLockChild()) && acceptViewChangeQC() {
			return newViewError("need change view")
		}
		return newFetchError(fmt.Sprintf("viewNumber higher than local(local:%d, msg:%d)", r.viewState.ViewNumber(), block.ViewNumber))
	}

	// if local epoch and viewNumber is the same with msg
	if err := acceptIndexBlock(); err != nil {
		return err
	}

	if r.viewState.IsDeadline() {
		return newCommonError(fmt.Sprintf("view's deadline is expire(over:%s)", time.Since(r.viewState.Deadline())))
	}
	return nil
}

// PrepareVote rules
// 1.Less than local viewNumber drop
// 2.Synchronization greater than local viewNumber
// 3.Lost more than the time window
func (r *baseSafetyRules) PrepareVoteRules(vote *protocols.PrepareVote) SafetyError {
	alreadyQCBlock := func() bool {
		return r.blockTree.FindBlockByHash(vote.BlockHash) != nil || vote.BlockNumber <= r.viewState.HighestPreCommitQCBlock().NumberU64()
	}

	existsPrepare := func() bool {
		prepare := r.viewState.ViewBlockByIndex(vote.BlockNumber)
		return prepare != nil && prepare.NumberU64() == vote.BlockNumber && prepare.Hash() == vote.BlockHash
	}

	doubtDuplicate := func() bool {
		//min:=r.viewState.MinViewVoteIndex()
		//for i := min; i <= r.viewState.MaxViewVoteIndex(); i++ {
		//	local := r.viewState.FindPrepareVote(vote.BlockNumber, vote.ValidatorIndex)
		//	if local != nil && local.BlockNumber == vote.BlockNumber && local.BlockHash != vote.BlockHash {
		//		return true
		//	}
		//}
		local := r.viewState.FindPrepareVote(vote.BlockNumber, vote.ValidatorIndex)
		if local != nil && local.BlockNumber == vote.BlockNumber && local.BlockHash != vote.BlockHash {
			return true
		}
		return false
	}

	acceptIndexVote := func() SafetyError {
		if vote.BlockIndex >= r.config.Sys.Amount {
			return newCommonError(fmt.Sprintf("voteIndex higher than amount(index:%d, amount:%d)", vote.BlockIndex, r.config.Sys.Amount))
		}
		if doubtDuplicate() {
			return nil
		}
		if r.viewState.FindPrepareVote(vote.BlockNumber, vote.ValidatorIndex) != nil {
			return newCommonError(fmt.Sprintf("prepare vote has exist(blockIndex:%d, validatorIndex:%d)", vote.BlockIndex, vote.ValidatorIndex))
		}
		if !existsPrepare() {
			return newFetchPrepareError(fmt.Sprintf("current index block  not existed,discard msg(number:%d)", vote.BlockNumber))
		}
		return nil
	}

	if alreadyQCBlock() {
		return newCommonError(fmt.Sprintf("current index block is already qc block,discard msg(index:%d,number:%d,hash:%s)", vote.BlockIndex, vote.BlockNumber, vote.BlockHash.String()))
	}
	if r.viewState.Epoch() != vote.Epoch {
		return r.changeEpochVoteRules(vote)
	}

	if r.viewState.BlockNumber() > vote.BlockNum() {
		return newCommonError(fmt.Sprintf("BlockNumber too low(local:%d, msg:%d)", r.viewState.BlockNumber(), vote.BlockNum()))
	}
	if r.viewState.BlockNumber() < vote.BlockNum() {
		if vote.BlockNum()>1{
			b, _ := r.blockTree.FindBlockAndQC(vote.ParentQC.BlockHash, vote.BlockNum()-1)
			if b == nil {
				//if vote.BlockNum()-r.viewState.BlockNumber()==1{
				//	return newCommonError(fmt.Sprintf("BlockNumber higher than local, but not find parent block(local:%d, msg:%d)", r.viewState.BlockNumber(), vote.BlockNum()))
				//}else {
				//	return newFetchError(fmt.Sprintf("BlockNumber higher than local(local:%d, msg:%d)", r.viewState.BlockNumber(), vote.BlockNum()))
				//}
				return newFetchError(fmt.Sprintf("BlockNumber higher than local(local:%d, msg:%d)", r.viewState.BlockNumber(), vote.BlockNum()))
			}
			//return newFetchError(fmt.Sprintf("BlockNumber higher than local(local:%d, msg:%d)", r.viewState.BlockNumber(), vote.BlockNum()))
		}
	}

	if r.viewState.ViewNumber() > vote.ViewNumber {
		return newCommonError(fmt.Sprintf("viewNumber too low(local:%d, msg:%d)", r.viewState.ViewNumber(), vote.ViewNumber))
	}

	if r.viewState.ViewNumber() < vote.ViewNumber {
		return newCommonError(fmt.Sprintf("viewNumber higher than local(local:%d, msg:%d)", r.viewState.ViewNumber(), vote.ViewNumber))
	}

	// if local epoch and viewNumber is the same with msg
	if err := acceptIndexVote(); err != nil {
		return err
	}

	if r.viewState.IsDeadline() {
		return newCommonError(fmt.Sprintf("view's deadline is expire(over:%d)", time.Since(r.viewState.Deadline())))
	}
	return nil
}

// PreCommit rules
// 1.Less than local viewNumber drop
// 2.Synchronization greater than local viewNumber
// 3.Lost more than the time window
func (r *baseSafetyRules) PreCommitRules(vote *protocols.PreCommit) SafetyError {
	alreadyPreCommitQCBlock := func() bool {
		return r.blockTree.FindBlockByHash(vote.BlockHash) != nil || vote.BlockNumber <= r.viewState.HighestPreCommitQCBlock().NumberU64()
	}

	existsPrepare := func() bool {
		prepare := r.viewState.ViewBlockByIndex(vote.BlockNumber)
		return prepare != nil && prepare.NumberU64() == vote.BlockNumber && prepare.Hash() == vote.BlockHash
	}

	doubtDuplicate := func() bool {
		local := r.viewState.FindPreCommit(vote.BlockNumber,vote.ValidatorIndex)
		if local != nil && local.BlockNumber == vote.BlockNumber && local.BlockHash != vote.BlockHash {
			return true
		}
		return false
	}

	acceptIndexVote := func() SafetyError {
		if vote.BlockIndex >= r.config.Sys.Amount {
			return newCommonError(fmt.Sprintf("voteIndex higher than amount(index:%d, amount:%d)", vote.BlockIndex, r.config.Sys.Amount))
		}
		if doubtDuplicate() {
			return nil
		}
		if r.viewState.FindPreCommit(vote.BlockNumber,vote.ValidatorIndex) != nil {
			return newCommonError(fmt.Sprintf("PreCommit vote has exist(blockIndex:%d, validatorIndex:%d)", vote.BlockIndex, vote.ValidatorIndex))
		}
		if !existsPrepare() {
			return newFetchPrepareError(fmt.Sprintf("current index block not existed,discard msg(number:%d)", vote.BlockNumber))
		}
		return nil
	}

	if alreadyPreCommitQCBlock() {
		return newCommonError(fmt.Sprintf("current index block is already PreCommit qc block,discard msg(index:%d,number:%d,hash:%s)", vote.BlockIndex, vote.BlockNumber, vote.BlockHash.String()))
	}
	if r.viewState.Epoch() != vote.Epoch {
		return r.changeEpochPreCommitRules(vote)
	}

	if r.viewState.BlockNumber() > vote.BlockNum() {
		return newCommonError(fmt.Sprintf("BlockNumber too low(local:%d, msg:%d)", r.viewState.BlockNumber(), vote.BlockNum()))
	}
	if r.viewState.BlockNumber() < vote.BlockNum() {
		if vote.BlockNum()>1{
			b, _ := r.blockTree.FindBlockAndQC(vote.ParentQC.BlockHash, vote.BlockNum()-1)
			if b == nil {
				//if vote.BlockNum()-r.viewState.BlockNumber()==1{
				//	return newCommonError(fmt.Sprintf("BlockNumber higher than local, but not find parent block(local:%d, msg:%d)", r.viewState.BlockNumber(), vote.BlockNum()))
				//}else {
				//	return newFetchError(fmt.Sprintf("BlockNumber higher than local(local:%d, msg:%d)", r.viewState.BlockNumber(), vote.BlockNum()))
				//}
				return newFetchError(fmt.Sprintf("BlockNumber higher than local(local:%d, msg:%d)", r.viewState.BlockNumber(), vote.BlockNum()))
			}
			//return newFetchError(fmt.Sprintf("BlockNumber higher than local(local:%d, msg:%d)", r.viewState.BlockNumber(), vote.BlockNum()))
		}
	}

	if r.viewState.ViewNumber() > vote.ViewNumber {
		return newCommonError(fmt.Sprintf("viewNumber too low(local:%d, msg:%d)", r.viewState.ViewNumber(), vote.ViewNumber))
	}

	if r.viewState.ViewNumber() < vote.ViewNumber {
		return newCommonError(fmt.Sprintf("viewNumber higher than local(local:%d, msg:%d)", r.viewState.ViewNumber(), vote.ViewNumber))
	}

	// if local epoch and viewNumber is the same with msg
	if err := acceptIndexVote(); err != nil {
		return err
	}

	if r.viewState.IsDeadline() {
		return newCommonError(fmt.Sprintf("view's deadline is expire(over:%d)", time.Since(r.viewState.Deadline())))
	}
	return nil
}

func (r *baseSafetyRules) changeEpochVoteRules(vote *protocols.PrepareVote) SafetyError {
	if r.viewState.Epoch() > vote.Epoch {
		return newCommonError(fmt.Sprintf("epoch too low(local:%d, msg:%d)", r.viewState.Epoch(), vote.Epoch))
	}

	return newFetchError("new epoch, need fetch blocks")
}

func (r *baseSafetyRules) changeEpochPreCommitRules(vote *protocols.PreCommit) SafetyError {
	if r.viewState.Epoch() > vote.Epoch {
		return newCommonError(fmt.Sprintf("epoch too low(local:%d, msg:%d)", r.viewState.Epoch(), vote.Epoch))
	}

	return newFetchError("new epoch, need fetch blocks")
}

// ViewChange rules
// 1.Less than local viewNumber drop
// 2.Synchronization greater than local viewNumber
func (r *baseSafetyRules) ViewChangeRules(viewChange *protocols.ViewChange) SafetyError {
	if r.viewState.Epoch() != viewChange.Epoch {
		return r.changeEpochViewChangeRules(viewChange)
	}

	if r.viewState.HighestCommitBlock().NumberU64() > viewChange.BlockNum() {
		return newCommonError(fmt.Sprintf("BlockNumber too low(local:%d, msg:%d)", r.viewState.HighestCommitBlock().NumberU64(), viewChange.BlockNum()))
	}
	if r.viewState.HighestCommitBlock().NumberU64() < viewChange.BlockNum() {
		if viewChange.BlockNum()>0{
			b, _ := r.blockTree.FindBlockAndQC(viewChange.PrepareQC.BlockHash, viewChange.BlockNum())
			if b == nil {
				//if viewChange.BlockNum()-r.viewState.HighestCommitBlock().NumberU64()==1{
				//	return newCommonError(fmt.Sprintf("BlockNumber higher than local, but not find the block(local:%d, msg:%d)", r.viewState.HighestCommitBlock().NumberU64(), viewChange.BlockNum()))
				//}else {
				//	return newFetchError(fmt.Sprintf("BlockNumber higher than local(local:%d, msg:%d)", r.viewState.HighestCommitBlock().NumberU64(), viewChange.BlockNum()))
				//}
				return newFetchError(fmt.Sprintf("BlockNumber higher than local(local:%d, msg:%d)", r.viewState.HighestCommitBlock().NumberU64(), viewChange.BlockNum()))
			}
			//return newFetchError(fmt.Sprintf("BlockNumber higher than local(local:%d, msg:%d)", r.viewState.HighestCommitBlock().NumberU64(), viewChange.BlockNum()))
		}
	}

	if r.viewState.ViewNumber() > viewChange.ViewNumber {
		return newCommonError(fmt.Sprintf("viewNumber too low(local:%d, msg:%d)", r.viewState.ViewNumber(), viewChange.ViewNumber))
	}

	if r.viewState.ViewNumber() < viewChange.ViewNumber {
		return newCommonError(fmt.Sprintf("viewNumber higher than local(local:%d, msg:%d)", r.viewState.ViewNumber(), viewChange.ViewNumber))
	}

	return nil
}

func (r *baseSafetyRules) changeEpochViewChangeRules(viewChange *protocols.ViewChange) SafetyError {
	if r.viewState.Epoch() > viewChange.Epoch {
		return newCommonError(fmt.Sprintf("epoch too low(local:%d, msg:%d)", r.viewState.Epoch(), viewChange.Epoch))
	}

	return newFetchError("new epoch, need fetch blocks")
}

func (r *baseSafetyRules) QCBlockRules(block *types.Block, qc *ctypes.QuorumCert) SafetyError {
	//if r.viewState.Epoch() > qc.Epoch || r.viewState.ViewNumber() > qc.ViewNumber {
	//	return newError(fmt.Sprintf("epoch or viewNumber too low(local:%s, msg:{Epoch:%d,ViewNumber:%d})", r.viewState.ViewString(), qc.Epoch, qc.ViewNumber))
	//}

	if b := r.blockTree.FindBlockByHash(block.ParentHash()); b == nil {
		return newError(fmt.Sprintf("not find parent qc block"))
	}
	if r.viewState.Epoch()+1 == qc.Epoch {
		return newViewError("need change view")
	}
	return nil
}

func NewSafetyRules(viewState *state.ViewState, blockTree *ctypes.BlockTree, config *ctypes.Config, validatorPool *validator.ValidatorPool) SafetyRules {
	return &baseSafetyRules{
		viewState:     viewState,
		blockTree:     blockTree,
		config:        config,
		validatorPool: validatorPool,
	}
}
