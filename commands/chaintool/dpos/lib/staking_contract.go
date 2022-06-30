package lib

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
