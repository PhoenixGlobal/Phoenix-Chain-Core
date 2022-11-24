package vm

import (
	"fmt"
	"math/big"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/configs"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common/vm"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/log"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/pos/plugin"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/pos/restricting"
)

const (
	TxCreateRestrictingPlan = 4000
	QueryRestrictingInfo    = 4100
)

type RestrictingContract struct {
	Plugin   *plugin.RestrictingPlugin
	Contract *Contract
	Evm      *EVM
}

func (rc *RestrictingContract) RequiredGas(input []byte) uint64 {
	if checkInputEmpty(input) {
		return 0
	}
	return configs.RestrictingPlanGas
}

func (rc *RestrictingContract) Run(input []byte) ([]byte, error) {
	if checkInputEmpty(input) {
		return nil, nil
	}
	return execPhoenixchainContract(input, rc.FnSigns())
}

func (rc *RestrictingContract) FnSigns() map[uint16]interface{} {
	return map[uint16]interface{}{
		// Set
		TxCreateRestrictingPlan: rc.createRestrictingPlan,

		// Get
		QueryRestrictingInfo: rc.getRestrictingInfo,
	}
}

func (rc *RestrictingContract) CheckGasPrice(gasPrice *big.Int, fcode uint16) error {
	return nil
}

// createRestrictingPlan is a PhoenixChain precompiled contract function, used for create a restricting plan
func (rc *RestrictingContract) createRestrictingPlan(account common.Address, plans []restricting.RestrictingPlan) ([]byte, error) {

	//sender := rc.Contract.Caller()
	from := rc.Contract.CallerAddress
	txHash := rc.Evm.StateDB.TxHash()
	blockNum := rc.Evm.BlockNumber
	blockHash := rc.Evm.BlockHash
	state := rc.Evm.StateDB

	log.Debug("Call createRestrictingPlan of RestrictingContract", "blockNumber", blockNum.Uint64(),
		"blockHash", blockHash.TerminalString(), "txHash", txHash.Hex(), "from", from.String(), "account", account.String())

	if !rc.Contract.UseGas(configs.CreateRestrictingPlanGas) {
		return nil, ErrOutOfGas
	}
	if !rc.Contract.UseGas(configs.ReleasePlanGas * uint64(len(plans))) {
		return nil, ErrOutOfGas
	}

	err := rc.Plugin.AddRestrictingRecord(from, account, blockNum.Uint64(), blockHash, plans, state, txHash)
	switch err.(type) {
	case nil:
		return txResultHandler(vm.RestrictingContractAddr, rc.Evm, "",
			"", TxCreateRestrictingPlan, common.NoErr)
	case *common.BizError:
		bizErr := err.(*common.BizError)
		return txResultHandler(vm.RestrictingContractAddr, rc.Evm, "createRestrictingPlan",
			bizErr.Error(), TxCreateRestrictingPlan, bizErr)
	default:
		log.Error("Failed to cal addRestrictingRecord on createRestrictingPlan", "blockNumber", blockNum.Uint64(),
			"blockHash", blockHash.TerminalString(), "txHash", txHash.Hex(), "error", err)
		return nil, err
	}
}

// createRestrictingPlan is a PhoenixChain precompiled contract function, used for getting restricting info.
// first output param is a slice of byte of restricting info;
// the secend output param is the result what plugin executed GetRestrictingInfo returns.
func (rc *RestrictingContract) getRestrictingInfo(account common.Address) ([]byte, error) {
	state := rc.Evm.StateDB

	result, err := rc.Plugin.GetRestrictingInfo(account, state)
	return callResultHandler(rc.Evm, fmt.Sprintf("getRestrictingInfo, account: %s", account.String()),
		result, err), nil
}
