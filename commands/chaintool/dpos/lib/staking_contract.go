package lib

import "math/big"

type StakingContract struct {
	executor *FunctionExecutor
}

func NewStakingContract(dposConfig *DposNetworkParameters, credentials *Credentials) *StakingContract {
	executor := &FunctionExecutor{
		httpEntry:    dposConfig.Url,
		chainId:      dposConfig.ChainId,
		contractAddr: dposConfig.StakingContract,
		credentials:  credentials,
	}
	return &StakingContract{executor}
}

func (sc *StakingContract) Staking(stakingParam StakingParam) (string, error) {
	f := NewFunction(STAKING_FUNC_TYPE, stakingParam.SubmitInputParameters())

	var receipt string
	err := sc.executor.SendWithResult(f, &receipt)
	return receipt, err
}


func (sc *StakingContract) AddStaking(nodeId string, stakingAmountType StakingAmountType, amount *big.Int) (string, error) {
	params := []interface{}{NodeId{HexStringId: nodeId}, UInt16{ValueInner: stakingAmountType.GetValue()}, UInt256{ValueInner: amount}}
	f := NewFunction(ADD_STAKING_FUNC_TYPE, params)

	var receipt string
	err := sc.executor.SendWithResult(f, &receipt)
	return receipt, err
}

func (sc *StakingContract) UnStaking(nodeId string) (string, error) {
	f := NewFunction(WITHDRAW_STAKING_FUNC_TYPE, []interface{}{NodeId{HexStringId: nodeId}})

	var receipt string
	err := sc.executor.SendWithResult(f, &receipt)
	return receipt, err
}

func (sc *StakingContract) UpdateStakingInfo(updateStakingParam UpdateStakingParam) (string, error) {
	f := NewFunction(UPDATE_STAKING_INFO_FUNC_TYPE, updateStakingParam.SubmitInputParameters())

	var receipt string
	err := sc.executor.SendWithResult(f, &receipt)
	return receipt, err
}