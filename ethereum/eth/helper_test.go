// This file contains some shares testing functionality, common to  multiple
// different files and modules being tested.

package eth

import (
	"Phoenix-Chain-Core/ethereum/eth/downloader"
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"
	"sort"
	"sync"
	"testing"

	"Phoenix-Chain-Core/ethereum/core/db/rawdb"

	"Phoenix-Chain-Core/consensus"
	"Phoenix-Chain-Core/pos/xcom"

	"Phoenix-Chain-Core/configs"
	"Phoenix-Chain-Core/ethereum/core"
	"Phoenix-Chain-Core/ethereum/core/types"
	"Phoenix-Chain-Core/ethereum/p2p"
	"Phoenix-Chain-Core/ethereum/p2p/discover"
	"Phoenix-Chain-Core/libs/common"
	"Phoenix-Chain-Core/libs/crypto"
	"Phoenix-Chain-Core/libs/ethdb"
	"Phoenix-Chain-Core/libs/event"
	_ "Phoenix-Chain-Core/pos/xcom"
)

var (
	testBankKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	testBank       = crypto.PubkeyToAddress(testBankKey.PublicKey)
)

// newTestProtocolManager creates a new protocol manager for testing purposes,
// with the given number of blocks already known, and potential notification
// channels for different events.
func newTestProtocolManager(mode downloader.SyncMode, blocks int, generator func(int, *core.BlockGen), newtx chan<- []*types.Transaction) (*ProtocolManager, ethdb.Database, error) {
	xcom.GetEc(xcom.DefaultTestNet)
	var (
		evmux = new(event.TypeMux)
		//	engine = pbft.New(params.GrapeChainConfig.Pbft, evmux, nil)
		engine = consensus.NewFaker()
		db     = rawdb.NewMemoryDatabase()
		gspec  = &core.Genesis{
			Config: configs.TestChainConfig,
			Alloc:  core.GenesisAlloc{testBank: {Balance: big.NewInt(1000000)}},
		}
		genesis = gspec.MustCommit(db)

		//blockchain, _ = core.NewBlockChain(db, nil, gspec.Config, engine, vm.Config{}, nil)
	)
	engine.InsertChain(genesis)
	//cache := core.NewBlockChainCache(blockchain)
	//
	//engine.SetBlockChainCache(cache)
	//
	//txpool := core.NewTxPool(core.DefaultTxPoolConfig, gspec.Config, core.NewTxPoolBlockChain(cache))
	//
	//engine.Start(blockchain, txpool, pbft.NewStaticAgency([]discover.Node{}))
	//chain, _ := core.GenerateChain(gspec.Config, genesis, engine, db, blocks, generator)
	chain := core.GenerateBlockChain2(gspec.Config, genesis, engine, db, blocks, generator)
	//if _, err := blockchain.InsertChain(chain); err != nil {
	//	panic(err)
	//}
	pm, err := NewProtocolManager(gspec.Config, mode, DefaultConfig.NetworkId, evmux, &testTxPool{added: newtx}, engine, chain, db, 1)
	if err != nil {
		return nil, nil, err
	}
	pm.Start(1000)
	return pm, db, nil
}

// newTestProtocolManagerMust creates a new protocol manager for testing purposes,
// with the given number of blocks already known, and potential notification
// channels for different events. In case of an error, the constructor force-
// fails the test.
func newTestProtocolManagerMust(t *testing.T, mode downloader.SyncMode, blocks int, generator func(int, *core.BlockGen), newtx chan<- []*types.Transaction) (*ProtocolManager, ethdb.Database) {
	pm, db, err := newTestProtocolManager(mode, blocks, generator, newtx)
	if err != nil {
		t.Fatalf("Failed to create protocol manager: %v", err)
	}
	return pm, db
}

// testTxPool is a fake, helper transaction pool for testing purposes
type testTxPool struct {
	txFeed event.Feed
	pool   []*types.Transaction        // Collection of all transactions
	added  chan<- []*types.Transaction // Notification channel for new transactions

	lock sync.RWMutex // Protects the transaction pool
}

// Has returns an indicator whether txpool has a transaction
// cached with the given hash.
func (p *testTxPool) Has(hash common.Hash) bool {
	p.lock.Lock()
	defer p.lock.Unlock()

	for _, tx := range p.pool {
		if tx.Hash() == hash {
			return true
		}
	}
	return false
}

// Get retrieves the transaction from local txpool with given
// tx hash.
func (p *testTxPool) Get(hash common.Hash) *types.Transaction {
	p.lock.Lock()
	defer p.lock.Unlock()

	for _, tx := range p.pool {
		if tx.Hash() == hash {
			return tx
		}
	}
	return nil
}

// AddRemotes appends a batch of transactions to the pool, and notifies any
// listeners if the addition channel is non nil
func (p *testTxPool) AddRemotes(txs []*types.Transaction) []error {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.pool = append(p.pool, txs...)
	if p.added != nil {
		p.added <- txs
	}
	return make([]error, len(txs))
}

// Pending returns all the transactions known to the pool
func (p *testTxPool) Pending() (map[common.Address]types.Transactions, error) {
	p.lock.RLock()
	defer p.lock.RUnlock()

	batches := make(map[common.Address]types.Transactions)
	for _, tx := range p.pool {
		from, _ := types.Sender(types.NewEIP155Signer(new(big.Int)), tx)
		batches[from] = append(batches[from], tx)
	}
	for _, batch := range batches {
		sort.Sort(types.TxByNonce(batch))
	}
	return batches, nil
}

func (p *testTxPool) SubscribeNewTxsEvent(ch chan<- core.NewTxsEvent) event.Subscription {
	return p.txFeed.Subscribe(ch)
}

// newTestTransaction create a new dummy transaction.
func newTestTransaction(from *ecdsa.PrivateKey, nonce uint64, datasize int) *types.Transaction {
	tx := types.NewTransaction(nonce, common.Address{}, big.NewInt(0), 100000, big.NewInt(0), make([]byte, datasize))
	tx, _ = types.SignTx(tx, types.NewEIP155Signer(new(big.Int)), from)
	return tx
}

// testPeer is a simulated peer to allow testing direct network calls.
type testPeer struct {
	net p2p.MsgReadWriter // Network layer reader/writer to simulate remote messaging
	app *p2p.MsgPipeRW    // Application layer reader/writer to simulate the local side
	*peer
}

// newTestPeer creates a new peer registered at the given protocol manager.
func newTestPeer(name string, version int, pm *ProtocolManager, shake bool) (*testPeer, <-chan error) {
	// Create a message pipe to communicate through
	app, net := p2p.MsgPipe()

	// Start the peer on a new thread
	var id discover.NodeID
	rand.Read(id[:])
	peer := pm.newPeer(version, p2p.NewPeer(id, name, nil), net)

	// Start the peer on a new thread
	errc := make(chan error, 1)
	go func() { errc <- pm.runPeer(peer) }()
	tp := &testPeer{app: app, net: net, peer: peer}
	// Execute any implicitly requested handshakes and return
	if shake {
		var (
			genesis = pm.blockchain.Genesis()
			head    = pm.blockchain.CurrentHeader()
		)
		tp.handshake(nil, head.Number, head.Hash(), genesis.Hash())
	}
	return tp, errc
}

// handshake simulates a trivial handshake that expects the same state from the
// remote side as we are simulating locally.
func (p *testPeer) handshake(t *testing.T, bn *big.Int, head common.Hash, genesis common.Hash) {
	msg := &statusData{
		ProtocolVersion: uint32(p.version),
		NetworkId:       DefaultConfig.NetworkId,
		CurrentBlock:    head,
		GenesisBlock:    genesis,
		BN:              bn,
	}
	if err := p2p.ExpectMsg(p.app, StatusMsg, msg); err != nil {
		t.Fatalf("status recv: %v", err)
	}
	if err := p2p.Send(p.app, StatusMsg, msg); err != nil {
		t.Fatalf("status send: %v", err)
	}
}

// close terminates the local side of the peer, notifying the remote protocol
// manager of termination.
func (p *testPeer) close() {
	p.app.Close()
}
