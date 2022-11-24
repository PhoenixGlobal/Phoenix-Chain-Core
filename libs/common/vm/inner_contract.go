package vm

import "github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common"

// the inner contract addr  table
var (
	RestrictingContractAddr    = common.HexToAddress("0x1000000000000000000000000000000000000001") // The PhoenixChain Precompiled contract addr for restricting
	StakingContractAddr        = common.HexToAddress("0x1000000000000000000000000000000000000002") // The PhoenixChain Precompiled contract addr for staking
	RewardManagerPoolAddr      = common.HexToAddress("0x1000000000000000000000000000000000000003") // The PhoenixChain Precompiled contract addr for reward
	SlashingContractAddr       = common.HexToAddress("0x1000000000000000000000000000000000000004") // The PhoenixChain Precompiled contract addr for slashing
	GovContractAddr            = common.HexToAddress("0x1000000000000000000000000000000000000005") // The PhoenixChain Precompiled contract addr for governance
	DelegateRewardPoolAddr     = common.HexToAddress("0x1000000000000000000000000000000000000006") // The PhoenixChain Precompiled contract addr for delegate reward
	ValidatorInnerContractAddr = common.HexToAddress("0x2000000000000000000000000000000000000000") // The PhoenixChain Precompiled contract addr for pbft inner
)

type PrecompiledContractCheck interface {
	IsPhoenixChainPrecompiledContract(address common.Address) bool
}

var PrecompiledContractCheckInstance PrecompiledContractCheck
