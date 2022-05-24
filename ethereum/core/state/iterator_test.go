package state

import (
	"bytes"
	"Phoenix-Chain-Core/libs/common/vm"
	"testing"

	"Phoenix-Chain-Core/libs/common"
	"Phoenix-Chain-Core/libs/ethdb"
)

var TestPhoenixChainPrecompiledContracts = map[common.Address]interface{}{
	vm.ValidatorInnerContractAddr: nil,
	// add by economic model
	vm.StakingContractAddr:     nil,
	vm.RestrictingContractAddr: nil,
	vm.SlashingContractAddr:    nil,
	vm.GovContractAddr:         nil,
	vm.RewardManagerPoolAddr:   nil,
	vm.DelegateRewardPoolAddr:  nil,
}

type TestPrecompiledContractCheck struct{}

func (pcc *TestPrecompiledContractCheck) IsPhoenixChainPrecompiledContract(address common.Address) bool {
	if _, ok := TestPhoenixChainPrecompiledContracts[address]; ok {
		return true
	}
	return false
}

// Tests that the node iterator indeed walks over the entire database contents.
func TestNodeIteratorCoverage(t *testing.T) {
	vm.PrecompiledContractCheckInstance = &TestPrecompiledContractCheck{}
	// Create some arbitrary test state to iterate
	db, root, _, valueKeys := makeTestState()

	state, err := New(root, db)
	if err != nil {
		t.Fatalf("failed to create state trie at %x: %v", root, err)
	}
	// Gather all the node hashes found by the iterator
	hashes := make(map[common.Hash]struct{})
	for it := NewNodeIterator(state); it.Next(); {
		if it.Hash != (common.Hash{}) {
			hashes[it.Hash] = struct{}{}
		}
	}
	// Cross check the iterated hashes and the database/nodepool content
	for hash := range hashes {
		if _, err := db.TrieDB().Node(hash); err != nil {
			t.Errorf("failed to retrieve reported node %x", hash)
		}
	}
	for _, hash := range db.TrieDB().Nodes() {
		if _, ok := hashes[hash]; !ok {
			if _, ok := valueKeys[string(hash.Bytes())]; !ok {
				t.Errorf("state entry not reported %x", hash)
			}
		}
	}
	it := db.TrieDB().DiskDB().(ethdb.Database).NewIterator()
	for it.Next() {
		key := it.Key()
		if bytes.HasPrefix(key, []byte("secure-key-")) {
			continue
		}
		if _, ok := hashes[common.BytesToHash(key)]; !ok {
			t.Errorf("state entry not reported %x", key)
		}
	}
	it.Release()
}
