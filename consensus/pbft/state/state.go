package state

import (
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/consensus/pbft/protocols"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common/math"

	ctypes "github.com/PhoenixGlobal/Phoenix-Chain-Core/consensus/pbft/types"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core/types"
)

const DefaultEpoch = 1
const DefaultViewNumber = 0

const ViewCacheLen = 10
const AddBlockNumberTimeInterval = 8

type PrepareVoteQueue struct {
	Votes []*protocols.PrepareVote `json:"votes"`
}

func newPrepareVoteQueue() *PrepareVoteQueue {
	return &PrepareVoteQueue{
		Votes: make([]*protocols.PrepareVote, 0),
	}
}

func (p *PrepareVoteQueue) Top() *protocols.PrepareVote {
	return p.Votes[0]
}

func (p *PrepareVoteQueue) Pop() *protocols.PrepareVote {
	v := p.Votes[0]
	p.Votes = p.Votes[1:]
	return v
}

func (p *PrepareVoteQueue) Push(vote *protocols.PrepareVote) {
	p.Votes = append(p.Votes, vote)
}

func (p *PrepareVoteQueue) Peek() []*protocols.PrepareVote {
	return p.Votes
}

func (p *PrepareVoteQueue) Empty() bool {
	return len(p.Votes) == 0
}

func (p *PrepareVoteQueue) Len() int {
	return len(p.Votes)
}

func (p *PrepareVoteQueue) reset() {
	p.Votes = make([]*protocols.PrepareVote, 0)
}

func (p *PrepareVoteQueue) Had(index uint32) bool {
	for _, p := range p.Votes {
		if p.BlockIndex == index {
			return true
		}
	}
	return false
}

type prepareVotes struct {
	Lock sync.Mutex
	Votes map[uint32]*protocols.PrepareVote `json:"votes"`
}

func newPrepareVotes() *prepareVotes {
	return &prepareVotes{
		Votes: make(map[uint32]*protocols.PrepareVote),
	}
}

func (p *prepareVotes) hadVote(vote *protocols.PrepareVote) bool {
	p.Lock.Lock()
	for _, v := range p.Votes {
		if v.MsgHash() == vote.MsgHash() {
			return true
		}
	}
	p.Lock.Unlock()
	return false
}

func (p *prepareVotes) len() int {
	return len(p.Votes)
}

func (p *prepareVotes) clear() {
	p.Votes = make(map[uint32]*protocols.PrepareVote)
}

func (p *prepareVotes) addPrepareVote(id uint32, vote *protocols.PrepareVote) {
	p.Lock.Lock()
	p.Votes[id] = vote
	p.Lock.Unlock()
}

type preCommits struct {
	Lock sync.Mutex
	PreCommits map[uint32]*protocols.PreCommit `json:"preCommits"`
}

func newPreCommits() *preCommits {
	return &preCommits{
		PreCommits: make(map[uint32]*protocols.PreCommit),
	}
}

func (p *preCommits) hadPreCommit(preCommit *protocols.PreCommit) bool {
	p.Lock.Lock()
	for _, v := range p.PreCommits {
		if v.MsgHash() == preCommit.MsgHash() {
			return true
		}
	}
	p.Lock.Unlock()
	return false
}

func (p *preCommits) len() int {
	return len(p.PreCommits)
}

func (p *preCommits) clear() {
	p.PreCommits = make(map[uint32]*protocols.PreCommit)
}

func (p *preCommits) addPreCommit(id uint32, preCommit *protocols.PreCommit) {
	p.Lock.Lock()
	p.PreCommits[id] = preCommit
	p.Lock.Unlock()
}

type viewBlocks struct {
	Lock sync.Mutex
	Blocks map[uint64]viewBlock `json:"blocks"`
	IsProduced bool `json:"isProduced"`
}

func (v *viewBlocks) MarshalJSON() ([]byte, error) {
	type viewBlocks struct {
		Hash   common.Hash `json:"hash"`
		Number uint64      `json:"number"`
		Index  uint32      `json:"blockIndex"`
	}

	vv := make(map[uint64]viewBlocks)
	v.Lock.Lock()
	for index, block := range v.Blocks {
		vv[index] = viewBlocks{
			Hash:   block.hash(),
			Number: block.number(),
			Index:  block.blockIndex(),
		}
	}
	v.Lock.Unlock()
	return json.Marshal(vv)
}

func (vb *viewBlocks) UnmarshalJSON(input []byte) error {
	type viewBlocks struct {
		Hash   common.Hash `json:"hash"`
		Number uint64      `json:"number"`
		Index  uint32      `json:"blockIndex"`
	}

	var vv map[uint64]viewBlocks
	err := json.Unmarshal(input, &vv)
	if err != nil {
		return err
	}

	vb.Blocks = make(map[uint64]viewBlock)
	vb.Lock.Lock()
	for k, v := range vv {
		vb.Blocks[k] = prepareViewBlock{
			pb: &protocols.PrepareBlock{
				BlockIndex: v.Index,
				Block:      types.NewSimplifiedBlock(v.Number, v.Hash),
			},
		}
	}
	vb.Lock.Unlock()
	return nil
}

func newViewBlocks() *viewBlocks {
	return &viewBlocks{
		Blocks: make(map[uint64]viewBlock),
		IsProduced:false,
	}
}

func (v *viewBlocks) index(i uint64) viewBlock {
	return v.Blocks[i]
}

func (v *viewBlocks) addBlock(block viewBlock) {
	v.Lock.Lock()
	defer v.Lock.Unlock()
	_,ok:=v.Blocks[block.number()]
	if !ok {
		if len(v.Blocks)>=ViewCacheLen{
			minNumber:=v.MinNumber()
			delete(v.Blocks,minNumber)
		}
	}
	v.Blocks[block.number()] = block
	v.IsProduced=true
}

func (v *viewBlocks) clear(blockNumber uint64) {
	//v.Lock.Lock()
	//defer v.Lock.Unlock()
	//delete(v.Blocks,blockNumber)
	v.IsProduced=false
}

func (v *viewBlocks) exist(blockNumber uint64) bool {
	_,ok:=v.Blocks[blockNumber]
	if ok{
		return true
	}
	return false
}

func (v *viewBlocks) len() int {
	return len(v.Blocks)
}

//func (v *viewBlocks) MaxIndex() uint32 {
//	max := uint32(math.MaxUint32)
//	for _, b := range v.Blocks {
//		if max == math.MaxUint32 || b.blockIndex() > max {
//			max = b.blockIndex()
//		}
//	}
//	return max
//}

func (v *viewBlocks) MaxNumber() uint64 {
	max := uint64(0)
	for _, b := range v.Blocks {
		if b.number() > max {
			max = b.number()
		}
	}
	return max
}

func (v *viewBlocks) MinNumber() uint64 {
	min := uint64(0)
	for _, b := range v.Blocks {
		if b.number() < min {
			min = b.number()
		}
	}
	return min
}

type viewQCs struct {
	Lock      sync.Mutex
	MaxNumber uint64                       `json:"maxNumber"`
	MinNumber uint64                       `json:"minNumber"`
	QCs      map[uint64]*ctypes.QuorumCert `json:"qcs"`
	IsProduced bool `json:"isProduced"`
}

func newViewQCs() *viewQCs {
	return &viewQCs{
		MinNumber: uint64(0),
		MaxNumber: uint64(0),
		QCs:      make(map[uint64]*ctypes.QuorumCert),
		IsProduced: false,
	}
}

func (v *viewQCs) index(i uint64) *ctypes.QuorumCert {
	return v.QCs[i]
}

func (v *viewQCs) addQC(qc *ctypes.QuorumCert) {
	v.Lock.Lock()
	defer v.Lock.Unlock()
	q,ok:=v.QCs[qc.BlockNumber]
	if ok && q.BlockHash==qc.BlockHash{
		return
	}

	if len(v.QCs)>=ViewCacheLen{
		minNumber:=v.minNumber()
		delete(v.QCs,minNumber)
	}
	v.QCs[qc.BlockNumber] = qc
	v.IsProduced=true
	if v.MaxNumber < qc.BlockNumber {
		v.MaxNumber = qc.BlockNumber
	}
	if v.MinNumber > qc.BlockNumber {
		v.MinNumber = qc.BlockNumber
	}
	return
}

//func (v *viewQCs) maxQCIndex() uint32 {
//
//}

func (v *viewQCs) maxQCNumber() uint64 {
	return v.MaxNumber
}

func (v *viewQCs) minNumber() uint64 {
	return v.MinNumber
}

func (v *viewQCs) clear(blockNumber uint64) {
	//if len(v.QCs)>=ViewCacheLen{
	//	v.QCs = make(map[uint64]*ctypes.QuorumCert)
	//	v.MaxNumber = math.MaxUint64
	//}
	v.Lock.Lock()
	delete(v.QCs,blockNumber)
	v.Lock.Unlock()
	v.IsProduced=false
}

func (v *viewQCs) len() int {
	return len(v.QCs)
}

type viewVotes struct {
	Lock      sync.Mutex
	Votes map[uint64]*prepareVotes `json:"votes"`
}

func newViewVotes() *viewVotes {
	return &viewVotes{
		Votes: make(map[uint64]*prepareVotes),
	}
}

func (v *viewVotes) addVote(id uint32, vote *protocols.PrepareVote) {
	v.Lock.Lock()
	if len(v.Votes)>=ViewCacheLen{
		minNumber:=v.MinNumber()
		delete(v.Votes,minNumber)
	}
	if ps, ok := v.Votes[vote.BlockNumber]; ok {
		//ps.Votes[id] = vote
		ps.addPrepareVote(id,vote)
		v.Votes[vote.BlockNumber] = ps
	} else {
		ps2 := newPrepareVotes()
		//ps.Votes[id] = vote
		ps2.addPrepareVote(id,vote)
		v.Votes[vote.BlockNumber] = ps2
	}
	v.Lock.Unlock()
}

func (v *viewVotes) index(i uint64) *prepareVotes {
	return v.Votes[i]
}

func (v *viewVotes) MaxNumber() uint64 {
	max := uint64(0)
	for index, _ := range v.Votes {
		if index > max {
			max = index
		}
	}
	return max
}

func (v *viewVotes) MinNumber() uint64 {
	min := uint64(0)
	for index, _ := range v.Votes {
		if index < min {
			min = index
		}
	}
	return min
}

func (v *viewVotes) clear(blockNumber uint64) {
	//if len(v.Votes)>=ViewCacheLen{
	//	v.Votes = make(map[uint64]*prepareVotes)
	//}
	v.Lock.Lock()
	delete(v.Votes,blockNumber)
	v.Lock.Unlock()
}

type viewPreCommits struct {
	Lock      sync.Mutex
	Votes map[uint64]*preCommits `json:"votes"`
}

func newViewPreCommits() *viewPreCommits {
	return &viewPreCommits{
		Votes: make(map[uint64]*preCommits),
	}
}

func (v *viewPreCommits) addPreCommit(id uint32, vote *protocols.PreCommit) {
	v.Lock.Lock()
	if len(v.Votes)>=ViewCacheLen{
		minNumber:=v.MinNumber()
		delete(v.Votes,minNumber)
	}
	if ps, ok := v.Votes[vote.BlockNumber]; ok {
		//ps.PreCommits[id] = vote
		ps.addPreCommit(id,vote)
		v.Votes[vote.BlockNumber] = ps
	} else {
		ps2 := newPreCommits()
		//ps.PreCommits[id] = vote
		ps2.addPreCommit(id,vote)
		v.Votes[vote.BlockNumber] = ps2
	}
	v.Lock.Unlock()
}

func (v *viewPreCommits) index(i uint64) *preCommits {
	return v.Votes[i]
}

func (v *viewPreCommits) MaxNumber() uint64 {
	max := uint64(0)
	for index, _ := range v.Votes {
		if index > max {
			max = index
		}
	}
	return max
}

func (v *viewPreCommits) MinNumber() uint64 {
	min := uint64(0)
	for index, _ := range v.Votes {
		if index < min {
			min = index
		}
	}
	return min
}

func (v *viewPreCommits) clear(blockNumber uint64) {
	//if len(v.Votes)>=ViewCacheLen{
	//	v.Votes = make(map[uint64]*preCommits)
	//}
	v.Lock.Lock()
	delete(v.Votes,blockNumber)
	v.Lock.Unlock()
}

type viewChanges struct {
	ViewChanges map[uint32]*protocols.ViewChange `json:"viewchanges"`
}

func newViewChanges() *viewChanges {
	return &viewChanges{
		ViewChanges: make(map[uint32]*protocols.ViewChange),
	}
}

func (v *viewChanges) addViewChange(id uint32, viewChange *protocols.ViewChange) {
	v.ViewChanges[id] = viewChange
}

func (v *viewChanges) len() int {
	return len(v.ViewChanges)
}

func (v *viewChanges) clear() {
	v.ViewChanges = make(map[uint32]*protocols.ViewChange)
}

type executing struct {
	// Block index of current view
	BlockIndex uint32 `json:"blockIndex"`
	// Whether to complete
	Finish bool `json:"finish"`
}

type view struct {
	epoch      uint64
	viewNumber uint64
	blockNumber uint64
	lastAddBlockNumberTime int64
	maxExecutedBlockNumber uint64
	Step       uint32

	// The status of the block is currently being executed,
	// Finish indicates whether the execution is complete,
	// and the next block can be executed asynchronously after the execution is completed.
	executing executing

	//finish     bool

	// viewchange received by the current view
	viewChanges *viewChanges

	// QC of the previous view
	lastViewChangeQC *ctypes.ViewChangeQC

	// This view has been sent to other verifiers for voting
	hadSendPrepareVote *PrepareVoteQueue

	//Pending Votes of current view, parent block need receive N-f prepareVotes
	pendingVote *PrepareVoteQueue

	//PrepareVote *protocols.PrepareVote
	hadSendVoteNumber uint64
	hadSendVoteLast *protocols.PrepareVote


	//PreCommit   *protocols.PreCommit
	hadSendPreCommitNumber uint64
	hadSendPreCommitLast   *protocols.PreCommit


	//Current view of the proposed block by the proposer
	viewBlocks *viewBlocks

	//preVote qcs
	viewQCs *viewQCs

	//preCommit qcs
	viewPreCommitQCs *viewQCs

	viewPreCommitQC *ctypes.QuorumCert
	//
	viewPrepareVoteQC *ctypes.QuorumCert

	//The current view generated by the vote
	//viewVotes *prepareVotes

	viewVotes *viewVotes

	//The current view generated by the vote
	//viewPreCommits *preCommits
	viewPreCommits *viewPreCommits
}

func newView() *view {
	return &view{
		executing:          executing{math.MaxUint32, false},
		blockNumber:        1,
		lastAddBlockNumberTime: time.Now().Unix(),
		maxExecutedBlockNumber: 0,
		viewChanges:        newViewChanges(),
		hadSendPrepareVote: newPrepareVoteQueue(),
		hadSendVoteNumber: 0,
		hadSendPreCommitNumber: 0,
		hadSendVoteLast :    nil,
		hadSendPreCommitLast:nil,
		pendingVote:        newPrepareVoteQueue(),
		viewBlocks:         newViewBlocks(),
		viewQCs:            newViewQCs(),
		viewPreCommitQCs:   newViewQCs(),
		//viewVotes:          newPrepareVotes(),
		viewVotes:          newViewVotes(),
		//viewPreCommits:     newPreCommits(),
		viewPreCommits:     newViewPreCommits(),
	}
}

func (v *view) Reset(blockNumber uint64) {
	atomic.StoreUint64(&v.epoch, 0)
	atomic.StoreUint64(&v.viewNumber, 0)
	atomic.StoreUint64(&v.blockNumber, 1)
	atomic.StoreUint64(&v.maxExecutedBlockNumber, 0)
	v.executing.BlockIndex = math.MaxUint32
	v.executing.Finish = false
	//v.finish=false
	v.viewChanges.clear()
	v.hadSendPrepareVote.reset()
	//atomic.StoreUint64(&v.hadSendVoteNumber, 0)
	//v.HadSendVote=nil
	//v.PrepareVote=nil
	//atomic.StoreUint64(&v.hadSendPreCommitNumber, 0)
	//v.HadSendPreCommit=nil
	//v.PreCommit=nil
	v.viewPrepareVoteQC=nil
	v.viewPreCommitQC=nil
	v.pendingVote.reset()
	v.viewBlocks.clear(blockNumber)
	v.viewQCs.clear(blockNumber)
	v.viewPreCommitQCs.clear(blockNumber)
	v.viewVotes.clear(blockNumber)
	v.viewPreCommits.clear(blockNumber)
}

func (vs *ViewState) Step() ctypes.RoundStepType{
	step:=atomic.LoadUint32(&vs.view.Step)
	return ctypes.RoundStepType(step)
}

func (v *view) ViewNumber() uint64 {
	return atomic.LoadUint64(&v.viewNumber)
}

func (v *view) BlockNumber() uint64 {
	return atomic.LoadUint64(&v.blockNumber)
}

func (v *view) LastAddBlockNumberTime() int64 {
	return atomic.LoadInt64(&v.lastAddBlockNumberTime)
}

func (v *view) MaxExecutedBlockNumber() uint64 {
	return atomic.LoadUint64(&v.maxExecutedBlockNumber)
}

func (v *view) AddBlockNumber(qc *ctypes.QuorumCert) {
	//blockNumber:=v.BlockNumber()
	atomic.StoreUint64(&v.blockNumber, qc.BlockNumber+1)
	atomic.StoreInt64(&v.lastAddBlockNumberTime, time.Now().Unix())
}

func (v *view) SetBlockNumber(blockNumber uint64) {
	//blockNumber:=v.BlockNumber()
	atomic.StoreUint64(&v.blockNumber, blockNumber)
	atomic.StoreInt64(&v.lastAddBlockNumberTime, time.Now().Unix())
}

func (v *view) HadSendVoteNumber() uint64 {
	return atomic.LoadUint64(&v.hadSendVoteNumber)
}

func (v *view) HadSendVoteLast() *protocols.PrepareVote {
	return v.hadSendVoteLast
}

func (v *view) HadSendPreCommitNumber() uint64 {
	return atomic.LoadUint64(&v.hadSendPreCommitNumber)
}

func (v *view) HadSendPreCommitLast() *protocols.PreCommit {
	return v.hadSendPreCommitLast
}

func (v *view)AddHadSendVoteNumber(number uint64,vote *protocols.PrepareVote)  {
	v.hadSendVoteLast=vote
	atomic.StoreUint64(&v.hadSendVoteNumber, number)
}

func (v *view)AddHadSendPreCommitNumber(number uint64,vote *protocols.PreCommit)  {
	v.hadSendPreCommitLast=vote
	atomic.StoreUint64(&v.hadSendPreCommitNumber, number)
}

func (v *view) SetMaxExecutedBlockNumber(blockNumber uint64) {
	if blockNumber>v.MaxExecutedBlockNumber(){
		atomic.StoreUint64(&v.maxExecutedBlockNumber, blockNumber)
	}
}

func (v *view) Epoch() uint64 {
	return atomic.LoadUint64(&v.epoch)
}

func (v *view) MarshalJSON() ([]byte, error) {
	type view struct {
		Epoch              uint64               `json:"epoch"`
		ViewNumber         uint64               `json:"viewNumber"`
		BlockNumber        uint64               `json:"blockNumber"`
		LastAddBlockNumberTime int64            `json:"lastAddBlockNumberTime"`
		Step               uint32               `json:"step"`
		Executing          executing            `json:"executing"`
		ViewChanges        *viewChanges         `json:"viewchange"`
		LastViewChangeQC   *ctypes.ViewChangeQC `json:"lastViewchange"`
		HadSendPrepareVote *PrepareVoteQueue    `json:"hadSendPrepareVote"`
		HadSendVoteNumber   uint64               `json:"hadSendVoteNumber"`
		HadSendPreCommitNumber  uint64           `json:"hadSendPreCommitNumber"`
		hadSendVoteLast   *protocols.PrepareVote  `json:"hadSendVoteLast"`
		hadSendPreCommitLast   *protocols.PreCommit `json:"hadSendPreCommitLast"`

		PendingVote        *PrepareVoteQueue    `json:"pendingPrepareVote"`
		//PrepareVote        *protocols.PrepareVote `json:"prepareVote"`
		//HadSendVote        *protocols.PrepareVote `json:"hadSendVote"`
		//PreCommit          *protocols.PreCommit   `json:"preCommit"`
		//HadSendPreCommit   *protocols.PreCommit   `json:"hadSendPreCommit"`
		ViewBlocks         *viewBlocks          `json:"viewBlocks"`
		ViewQCs            *viewQCs             `json:"viewQcs"`
		ViewPreCommitQC    *ctypes.QuorumCert   `json:"viewPreCommitQC"`
		ViewPrepareVoteQC  *ctypes.QuorumCert   `json:"viewPrepareVoteQC"`
		//ViewVotes          *prepareVotes           `json:"viewVotes"`
		//ViewPreCommits      *preCommits         `json:"viewPreCommits"`
		ViewVotes          *viewVotes           `json:"viewVotes"`
		ViewPreCommits     *viewPreCommits      `json:"viewPreCommits"`
	}
	vv := &view{
		Epoch:              atomic.LoadUint64(&v.epoch),
		ViewNumber:         atomic.LoadUint64(&v.viewNumber),
		BlockNumber:        atomic.LoadUint64(&v.blockNumber),
		LastAddBlockNumberTime: atomic.LoadInt64(&v.lastAddBlockNumberTime),
		Step:               atomic.LoadUint32(&v.Step),
		Executing:          v.executing,
		ViewChanges:        v.viewChanges,
		LastViewChangeQC:   v.lastViewChangeQC,
		HadSendPrepareVote: v.hadSendPrepareVote,
		PendingVote:        v.pendingVote,
		//PrepareVote:        v.PrepareVote,
		HadSendVoteNumber:   v.hadSendVoteNumber,
		hadSendVoteLast:     v.hadSendVoteLast,
		//PreCommit:          v.PreCommit,
		HadSendPreCommitNumber:   v.hadSendPreCommitNumber,
		hadSendPreCommitLast:   v.hadSendPreCommitLast,
		ViewBlocks:         v.viewBlocks,
		ViewQCs:            v.viewQCs,
		ViewPreCommitQC:    v.viewPreCommitQC,
		ViewPrepareVoteQC:  v.viewPrepareVoteQC,
		ViewVotes:          v.viewVotes,
		ViewPreCommits:     v.viewPreCommits,
	}

	return json.Marshal(vv)
}

func (v *view) UnmarshalJSON(input []byte) error {
	type view struct {
		Epoch              uint64               `json:"epoch"`
		ViewNumber         uint64               `json:"viewNumber"`
		BlockNumber        uint64               `json:"blockNumber"`
		LastAddBlockNumberTime int64			`json:"lastAddBlockNumberTime"`
		Step               uint32               `json:"step"`
		Executing          executing            `json:"executing"`
		ViewChanges        *viewChanges         `json:"viewchange"`
		LastViewChangeQC   *ctypes.ViewChangeQC `json:"lastViewchange"`
		HadSendPrepareVote *PrepareVoteQueue    `json:"hadSendPrepareVote"`
		PendingVote        *PrepareVoteQueue    `json:"pendingPrepareVote"`
		HadSendVoteNumber   uint64               `json:"hadSendVoteNumber"`
		HadSendVoteLast   *protocols.PrepareVote  `json:"hadSendVoteLast"`
		HadSendPreCommitLast   *protocols.PreCommit `json:"hadSendPreCommitLast"`
		HadSendPreCommitNumber  uint64           `json:"hadSendPreCommitNumber"`
		//PrepareVote        *protocols.PrepareVote `json:"prepareVote"`
		//HadSendVote        *protocols.PrepareVote `json:"hadSendVote"`
		//PreCommit          *protocols.PreCommit   `json:"preCommit"`
		//HadSendPreCommit   *protocols.PreCommit   `json:"hadSendPreCommit"`
		ViewBlocks         *viewBlocks          `json:"viewBlocks"`
		ViewQCs            *viewQCs             `json:"viewQcs"`
		ViewPreCommitQC    *ctypes.QuorumCert   `json:"viewPreCommitQC"`
		ViewPrepareVoteQC  *ctypes.QuorumCert   `json:"viewPrepareVoteQC"`
		//ViewVotes          *prepareVotes           `json:"viewVotes"`
		//ViewPreCommits      *preCommits         `json:"viewPreCommits"`
		ViewVotes          *viewVotes           `json:"viewVotes"`
		ViewPreCommits     *viewPreCommits      `json:"viewPreCommits"`
	}

	var vv view
	err := json.Unmarshal(input, &vv)
	if err != nil {
		return err
	}

	v.epoch = vv.Epoch
	v.viewNumber = vv.ViewNumber
	v.blockNumber = vv.BlockNumber
	v.lastAddBlockNumberTime=vv.LastAddBlockNumberTime
	v.Step = vv.Step
	v.executing = vv.Executing
	v.viewChanges = vv.ViewChanges
	v.lastViewChangeQC = vv.LastViewChangeQC
	v.hadSendPrepareVote = vv.HadSendPrepareVote
	v.pendingVote = vv.PendingVote
	//v.PrepareVote = vv.PrepareVote
	v.hadSendVoteNumber = vv.HadSendVoteNumber
	v.hadSendVoteLast = vv.HadSendVoteLast
	//v.PreCommit = vv.PreCommit
	v.hadSendPreCommitNumber = vv.HadSendPreCommitNumber
	v.hadSendPreCommitLast = vv.HadSendPreCommitLast
	v.viewBlocks = vv.ViewBlocks
	v.viewQCs = vv.ViewQCs
	v.viewPrepareVoteQC = vv.ViewPrepareVoteQC
	v.viewPreCommitQC = vv.ViewPreCommitQC
	v.viewVotes = vv.ViewVotes
	v.viewPreCommits = vv.ViewPreCommits
	return nil
}

//func (v *view) HadSendPrepareVote(vote *protocols.PrepareVote) bool {
//	return v.hadSendPrepareVote.hadVote(vote)
//}

//The block of current view, there two types, prepareBlock and block
type viewBlock interface {
	hash() common.Hash
	number() uint64
	blockIndex() uint32
	block() *types.Block
	//If prepareBlock is an implementation of viewBlock, return prepareBlock, otherwise nil
	prepareBlock() *protocols.PrepareBlock
}

type prepareViewBlock struct {
	pb *protocols.PrepareBlock
}

func (p prepareViewBlock) hash() common.Hash {
	return p.pb.Block.Hash()
}

func (p prepareViewBlock) number() uint64 {
	return p.pb.Block.NumberU64()
}

func (p prepareViewBlock) blockIndex() uint32 {
	return p.pb.BlockIndex
}

func (p prepareViewBlock) block() *types.Block {
	return p.pb.Block
}

func (p prepareViewBlock) prepareBlock() *protocols.PrepareBlock {
	return p.pb
}

type qcBlock struct {
	b  *types.Block
	qc *ctypes.QuorumCert   //PreCommitQC
}

func (q qcBlock) hash() common.Hash {
	return q.b.Hash()
}

func (q qcBlock) number() uint64 {
	return q.b.NumberU64()
}
func (q qcBlock) blockIndex() uint32 {
	if q.qc == nil {
		return 0
	}
	return q.qc.BlockIndex
}

func (q qcBlock) block() *types.Block {
	return q.b
}

func (q qcBlock) prepareBlock() *protocols.PrepareBlock {
	return nil
}

type ViewState struct {

	//Include ViewNumber, ViewChanges, prepareVote , proposal block of current view
	*view

	//*State

	highestQCBlock     atomic.Value   //PreVoteQC
	highestLockBlock   atomic.Value
	highestCommitBlock atomic.Value
	highestPreCommitQCBlock atomic.Value  //PreCommitQC

	//Set the timer of the view time window
	viewTimer *viewTimer

	blockTree *ctypes.BlockTree
}

func NewViewState(period uint64, blockTree *ctypes.BlockTree) *ViewState {
	return &ViewState{
		view:      newView(),
		viewTimer: newViewTimer(period),
		blockTree: blockTree,
	}
}

func (vs *ViewState) UpdateStep(step ctypes.RoundStepType) {
	stepInt:=uint32(step)
	atomic.StoreUint32(&vs.view.Step, stepInt)
}

func (vs *ViewState) ResetView(epoch uint64, viewNumber uint64, blockNumber uint64) {
	vs.view.Reset(blockNumber)
	atomic.StoreUint64(&vs.view.epoch, epoch)
	atomic.StoreUint64(&vs.view.viewNumber, viewNumber)
	atomic.StoreUint64(&vs.view.blockNumber, blockNumber)
	atomic.StoreUint32(&vs.view.Step, 0)
}

func (vs *ViewState) Epoch() uint64 {
	return vs.view.Epoch()
}

func (vs *ViewState) ViewNumber() uint64 {
	return vs.view.ViewNumber()
}

func (vs *ViewState) BlockNumber() uint64 {
	return vs.view.BlockNumber()
}

func (vs *ViewState) ViewString() string {
	return fmt.Sprintf("{Epoch:%d,ViewNumber:%d,BlockNumber:%d}", atomic.LoadUint64(&vs.view.epoch), atomic.LoadUint64(&vs.view.viewNumber),atomic.LoadUint64(&vs.view.blockNumber))
}

func (vs *ViewState) Deadline() time.Time {
	return vs.viewTimer.deadline
}

func (vs *ViewState) NextViewBlockNumber() uint64 {
	return vs.viewBlocks.MaxNumber() + 1
}

func (vs *ViewState) ExistBlock(number uint64) bool {
	return vs.viewBlocks.exist(number)
}

func (vs *ViewState) NextViewBlockIndex() uint32 {
	if vs.viewBlocks.IsProduced{
		return 1
	}
	return 0
}

func (vs *ViewState) MaxViewBlockNumber() uint64 {
	max := vs.viewBlocks.MaxNumber()
	//if max == math.MaxUint64 {
	//	return 0
	//}
	return max
}

func (vs *ViewState) MaxViewVotesNumber() uint64 {
	max := vs.viewVotes.MaxNumber()
	//if max == math.MaxUint64 {
	//	return 0
	//}
	return max
}

func (vs *ViewState) MaxViewPreCommitsNumber() uint64 {
	max := vs.viewPreCommits.MaxNumber()
	//if max == math.MaxUint64 {
	//	return 0
	//}
	return max
}

func (vs *ViewState) MinViewBlockNumber() uint64 {
	min := vs.viewBlocks.MinNumber()
	//if min == math.MaxUint64 {
	//	return 0
	//}
	return min
}

func (vs *ViewState) MaxQCNumber() uint64 {
	return vs.view.viewQCs.maxQCNumber()
}

func (vs *ViewState) MaxPreCommitQCNumber() uint64 {
	return vs.view.viewPreCommitQCs.maxQCNumber()
}

func (vs *ViewState) MaxQCIndex() uint32 {
	if vs.viewPreCommitQCs.IsProduced{
		return 0
	}
	return math.MaxUint32
}

func (vs *ViewState) SetPreCommitQC(viewPreCommitQC *ctypes.QuorumCert){
	vs.view.viewPreCommitQC=viewPreCommitQC
}

func (vs *ViewState) SetPrepareVoteQC(viewPrepareVoteQC *ctypes.QuorumCert){
	vs.view.viewPrepareVoteQC=viewPrepareVoteQC
}

//func (vs *ViewState) ViewPreCommitQC()*ctypes.QuorumCert{
//	return vs.view.viewPreCommitQC
//}

func (vs *ViewState) ViewPrepareVoteQC()*ctypes.QuorumCert{
	return vs.view.viewPrepareVoteQC
}

func (vs *ViewState) HadPrepareVoteQC(hash common.Hash, number uint64, viewNumber uint64)bool{
	qc := vs.viewQCs.index(number)
	if qc!=nil&&qc.BlockHash==hash&&qc.BlockNumber==number{
		return true
	}
	return false
}

func (vs *ViewState) HadPreCommitQC(hash common.Hash, number uint64, viewNumber uint64)bool{
	qc := vs.viewPreCommitQCs.index(number)
	if qc!=nil&&qc.BlockHash==hash&&qc.BlockNumber==number{
		return true
	}
	return false
}

func (vs *ViewState) ViewVoteSize() int {
	return len(vs.viewVotes.Votes)
}

func (vs *ViewState) MinViewVoteIndex() uint64 {
	min := vs.viewVotes.MinNumber()
	if min == math.MaxUint64 {
		return 0
	}
	return min
}

func (vs *ViewState) MaxViewVoteIndex() uint64 {
	max := vs.viewVotes.MaxNumber()
	if max == math.MaxUint64 {
		return 0
	}
	return max
}

func (vs *ViewState) PrepareVoteLenByNumber(blockNumber uint64) int {
	ps := vs.viewVotes.index(blockNumber)
	if ps != nil {
		return ps.len()
	}
	return 0
}

func (vs *ViewState) PreCommitLenByNumber(blockNumber uint64) int {
	ps := vs.viewPreCommits.index(blockNumber)
	if ps != nil {
		return ps.len()
	}
	return 0
}

// Find the block corresponding to the current view according to the index
func (vs *ViewState) ViewBlockByIndex(number uint64) *types.Block {
	if b := vs.view.viewBlocks.index(number); b != nil {
		return b.block()
	}
	return nil
}

func (vs *ViewState) PrepareBlockByIndex(number uint64) *protocols.PrepareBlock {
	if b := vs.view.viewBlocks.index(number); b != nil {
		return b.prepareBlock()
	}
	return nil
}

func (vs *ViewState) ViewBlockSize() int {
	return len(vs.viewBlocks.Blocks)
}

func (vs *ViewState) HadSendPrepareVote() *PrepareVoteQueue {
	return vs.view.hadSendPrepareVote
}

func (vs *ViewState) PendingPrepareVote() *PrepareVoteQueue {
	return vs.view.pendingVote
}

func (vs *ViewState) AllPrepareVoteByNumber(blockNumber uint64) map[uint32]*protocols.PrepareVote {
	ps := vs.viewVotes.index(blockNumber)
	if ps != nil {
		return ps.Votes
	}
	return nil
}

func (vs *ViewState) HadSendPreVoteByQC(qc *ctypes.QuorumCert) bool {
	if vs.hadSendVoteNumber<qc.BlockNumber{
		return false
	}
	if vs.hadSendVoteNumber>qc.BlockNumber{
		return true
	}
	hadSend:=vs.hadSendVoteLast
	if hadSend==nil{
		return false
	}
	if hadSend.Epoch>qc.Epoch||hadSend.Epoch==qc.Epoch&&hadSend.ViewNumber>=qc.ViewNumber{
		return true
	}
	return false
}

func (vs *ViewState) HadSendPreVote(vote *protocols.PrepareVote) bool {
	if vs.hadSendVoteNumber<vote.BlockNumber{
		return false
	}
	if vs.hadSendVoteNumber>vote.BlockNumber{
		return true
	}
	hadSend:=vs.hadSendVoteLast
	if hadSend==nil{
		return false
	}
	if hadSend.Epoch>vote.Epoch||hadSend.Epoch==vote.Epoch&&hadSend.ViewNumber>=vote.ViewNumber{
		return true
	}
	return false
}

func (vs *ViewState) HadSendPreCommit(vote *protocols.PreCommit) bool {
	if vs.hadSendPreCommitNumber<vote.BlockNumber{
		return false
	}
	if vs.hadSendPreCommitNumber>vote.BlockNumber{
		return true
	}
	hadSend:=vs.hadSendPreCommitLast
	if hadSend==nil{
		return false
	}
	if hadSend.Epoch>vote.Epoch||hadSend.Epoch==vote.Epoch&&hadSend.ViewNumber>=vote.ViewNumber{
		return true
	}
	return false
}

func (vs *ViewState) AllPreCommitByNumber(blockNumber uint64) map[uint32]*protocols.PreCommit {
	ps := vs.viewPreCommits.index(blockNumber)
	if ps != nil {
		return ps.PreCommits
	}
	return nil
}

func (vs *ViewState) FindPrepareVote(blockNumber uint64, validatorIndex uint32) *protocols.PrepareVote {
	ps := vs.viewVotes.index(blockNumber)
	if ps != nil {
		if v, ok := ps.Votes[validatorIndex]; ok {
			return v
		}
	}
	return nil
}

func (vs *ViewState) FindPreCommit(blockNumber uint64,validatorIndex uint32) *protocols.PreCommit {
	ps := vs.viewPreCommits.index(blockNumber)
	if ps != nil {
		if v, ok := ps.PreCommits[validatorIndex]; ok {
			return v
		}
	}
	return nil
}

func (vs *ViewState) AllViewChange() map[uint32]*protocols.ViewChange {
	return vs.viewChanges.ViewChanges
}

// Returns the block index being executed, has it been completed
func (vs *ViewState) Executing() (uint32, bool) {
	return vs.view.executing.BlockIndex, vs.view.executing.Finish
}

//// Returns Executing block status
//func (vs *ViewState) ExecutingStatus() bool {
//	return vs.view.finish
//}

func (vs *ViewState) SetLastViewChangeQC(qc *ctypes.ViewChangeQC) {
	vs.view.lastViewChangeQC = qc
}

func (vs *ViewState) LastViewChangeQC() *ctypes.ViewChangeQC {
	return vs.view.lastViewChangeQC
}

// Set Executing block status
func (vs *ViewState) SetExecuting(index uint32, finish bool) {
	vs.view.executing.BlockIndex, vs.view.executing.Finish = index, finish
}

//// Set Executing block status
//func (vs *ViewState) SetFinish(finish bool) {
//	vs.view.finish = finish
//}

func (vs *ViewState) ViewBlockAndQC(blockNumber uint64) (*types.Block, *ctypes.QuorumCert) {
	//qc := vs.viewPrepareVoteQC
	qc := vs.viewQCs.index(blockNumber)
	if b := vs.view.viewBlocks.index(blockNumber); b != nil {
		return b.block(), qc
	}
	return nil, qc
}

func (vs *ViewState) ViewBlockAndPreCommitQC(blockNumber uint64) (*types.Block, *ctypes.QuorumCert) {
	//qc := vs.viewPreCommitQC
	qc := vs.viewPreCommitQCs.index(blockNumber)
	if b := vs.view.viewBlocks.index(blockNumber); b != nil {
		return b.block(), qc
	}
	return nil, qc
}

func (vs *ViewState) AddPrepareBlock(pb *protocols.PrepareBlock) {
	vs.view.viewBlocks.addBlock(&prepareViewBlock{pb})
}

func (vs *ViewState) AddQCBlock(block *types.Block, qc *ctypes.QuorumCert) {
	vs.view.viewBlocks.addBlock(&qcBlock{b: block, qc: qc})
}

func (vs *ViewState) AddQC(qc *ctypes.QuorumCert) {
	if qc.Step==ctypes.RoundStepPrepareVote{
		vs.view.viewQCs.addQC(qc)
	}else {
		qc2:=&ctypes.QuorumCert{
			Epoch:        qc.Epoch,
			ViewNumber:   qc.ViewNumber,
			BlockHash:    qc.BlockHash,
			BlockNumber:  qc.BlockNumber,
			BlockIndex:   qc.BlockIndex,
			Step:         ctypes.RoundStepPrepareVote,
			Signature:    qc.Signature,
			ValidatorSet: qc.ValidatorSet,
		}
		vs.view.viewQCs.addQC(qc2)
	}
}

func (vs *ViewState) AddPreCommitQC(qc *ctypes.QuorumCert) {
	vs.view.viewPreCommitQCs.addQC(qc)
}

func (vs *ViewState) AddPrepareVote(id uint32, vote *protocols.PrepareVote) {
	vs.view.viewVotes.addVote(id, vote)
	//vs.view.viewVotes.addPrepareVote(id, vote)qc
	//oldPrepareVote:=vs.view.PrepareVote
	//if oldPrepareVote==nil{
	//	vs.view.PrepareVote=vote
	//	return
	//}
	//if oldPrepareVote.Epoch<vote.Epoch{
	//	vs.view.PrepareVote=vote
	//	return
	//}
	//if oldPrepareVote.BlockNumber<vote.BlockNumber||(oldPrepareVote.BlockNumber==vote.BlockNumber&&oldPrepareVote.ViewNumber<vote.ViewNumber){
	//	vs.view.PrepareVote=vote
	//}
}
func (vs *ViewState) AddPreCommit(id uint32, vote *protocols.PreCommit) {
	vs.view.viewPreCommits.addPreCommit(id, vote)
	//oldPreCommit:=vs.view.PreCommit
	//if oldPreCommit==nil{
	//	vs.view.PreCommit=vote
	//	return
	//}
	//if oldPreCommit.Epoch<vote.Epoch{
	//	vs.view.PreCommit=vote
	//	return
	//}
	//if oldPreCommit.BlockNumber<vote.BlockNumber||(oldPreCommit.BlockNumber==vote.BlockNumber&&oldPreCommit.ViewNumber<vote.ViewNumber){
	//	vs.view.PreCommit=vote
	//}
}

func (vs *ViewState) AddViewChange(id uint32, viewChange *protocols.ViewChange) {
	vs.view.viewChanges.addViewChange(id, viewChange)
}

func (vs *ViewState) ViewChangeByIndex(index uint32) *protocols.ViewChange {
	return vs.view.viewChanges.ViewChanges[index]
}

func (vs *ViewState) ViewChangeLen() int {
	return vs.view.viewChanges.len()
}

func (vs *ViewState) HighestBlockString() string {
	qc := vs.HighestPreCommitQCBlock()
	lock := vs.HighestLockBlock()
	commit := vs.HighestCommitBlock()
	return fmt.Sprintf("{HighestQC:{hash:%s,number:%d},HighestLock:{hash:%s,number:%d},HighestCommit:{hash:%s,number:%d}}",
		qc.Hash().TerminalString(), qc.NumberU64(),
		lock.Hash().TerminalString(), lock.NumberU64(),
		commit.Hash().TerminalString(), commit.NumberU64())
}

func (vs *ViewState) HighestExecutedBlock() *types.Block {
	var block *types.Block
	maxNumber:=vs.view.MaxExecutedBlockNumber()
	if maxNumber==0{
		block = vs.HighestPreCommitQCBlock()
		return block
	}
	block = vs.viewBlocks.index(maxNumber).block()

	//if vs.executing.BlockIndex == math.MaxUint32 || (vs.executing.BlockIndex == 0 && !vs.executing.Finish) {
	//	block := vs.HighestQCBlock()
	//	if vs.lastViewChangeQC != nil {
	//		_, _, _, _, hash, _ := vs.lastViewChangeQC.MaxBlock()
	//		// fixme insertQCBlock should also change the state of executing
	//		if b := vs.blockTree.FindBlockByHash(hash); b != nil {
	//			if b.NumberU64()>block.NumberU64(){
	//				block = b
	//			}
	//		}
	//	}
	//	return block
	//}
	//
	//var block *types.Block
	//if vs.executing.Finish {
	//	block = vs.viewBlocks.index(vs.viewBlocks.MaxNumber()).block()
	//} else {
	//	//if vs.executing.BlockIndex<=0{
	//	//	return block
	//	//}
	//	block = vs.HighestPreCommitQCBlock()
	//	//block = vs.viewBlocks.index(vs.executing.BlockIndex-1).block()
	//}
	return block
}

func (vs *ViewState) FindBlock(hash common.Hash, number uint64) *types.Block {
	for _, b := range vs.viewBlocks.Blocks {
		if b.hash() == hash && b.number() == number {
			return b.block()
		}
	}
	return nil
}

func (vs *ViewState) SetHighestQCBlock(ext *types.Block) {
	vs.highestQCBlock.Store(ext)
}

func (vs *ViewState) HighestQCBlock() *types.Block {
	if v := vs.highestQCBlock.Load(); v == nil {
		panic("Get highest qc block failed")
	} else {
		return v.(*types.Block)
	}
}

func (vs *ViewState) SetHighestPreCommitQCBlock(ext *types.Block) {
	vs.highestPreCommitQCBlock.Store(ext)
}

func (vs *ViewState) HighestPreCommitQCBlock() *types.Block {
	if v := vs.highestPreCommitQCBlock.Load(); v == nil {
		panic("Get highest PreCommit qc block failed")
	} else {
		return v.(*types.Block)
	}
}

func (vs *ViewState) SetHighestLockBlock(ext *types.Block) {
	vs.highestLockBlock.Store(ext)
}

func (vs *ViewState) HighestLockBlock() *types.Block {
	if v := vs.highestLockBlock.Load(); v == nil {
		panic("Get highest lock block failed")
	} else {
		return v.(*types.Block)
	}
}

func (vs *ViewState) SetHighestCommitBlock(ext *types.Block) {
	vs.highestCommitBlock.Store(ext)
}

func (vs *ViewState) HighestCommitBlock() *types.Block {
	if v := vs.highestCommitBlock.Load(); v == nil {
		panic("Get highest commit block failed")
	} else {
		return v.(*types.Block)
	}
}

func (vs *ViewState) IsDeadline() bool {
	return vs.viewTimer.isDeadline()
}

func (vs *ViewState) ViewTimeout() <-chan time.Time {
	return vs.viewTimer.timerChan()
}

func (vs *ViewState) SetViewTimer(viewInterval uint64) {
	vs.viewTimer.setupTimer(viewInterval)
}

func (vs *ViewState) String() string {
	return fmt.Sprintf("")
}

func (vs *ViewState) MarshalJSON() ([]byte, error) {
	type hashNumber struct {
		Hash   common.Hash `json:"hash"`
		Number uint64      `json:"number"`
	}
	type state struct {
		View               *view      `json:"view"`
		HighestQCBlock     hashNumber `json:"highestQCBlock"`
		HighestLockBlock   hashNumber `json:"highestLockBlock"`
		HighestCommitBlock hashNumber `json:"highestCommitBlock"`
		highestPreCommitQCBlock hashNumber `json:"highestPreCommitQCBlock"`
	}

	s := &state{
		View:               vs.view,
		HighestQCBlock:     hashNumber{Hash: vs.HighestQCBlock().Hash(), Number: vs.HighestQCBlock().NumberU64()},
		HighestLockBlock:   hashNumber{Hash: vs.HighestLockBlock().Hash(), Number: vs.HighestLockBlock().NumberU64()},
		HighestCommitBlock: hashNumber{Hash: vs.HighestCommitBlock().Hash(), Number: vs.HighestCommitBlock().NumberU64()},
		highestPreCommitQCBlock: hashNumber{Hash: vs.HighestPreCommitQCBlock().Hash(), Number: vs.HighestPreCommitQCBlock().NumberU64()},
	}
	return json.Marshal(s)
}

func (vs *ViewState) UnmarshalJSON(input []byte) error {
	type hashNumber struct {
		Hash   common.Hash `json:"hash"`
		Number uint64      `json:"number"`
	}
	type state struct {
		View               *view      `json:"view"`
		HighestQCBlock     hashNumber `json:"highestQCBlock"`
		HighestLockBlock   hashNumber `json:"highestLockBlock"`
		HighestCommitBlock hashNumber `json:"highestCommitBlock"`
		highestPreCommitQCBlock hashNumber `json:"highestPreCommitQCBlock"`
	}

	var s state
	err := json.Unmarshal(input, &s)
	if err != nil {
		return err
	}

	vs.view = s.View
	vs.SetHighestQCBlock(types.NewSimplifiedBlock(s.HighestQCBlock.Number, s.HighestQCBlock.Hash))
	vs.SetHighestLockBlock(types.NewSimplifiedBlock(s.HighestLockBlock.Number, s.HighestLockBlock.Hash))
	vs.SetHighestPreCommitQCBlock(types.NewSimplifiedBlock(s.highestPreCommitQCBlock.Number, s.highestPreCommitQCBlock.Hash))
	vs.SetHighestCommitBlock(types.NewSimplifiedBlock(s.HighestCommitBlock.Number, s.HighestCommitBlock.Hash))
	return nil
}
