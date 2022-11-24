package vm

import (
	"math/big"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core/types"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common"
)

// StateDB is an EVM database for full state querying.
type StateDB interface {
	CreateAccount(common.Address)

	SubBalance(common.Address, *big.Int)
	AddBalance(common.Address, *big.Int)
	GetBalance(common.Address) *big.Int

	GetNonce(common.Address) uint64
	SetNonce(common.Address, uint64)

	GetCodeHash(common.Address) common.Hash
	GetCode(common.Address) []byte
	SetCode(common.Address, []byte)
	GetCodeSize(common.Address) int

	AddRefund(uint64)
	SubRefund(uint64)
	GetRefund() uint64

	// todo: hash -> bytes
	GetCommittedState(common.Address, []byte) []byte
	//GetState(common.Address, common.Hash) common.Hash
	//SetState(common.Address, common.Hash, common.Hash)
	GetState(common.Address, []byte) []byte
	SetState(common.Address, []byte, []byte)

	Suicide(common.Address) bool
	HasSuicided(common.Address) bool

	// Exist reports whether the given account exists in state.
	// Notably this should also return true for suicided accounts.
	Exist(common.Address) bool
	// Empty returns whether the given account is empty. Empty
	// is defined according to EIP161 (balance = nonce = code = 0).
	Empty(common.Address) bool

	RevertToSnapshot(int)
	Snapshot() int

	AddLog(*types.Log)
	GetLogs(hash common.Hash) []*types.Log
	AddPreimage(common.Hash, []byte)

	ForEachStorage(common.Address, func([]byte, []byte) bool)
	MigrateStorage(from, to common.Address)
	//dpos add
	TxHash() common.Hash
	TxIdx() uint32

	IntermediateRoot(deleteEmptyObjects bool) common.Hash

	GetCurrentActiveVersion() uint32
}

// MerkleProof
type ProofList [][]byte

func (n *ProofList) Put(key []byte, value []byte) error {
	*n = append(*n, value)
	return nil
}

// CallContext provides a basic interface for the EVM calling conventions. The EVM
// depends on this context being implemented for doing subcalls and initialising new EVM contracts.
type CallContext interface {
	// Call another contract
	Call(env *EVM, me ContractRef, addr common.Address, data []byte, gas, value *big.Int) ([]byte, error)
	// Take another's contract code and execute within our own context
	CallCode(env *EVM, me ContractRef, addr common.Address, data []byte, gas, value *big.Int) ([]byte, error)
	// Same as CallCode except sender and value is propagated from parent to child scope
	DelegateCall(env *EVM, me ContractRef, addr common.Address, data []byte, gas *big.Int) ([]byte, error)
	// Create a new contract
	Create(env *EVM, me ContractRef, data []byte, gas, value *big.Int) ([]byte, common.Address, error)
}
