package pbft

import (
	"bytes"
	"container/list"
	"crypto/elliptic"
	"encoding/json"
	"fmt"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core/types/pbfttypes"
	"strings"
	"sync/atomic"

	mapset "github.com/deckarep/golang-set"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common/hexutil"

	"github.com/pkg/errors"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/crypto/bls"

	"reflect"
	"sync"
	"time"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/configs"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/consensus"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/consensus/pbft/evidence"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/consensus/pbft/executor"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/consensus/pbft/fetcher"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/consensus/pbft/network"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/consensus/pbft/protocols"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/consensus/pbft/rules"
	cstate "github.com/PhoenixGlobal/Phoenix-Chain-Core/consensus/pbft/state"
	ctypes "github.com/PhoenixGlobal/Phoenix-Chain-Core/consensus/pbft/types"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/consensus/pbft/utils"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/consensus/pbft/validator"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/consensus/pbft/wal"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core/state"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core/types"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/node"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/p2p"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/p2p/discover"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/crypto"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/event"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/log"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/rpc"
)

const (
	pbftVersion = 1

	maxStatQueuesSize      = 200
	syncCacheTimeout       = 200 * time.Millisecond
	checkBlockSyncInterval = 100 * time.Millisecond
)

var (
	ErrorUnKnowBlock  = errors.New("unknown block")
	ErrorEngineBusy   = errors.New("pbft engine busy")
	ErrorNotRunning   = errors.New("pbft is not running")
	ErrorNotValidator = errors.New("current node not a validator")
	ErrorTimeout      = errors.New("timeout")
)

type HandleError interface {
	error
	AuthFailed() bool
}

type handleError struct {
	err error
}

func (e handleError) Error() string {
	return e.err.Error()
}

func (e handleError) AuthFailed() bool {
	return false
}

type authFailedError struct {
	err error
}

func (e authFailedError) Error() string {
	return e.err.Error()
}

func (e authFailedError) AuthFailed() bool {
	return true
}

// Pbft is the core structure of the consensus engine
// and is responsible for handling consensus logic.
type Pbft struct {
	config           ctypes.Config
	eventMux         *event.TypeMux
	closeOnce        sync.Once
	exitCh           chan struct{}
	txPool           consensus.TxPoolReset
	blockChain       consensus.ChainReader
	blockCacheWriter consensus.BlockCacheWriter
	peerMsgCh        chan *ctypes.MsgInfo
	syncMsgCh        chan *ctypes.MsgInfo
	evPool           evidence.EvidencePool
	log              log.Logger
	network          *network.EngineManager

	start    int32
	syncing  int32
	fetching int32

	// Commit block error
	commitErrCh chan error
	// Async call channel
	asyncCallCh chan func()

	fetcher *fetcher.Fetcher

	// Synchronizing request cache to prevent multiple repeat requests
	syncingCache *ctypes.SyncCache
	// Control the current view state
	state *cstate.ViewState

	// Block asyncExecutor, the block responsible for executing the current view
	asyncExecutor     executor.AsyncBlockExecutor
	executeStatusHook func(s *executor.BlockExecuteStatus)

	// Verification security rules for proposed blocks and viewchange
	safetyRules rules.SafetyRules

	// Determine when to allow voting
	voteRules rules.VoteRules

	// Validator pool
	validatorPool *validator.ValidatorPool

	// Store blocks that are not committed
	blockTree *ctypes.BlockTree

	csPool *ctypes.CSMsgPool

	// wal
	nodeServiceContext *node.ServiceContext
	wal                wal.Wal
	bridge             Bridge

	loading                   int32
	updateChainStateHook      pbfttypes.UpdateChainStateFn
	updateChainStateDelayHook func(qcState, lockState, commitState *protocols.State)

	// Record the number of peer requests for obtaining pbft information.
	queues     map[string]int // Per peer message counts to prevent memory exhaustion.
	queuesLock sync.RWMutex

	// Record message repetitions.
	statQueues       map[common.Hash]map[string]int
	statQueuesLock   sync.RWMutex
	messageHashCache mapset.Set

	// Delay time of each node
	netLatencyMap  map[string]*list.List
	netLatencyLock sync.RWMutex

	//test
	insertBlockQCHook  func(block *types.Block, qc *ctypes.QuorumCert)
	executeFinishHook  func(index uint32)
	consensusNodesMock func() ([]discover.NodeID, error)
}

// New returns a new PBFT.
func New(sysConfig *configs.PbftConfig, optConfig *ctypes.OptionsConfig, eventMux *event.TypeMux, ctx *node.ServiceContext) *Pbft {
	pbft := &Pbft{
		config:             ctypes.Config{Sys: sysConfig, Option: optConfig},
		eventMux:           eventMux,
		exitCh:             make(chan struct{}),
		peerMsgCh:          make(chan *ctypes.MsgInfo, optConfig.PeerMsgQueueSize),
		syncMsgCh:          make(chan *ctypes.MsgInfo, optConfig.PeerMsgQueueSize),
		csPool:             ctypes.NewCSMsgPool(),
		log:                log.New(),
		start:              0,
		syncing:            0,
		fetching:           0,
		commitErrCh:        make(chan error, 1),
		asyncCallCh:        make(chan func(), optConfig.PeerMsgQueueSize),
		fetcher:            fetcher.NewFetcher(),
		syncingCache:       ctypes.NewSyncCache(syncCacheTimeout),
		nodeServiceContext: ctx,
		queues:             make(map[string]int),
		statQueues:         make(map[common.Hash]map[string]int),
		messageHashCache:   mapset.NewSet(),
		netLatencyMap:      make(map[string]*list.List),
	}

	if evPool, err := evidence.NewEvidencePool(ctx, optConfig.EvidenceDir); err == nil {
		pbft.evPool = evPool
	} else {
		return nil
	}

	return pbft
}

// Start starts consensus engine.
func (pbft *Pbft) Start(chain consensus.ChainReader, blockCacheWriter consensus.BlockCacheWriter, txPool consensus.TxPoolReset, agency consensus.Agency) error {
	pbft.log.Info("~ Start pbft consensus")
	pbft.blockChain = chain
	pbft.txPool = txPool
	pbft.blockCacheWriter = blockCacheWriter
	pbft.asyncExecutor = executor.NewAsyncExecutor(blockCacheWriter.Execute)

	//Initialize block tree
	block := chain.GetBlock(chain.CurrentHeader().Hash(), chain.CurrentHeader().Number.Uint64())
	//block := chain.CurrentBlock()
	isGenesis := func() bool {
		return block.NumberU64() == 0
	}

	var qc *ctypes.QuorumCert
	if !isGenesis() {
		var err error
		_, qc, err = ctypes.DecodeExtra(block.ExtraData())

		if err != nil {
			pbft.log.Error("It's not genesis", "err", err)
			return errors.Wrap(err, fmt.Sprintf("start pbft failed"))
		}
	}

	pbft.log.Info("Pbft engine init status","initialize block number",block.NumberU64(),"qc",qc.String())

	pbft.blockTree = ctypes.NewBlockTree(block, qc)
	utils.SetTrue(&pbft.loading)

	//Initialize view state
	pbft.state = cstate.NewViewState(pbft.config.Sys.Period, pbft.blockTree)
	pbft.state.SetHighestQCBlock(block)
	pbft.state.SetHighestLockBlock(block)
	pbft.state.SetHighestPreCommitQCBlock(block)
	pbft.state.SetHighestCommitBlock(block)

	// init handler and router to process message.
	// pbft -> handler -> router.
	pbft.network = network.NewEngineManger(pbft) // init engineManager as handler.
	// Start the handler to process the message.
	go pbft.network.Start()

	if pbft.config.Option.NodePriKey == nil {
		pbft.config.Option.NodePriKey = pbft.nodeServiceContext.NodePriKey()
		pbft.config.Option.NodeID = discover.PubkeyID(&pbft.config.Option.NodePriKey.PublicKey)
	}

	if isGenesis() {
		pbft.validatorPool = validator.NewValidatorPool(agency, block.NumberU64(), cstate.DefaultEpoch, pbft.config.Option.NodeID)
		pbft.changeView(cstate.DefaultEpoch, cstate.DefaultViewNumber, block, qc, nil)
	} else {
		pbft.validatorPool = validator.NewValidatorPool(agency, block.NumberU64(), qc.Epoch, pbft.config.Option.NodeID)
		pbft.changeView(qc.Epoch, qc.ViewNumber, block, qc, nil)
	}

	// Initialize current view
	if qc != nil {
		pbft.initStateBlockNumberByQC(qc)
		pbft.state.SetExecuting(qc.BlockIndex, true)
		pbft.state.SetMaxExecutedBlockNumber(qc.BlockNumber)
		//pbft.state.SetFinish(true)
		pbft.state.AddQCBlock(block, qc)
		pbft.state.AddQC(qc)
		pbft.state.AddPreCommitQC(qc)
	}else {
		if !isGenesis(){
			pbft.initStateBlockNumberByBlockNumber(block.NumberU64())
		}
	}

	// try change view again
	pbft.tryChangeView()

	//Initialize rules
	pbft.safetyRules = rules.NewSafetyRules(pbft.state, pbft.blockTree, &pbft.config, pbft.validatorPool)
	pbft.voteRules = rules.NewVoteRules(pbft.state)

	// load consensus state
	if err := pbft.LoadWal(); err != nil {
		pbft.log.Error("Load wal failed", "err", err)
		return err
	}
	utils.SetFalse(&pbft.loading)

	go pbft.receiveLoop()

	pbft.fetcher.Start()

	utils.SetTrue(&pbft.start)
	pbft.log.Info("Pbft engine start")
	return nil
}

// ReceiveMessage Entrance: The messages related to the consensus are entered from here.
// The message sent from the peer node is sent to the PBFT message queue and
// there is a loop that will distribute the incoming message.
func (pbft *Pbft) ReceiveMessage(msg *ctypes.MsgInfo) error {
	if !pbft.running() {
		//pbft.log.Trace("Pbft not running, stop process message", "fecthing", utils.True(&pbft.fetching), "syncing", utils.True(&pbft.syncing))
		return nil
	}

	invalidMsg := func(epoch, blockNumber uint64) bool {
		if epoch >= pbft.state.Epoch() && blockNumber+1 >= pbft.state.BlockNumber() {
			return false
		}
		return true
	}

	cMsg := msg.Msg.(ctypes.ConsensusMsg)
	if invalidMsg(cMsg.EpochNum(), cMsg.BlockNum()) {
		pbft.log.Debug("Invalid msg", "peer", msg.PeerID, "type", reflect.TypeOf(msg.Msg), "msg", msg.Msg.String())
		return nil
	}

	err := pbft.recordMessage(msg)
	//pbft.log.Debug("Record message", "type", fmt.Sprintf("%T", msg.Msg), "msgHash", msg.Msg.MsgHash(), "duration", time.Since(begin))
	if err != nil {
		pbft.log.Warn("ReceiveMessage failed", "err", err)
		return err
	}

	// Repeat filtering on consensus messages.
	// First check.
	if pbft.network.ContainsHistoryMessageHash(msg.Msg.MsgHash()) {
		//pbft.log.Trace("Processed message for ReceiveMessage, no need to process", "msgHash", msg.Msg.MsgHash())
		pbft.forgetMessage(msg.PeerID)
		return nil
	}

	select {
	case pbft.peerMsgCh <- msg:
		pbft.log.Debug("Received message from peer", "peer", msg.PeerID, "type", fmt.Sprintf("%T", msg.Msg), "msgHash", msg.Msg.MsgHash(), "BHash", msg.Msg.BHash(), "msg", msg.String(), "peerMsgCh", len(pbft.peerMsgCh))
	case <-pbft.exitCh:
		pbft.log.Warn("Pbft exit")
	default:
		pbft.log.Debug("peerMsgCh is full, discard", "peerMsgCh", len(pbft.peerMsgCh))
	}
	return nil
}

// recordMessage records the number of messages sent by each node,
// mainly to prevent Dos attacks
func (pbft *Pbft) recordMessage(msg *ctypes.MsgInfo) error {
	pbft.queuesLock.Lock()
	defer pbft.queuesLock.Unlock()
	count := pbft.queues[msg.PeerID] + 1
	if int64(count) > pbft.config.Option.MaxQueuesLimit {
		log.Warn("Discarded message, exceeded allowance for the layer of pbft", "peer", msg.PeerID, "msgHash", msg.Msg.MsgHash().TerminalString())
		// Need further confirmation.
		// todo: Is the program exiting or dropping the message here?
		return fmt.Errorf("execeed max queues limit")
	}
	pbft.queues[msg.PeerID] = count
	return nil
}

// forgetMessage clears the record after the message processing is completed.
func (pbft *Pbft) forgetMessage(peerID string) error {
	pbft.queuesLock.Lock()
	defer pbft.queuesLock.Unlock()
	// After the message is processed, the counter is decremented by one.
	// If it is reduced to 0, the mapping relationship of the corresponding
	// node will be deleted.
	pbft.queues[peerID]--
	if pbft.queues[peerID] == 0 {
		delete(pbft.queues, peerID)
	}
	return nil
}

// statMessage statistics record of duplicate messages.
func (pbft *Pbft) statMessage(msg *ctypes.MsgInfo) error {
	if msg == nil {
		return fmt.Errorf("invalid msg")
	}
	pbft.statQueuesLock.Lock()
	defer pbft.statQueuesLock.Unlock()

	for pbft.messageHashCache.Cardinality() >= maxStatQueuesSize {
		msgHash := pbft.messageHashCache.Pop().(common.Hash)
		// Printout.
		var bf bytes.Buffer
		for k, v := range pbft.statQueues[msgHash] {
			bf.WriteString(fmt.Sprintf("{%s:%d},", k, v))
		}
		output := strings.TrimSuffix(bf.String(), ",")
		pbft.log.Debug("Statistics sync message", "msgHash", msgHash, "stat", output)
		// remove the key from map.
		delete(pbft.statQueues, msgHash)
	}
	// Reset the variable if there is a difference
	// between the map and the set data.
	if len(pbft.statQueues) > maxStatQueuesSize {
		pbft.messageHashCache.Clear()
		pbft.statQueues = make(map[common.Hash]map[string]int)
	}

	hash := msg.Msg.MsgHash()
	if _, ok := pbft.statQueues[hash]; ok {
		if _, exists := pbft.statQueues[hash][msg.PeerID]; exists {
			pbft.statQueues[hash][msg.PeerID]++
		} else {
			pbft.statQueues[hash][msg.PeerID] = 1
		}
	} else {
		pbft.statQueues[hash] = map[string]int{
			msg.PeerID: 1,
		}
		pbft.messageHashCache.Add(hash)
	}
	return nil
}

// ReceiveSyncMsg is used to receive messages that are synchronized from other nodes.
//
// Possible message types are:
//  PrepareBlockVotesMsg/GetLatestStatusMsg/LatestStatusMsg/
func (pbft *Pbft) ReceiveSyncMsg(msg *ctypes.MsgInfo) error {
	// If the node is synchronizing the block, discard sync msg directly and do not count the msg
	// When the syncMsgCh channel is congested, it is easy to cause a message backlog
	if utils.True(&pbft.syncing) {
		pbft.log.Debug("Currently syncing, consensus message pause, discard sync msg")
		return nil
	}

	err := pbft.recordMessage(msg)
	if err != nil {
		pbft.log.Warn("ReceiveMessage failed", "err", err)
		return err
	}

	// message stat.
	pbft.statMessage(msg)

	// Non-core consensus messages are temporarily not filtered repeatedly.
	select {
	case pbft.syncMsgCh <- msg:
		pbft.log.Debug("Receive synchronization related messages from peer", "peer", msg.PeerID, "type", fmt.Sprintf("%T", msg.Msg), "msgHash", msg.Msg.MsgHash(), "BHash", msg.Msg.BHash(), "msg", msg.Msg.String(), "syncMsgCh", len(pbft.syncMsgCh))
	case <-pbft.exitCh:
		pbft.log.Warn("Pbft exit")
	default:
		pbft.log.Debug("syncMsgCh is full, discard", "syncMsgCh", len(pbft.syncMsgCh))
	}
	return nil
}

// LoadWal tries to recover consensus state and view msg from the wal.
func (pbft *Pbft) LoadWal() (err error) {
	// init wal and load wal state
	var context *node.ServiceContext
	if pbft.config.Option.WalMode {
		context = pbft.nodeServiceContext
	}
	if pbft.wal, err = wal.NewWal(context, ""); err != nil {
		return err
	}
	if pbft.bridge, err = NewBridge(context, pbft); err != nil {
		return err
	}

	// load consensus chainState
	if err = pbft.wal.LoadChainState(pbft.recoveryChainState); err != nil {
		pbft.log.Error(err.Error())
		return err
	}
	// load consensus message
	if err = pbft.wal.Load(pbft.recoveryMsg); err != nil {
		pbft.log.Error(err.Error())
		return err
	}
	return nil
}

// receiveLoop receives all consensus related messages, all processing logic in the same goroutine
func (pbft *Pbft) receiveLoop() {

	// Responsible for handling consensus message logic.
	consensusMessageHandler := func(msg *ctypes.MsgInfo) {
		if !pbft.network.ContainsHistoryMessageHash(msg.Msg.MsgHash()) {
			err := pbft.handleConsensusMsg(msg)
			if err == nil {
				pbft.network.MarkHistoryMessageHash(msg.Msg.MsgHash())
				if err := pbft.network.Forwarding(msg.PeerID, msg.Msg); err != nil {
					pbft.log.Debug("Forward message failed", "err", err)
				}
			} else if e, ok := err.(HandleError); ok && e.AuthFailed() {
				// If the verification signature is abnormal,
				// the peer node is added to the local blacklist
				// and disconnected.
				pbft.log.Error("Verify signature failed, will add to blacklist", "peerID", msg.PeerID, "err", err)
				pbft.network.MarkBlacklist(msg.PeerID)
				pbft.network.RemovePeer(msg.PeerID)
			}
		} else {
			//pbft.log.Trace("The message has been processed, discard it", "msgHash", msg.Msg.MsgHash(), "peerID", msg.PeerID)
		}
		pbft.forgetMessage(msg.PeerID)
	}

	// channel Divided into read-only type, writable type
	// Read-only is the channel that gets the current PBFT status.
	// Writable type is the channel that affects the consensus state.
	for {
		select {
		case msg := <-pbft.peerMsgCh:
			consensusMessageHandler(msg)
		default:
		}
		select {
		case msg := <-pbft.peerMsgCh:
			// Forward the message before processing the message.
			consensusMessageHandler(msg)
		case msg := <-pbft.syncMsgCh:
			if err := pbft.handleSyncMsg(msg); err != nil {
				if err, ok := err.(HandleError); ok {
					if err.AuthFailed() {
						pbft.log.Error("Verify signature failed to sync message, will add to blacklist", "peerID", msg.PeerID)
						pbft.network.MarkBlacklist(msg.PeerID)
						pbft.network.RemovePeer(msg.PeerID)
					}
				}
			}
			pbft.forgetMessage(msg.PeerID)
		case msg := <-pbft.asyncExecutor.ExecuteStatus():
			pbft.onAsyncExecuteStatus(msg)
			if pbft.executeStatusHook != nil {
				pbft.executeStatusHook(msg)
			}

		case fn := <-pbft.asyncCallCh:
			fn()

		case <-pbft.state.ViewTimeout():
			pbft.OnViewTimeout()
		case err := <-pbft.commitErrCh:
			pbft.OnCommitError(err)
		}
	}
}

// Handling consensus messages, there are three main types of messages. prepareBlock, prepareVote, viewChange
func (pbft *Pbft) handleConsensusMsg(info *ctypes.MsgInfo) error {
	if !pbft.running() {
		pbft.log.Debug("Consensus message pause", "syncing", atomic.LoadInt32(&pbft.syncing), "fetching", atomic.LoadInt32(&pbft.fetching), "peerMsgCh", len(pbft.peerMsgCh))
		return &handleError{fmt.Errorf("consensus message pause, ignore message")}
	}
	msg, id := info.Msg, info.PeerID
	var err error

	switch msg := msg.(type) {
	case *protocols.PrepareBlock:
		pbft.csPool.AddPrepareBlock(msg.BlockIndex, ctypes.NewInnerMsgInfo(info.Msg, info.PeerID))
		err = pbft.OnPrepareBlock(id, msg)
	case *protocols.PrepareVote:
		pbft.csPool.AddPrepareVote(msg.BlockIndex, msg.ValidatorIndex, ctypes.NewInnerMsgInfo(info.Msg, info.PeerID))
		err = pbft.OnPrepareVote(id, msg)
	case *protocols.PreCommit:
		pbft.csPool.AddPreCommit(msg.BlockIndex, msg.ValidatorIndex, ctypes.NewInnerMsgInfo(info.Msg, info.PeerID))
		err = pbft.OnPreCommit(id, msg)
	case *protocols.ViewChange:
		err = pbft.OnViewChange(id, msg)
	}

	if err != nil {
		pbft.log.Debug("Handle msg Failed", "error", err, "type", reflect.TypeOf(msg), "peer", id, "err", err, "peerMsgCh", len(pbft.peerMsgCh))
	}
	return err
}

// Behind the node will be synchronized by synchronization message
func (pbft *Pbft) handleSyncMsg(info *ctypes.MsgInfo) error {
	if utils.True(&pbft.syncing) {
		pbft.log.Debug("Currently syncing, consensus message pause", "syncMsgCh", len(pbft.syncMsgCh))
		return nil
	}
	msg, id := info.Msg, info.PeerID
	var err error
	if !pbft.fetcher.MatchTask(id, msg) {
		switch msg := msg.(type) {
		case *protocols.GetPrepareBlock:
			err = pbft.OnGetPrepareBlock(id, msg)

		case *protocols.GetBlockQuorumCert:
			err = pbft.OnGetBlockQuorumCert(id, msg)

		case *protocols.BlockQuorumCert:
			pbft.csPool.AddPrepareQC(msg.BlockQC.Epoch, msg.BlockQC.BlockNumber, msg.BlockQC.BlockIndex, ctypes.NewInnerMsgInfo(info.Msg, info.PeerID))
			err = pbft.OnBlockQuorumCert(id, msg)

		case *protocols.GetBlockPreCommitQuorumCert:
			err = pbft.OnGetBlockPreCommitQuorumCert(id, msg)

		case *protocols.BlockPreCommitQuorumCert:
			pbft.csPool.AddPreCommitQC(msg.BlockQC.Epoch, msg.BlockQC.BlockNumber, msg.BlockQC.BlockIndex, ctypes.NewInnerMsgInfo(info.Msg, info.PeerID))
			err = pbft.OnBlockPreCommitQuorumCert(id, msg)

		case *protocols.GetPrepareVote:
			err = pbft.OnGetPrepareVote(id, msg)

		case *protocols.GetPreCommit:
			err = pbft.OnGetPreCommit(id, msg)

		case *protocols.PrepareVotes:
			err = pbft.OnPrepareVotes(id, msg)

		case *protocols.PreCommits:
			err = pbft.OnPreCommits(id, msg)

		case *protocols.GetQCBlockList:
			err = pbft.OnGetQCBlockList(id, msg)

		case *protocols.GetLatestStatus:
			err = pbft.OnGetLatestStatus(id, msg)

		case *protocols.LatestStatus:
			err = pbft.OnLatestStatus(id, msg)

		case *protocols.PrepareBlockHash:
			err = pbft.OnPrepareBlockHash(id, msg)

		case *protocols.GetViewChange:
			err = pbft.OnGetViewChange(id, msg)

		case *protocols.ViewChangeQuorumCert:
			err = pbft.OnViewChangeQuorumCert(id, msg)

		case *protocols.ViewChanges:
			err = pbft.OnViewChanges(id, msg)
		}
	}
	return err
}

// running returns whether the consensus engine is running.
func (pbft *Pbft) running() bool {
	return utils.False(&pbft.syncing) && utils.False(&pbft.fetching)
}

// Author returns the current node's Author.
func (pbft *Pbft) Author(header *types.Header) (common.Address, error) {
	return header.Coinbase, nil
}

// VerifyHeader verify the validity of the block header.
func (pbft *Pbft) VerifyHeader(chain consensus.ChainReader, header *types.Header, seal bool) error {
	if header.Number == nil {
		pbft.log.Error("Verify header fail, unknown block")
		return ErrorUnKnowBlock
	}

	//pbft.log.Trace("Verify header", "number", header.Number, "hash", header.Hash, "seal", seal)
	if len(header.Extra) < consensus.ExtraSeal+int(configs.MaximumExtraDataSize) {
		pbft.log.Error("Verify header fail, missing signature", "number", header.Number, "hash", header.Hash)
		return fmt.Errorf("verify header fail, missing signature, number:%d, hash:%s", header.Number.Uint64(), header.Hash().String())
	}

	if header.IsInvalid() {
		pbft.log.Error("Verify header fail, Extra field is too long", "number", header.Number, "hash", header.CacheHash())
		return fmt.Errorf("verify header fail, Extra field is too long, number:%d, hash:%s", header.Number.Uint64(), header.CacheHash().String())
	}

	if err := pbft.validatorPool.VerifyHeader(header); err != nil {
		pbft.log.Error("Verify header fail", "number", header.Number, "hash", header.Hash(), "err", err)
		return fmt.Errorf("verify header fail, number:%d, hash:%s, err:%s", header.Number.Uint64(), header.Hash().String(), err.Error())
	}
	return nil
}

// VerifyHeaders is used to verify the validity of block headers in batch.
func (pbft *Pbft) VerifyHeaders(chain consensus.ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	//pbft.log.Trace("Verify headers", "total", len(headers))

	abort := make(chan struct{})
	results := make(chan error, len(headers))

	go func() {
		for _, header := range headers {
			err := pbft.VerifyHeader(chain, header, false)

			select {
			case <-abort:
				return
			case results <- err:
			}
		}
	}()
	return abort, results
}

// VerifySeal implements consensus.Engine, checking whether the signature contained
// in the header satisfies the consensus protocol requirements.
func (pbft *Pbft) VerifySeal(chain consensus.ChainReader, header *types.Header) error {
	//pbft.log.Trace("Verify seal", "hash", header.Hash(), "number", header.Number)
	if header.Number.Uint64() == 0 {
		return ErrorUnKnowBlock
	}
	return nil
}

// Prepare implements consensus.Engine, preparing all the consensus fields of the
// header of running the transactions on top.
func (pbft *Pbft) Prepare(chain consensus.ChainReader, header *types.Header) error {
	pbft.log.Debug("Prepare", "hash", header.Hash(), "number", header.Number.Uint64())

	//header.Extra[0:31] to store block's version info etc. and right pad with 0x00;
	//header.Extra[32:] to store block's sign of producer, the length of sign is 65.
	if len(header.Extra) < 32 {
		pbft.log.Debug("Prepare, add header-extra byte 0x00 till 32 bytes", "extraLength", len(header.Extra))
		header.Extra = append(header.Extra, bytes.Repeat([]byte{0x00}, 32-len(header.Extra))...)
	}
	header.Extra = header.Extra[:32]

	//init header.Extra[32: 32+65]
	header.Extra = append(header.Extra, make([]byte, consensus.ExtraSeal)...)
	pbft.log.Debug("Prepare, add header-extra ExtraSeal bytes(0x00)", "extraLength", len(header.Extra))
	return nil
}

// Finalize implements consensus.Engine, no block
// rewards given, and returns the final block.
func (pbft *Pbft) Finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, receipts []*types.Receipt) (*types.Block, error) {
	header.Root = state.IntermediateRoot(true)
	pbft.log.Debug("Finalize block", "hash", header.Hash(), "number", header.Number, "txs", len(txs), "receipts", len(receipts), "root", header.Root.String())
	return types.NewBlock(header, txs, receipts), nil
}

// Seal is used to generate a block, and block data is
// passed to the execution channel.
func (pbft *Pbft) Seal(chain consensus.ChainReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}, complete chan<- struct{}) error {
	pbft.log.Info("Seal block", "number", block.Number(), "parentHash", block.ParentHash())
	header := block.Header()
	if block.NumberU64() == 0 {
		return ErrorUnKnowBlock
	}

	sign, err := pbft.signFn(header.SealHash().Bytes())
	if err != nil {
		pbft.log.Error("Seal block sign fail", "number", block.Number(), "parentHash", block.ParentHash(), "err", err)
		return err
	}

	copy(header.Extra[len(header.Extra)-consensus.ExtraSeal:], sign[:])

	sealBlock := block.WithSeal(header)

	pbft.asyncCallCh <- func() {
		pbft.OnSeal(sealBlock, results, stop, complete)
	}
	return nil
}

func (pbft *Pbft) DealBlock(block *types.Block)(b *types.Block,e error) {
	header := block.Header()
	if block.NumberU64() == 0 {
		e=ErrorUnKnowBlock
		return
	}

	sign, err := pbft.signFn(header.SealHash().Bytes())
	if err != nil {
		pbft.log.Error("Seal block sign fail", "number", block.Number(), "parentHash", block.ParentHash(), "err", err)
		e=err
		return
	}

	copy(header.Extra[len(header.Extra)-consensus.ExtraSeal:], sign[:])

	b = block.WithSeal(header)
	return
}

// OnSeal is used to process the blocks that have already been generated.
func (pbft *Pbft) OnSeal(block *types.Block, results chan<- *types.Block, stop <-chan struct{}, complete chan<- struct{}) {
	// When the seal ends, the completion signal will be passed to the goroutine of worker
	defer func() { complete <- struct{}{} }()

	if pbft.state.HighestPreCommitQCBlock().Hash() != block.ParentHash() {
		pbft.log.Warn("Futile block cause highest preCommitQC block changed", "number", block.Number(), "parentHash", block.ParentHash(),
			"qcNumber", pbft.state.HighestPreCommitQCBlock().Number(), "qcHash", pbft.state.HighestPreCommitQCBlock().Hash())
		return
	}

	me, err := pbft.validatorPool.GetValidatorByNodeID(pbft.state.Epoch(), pbft.NodeID())
	if err != nil {
		pbft.log.Warn("Can not got the validator, seal fail", "epoch", pbft.state.Epoch(), "nodeID", pbft.NodeID())
		return
	}
	numValidators := pbft.validatorPool.Len(pbft.state.Epoch())
	//currentProposer := pbft.state.BlockNumber() % uint64(numValidators)
	//currentProposer := (pbft.state.BlockNumber()-1)% uint64(numValidators)
	currentProposer := pbft.CalCurrentProposer(numValidators)
	if currentProposer != uint64(me.Index) {
		pbft.log.Warn("You are not the current proposer", "index", me.Index, "currentProposer", currentProposer)
		return
	}

	prepareBlock := &protocols.PrepareBlock{
		Epoch:         pbft.state.Epoch(),
		ViewNumber:    pbft.state.ViewNumber(),
		Block:         block,
		BlockIndex:    0,
		ProposalIndex: uint32(me.Index),
	}

	// Next index is equal zero, This view does not produce a block.
	//if pbft.state.NextViewBlockIndex() == 0 {
	//	//parentBlock, parentQC := pbft.blockTree.FindBlockAndQC(block.ParentHash(), block.NumberU64()-1)
	//	//if parentBlock == nil {
	//	//	pbft.log.Error("Can not find parent block", "number", block.Number(), "parentHash", block.ParentHash())
	//	//	return
	//	//}
	//	//prepareBlock.PrepareQC = parentQC
	//	prepareBlock.ViewChangeQC = pbft.state.LastViewChangeQC()
	//}

	prepareBlock.ViewChangeQC = pbft.state.LastViewChangeQC()

	if err := pbft.signMsgByBls(prepareBlock); err != nil {
		pbft.log.Error("Sign PrepareBlock failed", "err", err, "hash", block.Hash(), "number", block.NumberU64())
		return
	}

	//pbft.state.SetExecuting(prepareBlock.BlockIndex, true)
	if err := pbft.OnPrepareBlock("", prepareBlock); err != nil {
		pbft.log.Error("Check Seal Block failed", "err", err, "hash", block.Hash(), "number", block.NumberU64())
		pbft.state.SetExecuting(prepareBlock.BlockIndex, false)
		return
	}

	// write sendPrepareBlock info to wal
	if !pbft.isLoading() {
		pbft.bridge.SendPrepareBlock(prepareBlock)
	}
	pbft.network.Broadcast(prepareBlock)
	pbft.log.Info("Broadcast PrepareBlock", "prepareBlock", prepareBlock.String())

	//if err := pbft.signBlock(block.Hash(), block.NumberU64(), prepareBlock.BlockIndex); err != nil {
	//	pbft.log.Error("Sign PrepareBlock failed", "err", err, "hash", block.Hash(), "number", block.NumberU64())
	//	return
	//}

	pbft.txPool.Reset(block)

	//pbft.findQCBlock()

	pbft.validatorPool.Flush(prepareBlock.Block.Header())

	// Record the number of blocks.
	minedCounter.Inc(1)
	preBlock := pbft.blockTree.FindBlockByHash(block.ParentHash())
	if preBlock != nil {
		blockMinedGauage.Update(common.Millis(time.Now()) - int64(preBlock.Time()))
	}
	go func() {
		select {
		case <-stop:
			return
		case results <- block:
			blockProduceMeter.Mark(1)
		default:
			pbft.log.Warn("Sealing result channel is not ready by miner", "sealHash", block.Header().SealHash())
		}
	}()
}

// ReOnSeal is used to process the blocks that have already been generated.
func (pbft *Pbft) ReOnSeal(blockNumber uint64) {
	pbft.log.Info("ReOnSeal", "blockNumber", blockNumber,"ViewNumber",pbft.state.ViewNumber())
	block:=pbft.state.ViewBlockByIndex(blockNumber)
	me, err := pbft.validatorPool.GetValidatorByNodeID(pbft.state.Epoch(), pbft.NodeID())
	if err != nil {
		pbft.log.Warn("Can not got the validator, seal fail", "epoch", pbft.state.Epoch(), "nodeID", pbft.NodeID())
		return
	}
	numValidators := pbft.validatorPool.Len(pbft.state.Epoch())
	//currentProposer := pbft.state.BlockNumber() % uint64(numValidators)
	//currentProposer := (pbft.state.BlockNumber()-1)% uint64(numValidators)
	currentProposer := pbft.CalCurrentProposer(numValidators)
	if currentProposer != uint64(me.Index) {
		pbft.log.Warn("You are not the current proposer", "index", me.Index, "currentProposer", currentProposer)
		return
	}
	prepareBlock := &protocols.PrepareBlock{
		Epoch:         pbft.state.Epoch(),
		ViewNumber:    pbft.state.ViewNumber(),
		Block:         block,
		BlockIndex:    0,
		ProposalIndex: uint32(me.Index),
	}
	prepareBlock.ViewChangeQC = pbft.state.LastViewChangeQC()

	if err := pbft.signMsgByBls(prepareBlock); err != nil {
		pbft.log.Error("Sign PrepareBlock failed", "err", err, "hash", block.Hash(), "number", block.NumberU64())
		return
	}

	//pbft.state.SetExecuting(prepareBlock.BlockIndex, true)
	if err := pbft.OnPrepareBlock("", prepareBlock); err != nil {
		pbft.log.Error("Check Seal Block failed", "err", err, "hash", block.Hash(), "number", block.NumberU64())
		pbft.state.SetExecuting(prepareBlock.BlockIndex, false)
		return
	}

	// write sendPrepareBlock info to wal
	if !pbft.isLoading() {
		pbft.bridge.SendPrepareBlock(prepareBlock)
	}
	pbft.network.Broadcast(prepareBlock)
	pbft.log.Info("ReOnSeal broadcast PrepareBlock", "prepareBlock", prepareBlock.String())

	//if err := pbft.signBlock(block.Hash(), block.NumberU64(), prepareBlock.BlockIndex); err != nil {
	//	pbft.log.Error("Sign PrepareBlock failed", "err", err, "hash", block.Hash(), "number", block.NumberU64())
	//	return
	//}

	pbft.txPool.Reset(block)

	//pbft.findQCBlock()

	pbft.validatorPool.Flush(prepareBlock.Block.Header())

	// Record the number of blocks.
	preBlock := pbft.blockTree.FindBlockByHash(block.ParentHash())
	if preBlock != nil {
		blockMinedGauage.Update(common.Millis(time.Now()) - int64(preBlock.Time()))
	}
}

// SealHash returns the hash of a block prior to it being sealed.
func (pbft *Pbft) SealHash(header *types.Header) common.Hash {
	pbft.log.Debug("Seal hash", "hash", header.Hash(), "number", header.Number)
	return header.SealHash()
}

// APIs returns a list of APIs provided by the consensus engine.
func (pbft *Pbft) APIs(chain consensus.ChainReader) []rpc.API {
	return []rpc.API{
		{
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewDebugConsensusAPI(pbft),
			Public:    true,
		},
		{
			Namespace: "phoenixchain",
			Version:   "1.0",
			Service:   NewPublicPhoenixchainConsensusAPI(pbft),
			Public:    true,
		},
		{
			Namespace: "admin",
			Version:   "1.0",
			Service:   NewPublicAdminConsensusAPI(pbft),
			Public:    true,
		},
	}
}

// Protocols return consensus engine to provide protocol information.
func (pbft *Pbft) Protocols() []p2p.Protocol {
	return pbft.network.Protocols()
}

// NextBaseBlock is used to calculate the next block.
func (pbft *Pbft) NextBaseBlock() *types.Block {
	result := make(chan *types.Block, 1)
	pbft.asyncCallCh <- func() {
		block := pbft.state.HighestPreCommitQCBlock()
		pbft.log.Debug("Base block", "hash", block.Hash(), "number", block.Number())
		result <- block
	}
	return <-result
}

// InsertChain is used to insert the block into the chain.
func (pbft *Pbft) InsertChain(block *types.Block) error {
	if block.NumberU64() <= pbft.state.HighestCommitBlock().NumberU64() || pbft.HasBlock(block.Hash(), block.NumberU64()) {
		pbft.log.Debug("The inserted block has exists in chain",
			"number", block.Number(), "hash", block.Hash(),
			"CommittedNumber", pbft.state.HighestCommitBlock().Number(),
			"CommittedHash", pbft.state.HighestCommitBlock().Hash())
		return nil
	}

	t := time.Now()
	pbft.log.Info("Insert chain", "number", block.Number(), "hash", block.Hash(), "time", common.Beautiful(t))

	// Verifies block
	_, qc, err := ctypes.DecodeExtra(block.ExtraData())
	if err != nil {
		pbft.log.Error("Decode block extra date fail", "number", block.Number(), "hash", block.Hash())
		return errors.New("failed to decode block extra data")
	}

	parent := pbft.GetBlock(block.ParentHash(), block.NumberU64()-1)
	if parent == nil {
		pbft.log.Warn("Not found the inserted block's parent block",
			"number", block.Number(), "hash", block.Hash(),
			"parentHash", block.ParentHash(),
			"lockedNumber", pbft.state.HighestLockBlock().Number(),
			"lockedHash", pbft.state.HighestLockBlock().Hash(),
			"qcNumber", pbft.state.HighestPreCommitQCBlock().Number(),
			"qcHash", pbft.state.HighestPreCommitQCBlock().Hash())
		return errors.New("orphan block")
	}

	err = pbft.blockCacheWriter.Execute(block, parent)
	if err != nil {
		pbft.log.Error("Executing block failed", "number", block.Number(), "hash", block.Hash(), "parent", parent.Hash(), "parentHash", block.ParentHash(), "err", err)
		return errors.New("failed to executed block")
	}

	result := make(chan error, 1)
	pbft.asyncCallCh <- func() {
		if pbft.HasBlock(block.Hash(), block.NumberU64()) {
			pbft.log.Debug("The inserted block has exists in block tree", "number", block.Number(), "hash", block.Hash(), "time", common.Beautiful(t), "duration", time.Since(t))
			result <- nil
			return
		}

		if err := pbft.verifyPrepareQC(block.NumberU64(), block.Hash(), qc,0); err != nil {
			pbft.log.Error("Verify prepare QC fail", "number", block.Number(), "hash", block.Hash(), "err", err)
			result <- err
			return
		}
		result <- pbft.OnInsertQCBlock([]*types.Block{block}, []*ctypes.QuorumCert{qc})
	}

	timer := time.NewTimer(checkBlockSyncInterval)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			if pbft.HasBlock(block.Hash(), block.NumberU64()) {
				pbft.log.Debug("The inserted block has exists in block tree", "number", block.Number(), "hash", block.Hash(), "time", common.Beautiful(t))
				return nil
			}
			timer.Reset(checkBlockSyncInterval)
		case err := <-result:
			return err
		}
	}
}

// HasBlock check if the specified block exists in block tree.
func (pbft *Pbft) HasBlock(hash common.Hash, number uint64) bool {
	// Can only be invoked after startup
	qcBlock := pbft.state.HighestPreCommitQCBlock()
	return qcBlock.NumberU64() > number || (qcBlock.NumberU64() == number && qcBlock.Hash() == hash)
}

// Status returns the status data of the consensus engine.
func (pbft *Pbft) Status() []byte {
	status := make(chan []byte, 1)
	pbft.asyncCallCh <- func() {
		s := &Status{
			Tree:      pbft.blockTree,
			State:     pbft.state,
			Validator: pbft.IsConsensusNode(),
		}
		b, _ := json.Marshal(s)
		status <- b
	}
	return <-status
}

// GetPrepareQC returns the QC data of the specified block height.
func (pbft *Pbft) GetPrepareQC(number uint64) *ctypes.QuorumCert {
	pbft.log.Debug("get prepare QC")
	if header := pbft.blockChain.GetHeaderByNumber(number); header != nil {
		if block := pbft.blockChain.GetBlock(header.Hash(), number); block != nil {
			if _, qc, err := ctypes.DecodeExtra(block.ExtraData()); err == nil {
				return qc
			}
		}
	}
	return &ctypes.QuorumCert{}
}

// GetBlockByHash get the specified block by hash.
func (pbft *Pbft) GetBlockByHash(hash common.Hash) *types.Block {
	result := make(chan *types.Block, 1)
	pbft.asyncCallCh <- func() {
		block := pbft.blockTree.FindBlockByHash(hash)
		if block == nil {
			header := pbft.blockChain.GetHeaderByHash(hash)
			if header != nil {
				block = pbft.blockChain.GetBlock(header.Hash(), header.Number.Uint64())
			}
		}
		result <- block
	}
	return <-result
}

// GetBlockByHash get the specified block by hash and number.
func (pbft *Pbft) GetBlockByHashAndNum(hash common.Hash, number uint64) *types.Block {

	callBlock := func() *types.Block {
		// First extract from the confirmed block.
		block := pbft.blockTree.FindBlockByHash(hash)
		if block == nil {
			// Extract from view state.
			block = pbft.state.FindBlock(hash, number)
		}
		if block == nil {
			header := pbft.blockChain.GetHeaderByHash(hash)
			if header != nil {
				block = pbft.blockChain.GetBlock(header.Hash(), header.Number.Uint64())
			}
		}
		return block
	}
	if !pbft.isStart() {
		return callBlock()
	}
	result := make(chan *types.Block, 1)
	pbft.asyncCallCh <- func() {
		result <- callBlock()
	}
	return <-result
}

// CurrentBlock get the current lock block.
func (pbft *Pbft) CurrentBlock() *types.Block {
	var block *types.Block
	pbft.checkStart(func() {
		block = pbft.state.HighestLockBlock()
	})
	return block
}

func (pbft *Pbft) checkStart(exe func()) {
	pbft.log.Debug("Pbft status", "start", pbft.start)
	if utils.True(&pbft.start) {
		exe()
	}
}

// FastSyncCommitHead processes logic that performs fast synchronization.
func (pbft *Pbft) FastSyncCommitHead(block *types.Block) error {
	pbft.log.Info("Fast sync commit head", "number", block.Number(), "hash", block.Hash())

	result := make(chan error, 1)
	pbft.asyncCallCh <- func() {
		_, qc, err := ctypes.DecodeExtra(block.ExtraData())
		if err != nil {
			pbft.log.Warn("Decode block extra data fail", "number", block.Number(), "hash", block.Hash())
			result <- errors.New("failed to decode block extra data")
			return
		}

		pbft.validatorPool.Reset(block.NumberU64(), qc.Epoch)

		pbft.blockTree.Reset(block, qc)
		pbft.changeView(qc.Epoch, qc.ViewNumber, block, qc, nil)

		pbft.state.SetHighestQCBlock(block)
		pbft.state.SetHighestPreCommitQCBlock(block)
		pbft.state.SetHighestLockBlock(block)
		pbft.state.SetHighestCommitBlock(block)

		result <- err
	}
	return <-result
}

// Close turns off the consensus engine.
func (pbft *Pbft) Close() error {
	pbft.log.Info("Close pbft consensus")
	utils.SetFalse(&pbft.start)
	pbft.closeOnce.Do(func() {
		// Short circuit if the exit channel is not allocated.
		if pbft.exitCh == nil {
			return
		}
		close(pbft.exitCh)
	})
	if pbft.asyncExecutor != nil {
		pbft.asyncExecutor.Stop()
	}
	pbft.bridge.Close()
	return nil
}

// ConsensusNodes returns to the list of consensus nodes.
func (pbft *Pbft) ConsensusNodes() ([]discover.NodeID, error) {
	if pbft.consensusNodesMock != nil {
		return pbft.consensusNodesMock()
	}
	return pbft.validatorPool.ValidatorList(pbft.state.Epoch()), nil
}

// ShouldSeal check if we can seal block.
func (pbft *Pbft) ShouldSeal(curTime time.Time) (bool, error) {
	if pbft.isLoading() || !pbft.isStart() || !pbft.running() {
		//pbft.log.Trace("Should seal fail, pbft not running", "curTime", common.Beautiful(curTime))
		return false, nil
	}

	result := make(chan error, 2)
	pbft.asyncCallCh <- func() {
		pbft.OnShouldSeal(result)
	}
	select {
	case err := <-result:
		if err == nil {
			masterCounter.Inc(1)
		}
		//pbft.log.Trace("Should seal", "curTime", common.Beautiful(curTime), "err", err)
		return err == nil, err
	case <-time.After(50 * time.Millisecond):
		result <- ErrorTimeout
		//pbft.log.Trace("Should seal timeout", "curTime", common.Beautiful(curTime), "asyncCallCh", len(pbft.asyncCallCh))
		return false, ErrorEngineBusy
	}
}

func (pbft *Pbft) CalCurrentProposer(numValidators int) uint64 {
	var currentProposer uint64
	//if pbft.IsAddBlockTimeOver(){
	//	timeOverInterval:=time.Now().Unix()-pbft.state.LastAddBlockNumberTime()-cstate.AddBlockNumberTimeInterval
	//	modInt:=(timeOverInterval/cstate.AddBlockNumberTimeInterval)% int64(numValidators)
	//	currentProposer = (pbft.state.BlockNumber()+uint64(modInt)) % uint64(numValidators)
	//	pbft.log.Debug("CalCurrentProposer,IsAddBlockTimeOver", "old currentProposer",  (pbft.state.BlockNumber()-1) % uint64(numValidators), "new currentProposer", currentProposer)
	//}else {
	//	currentProposer = (pbft.state.BlockNumber()-1) % uint64(numValidators)
	//}
	currentProposer = (pbft.state.BlockNumber()-1+pbft.state.ViewNumber()/10) % uint64(numValidators)
	return currentProposer
}

func (pbft *Pbft) CalCurrentProposerWithBlockNumber(blockNumber uint64,numValidators int) uint64 {
	var currentProposer uint64
	//if pbft.IsAddBlockTimeOver(){
	//	currentProposer = blockNumber % uint64(numValidators)
	//	pbft.log.Debug("CalCurrentProposerWithBlockNumber,IsAddBlockTimeOver", "old currentProposer",  (blockNumber-1) % uint64(numValidators), "new currentProposer", currentProposer)
	//}else {
	//	currentProposer = (blockNumber-1) % uint64(numValidators)
	//}
	currentProposer = (blockNumber-1+pbft.state.ViewNumber()/10) % uint64(numValidators)
	return currentProposer
}

// OnShouldSeal determines whether the current condition
// of the block is satisfied.
func (pbft *Pbft) OnShouldSeal(result chan error) {
	select {
	case <-result:
		pbft.log.Trace("Should seal timeout")
		return
	default:
	}

	if !pbft.running() {
		result <- ErrorNotRunning
		return
	}

	if pbft.state.IsDeadline() {
		result <- fmt.Errorf("view timeout: %s", common.Beautiful(pbft.state.Deadline()))
		return
	}
	highestPreCommitQCBlockNumber := pbft.state.HighestPreCommitQCBlock().NumberU64()
	if !pbft.validatorPool.IsValidator(pbft.state.Epoch(), pbft.config.Option.NodeID) {
		result <- ErrorNotValidator
		return
	}

	numValidators := pbft.validatorPool.Len(pbft.state.Epoch())
	//currentProposer := (pbft.state.BlockNumber()-1) % uint64(numValidators)
	currentProposer:=pbft.CalCurrentProposer(numValidators)
	validator, err := pbft.validatorPool.GetValidatorByNodeID(pbft.state.Epoch(), pbft.config.Option.NodeID)
	if err != nil {
		pbft.log.Error("Should seal fail", "err", err)
		result <- err
		return
	}

	if currentProposer != uint64(validator.Index) {
		result <- errors.New("current node not the proposer")
		return
	}

	if pbft.state.NextViewBlockIndex() >= pbft.config.Sys.Amount {
		result <- errors.New("produce block over limit")
		return
	}

	qcBlock := pbft.state.HighestPreCommitQCBlock()
	_, qc := pbft.blockTree.FindBlockAndQC(qcBlock.Hash(), qcBlock.NumberU64())
	if pbft.validatorPool.ShouldSwitch(highestPreCommitQCBlockNumber) && qc != nil && qc.Epoch == pbft.state.Epoch() {
		pbft.log.Debug("New epoch, waiting for view's timeout", "highestPreCommitQCBlock", highestPreCommitQCBlockNumber, "index", validator.Index)
		result <- errors.New("new epoch,waiting for view's timeout")
		return
	}

	rtt := pbft.avgRTT()
	if pbft.state.Deadline().Sub(time.Now()) <= rtt {
		pbft.log.Debug("Not enough time to propagated block, stopped sealing", "deadline", pbft.state.Deadline(), "interval", pbft.state.Deadline().Sub(time.Now()), "rtt", rtt)
		result <- errors.New("not enough time to propagated block, stopped sealing")
		return
	}

	blockNumber:=pbft.state.BlockNumber()
	if pbft.state.ExistBlock(blockNumber) {
		prepareBlock:=pbft.state.PrepareBlockByIndex(blockNumber)
		proposer := pbft.currentProposer()
		if uint32(proposer.Index) == prepareBlock.NodeIndex() {
			pbft.ReOnSeal(blockNumber)
			result <- errors.New("do not repeat producing block")
			return
		}
	}

	proposerIndexGauage.Update(int64(currentProposer))
	validatorCountGauage.Update(int64(numValidators))
	result <- nil
}

// CalcBlockDeadline return the deadline of the block.
func (pbft *Pbft) CalcBlockDeadline(timePoint time.Time) time.Time {
	produceInterval := time.Duration(pbft.config.Sys.Period/uint64(pbft.config.Sys.Amount)) * time.Millisecond
	rtt := pbft.avgRTT()
	executeTime := (produceInterval - rtt) / 2
	pbft.log.Debug("Calc block deadline", "timePoint", timePoint, "stateDeadline", pbft.state.Deadline(), "produceInterval", produceInterval, "rtt", rtt, "executeTime", executeTime)
	if pbft.state.Deadline().Sub(timePoint) > produceInterval {
		return timePoint.Add(produceInterval - rtt - executeTime)
	}
	return pbft.state.Deadline()
}

func (pbft *Pbft) IsAddBlockTimeOver() bool {
	if time.Now().Unix()-pbft.state.LastAddBlockNumberTime()>cstate.AddBlockNumberTimeInterval{
		pbft.log.Debug("AddBlockNumber is time over", "LastAddBlockNumberTime", pbft.state.LastAddBlockNumberTime(), "now time", time.Now().Unix())
		return true
	}else {
		return false
	}
}

// CalcNextBlockTime returns the deadline  of the next block.
func (pbft *Pbft) CalcNextBlockTime(blockTime time.Time) time.Time {
	produceInterval := time.Duration(pbft.config.Sys.Period/uint64(pbft.config.Sys.Amount)) * time.Millisecond
	rtt := pbft.avgRTT()
	executeTime := (produceInterval - rtt) / 2
	pbft.log.Debug("Calc next block time",
		"blockTime", blockTime, "now", time.Now(), "produceInterval", produceInterval,
		"period", pbft.config.Sys.Period, "amount", pbft.config.Sys.Amount,
		"interval", time.Since(blockTime), "rtt", rtt, "executeTime", executeTime)
	if time.Since(blockTime) < produceInterval {
		return blockTime.Add(executeTime + rtt)
	}
	// Commit new block immediately.
	return blockTime.Add(produceInterval)
}

// IsConsensusNode returns whether the current node is a consensus node.
func (pbft *Pbft) IsConsensusNode() bool {
	return pbft.validatorPool.IsValidator(pbft.state.Epoch(), pbft.config.Option.NodeID)
}

// GetBlock returns the block corresponding to the specified number and hash.
func (pbft *Pbft) GetBlock(hash common.Hash, number uint64) *types.Block {
	result := make(chan *types.Block, 1)
	pbft.asyncCallCh <- func() {
		block, _ := pbft.blockTree.FindBlockAndQC(hash, number)
		result <- block
	}
	return <-result
}

// GetBlockWithoutLock returns the block corresponding to the specified number and hash.
func (pbft *Pbft) GetBlockWithoutLock(hash common.Hash, number uint64) *types.Block {
	block, _ := pbft.blockTree.FindBlockAndQC(hash, number)
	if block == nil {
		if eb := pbft.state.FindBlock(hash, number); eb != nil {
			block = eb
		} else {
			pbft.log.Debug("Get block failed", "hash", hash, "number", number)
		}
	}
	return block
}

// IsSignedBySelf returns the verification result , and the result is
// to determine whether the block information is the signature of the current node.
func (pbft *Pbft) IsSignedBySelf(sealHash common.Hash, header *types.Header) bool {
	return pbft.verifySelfSigned(sealHash.Bytes(), header.Signature())
}

// TracingSwitch will be abandoned in the future.
func (pbft *Pbft) TracingSwitch(flag int8) {
	panic("implement me")
}

// Config returns the configuration information of the consensus engine.
func (pbft *Pbft) Config() *ctypes.Config {
	return &pbft.config
}

// HighestCommitBlockBn returns the highest submitted block number of the current node.
func (pbft *Pbft) HighestCommitBlockBn() (uint64, common.Hash) {
	return pbft.state.HighestCommitBlock().NumberU64(), pbft.state.HighestCommitBlock().Hash()
}

// HighestLockBlockBn returns the highest locked block number of the current node.
func (pbft *Pbft) HighestLockBlockBn() (uint64, common.Hash) {
	return pbft.state.HighestLockBlock().NumberU64(), pbft.state.HighestLockBlock().Hash()
}

// HighestQCBlockBn return the highest QC block number of the current node.
func (pbft *Pbft) HighestQCBlockBn() (uint64, common.Hash) {
	return pbft.state.HighestPreCommitQCBlock().NumberU64(), pbft.state.HighestPreCommitQCBlock().Hash()
}

// HighestPreCommitQCBlockBn return the highest PreCommit QC block number of the current node.
func (pbft *Pbft) HighestPreCommitQCBlockBn() (uint64, common.Hash) {
	return pbft.state.HighestPreCommitQCBlock().NumberU64(), pbft.state.HighestPreCommitQCBlock().Hash()
}

func (pbft *Pbft) threshold(num int) int {
	return num - (num-1)/3
}

func (pbft *Pbft) commitBlock(commitBlock *types.Block, commitQC *ctypes.QuorumCert, lockBlock *types.Block, qcBlock *types.Block) {
	extra, err := ctypes.EncodeExtra(byte(pbftVersion), commitQC)
	if err != nil {
		pbft.log.Error("Encode extra error", "number", commitBlock.Number(), "hash", commitBlock.Hash(), "pbftVersion", pbftVersion)
		return
	}

	lockBlock, lockQC := pbft.blockTree.FindBlockAndQC(lockBlock.Hash(), lockBlock.NumberU64())
	qcBlock, qcQC := pbft.blockTree.FindBlockAndQC(qcBlock.Hash(), qcBlock.NumberU64())

	qcState := &protocols.State{Block: qcBlock, QuorumCert: qcQC}
	lockState := &protocols.State{Block: lockBlock, QuorumCert: lockQC}
	commitState := &protocols.State{Block: commitBlock, QuorumCert: commitQC}
	if pbft.updateChainStateDelayHook != nil {
		go pbft.updateChainStateDelayHook(qcState, lockState, commitState)
		return
	}
	if pbft.updateChainStateHook != nil {
		pbft.updateChainStateHook(qcState, lockState, commitState)
	}
	pbft.log.Info("CommitBlock, send consensus result to worker", "number", commitBlock.Number(), "hash", commitBlock.Hash())
	cpy := types.NewBlockWithHeader(commitBlock.Header()).WithBody(commitBlock.Transactions(), commitBlock.ExtraData())
	pbft.eventMux.Post(pbfttypes.PbftResult{
		Block:              cpy,
		ExtraData:          extra,
		SyncState:          pbft.commitErrCh,
		ChainStateUpdateCB: func() { pbft.bridge.UpdateChainState(qcState, lockState, commitState) },
	})
}

// Evidences implements functions in API.
func (pbft *Pbft) Evidences() string {
	evs := pbft.evPool.Evidences()
	if len(evs) == 0 {
		return "{}"
	}
	evds := evidence.ClassifyEvidence(evs)
	js, err := json.MarshalIndent(evds, "", "  ")
	if err != nil {
		return ""
	}
	return string(js)
}

func (pbft *Pbft) verifySelfSigned(m []byte, sig []byte) bool {
	recPubKey, err := crypto.Ecrecover(m, sig)
	if err != nil {
		return false
	}

	pubKey := pbft.config.Option.NodePriKey.PublicKey
	pbytes := elliptic.Marshal(pubKey.Curve, pubKey.X, pubKey.Y)
	return bytes.Equal(pbytes, recPubKey)
}

// signFn use private key to sign byte slice.
func (pbft *Pbft) signFn(m []byte) ([]byte, error) {
	return crypto.Sign(m, pbft.config.Option.NodePriKey)
}

// signFn use bls private key to sign byte slice.
func (pbft *Pbft) signFnByBls(m []byte) ([]byte, error) {
	sign := pbft.config.Option.BlsPriKey.Sign(string(m))
	return sign.Serialize(), nil
}

// signMsg use bls private key to sign msg.
func (pbft *Pbft) signMsgByBls(msg ctypes.ConsensusMsg) error {
	buf, err := msg.CannibalizeBytes()
	if err != nil {
		return err
	}
	sign, err := pbft.signFnByBls(buf)
	if err != nil {
		return err
	}
	msg.SetSign(sign)
	return nil
}

func (pbft *Pbft) isLoading() bool {
	return utils.True(&pbft.loading)
}

func (pbft *Pbft) isStart() bool {
	return utils.True(&pbft.start)
}

func (pbft *Pbft) isCurrentValidator() (*pbfttypes.ValidateNode, error) {
	return pbft.validatorPool.GetValidatorByNodeID(pbft.state.Epoch(), pbft.config.Option.NodeID)
}

func (pbft *Pbft) currentProposer() *pbfttypes.ValidateNode {
	length := pbft.validatorPool.Len(pbft.state.Epoch())
	//currentProposer := pbft.state.BlockNumber() % uint64(length)
	//currentProposer := (pbft.state.BlockNumber()-1) % uint64(length)
	currentProposer:=pbft.CalCurrentProposer(length)
	validator, _ := pbft.validatorPool.GetValidatorByIndex(pbft.state.Epoch(), uint32(currentProposer))
	return validator
}

func (pbft *Pbft) isProposer(epoch, blockNumber uint64, nodeIndex uint32) bool {
	if err := pbft.validatorPool.EnableVerifyEpoch(epoch); err != nil {
		return false
	}
	length := pbft.validatorPool.Len(epoch)
	var index uint64
	if blockNumber>0{
		//index = (blockNumber-1) % uint64(length)
		index = pbft.CalCurrentProposerWithBlockNumber(blockNumber,length)
	}else {
		index = 0
	}
	if validator, err := pbft.validatorPool.GetValidatorByIndex(epoch, uint32(index)); err == nil {
		return validator.Index == nodeIndex
	}
	return false
}

func (pbft *Pbft) currentValidatorLen() int {
	return pbft.validatorPool.Len(pbft.state.Epoch())
}

func (pbft *Pbft) verifyConsensusSign(msg ctypes.ConsensusMsg) error {
	if err := pbft.validatorPool.EnableVerifyEpoch(msg.EpochNum()); err != nil {
		return err
	}
	digest, err := msg.CannibalizeBytes()
	if err != nil {
		return errors.Wrap(err, "get msg's cannibalize bytes failed")
	}

	// Verify consensus msg signature
	if err := pbft.validatorPool.Verify(msg.EpochNum(), msg.NodeIndex(), digest, msg.Sign()); err != nil {
		return authFailedError{err: err}
	}
	return nil
}

func (pbft *Pbft) checkViewChangeQC(pb *protocols.PrepareBlock) error {
	//// check if the prepareBlock must take viewChangeQC
	//needViewChangeQC := func(pb *protocols.PrepareBlock) bool {
	//	_, localQC := pbft.blockTree.FindBlockAndQC(pb.Block.ParentHash(), pb.Block.NumberU64()-1)
	//	if localQC != nil && pbft.validatorPool.EqualSwitchPoint(localQC.BlockNumber) {
	//		return false
	//	}
	//	return localQC != nil && localQC.BlockIndex < pbft.config.Sys.Amount-1
	//}
	// check if the prepareBlock base on viewChangeQC maxBlock
	baseViewChangeQC := func(pb *protocols.PrepareBlock) bool {
		_, _, _, _, hash, number := pb.ViewChangeQC.MaxBlock()
		return pb.Block.NumberU64() == number+1 && pb.Block.ParentHash() == hash
	}

	//if needViewChangeQC(pb) && pb.ViewChangeQC == nil {
	//	return authFailedError{err: fmt.Errorf("prepareBlock need ViewChangeQC")}
	//}
	if pb.ViewChangeQC != nil {
		if !baseViewChangeQC(pb) {
			return authFailedError{err: fmt.Errorf("prepareBlock is not based on viewChangeQC maxBlock, viewchangeQC:%s, PrepareBlock:%s", pb.ViewChangeQC.String(), pb.String())}
		}
		if err := pbft.verifyViewChangeQC(pb.ViewChangeQC); err != nil {
			return err
		}
	}
	return nil
}

// check if the consensusMsg must take prepareQC or need not take prepareQC
func (pbft *Pbft) checkPrepareQC(msg ctypes.ConsensusMsg) error {
	switch cm := msg.(type) {
	case *protocols.PrepareBlock:
		if (cm.BlockNum() == 1 || cm.ViewNumber == 0) && cm.PrepareQC != nil {
			return authFailedError{err: fmt.Errorf("prepareBlock need not take PrepareQC, prepare:%s", cm.String())}
		}
		//if cm.BlockIndex == 0 && cm.BlockNum() != 1 && cm.PrepareQC == nil {
		//	return authFailedError{err: fmt.Errorf("prepareBlock need take PrepareQC, prepare:%s", cm.String())}
		//}
	case *protocols.PrepareVote:
		//if cm.BlockNum() == 1 && cm.ParentQC != nil {
		//	return authFailedError{err: fmt.Errorf("prepareVote need not take PrepareQC, vote:%s", cm.String())}
		//}
		if cm.BlockNum() > 1 && cm.ParentQC == nil {
			return authFailedError{err: fmt.Errorf("prepareVote need take PrepareQC, vote:%s", cm.String())}
		}
	case *protocols.PreCommit:
		//if cm.BlockNum() == 1 && cm.ParentQC != nil {
		//	return authFailedError{err: fmt.Errorf("preCommit need not take ParentQC, vote:%s", cm.String())}
		//}
		if cm.BlockNum() > 1 && cm.ParentQC == nil {
			return authFailedError{err: fmt.Errorf("preCommit need take ParentQC, vote:%s", cm.String())}
		}
	case *protocols.ViewChange:
		//if cm.BlockNumber == 0 && cm.PrepareQC != nil {
		//	return authFailedError{err: fmt.Errorf("viewChange need not take PrepareQC, viewChange:%s", cm.String())}
		//}
		if cm.BlockNumber > 1 && cm.PrepareQC == nil {
			return authFailedError{err: fmt.Errorf("viewChange need take PrepareQC, viewChange:%s", cm.String())}
		}
	default:
		return authFailedError{err: fmt.Errorf("invalid consensusMsg")}
	}
	return nil
}

func (pbft *Pbft) doubtDuplicate(msg ctypes.ConsensusMsg, node *pbfttypes.ValidateNode) error {
	switch cm := msg.(type) {
	case *protocols.PrepareBlock:
		if err := pbft.evPool.AddPrepareBlock(cm, node); err != nil {
			if _, ok := err.(*evidence.DuplicatePrepareBlockEvidence); ok {
				pbft.log.Warn("Receive DuplicatePrepareBlockEvidence msg", "err", err.Error())
				return err
			}
		}
	case *protocols.PrepareVote:
		if err := pbft.evPool.AddPrepareVote(cm, node); err != nil {
			if _, ok := err.(*evidence.DuplicatePrepareVoteEvidence); ok {
				pbft.log.Warn("Receive DuplicatePrepareVoteEvidence msg", "err", err.Error())
				return err
			}
		}
	case *protocols.PreCommit:
		if err := pbft.evPool.AddPreCommit(cm, node); err != nil {
			if _, ok := err.(*evidence.DuplicatePrecommitEvidence); ok {
				pbft.log.Warn("Receive DuplicatePrecommitEvidence msg", "err", err.Error())
				return err
			}
		}
	case *protocols.ViewChange:
		if err := pbft.evPool.AddViewChange(cm, node); err != nil {
			if _, ok := err.(*evidence.DuplicateViewChangeEvidence); ok {
				pbft.log.Warn("Receive DuplicateViewChangeEvidence msg", "err", err.Error())
				return err
			}
		}
	default:
		return authFailedError{err: fmt.Errorf("invalid consensusMsg")}
	}
	return nil
}

func (pbft *Pbft) verifyConsensusMsg(msg ctypes.ConsensusMsg) (*pbfttypes.ValidateNode, error) {
	// check if the consensusMsg must take prepareQC, Otherwise maybe panic
	if err := pbft.checkPrepareQC(msg); err != nil {
		return nil, err
	}
	// Verify consensus msg signature
	if err := pbft.verifyConsensusSign(msg); err != nil {
		return nil, err
	}

	// Get validator of signer
	vnode, err := pbft.validatorPool.GetValidatorByIndex(msg.EpochNum(), msg.NodeIndex())

	if err != nil {
		return nil, authFailedError{err: errors.Wrap(err, "get validator failed")}
	}

	// check that the consensusMsg is duplicate
	if err = pbft.doubtDuplicate(msg, vnode); err != nil {
		return nil, err
	}

	var (
		prepareQC *ctypes.QuorumCert
		oriNumber uint64
		oriHash   common.Hash
	)

	switch cm := msg.(type) {
	case *protocols.PrepareBlock:
		proposer := pbft.currentProposer()
		if uint32(proposer.Index) != msg.NodeIndex() {
			return nil, fmt.Errorf("current proposer index:%d, prepare block author index:%d", proposer.Index, msg.NodeIndex())
		}
		// BlockNum equal 1, the parent's block is genesis, doesn't has prepareQC
		// BlockIndex is not equal 0, this is not first block of current proposer
		if cm.BlockNum() == 1 || cm.BlockIndex != 0 {
			return vnode, nil
		}
		if err := pbft.checkViewChangeQC(cm); err != nil {
			return nil, err
		}
		//prepareQC = cm.PrepareQC
		//_, localQC := pbft.blockTree.FindBlockAndQC(cm.Block.ParentHash(), cm.Block.NumberU64()-1)
		//if localQC == nil {
		//	return nil, fmt.Errorf("parentBlock and parentQC not exists,number:%d,hash:%s", cm.Block.NumberU64()-1, cm.Block.ParentHash())
		//}
		//oriNumber = localQC.BlockNumber
		//oriHash = localQC.BlockHash

	case *protocols.PrepareVote:
		if cm.BlockNum() == 1 {
			return vnode, nil
		}
		prepareQC = cm.ParentQC

		_, localQC := pbft.blockTree.FindBlockAndQC(prepareQC.BlockHash, cm.BlockNumber-1)
		if localQC == nil {
			return nil, fmt.Errorf("parentBlock and parentQC not exists,number:%d,hash:%s", cm.BlockNumber-1, prepareQC.BlockHash.String())
		}
		oriNumber = localQC.BlockNumber
		oriHash = localQC.BlockHash


		if err := pbft.verifyPrepareQC(oriNumber, oriHash, prepareQC,ctypes.RoundStepPrepareVote); err != nil {
			return nil, err
		}

	case *protocols.PreCommit:
		if cm.BlockNum() == 1 {
			return vnode, nil
		}
		prepareQC = cm.ParentQC

		_, localQC := pbft.blockTree.FindBlockAndQC(prepareQC.BlockHash, cm.BlockNumber-1)
		if localQC == nil {
			return nil, fmt.Errorf("parentBlock and parentQC not exists,number:%d,hash:%s", cm.BlockNumber-1, prepareQC.BlockHash.String())
		}
		oriNumber = localQC.BlockNumber
		oriHash = localQC.BlockHash


		if err := pbft.verifyPrepareQC(oriNumber, oriHash, prepareQC,ctypes.RoundStepPreCommit); err != nil {
			return nil, err
		}

	case *protocols.ViewChange:
		// Genesis block doesn't has prepareQC
		if cm.BlockNumber == 0 {
			return vnode, nil
		}
		prepareQC = cm.PrepareQC
		oriNumber = cm.BlockNumber
		oriHash = cm.BlockHash

		if err := pbft.verifyPrepareQC(oriNumber, oriHash, prepareQC,0); err != nil {
			return nil, err
		}
	}

	return vnode, nil
}

func (pbft *Pbft) Pause() {
	pbft.log.Info("Pause pbft consensus")
	utils.SetTrue(&pbft.syncing)
}

func (pbft *Pbft) Resume() {
	pbft.log.Info("Resume pbft consensus")
	utils.SetFalse(&pbft.syncing)
}

func (pbft *Pbft) Syncing() bool {
	return utils.True(&pbft.syncing)
}

func (pbft *Pbft) generatePrepareQC(votes map[uint32]*protocols.PrepareVote) *ctypes.QuorumCert {
	if len(votes) == 0 {
		return nil
	}

	var vote *protocols.PrepareVote

	for _, v := range votes {
		if vote==nil{
			vote = v
		}else {
			if v.ViewNumber>vote.ViewNumber{
				vote = v
			}
		}
	}

	// Validator set prepareQC is the same as highestQC
	total := pbft.validatorPool.Len(pbft.state.Epoch())

	vSet := utils.NewBitArray(uint32(total))
	vSet.SetIndex(vote.NodeIndex(), true)

	var aggSig bls.Sign
	if err := aggSig.Deserialize(vote.Sign()); err != nil {
		return nil
	}

	qc := &ctypes.QuorumCert{
		Epoch:        vote.Epoch,
		ViewNumber:   vote.ViewNumber,
		BlockHash:    vote.BlockHash,
		BlockNumber:  vote.BlockNumber,
		BlockIndex:   vote.BlockIndex,
		Step:         ctypes.RoundStepPrepareVote,
		ValidatorSet: utils.NewBitArray(vSet.Size()),
	}
	for _, p := range votes {
		//Check whether two votes are equal
		if !vote.EqualState(p) {
			pbft.log.Error(fmt.Sprintf("QuorumCert isn't same  vote1:%s vote2:%s", vote.String(), p.String()))
			return nil
		}
		if p.NodeIndex() != vote.NodeIndex() {
			var sig bls.Sign
			err := sig.Deserialize(p.Sign())
			if err != nil {
				return nil
			}

			aggSig.Add(&sig)
			vSet.SetIndex(p.NodeIndex(), true)
		}
	}
	qc.Signature.SetBytes(aggSig.Serialize())
	qc.ValidatorSet.Update(vSet)
	log.Debug("Generate prepare qc", "hash", vote.BlockHash, "number", vote.BlockNumber, "qc", qc.String())
	return qc
}

func (pbft *Pbft) generatePreCommitQC(votes map[uint32]*protocols.PreCommit) *ctypes.QuorumCert {
	if len(votes) == 0 {
		return nil
	}

	var vote *protocols.PreCommit

	for _, v := range votes {
		if vote==nil{
			vote = v
		}else {
			if v.ViewNumber>vote.ViewNumber{
				vote = v
			}
		}
	}

	// Validator set prepareQC is the same as highestQC
	total := pbft.validatorPool.Len(pbft.state.Epoch())

	vSet := utils.NewBitArray(uint32(total))
	vSet.SetIndex(vote.NodeIndex(), true)

	var aggSig bls.Sign
	if err := aggSig.Deserialize(vote.Sign()); err != nil {
		return nil
	}

	qc := &ctypes.QuorumCert{
		Epoch:        vote.Epoch,
		ViewNumber:   vote.ViewNumber,
		BlockHash:    vote.BlockHash,
		BlockNumber:  vote.BlockNumber,
		BlockIndex:   vote.BlockIndex,
		Step:         ctypes.RoundStepPreCommit,
		ValidatorSet: utils.NewBitArray(vSet.Size()),
	}
	for _, p := range votes {
		//Check whether two votes are equal
		if !vote.EqualState(p) {
			pbft.log.Error(fmt.Sprintf("QuorumCert isn't same  vote1:%s vote2:%s", vote.String(), p.String()))
			return nil
		}
		if p.NodeIndex() != vote.NodeIndex() {
			var sig bls.Sign
			err := sig.Deserialize(p.Sign())
			if err != nil {
				return nil
			}

			aggSig.Add(&sig)
			vSet.SetIndex(p.NodeIndex(), true)
		}
	}
	qc.Signature.SetBytes(aggSig.Serialize())
	qc.ValidatorSet.Update(vSet)
	log.Debug("Generate preCommit qc", "hash", vote.BlockHash, "number", vote.BlockNumber, "qc", qc.String())
	return qc
}

func (pbft *Pbft) generateViewChangeQC(viewChanges map[uint32]*protocols.ViewChange) *ctypes.ViewChangeQC {
	type ViewChangeQC struct {
		cert   *ctypes.ViewChangeQuorumCert
		aggSig *bls.Sign
		ba     *utils.BitArray
	}

	total := uint32(pbft.validatorPool.Len(pbft.state.Epoch()))

	qcs := make(map[common.Hash]*ViewChangeQC)

	for _, v := range viewChanges {
		var aggSig bls.Sign
		if err := aggSig.Deserialize(v.Sign()); err != nil {
			return nil
		}

		if vc, ok := qcs[v.BlockHash]; !ok {
			blockEpoch, blockView := uint64(0), uint64(0)
			if v.PrepareQC != nil {
				blockEpoch, blockView = v.PrepareQC.Epoch, v.PrepareQC.ViewNumber
			}
			qc := &ViewChangeQC{
				cert: &ctypes.ViewChangeQuorumCert{
					Epoch:           v.Epoch,
					ViewNumber:      v.ViewNumber,
					BlockHash:       v.BlockHash,
					BlockNumber:     v.BlockNumber,
					BlockEpoch:      blockEpoch,
					BlockViewNumber: blockView,
					ValidatorSet:    utils.NewBitArray(total),
				},
				aggSig: &aggSig,
				ba:     utils.NewBitArray(total),
			}
			qc.ba.SetIndex(v.NodeIndex(), true)
			qcs[v.BlockHash] = qc
		} else {
			vc.aggSig.Add(&aggSig)
			vc.ba.SetIndex(v.NodeIndex(), true)
		}
	}

	qc := &ctypes.ViewChangeQC{QCs: make([]*ctypes.ViewChangeQuorumCert, 0)}
	for _, q := range qcs {
		q.cert.Signature.SetBytes(q.aggSig.Serialize())
		q.cert.ValidatorSet.Update(q.ba)
		qc.QCs = append(qc.QCs, q.cert)
	}
	log.Debug("Generate view change qc", "qc", qc.String())
	return qc
}

func (pbft *Pbft) verifyPrepareQC(oriNum uint64, oriHash common.Hash, qc *ctypes.QuorumCert,step ctypes.RoundStepType) error {
	if qc == nil {
		return fmt.Errorf("verify prepare qc failed,qc is nil")
	}
	if err := pbft.validatorPool.EnableVerifyEpoch(qc.Epoch); err != nil {
		return err
	}
	// check signature number
	threshold := pbft.threshold(pbft.validatorPool.Len(qc.Epoch))

	signsTotal := qc.Len()
	if signsTotal < threshold {
		return authFailedError{err: fmt.Errorf("block qc has small number of signature total:%d, threshold:%d", signsTotal, threshold)}
	}
	// check if the corresponding block QC
	if oriNum != qc.BlockNumber || oriHash != qc.BlockHash {
		return fmt.Errorf("verify prepare qc failed,not the corresponding qc,oriNum:%d,oriHash:%s,qcNum:%d,qcHash:%s",
				oriNum, oriHash.String(), qc.BlockNumber, qc.BlockHash.String())}

	var cb []byte
	var err error
	if cb, err = qc.CannibalizeBytes(); err != nil {
		return err
	}
	if err = pbft.validatorPool.VerifyAggSigByBA(qc.Epoch, qc.ValidatorSet, cb, qc.Signature.Bytes()); err != nil {
		pbft.log.Error("Verify failed", "qc", qc.String(), "validators", pbft.validatorPool.Validators(qc.Epoch).String())
		return authFailedError{err: fmt.Errorf("verify prepare qc failed: %v", err)}
	}
	return nil
}

func (pbft *Pbft) validateViewChangeQC(viewChangeQC *ctypes.ViewChangeQC) error {

	vcEpoch, _, _, _, _, _ := viewChangeQC.MaxBlock()

	maxLimit := pbft.validatorPool.Len(vcEpoch)
	if len(viewChangeQC.QCs) > maxLimit {
		return fmt.Errorf("viewchangeQC exceed validator max limit, total:%d, threshold:%d", len(viewChangeQC.QCs), maxLimit)
	}
	// the threshold of validator on current epoch
	threshold := pbft.threshold(maxLimit)
	// check signature number
	signsTotal := viewChangeQC.Len()
	if signsTotal < threshold {
		return fmt.Errorf("viewchange has small number of signature total:%d, threshold:%d", signsTotal, threshold)
	}

	var err error
	epoch := uint64(0)
	viewNumber := uint64(0)
	existHash := make(map[common.Hash]interface{})
	for i, vc := range viewChangeQC.QCs {
		// Check if it is the same view
		if i == 0 {
			epoch = vc.Epoch
			viewNumber = vc.ViewNumber
		} else if epoch != vc.Epoch || viewNumber != vc.ViewNumber {
			err = fmt.Errorf("has multiple view messages")
			break
		}
		// check if it has the same blockHash
		if _, ok := existHash[vc.BlockHash]; ok {
			err = fmt.Errorf("has duplicated blockHash")
			break
		}
		existHash[vc.BlockHash] = struct{}{}
	}
	return err
}

func (pbft *Pbft) verifyViewChangeQC(viewChangeQC *ctypes.ViewChangeQC) error {
	vcEpoch, _, _, _, _, _ := viewChangeQC.MaxBlock()

	if err := pbft.validatorPool.EnableVerifyEpoch(vcEpoch); err != nil {
		return err
	}

	// check parameter validity
	if err := pbft.validateViewChangeQC(viewChangeQC); err != nil {
		return err
	}

	var err error
	for _, vc := range viewChangeQC.QCs {
		var cb []byte
		if cb, err = vc.CannibalizeBytes(); err != nil {
			err = fmt.Errorf("get cannibalize bytes failed")
			break
		}

		if err = pbft.validatorPool.VerifyAggSigByBA(vc.Epoch, vc.ValidatorSet, cb, vc.Signature.Bytes()); err != nil {
			pbft.log.Debug("verify failed", "qc", vc.String(), "validators", pbft.validatorPool.Validators(vc.Epoch).String())

			err = authFailedError{err: fmt.Errorf("verify viewchange qc failed:number:%d,validators:%s,msg:%s,signature:%s,err:%v",
				vc.BlockNumber, vc.ValidatorSet.String(), hexutil.Encode(cb), vc.Signature.String(), err)}
			break
		}
	}

	return err
}

// NodeID returns the ID value of the current node
func (pbft *Pbft) NodeID() discover.NodeID {
	return pbft.config.Option.NodeID
}

func (pbft *Pbft) avgRTT() time.Duration {
	produceInterval := time.Duration(pbft.config.Sys.Period/uint64(pbft.config.Sys.Amount)) * time.Millisecond
	rtt := pbft.AvgLatency() * 2
	if rtt == 0 || rtt >= produceInterval {
		rtt = pbft.DefaultAvgLatency() * 2
	}
	return rtt
}

func (pbft *Pbft) GetSchnorrNIZKProve() (*bls.SchnorrProof, error) {
	return pbft.config.Option.BlsPriKey.MakeSchnorrNIZKP()
}

func (pbft *Pbft) DecodeExtra(extra []byte) (common.Hash, uint64, error) {
	_, qc, err := ctypes.DecodeExtra(extra)
	if err != nil {
		return common.Hash{}, 0, err
	}
	return qc.BlockHash, qc.BlockNumber, nil
}
