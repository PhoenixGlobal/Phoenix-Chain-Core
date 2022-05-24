package vm

import (
	"fmt"
	"math/big"

	"Phoenix-Chain-Core/ethereum/p2p/discover"

	"Phoenix-Chain-Core/libs/common/consensus"

	"Phoenix-Chain-Core/libs/common/vm"

	"Phoenix-Chain-Core/libs/common/hexutil"
	"Phoenix-Chain-Core/libs/log"
	"Phoenix-Chain-Core/configs"

	"Phoenix-Chain-Core/libs/common"
	"Phoenix-Chain-Core/pos/plugin"
)

const (
	TxReportDuplicateSign = 3000
	CheckDuplicateSign    = 3001
)

type SlashingContract struct {
	Plugin   *plugin.SlashingPlugin
	Contract *Contract
	Evm      *EVM
}

func (sc *SlashingContract) RequiredGas(input []byte) uint64 {
	if checkInputEmpty(input) {
		return 0
	}
	return configs.SlashingGas
}

func (sc *SlashingContract) Run(input []byte) ([]byte, error) {
	if checkInputEmpty(input) {
		return nil, nil
	}
	return execPhoenixchainContract(input, sc.FnSigns())
}

func (sc *SlashingContract) FnSigns() map[uint16]interface{} {
	return map[uint16]interface{}{
		// Set
		TxReportDuplicateSign: sc.reportDuplicateSign,
		// Get
		CheckDuplicateSign: sc.checkDuplicateSign,
	}
}

func (sc *SlashingContract) CheckGasPrice(gasPrice *big.Int, fcode uint16) error {
	return nil
}

// Report the double signing behavior of the node
func (sc *SlashingContract) reportDuplicateSign(dupType uint8, data string) ([]byte, error) {

	txHash := sc.Evm.StateDB.TxHash()
	blockNumber := sc.Evm.BlockNumber
	blockHash := sc.Evm.BlockHash
	from := sc.Contract.CallerAddress

	if !sc.Contract.UseGas(configs.ReportDuplicateSignGas) {
		return nil, ErrOutOfGas
	}

	if !sc.Contract.UseGas(configs.DuplicateEvidencesGas) {
		return nil, ErrOutOfGas
	}

	log.Debug("Call reportDuplicateSign", "blockNumber", blockNumber, "blockHash", blockHash.Hex(),
		"TxHash", txHash.Hex(), "from", from.String())
	evidence, err := sc.Plugin.DecodeEvidence(consensus.EvidenceType(dupType), data)
	if nil != err {
		return txResultHandler(vm.SlashingContractAddr, sc.Evm, "reportDuplicateSign",
			common.InvalidParameter.Wrap(err.Error()).Error(),
			TxReportDuplicateSign, common.InvalidParameter)
	}

	if txHash == common.ZeroHash {
		return nil, nil
	}

	if err := sc.Plugin.Slash(evidence, blockHash, blockNumber.Uint64(), sc.Evm.StateDB, from); nil != err {
		if bizErr, ok := err.(*common.BizError); ok {
			return txResultHandler(vm.SlashingContractAddr, sc.Evm, "reportDuplicateSign",
				bizErr.Error(), TxReportDuplicateSign, bizErr)
		} else {
			return nil, err
		}
	}
	return txResultHandler(vm.SlashingContractAddr, sc.Evm, "",
		"", TxReportDuplicateSign, common.NoErr)
}

// Check if the node has double sign behavior at a certain block height
func (sc *SlashingContract) checkDuplicateSign(dupType uint8, nodeId discover.NodeID, blockNumber uint64) ([]byte, error) {
	log.Info("checkDuplicateSign exist", "blockNumber", blockNumber, "nodeId", nodeId.TerminalString(), "dupType", dupType)
	txHash, err := sc.Plugin.CheckDuplicateSign(nodeId, blockNumber, consensus.EvidenceType(dupType), sc.Evm.StateDB)
	var data string

	if nil != err {
		return callResultHandler(sc.Evm, fmt.Sprintf("checkDuplicateSign, duplicateSignBlockNum: %d, nodeId: %s, dupType: %d",
			blockNumber, nodeId, dupType), data, common.InternalError.Wrap(err.Error())), nil
	}
	if len(txHash) > 0 {
		data = hexutil.Encode(txHash)
	}
	return callResultHandler(sc.Evm, fmt.Sprintf("checkDuplicateSign, duplicateSignBlockNum: %d, nodeId: %s, dupType: %d",
		blockNumber, nodeId, dupType), data, nil), nil
}
