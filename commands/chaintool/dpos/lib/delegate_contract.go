package lib

import "math/big"

type DelegateContract struct {
	executor *FunctionExecutor
}

func NewDelegateContract(dposConfig *DposNetworkParameters, credentials *Credentials) *DelegateContract {
	executor := &FunctionExecutor{
		httpEntry:    dposConfig.Url,
		chainId:      dposConfig.ChainId,
		contractAddr: dposConfig.StakingContract,
		credentials:  credentials,
	}
	return &DelegateContract{executor}
}

func (dc *DelegateContract) Delegate(nodeId string, stakingAmountType StakingAmountType, amount *big.Int) (string, error) {
	params := []interface{}{UInt16{ValueInner: stakingAmountType.GetValue()}, NodeId{HexStringId: nodeId}, UInt256{ValueInner: amount}}
	function := NewFunction(DELEGATE_FUNC_TYPE, params)

	var receipt string
	err := dc.executor.SendWithResult(function, &receipt)
	return receipt, err
}

func (dc *DelegateContract) UnDelegate(nodeId string, stakingBlockNum *big.Int, amount *big.Int) (string, error) {
	params := []interface{}{UInt64{ValueInner: stakingBlockNum}, NodeId{HexStringId: nodeId}, UInt256{ValueInner: amount}}
	function := NewFunction(WITHDRAW_DELEGATE_FUNC_TYPE, params)

	var receipt string
	err := dc.executor.SendWithResult(function, &receipt)
	return receipt, err
}