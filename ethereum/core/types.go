package core

import (
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common"
	"time"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core/state"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core/types"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core/vm"
)

// Validator is an interface which defines the standard for block validation. It
// is only responsible for validating block contents, as the header validation is
// done by the specific consensus engines.
//
type Validator interface {
	// ValidateBody validates the given block's content.
	ValidateBody(block *types.Block) error

	// ValidateState validates the given statedb and optionally the receipts and
	// gas used.
	ValidateState(block *types.Block, state *state.StateDB, receipts types.Receipts, usedGas uint64) error
}

// Processor is an interface for processing blocks using a given initial state.
//
// Process takes the block to be processed and the statedb upon which the
// initial state is based. It should return the receipts generated, amount
// of gas used in the process and return an error if any of the internal rules
// failed.
type Processor interface {
	Process(block *types.Block, statedb *state.StateDB, cfg vm.Config) (types.Receipts, []*types.Log, uint64, error)
}

type Committer interface {
	CommitTransactions(header *types.Header, txs *types.TransactionsByPriceAndNonce, interrupt *int32, timestamp int64, blockDeadline time.Time, tempContractCache map[common.Address]struct{}) (bool, bool)
}
