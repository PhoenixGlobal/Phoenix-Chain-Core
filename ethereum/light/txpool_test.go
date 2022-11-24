package light

import (
	"context"
	"math"
	"math/big"
	"testing"
	"time"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/consensus"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core/types"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core/db/rawdb"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core/vm"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/configs"
)

type testTxRelay struct {
	send, discard, mined chan int
}

func (self *testTxRelay) Send(txs types.Transactions) {
	self.send <- len(txs)
}

func (self *testTxRelay) NewHead(head common.Hash, mined []common.Hash, rollback []common.Hash) {
	m := len(mined)
	if m != 0 {
		self.mined <- m
	}
}

func (self *testTxRelay) Discard(hashes []common.Hash) {
	self.discard <- len(hashes)
}

const poolTestTxs = 1000
const poolTestBlocks = 100

// test tx 0..n-1
var testTx [poolTestTxs]*types.Transaction

// txs sent before block i
func sentTx(i int) int {
	return int(math.Pow(float64(i)/float64(poolTestBlocks), 0.9) * poolTestTxs)
}

// txs included in block i or before that (minedTx(i) <= sentTx(i))
func minedTx(i int) int {
	return int(math.Pow(float64(i)/float64(poolTestBlocks), 1.1) * poolTestTxs)
}

func txPoolTestChainGen(i int, block *core.BlockGen) {
	s := minedTx(i)
	e := minedTx(i + 1)
	for i := s; i < e; i++ {
		block.AddTx(testTx[i])
	}
}

func TestTxPool(t *testing.T) {
	for i := range testTx {
		testTx[i], _ = types.SignTx(types.NewTransaction(uint64(i), acc1Addr, big.NewInt(10000), configs.TxGas, nil, nil), types.NewEIP155Signer(configs.TestChainConfig.ChainID), testBankKey)
	}

	var (
		sdb     = rawdb.NewMemoryDatabase()
		ldb     = rawdb.NewMemoryDatabase()
		gspec   = core.Genesis{Alloc: core.GenesisAlloc{testBankAddress: {Balance: testBankFunds}}}
		genesis = gspec.MustCommit(sdb)
	)
	gspec.MustCommit(ldb)
	// Assemble the test environment
	blockchain, _ := core.NewBlockChain(sdb, nil, configs.TestChainConfig, consensus.NewFaker(), vm.Config{}, nil)
	gchain, _ := core.GenerateChain(configs.TestChainConfig, genesis, consensus.NewFaker(), sdb, poolTestBlocks, txPoolTestChainGen)
	if _, err := blockchain.InsertChain(gchain); err != nil {
		panic(err)
	}

	odr := &testOdr{sdb: sdb, ldb: ldb, indexerConfig: TestClientIndexerConfig}
	relay := &testTxRelay{
		send:    make(chan int, 1),
		discard: make(chan int, 1),
		mined:   make(chan int, 1),
	}
	lightchain, _ := NewLightChain(odr, configs.TestChainConfig, consensus.NewFaker())
	txPermanent = 50
	pool := NewTxPool(configs.TestChainConfig, lightchain, relay)
	//TODO: cftf.faker do not supported
	relay.mined <- 1

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	for ii, block := range gchain {
		i := ii + 1
		s := sentTx(i - 1)
		e := sentTx(i)
		for i := s; i < e; i++ {
			pool.Add(ctx, testTx[i])
			got := <-relay.send
			exp := 1
			if got != exp {
				t.Errorf("relay.Send expected len = %d, got %d", exp, got)
			}
		}

		if _, err := lightchain.InsertHeaderChain([]*types.Header{block.Header()}, 1); err != nil {
			panic(err)
		}
	}
}
