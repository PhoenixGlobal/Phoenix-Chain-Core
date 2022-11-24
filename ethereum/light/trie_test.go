package light

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/consensus"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core/db/rawdb"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core/state"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core/vm"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/configs"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/trie"
	"github.com/davecgh/go-spew/spew"
)

func TestNodeIterator(t *testing.T) {
	var (
		fulldb  = rawdb.NewMemoryDatabase()
		lightdb = rawdb.NewMemoryDatabase()
		gspec   = core.Genesis{Alloc: core.GenesisAlloc{testBankAddress: {Balance: testBankFunds}}}
		genesis = gspec.MustCommit(fulldb)
	)
	gspec.MustCommit(lightdb)
	blockchain, _ := core.NewBlockChain(fulldb, nil, configs.TestChainConfig, consensus.NewFaker(), vm.Config{}, nil)
	gchain, _ := core.GenerateChain(configs.TestChainConfig, genesis, consensus.NewFaker(), fulldb, 4, testChainGen)
	if _, err := blockchain.InsertChain(gchain); err != nil {
		panic(err)
	}

	ctx := context.Background()
	odr := &testOdr{sdb: fulldb, ldb: lightdb, indexerConfig: TestClientIndexerConfig}
	head := blockchain.CurrentHeader()
	lightTrie, _ := NewStateDatabase(ctx, head, odr).OpenTrie(head.Root)
	fullTrie, _ := state.NewDatabase(fulldb).OpenTrie(head.Root)
	if err := diffTries(fullTrie, lightTrie); err != nil {
		t.Fatal(err)
	}
}

func diffTries(t1, t2 state.Trie) error {
	i1 := trie.NewIterator(t1.NodeIterator(nil))
	i2 := trie.NewIterator(t2.NodeIterator(nil))
	for i1.Next() && i2.Next() {
		if !bytes.Equal(i1.Key, i2.Key) {
			spew.Dump(i2)
			return fmt.Errorf("tries have different keys %x, %x", i1.Key, i2.Key)
		}
		if !bytes.Equal(i1.Value, i2.Value) {
			return fmt.Errorf("tries differ at key %x", i1.Key)
		}
	}
	switch {
	case i1.Err != nil:
		return fmt.Errorf("full trie iterator error: %v", i1.Err)
	case i2.Err != nil:
		return fmt.Errorf("light trie iterator error: %v", i1.Err)
	case i1.Next():
		return fmt.Errorf("full trie iterator has more k/v pairs")
	case i2.Next():
		return fmt.Errorf("light trie iterator has more k/v pairs")
	}
	return nil
}
