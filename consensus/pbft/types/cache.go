package types

import (
	"sync"
	"time"
)

type SyncCache struct {
	lock    sync.RWMutex
	items   map[interface{}]time.Time
	timeout time.Duration
}

func NewSyncCache(timeout time.Duration) *SyncCache {
	cache := &SyncCache{
		items:   make(map[interface{}]time.Time),
		timeout: timeout,
	}
	return cache
}

func (s *SyncCache) Add(v interface{}) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	s.items[v] = time.Now()
}

func (s *SyncCache) AddOrReplace(v interface{}) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if t, ok := s.items[v]; ok {
		if time.Since(t) < s.timeout {
			return false
		}
	}
	s.items[v] = time.Now()
	return true
}

func (s *SyncCache) Remove(v interface{}) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	delete(s.items, v)
}

func (s *SyncCache) Purge() {
	s.lock.RLock()
	defer s.lock.RUnlock()
	s.items = make(map[interface{}]time.Time)
}
func (s *SyncCache) Len() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.items)
}

type viewCache struct {
	prepareBlocks map[uint32]*MsgInfo
	prepareVotes  map[uint32]map[uint32]*MsgInfo
	preCommits    map[uint32]map[uint32]*MsgInfo
	prepareQC     map[uint32]*MsgInfo
	preCommitQC     map[uint32]*MsgInfo
	blockMetric   map[uint32]uint32
	preCommitMetric    map[uint32]map[uint32]uint32
	voteMetric    map[uint32]map[uint32]uint32
	qcMetric      map[uint32]uint32
	preCommitQcMetric      map[uint32]uint32
}

func newViewCache() *viewCache {
	return &viewCache{
		prepareBlocks: make(map[uint32]*MsgInfo),
		prepareVotes:  make(map[uint32]map[uint32]*MsgInfo),
		preCommits:  make(map[uint32]map[uint32]*MsgInfo),
		prepareQC:     make(map[uint32]*MsgInfo),
		preCommitQC:    make(map[uint32]*MsgInfo),
		blockMetric:   make(map[uint32]uint32),
		preCommitMetric:    make(map[uint32]map[uint32]uint32),
		voteMetric:    make(map[uint32]map[uint32]uint32),
		qcMetric:      make(map[uint32]uint32),
		preCommitQcMetric:      make(map[uint32]uint32),
	}
}

// Add prepare block to cache.
func (v *viewCache) addPrepareBlock(blockIndex uint32, msg *MsgInfo) {
	v.prepareBlocks[blockIndex] = msg
}

// Get prepare block and clear it from the cache
func (v *viewCache) getPrepareBlock(index uint32) *MsgInfo {
	if m, ok := v.prepareBlocks[index]; ok {
		v.addBlockMetric(index)
		delete(v.prepareBlocks, index)
		return m
	}
	return nil
}

func (v *viewCache) addBlockMetric(index uint32) {
	if m, ok := v.blockMetric[index]; ok {
		v.blockMetric[index] = m + 1
	} else {
		v.blockMetric[index] = 1
	}
}

// Add prepare block to cache.
func (v *viewCache) addPrepareQC(blockIndex uint32, msg *MsgInfo) {
	v.prepareQC[blockIndex] = msg
}

// Get prepare QC and clear it from the cache
func (v *viewCache) getPrepareQC(index uint32) *MsgInfo {
	if m, ok := v.prepareQC[index]; ok {
		v.addQCMetric(index)
		delete(v.prepareQC, index)
		return m
	}
	return nil
}

// Add prepare block to cache.
func (v *viewCache) addPreCommitQC(blockIndex uint32, msg *MsgInfo) {
	v.preCommitQC[blockIndex] = msg
}

// Get prepare QC and clear it from the cache
func (v *viewCache) getPreCommitQC(index uint32) *MsgInfo {
	if m, ok := v.preCommitQC[index]; ok {
		v.addPreCommitQCMetric(index)
		delete(v.prepareQC, index)
		return m
	}
	return nil
}

func (v *viewCache) getBlockMetric(index uint32) uint32 {
	if m, ok := v.blockMetric[index]; ok {
		return m
	}
	return 0
}

func (v *viewCache) addQCMetric(index uint32) {
	if m, ok := v.qcMetric[index]; ok {
		v.qcMetric[index] = m + 1
	} else {
		v.qcMetric[index] = 1
	}
}

func (v *viewCache) addPreCommitQCMetric(index uint32) {
	if m, ok := v.qcMetric[index]; ok {
		v.preCommitQcMetric[index] = m + 1
	} else {
		v.preCommitQcMetric[index] = 1
	}
}

func (v *viewCache) getQCMetric(index uint32) uint32 {
	if m, ok := v.qcMetric[index]; ok {
		return m
	}
	return 0
}

func (v *viewCache) getPreCommitQCMetric(index uint32) uint32 {
	if m, ok := v.preCommitQcMetric[index]; ok {
		return m
	}
	return 0
}

// Add prepare votes to cache.
func (v *viewCache) addPrepareVote(blockIndex uint32, validatorIndex uint32, msg *MsgInfo) {
	if votes, ok := v.prepareVotes[blockIndex]; ok {
		votes[validatorIndex] = msg
	} else {
		votes := make(map[uint32]*MsgInfo)
		votes[validatorIndex] = msg
		v.prepareVotes[blockIndex] = votes
	}
}

// Get prepare vote and clear it from the cache
func (v *viewCache) getPrepareVote(blockIndex uint32, validatorIndex uint32) *MsgInfo {
	if p, ok := v.prepareVotes[blockIndex]; ok {
		if m, ok := p[validatorIndex]; ok {
			v.addVoteMetric(blockIndex, validatorIndex)
			delete(p, validatorIndex)
			return m
		}
	}
	return nil
}

// Add precommits to cache.
func (v *viewCache) addPrecommit(blockIndex uint32, validatorIndex uint32, msg *MsgInfo) {
	if votes, ok := v.preCommits[blockIndex]; ok {
		votes[validatorIndex] = msg
	} else {
		precommits := make(map[uint32]*MsgInfo)
		precommits[validatorIndex] = msg
		v.preCommits[blockIndex] = precommits
	}
}

// Get precommit and clear it from the cache
func (v *viewCache) getPrecommit(blockIndex uint32, validatorIndex uint32) *MsgInfo {
	if p, ok := v.preCommits[blockIndex]; ok {
		if m, ok := p[validatorIndex]; ok {
			v.addPrecommitMetric(blockIndex, validatorIndex)
			delete(p, validatorIndex)
			return m
		}
	}
	return nil
}

func (v *viewCache) addPrecommitMetric(blockIndex uint32, validatorIndex uint32) {
	if preCommits, ok := v.preCommitMetric[blockIndex]; ok {
		if m, ok := preCommits[validatorIndex]; ok {
			preCommits[validatorIndex] = m + 1
		} else {
			preCommits[validatorIndex] = 1
		}
	} else {
		preCommits := make(map[uint32]uint32)
		preCommits[validatorIndex] = 1
		v.preCommitMetric[blockIndex] = preCommits
	}
}

func (v *viewCache) getPrecommitMetric(blockIndex uint32, validatorIndex uint32) uint32 {
	if preCommits, ok := v.preCommitMetric[blockIndex]; ok {
		if m, ok := preCommits[validatorIndex]; ok {
			return m
		}
	}
	return 0
}

func (v *viewCache) addVoteMetric(blockIndex uint32, validatorIndex uint32) {
	if votes, ok := v.voteMetric[blockIndex]; ok {
		if m, ok := votes[validatorIndex]; ok {
			votes[validatorIndex] = m + 1
		} else {
			votes[validatorIndex] = 1
		}
	} else {
		votes := make(map[uint32]uint32)
		votes[validatorIndex] = 1
		v.voteMetric[blockIndex] = votes
	}
}

func (v *viewCache) getVoteMetric(blockIndex uint32, validatorIndex uint32) uint32 {
	if votes, ok := v.voteMetric[blockIndex]; ok {
		if m, ok := votes[validatorIndex]; ok {
			return m
		}
	}
	return 0
}

type epochCache struct {
	views map[uint64]*viewCache  //map[blockNumber]*viewCache
}

func newEpochCache() *epochCache {
	return &epochCache{
		views: make(map[uint64]*viewCache),
	}
}

func (e *epochCache) matchViewCache(blockNumber uint64) *viewCache {
	for k, v := range e.views {
		if k == blockNumber {
			return v
		}
	}
	newView := newViewCache()
	e.views[blockNumber] = newView
	return newView
}

func (e *epochCache) findViewCache(blockNumber uint64) *viewCache {
	for k, v := range e.views {
		if k == blockNumber {
			return v
		}
	}
	return nil
}

func (e *epochCache) purge(blockNumber uint64) {
	for k, _ := range e.views {
		if k < blockNumber {
			delete(e.views, k)
		}
	}
}

type CSMsgPool struct {
	epochs   map[uint64]*epochCache
	minEpoch uint64
	minBlockNumber  uint64
}

func NewCSMsgPool() *CSMsgPool {
	return &CSMsgPool{
		epochs:   make(map[uint64]*epochCache),
		minEpoch: 0,
		minBlockNumber:  0,
	}
}

func (cs *CSMsgPool) invalidEpochView(epoch, blockNumber uint64) bool {
	if (cs.minEpoch == epoch||cs.minEpoch+1 == epoch) && (cs.minBlockNumber == blockNumber||cs.minBlockNumber+1 == blockNumber) {
		return false
	}
	return true
}

// Add prepare block to cache.
func (cs *CSMsgPool) AddPrepareBlock(blockIndex uint32, msg *MsgInfo) {
	if csMsg, ok := msg.Msg.(ConsensusMsg); ok {
		if cs.invalidEpochView(csMsg.EpochNum(), csMsg.BlockNum()) || msg.Inner {
			return
		}
		cs.matchEpochCache(csMsg.EpochNum()).
			matchViewCache(csMsg.BlockNum()).
			addPrepareBlock(blockIndex, msg)
	}
}

// Get prepare block and clear it from the cache
func (cs *CSMsgPool) GetPrepareBlock(epoch, blockNumber uint64, index uint32) *MsgInfo {
	if cs.invalidEpochView(epoch, blockNumber) {
		return nil
	}

	viewCache := cs.findViewCache(epoch, blockNumber)
	if viewCache != nil {
		return viewCache.getPrepareBlock(index)
	}
	return nil
}

// Add prepare QC to cache.
func (cs *CSMsgPool) AddPrepareQC(epoch, blockNumber uint64, blockIndex uint32, msg *MsgInfo) {
	if cs.invalidEpochView(epoch, blockNumber) || msg.Inner {
		return
	}

	cs.matchEpochCache(epoch).
		matchViewCache(blockNumber).
		addPrepareQC(blockIndex, msg)
}

// Get prepare QC and clear it from the cache
func (cs *CSMsgPool) GetPrepareQC(epoch, blockNumber uint64, index uint32) *MsgInfo {
	if cs.invalidEpochView(epoch, blockNumber) {
		return nil
	}

	viewCache := cs.findViewCache(epoch, blockNumber)
	if viewCache != nil {
		return viewCache.getPrepareQC(index)
	}
	return nil
}

// Add prepare QC to cache.
func (cs *CSMsgPool) AddPreCommitQC(epoch, blockNumber uint64, blockIndex uint32, msg *MsgInfo) {
	if cs.invalidEpochView(epoch, blockNumber) || msg.Inner {
		return
	}

	cs.matchEpochCache(epoch).
		matchViewCache(blockNumber).
		addPreCommitQC(blockIndex, msg)
}

// Get prepare QC and clear it from the cache
func (cs *CSMsgPool) GetPreCommitQC(epoch, blockNumber uint64, index uint32) *MsgInfo {
	if cs.invalidEpochView(epoch, blockNumber) {
		return nil
	}

	viewCache := cs.findViewCache(epoch, blockNumber)
	if viewCache != nil {
		return viewCache.getPreCommitQC(index)
	}
	return nil
}

// Add prepare votes to cache.
func (cs *CSMsgPool) AddPrepareVote(blockIndex uint32, validatorIndex uint32, msg *MsgInfo) {
	if csMsg, ok := msg.Msg.(ConsensusMsg); ok {
		if cs.invalidEpochView(csMsg.EpochNum(), csMsg.BlockNum()) || msg.Inner {
			return
		}
		cs.matchEpochCache(csMsg.EpochNum()).
			matchViewCache(csMsg.BlockNum()).
			addPrepareVote(blockIndex, validatorIndex, msg)
	}
}

// Get prepare vote and clear it from the cache
func (cs *CSMsgPool) GetPrepareVote(epoch, blockNumber uint64, blockIndex uint32, validatorIndex uint32) *MsgInfo {
	if cs.invalidEpochView(epoch, blockNumber) {
		return nil
	}

	viewCache := cs.findViewCache(epoch, blockNumber)
	if viewCache != nil {
		return viewCache.getPrepareVote(blockIndex, validatorIndex)
	}
	return nil
}

// Add prepare votes to cache.
func (cs *CSMsgPool) AddPreCommit(blockIndex uint32, validatorIndex uint32, msg *MsgInfo) {
	if csMsg, ok := msg.Msg.(ConsensusMsg); ok {
		if cs.invalidEpochView(csMsg.EpochNum(), csMsg.BlockNum()) || msg.Inner {
			return
		}
		cs.matchEpochCache(csMsg.EpochNum()).
			matchViewCache(csMsg.BlockNum()).
			addPrecommit(blockIndex, validatorIndex, msg)
	}
}

// Get prepare vote and clear it from the cache
func (cs *CSMsgPool) GetPreCommit(epoch, blockNumber uint64, blockIndex uint32, validatorIndex uint32) *MsgInfo {
	if cs.invalidEpochView(epoch, blockNumber) {
		return nil
	}

	viewCache := cs.findViewCache(epoch, blockNumber)
	if viewCache != nil {
		return viewCache.getPrecommit(blockIndex, validatorIndex)
	}
	return nil
}

func (cs *CSMsgPool) Purge(epoch, blockNumber uint64) {

	for k, v := range cs.epochs {
		if k < epoch {
			delete(cs.epochs, k)
		} else {
			v.purge(blockNumber)
		}
	}
	cs.minEpoch = epoch
	cs.minBlockNumber = blockNumber
}

func (cs *CSMsgPool) matchEpochCache(epoch uint64) *epochCache {
	for k, v := range cs.epochs {
		if k == epoch {
			return v
		}
	}
	newEpoch := newEpochCache()
	cs.epochs[epoch] = newEpoch
	return newEpoch
}

func (cs *CSMsgPool) findEpochCache(epoch uint64) *epochCache {
	for k, v := range cs.epochs {
		if k == epoch {
			return v
		}
	}
	return nil
}

func (cs *CSMsgPool) findViewCache(epoch, blockNumber uint64) *viewCache {
	epochCache := cs.findEpochCache(epoch)
	if epochCache != nil {
		return epochCache.findViewCache(blockNumber)
	}
	return nil
}
