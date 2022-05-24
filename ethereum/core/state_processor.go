package core

import (
	"bytes"
	"strconv"

	"Phoenix-Chain-Core/libs/common"
	"Phoenix-Chain-Core/consensus"
	"Phoenix-Chain-Core/ethereum/core/db/snapshotdb"
	"Phoenix-Chain-Core/ethereum/core/state"
	"Phoenix-Chain-Core/ethereum/core/types"
	"Phoenix-Chain-Core/ethereum/core/vm"
	"Phoenix-Chain-Core/libs/crypto"
	"Phoenix-Chain-Core/libs/log"
	"Phoenix-Chain-Core/configs"
	"Phoenix-Chain-Core/libs/rlp"
)

// StateProcessor is a basic Processor, which takes care of transitioning
// state from one point to another.
//
// StateProcessor implements Processor.
type StateProcessor struct {
	config *configs.ChainConfig // Chain configuration options
	bc     *BlockChain          // Canonical block chain
	engine consensus.Engine     // Consensus engine used for block rewards
}

// NewStateProcessor initialises a new StateProcessor.
func NewStateProcessor(config *configs.ChainConfig, bc *BlockChain, engine consensus.Engine) *StateProcessor {
	return &StateProcessor{
		config: config,
		bc:     bc,
		engine: engine,
	}
}

// Process processes the state changes according to the Ethereum rules by running
// the transaction messages using the statedb and applying any rewards to
// the processor (coinbase).
//
// Process returns the receipts and logs accumulated during the process and
// returns the amount of gas that was used in the process. If any of the
// transactions failed to execute due to insufficient gas it will return an error.
func (p *StateProcessor) Process(block *types.Block, statedb *state.StateDB, cfg vm.Config) (types.Receipts, []*types.Log, uint64, error) {
	var (
		receipts types.Receipts
		usedGas  = new(uint64)
		header   = block.Header()
		allLogs  []*types.Log
		gp       = new(GasPool).AddGas(block.GasLimit())
	)

	if bcr != nil {
		// BeginBlocker()
		if err := bcr.BeginBlocker(header, statedb); nil != err {
			log.Error("Failed to call BeginBlocker on StateProcessor", "blockNumber", block.Number(),
				"blockHash", block.Hash(), "err", err)
			return nil, nil, 0, err
		}
	}

	// Iterate over and process the individual transactions
	for i, tx := range block.Transactions() {
		statedb.Prepare(tx.Hash(), block.Hash(), i)
		//preUsedGas := uint64(0)

		receipt, _, err := ApplyTransaction(p.config, p.bc, gp, statedb, header, tx, usedGas, cfg)
		if err != nil {
			log.Error("Failed to execute tx on StateProcessor", "blockNumber", block.Number(),
				"blockHash", block.Hash().TerminalString(), "txHash", tx.Hash().String(), "err", err)
			return nil, nil, 0, err
		}
		//log.Debug("tx process success", "txHash", tx.Hash().Hex(), "txTo", tx.To().Hex(), "dataLength", len(tx.Data()), "toCodeSize", statedb.GetCodeSize(*tx.To()), "txUsedGas", *usedGas-preUsedGas)
		receipts = append(receipts, receipt)
		allLogs = append(allLogs, receipt.Logs...)
	}

	if bcr != nil {
		// EndBlocker()
		if err := bcr.EndBlocker(header, statedb); nil != err {
			log.Error("Failed to call EndBlocker on StateProcessor", "blockNumber", block.Number(),
				"blockHash", block.Hash().TerminalString(), "err", err)
			return nil, nil, 0, err
		}
	}

	// Finalize the block, applying any consensus engine specific extras (e.g. block rewards)
	p.engine.Finalize(p.bc, header, statedb, block.Transactions(), receipts)

	return receipts, allLogs, *usedGas, nil
}

// ApplyTransaction attempts to apply a transaction to the given state database
// and uses the input parameters for its environment. It returns the receipt
// for the transaction, gas used and an error if the transaction failed,
// indicating the block was invalid.
func ApplyTransaction(config *configs.ChainConfig, bc ChainContext, gp *GasPool,
	statedb *state.StateDB, header *types.Header, tx *types.Transaction,
	usedGas *uint64, cfg vm.Config) (*types.Receipt, uint64, error) {

	msg, err := tx.AsMessage(types.NewEIP155Signer(config.ChainID))

	if err != nil {
		return nil, 0, err
	}
	// Create a new context to be used in the EVM environment
	context := NewEVMContext(msg, header, bc)
	// Create a new environment which holds all relevant information
	// about the transaction and calling mechanisms.
	vmenv := vm.NewEVM(context, snapshotdb.Instance(), statedb, config, cfg)

	log.Trace("execute tx start", "blockNumber", header.Number, "txHash", tx.Hash().String())

	// Apply the transaction to the current state (included in the env)
	result, err := ApplyMessage(vmenv, msg, gp)
	if err != nil {
		return nil, 0, err
	}
	// Update the state with pending changes
	statedb.Finalise(true)

	var root []byte
	*usedGas += result.UsedGas

	// Create a new receipt for the transaction, storing the intermediate root and gas used by the tx
	// based on the eip phase, we're passing whether the root touch-delete accounts.
	receipt := types.NewReceipt(root, result.Failed(), *usedGas)
	receipt.TxHash = tx.Hash()
	receipt.GasUsed = result.UsedGas
	// if the transaction created a contract, store the creation address in the receipt.
	if msg.To() == nil {
		receipt.ContractAddress = crypto.CreateAddress(vmenv.Context.Origin, tx.Nonce())
	}
	// Set the receipt logs
	if result.Failed() {
		if bizError, ok := result.Err.(*common.BizError); ok {
			buf := new(bytes.Buffer)
			res := strconv.Itoa(int(bizError.Code))
			if err := rlp.Encode(buf, [][]byte{[]byte(res)}); nil != err {
				log.Error("Cannot RlpEncode the log data", "data", bizError.Code, "err", err)
				return nil, 0, err
			}
			receipt.Logs = []*types.Log{
				&types.Log{
					Address:     *msg.To(),
					Topics:      nil,
					Data:        buf.Bytes(),
					BlockNumber: header.Number.Uint64(),
				},
			}
		} else {
			receipt.Logs = statedb.GetLogs(tx.Hash())
		}
	} else {
		receipt.Logs = statedb.GetLogs(tx.Hash())
	}
	//create a bloom for filtering
	receipt.Bloom = types.CreateBloom(types.Receipts{receipt})
	receipt.BlockHash = statedb.BlockHash()
	receipt.BlockNumber = header.Number
	receipt.TransactionIndex = uint(statedb.TxIndex())
	return receipt, result.UsedGas, err
}
