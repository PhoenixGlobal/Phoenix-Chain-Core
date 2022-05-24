package core

import (
	"fmt"

	"Phoenix-Chain-Core/libs/common"
	"Phoenix-Chain-Core/pos/gov"

	"Phoenix-Chain-Core/libs/log"

	"Phoenix-Chain-Core/configs"
	"Phoenix-Chain-Core/consensus"
	"Phoenix-Chain-Core/ethereum/core/state"
	"Phoenix-Chain-Core/ethereum/core/types"
)

// BlockValidator is responsible for validating block headers and
// processed state.
//
// BlockValidator implements Validator.
type BlockValidator struct {
	config *configs.ChainConfig // Chain configuration options
	bc     *BlockChain          // Canonical block chain
	engine consensus.Engine     // Consensus engine used for validating
}

// NewBlockValidator returns a new block validator which is safe for re-use
func NewBlockValidator(config *configs.ChainConfig, blockchain *BlockChain, engine consensus.Engine) *BlockValidator {
	validator := &BlockValidator{
		config: config,
		engine: engine,
		bc:     blockchain,
	}
	return validator
}

// ValidateBody verifies the block
// header's transaction. The headers are assumed to be already
// validated at this point.
func (v *BlockValidator) ValidateBody(block *types.Block) error {
	// Check whether the block's known, and if not, that it's linkable
	if v.bc.HasBlockAndState(block.Hash(), block.NumberU64()) {
		return ErrKnownBlock
	}
	//if !v.bc.HasBlockAndState(block.ParentHash(), block.NumberU64()-1) {
	//	if !v.bc.HasBlock(block.ParentHash(), block.NumberU64()-1) {
	//		return consensus.ErrUnknownAncestor
	//	}
	//	return consensus.ErrPrunedAncestor
	//}
	// Header validity is known at this point, check the transactions
	header := block.Header()
	if hash := types.DeriveSha(block.Transactions()); hash != header.TxHash {
		return fmt.Errorf("transaction root hash mismatch: have %x, want %x", hash, header.TxHash)
	}
	return nil
}

// ValidateState validates the various changes that happen after a state
// transition, such as amount of used gas, the receipt roots and the state root
// itself. ValidateState returns a database batch if the validation was a success
// otherwise nil and an error is returned.
func (v *BlockValidator) ValidateState(block *types.Block, statedb *state.StateDB, receipts types.Receipts, usedGas uint64) error {
	header := block.Header()
	if block.GasUsed() != usedGas {
		return fmt.Errorf("invalid gas used (remote: %d local: %d)", block.GasUsed(), usedGas)
	}
	// Validate the received block's bloom with the one derived from the generated receipts.
	// For valid blocks this should always validate to true.
	rbloom := types.CreateBloom(receipts)
	if rbloom != header.Bloom {
		return fmt.Errorf("invalid bloom (remote: %x  local: %x)", header.Bloom, rbloom)
	}
	// Tre receipt Trie's root (R = (Tr [[H1, R1], ... [Hn, R1]]))
	receiptSha := types.DeriveSha(receipts)
	if receiptSha != header.ReceiptHash {
		return fmt.Errorf("invalid receipt root hash (remote: %x local: %x)", header.ReceiptHash, receiptSha)
	}
	// Validate the state root against the received state root and throw
	// an error if they don't match.
	if root := statedb.IntermediateRoot(true); header.Root != root {
		return fmt.Errorf("invalid merkle root (remote: %x local: %x)", header.Root, root)
	}
	return nil
}

// CalcGasLimit computes the gas limit of the next block after parent. It aims
// to keep the baseline gas above the provided floor, and increase it towards the
// ceil if the blocks are full. If the ceil is exceeded, it will always decrease
// the gas allowance.
func CalcGasLimit(parent *types.Block, gasFloor /*, gasCeil*/ uint64) uint64 {

	var gasCeil uint64

	govGasCeil, err := gov.GovernMaxBlockGasLimit(parent.Number().Uint64()+1, common.ZeroHash)
	if nil != err {
		log.Error("cannot find GasLimit from govern", "err", err)
		gasCeil = uint64(configs.DefaultMinerGasCeil)
	} else {
		gasCeil = uint64(govGasCeil)
	}

	if gasFloor > gasCeil {
		gasFloor = gasCeil
	}

	contrib := (parent.GasUsed() + parent.GasUsed()/2) / configs.GasLimitBoundDivisor

	decay := parent.GasLimit()/configs.GasLimitBoundDivisor - 1

	/*
		strategy: gasLimit of block-to-mine is set based on parent's
		gasUsed value.  if parentGasUsed > parentGasLimit * (2/3) then we
		increase it, otherwise lower it (or leave it unchanged if it's right
		at that usage) the amount increased/decreased depends on how far away
		from parentGasLimit * (2/3) parentGasUsed is.
	*/
	limit := parent.GasLimit() - decay + contrib

	// If we're outside our allowed gas range, we try to hone towards them
	if limit < gasFloor {
		limit = parent.GasLimit() + decay
		if limit > gasFloor {
			limit = gasFloor
		}
	} else if limit > gasCeil {
		limit = parent.GasLimit() - decay
		if limit < gasCeil {
			limit = gasCeil
		}
	}
	log.Info("Call CalcGasLimit", "blockNumber", parent.Number().Uint64()+1, "gasFloor", gasFloor, "gasCeil", gasCeil, "parentLimit", parent.GasLimit(), "limit", limit)
	return limit
}
