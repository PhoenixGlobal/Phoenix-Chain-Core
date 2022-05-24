// Package miner implements Ethereum block creation and mining.
package miner

import (
	"Phoenix-Chain-Core/libs/common/hexutil"
	"Phoenix-Chain-Core/ethereum/eth/downloader"
	"math/big"
	"sync/atomic"
	"time"

	"Phoenix-Chain-Core/consensus"
	"Phoenix-Chain-Core/ethereum/core"
	"Phoenix-Chain-Core/ethereum/core/state"
	"Phoenix-Chain-Core/ethereum/core/types"
	"Phoenix-Chain-Core/libs/event"
	"Phoenix-Chain-Core/libs/log"
	"Phoenix-Chain-Core/configs"
)

// Backend wraps all methods required for mining.
type Backend interface {
	BlockChain() *core.BlockChain
	TxPool() *core.TxPool
}

// Config is the configuration parameters of mining.
type Config struct {
	ExtraData hexutil.Bytes  `toml:",omitempty"` // Block extra data set by the miner
	GasFloor  uint64         // Target gas floor for mined blocks.
	GasPrice  *big.Int       // Minimum gas price for mining a transaction
	Recommit  time.Duration  // The time interval for miner to re-create mining work.
	Noverify  bool           // Disable remote mining solution verification(only useful in ethash).
}

// Miner creates blocks and searches for proof-of-work values.
type Miner struct {
	mux    *event.TypeMux
	worker *worker
	eth    Backend
	engine consensus.Engine
	exitCh chan struct{}

	canStart    int32 // can start indicates whether we can start the mining operation
	shouldStart int32 // should start indicates whether we should start after sync
}

func New(eth Backend,  config *Config, chainConfig *configs.ChainConfig, miningConfig *core.MiningConfig, mux *event.TypeMux,
	engine consensus.Engine, isLocalBlock func(block *types.Block) bool,
	blockChainCache *core.BlockChainCache, vmTimeout uint64) *Miner {
	miner := &Miner{
		eth:      eth,
		mux:      mux,
		engine:   engine,
		exitCh:   make(chan struct{}),
		worker:   newWorker(config, chainConfig, miningConfig, engine, eth, mux, isLocalBlock, blockChainCache, vmTimeout),
		canStart: 1,
	}
	go miner.update()

	return miner
}

// update keeps track of the downloader events. Please be aware that this is a one shot type of update loop.
// It's entered once and as soon as `Done` or `Failed` has been broadcasted the events are unregistered and
// the loop is exited. This to prevent a major security vuln where external parties can DOS you with blocks
// and halt your mining operation for as long as the DOS continues.
func (self *Miner) update() {
	events := self.mux.Subscribe(downloader.StartEvent{}, downloader.DoneEvent{}, downloader.FailedEvent{})
	defer events.Unsubscribe()

	for {
		select {
		case ev := <-events.Chan():
			if ev == nil {
				return
			}
			switch ev.Data.(type) {
			case downloader.StartEvent:
				atomic.StoreInt32(&self.canStart, 0)
				if self.Mining() {
					self.Stop()
					atomic.StoreInt32(&self.shouldStart, 1)
					log.Info("Mining aborted due to sync")
				}
			case downloader.DoneEvent, downloader.FailedEvent:
				shouldStart := atomic.LoadInt32(&self.shouldStart) == 1

				atomic.StoreInt32(&self.canStart, 1)
				atomic.StoreInt32(&self.shouldStart, 0)
				if shouldStart {
					self.Start()
				}

				// stop immediately and ignore all further pending events
				return
			}
		case <-self.exitCh:
			return
		}
	}
}

func (self *Miner) Start() {
	atomic.StoreInt32(&self.shouldStart, 1)

	if atomic.LoadInt32(&self.canStart) == 0 {
		log.Info("Network syncing, will start miner afterwards")
		return
	}

	self.worker.start()
	if bft, ok := self.engine.(consensus.Bft); ok {
		bft.Resume()
	}
}

func (self *Miner) Stop() {
	self.worker.stop()
	if bft, ok := self.engine.(consensus.Bft); ok {
		bft.Pause()
	}

	atomic.StoreInt32(&self.shouldStart, 0)
}

func (self *Miner) Close() {
	self.worker.close()
	close(self.exitCh)
}

func (self *Miner) Mining() bool {
	return self.worker.isRunning()
}

//func (self *Miner) SetExtra(extra []byte) error {
//	if uint64(len(extra)) > params.MaximumExtraDataSize {
//		return fmt.Errorf("Extra exceeds max length. %d > %v", len(extra), params.MaximumExtraDataSize)
//	}
//	self.worker.setExtra(extra)
//	return nil
//}

// SetRecommitInterval sets the interval for sealing work resubmitting.
func (self *Miner) SetRecommitInterval(interval time.Duration) {
	self.worker.setRecommitInterval(interval)
}

// Pending returns the currently pending block and associated state.
func (self *Miner) Pending() (*types.Block, *state.StateDB) {
	return self.worker.pending()
}

// PendingBlock returns the currently pending block.
//
// Note, to access both the pending block and the pending state
// simultaneously, please use Pending(), as the pending state can
// change between multiple method calls
func (self *Miner) PendingBlock() *types.Block {
	return self.worker.pendingBlock()
}
