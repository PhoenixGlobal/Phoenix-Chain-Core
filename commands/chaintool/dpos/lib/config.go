package lib

import (
	"Phoenix-Chain-Core/commands/chaintool/core"
	"Phoenix-Chain-Core/configs"
	"Phoenix-Chain-Core/ethereum/core/types"
	"Phoenix-Chain-Core/libs/common/vm"
	"math/big"
)

type Config struct {
	Url     string
	ChainId *big.Int
}

type DposNetworkParameters struct {
	Config
	StakingContract         string
}
type TransactionHash string

type TransactionReceipt types.Receipt


const PrivateKey = "f77192cea69effdf050308a18f16dc0213d061d553363adbbb3219381186e5a9"

const MainFanAccount = "0xf35894c6b663e1E0934812c81CF4984898D5d8d2"

var KeyCredentials, _ = NewCredential(PrivateKey)


var (
	DefaultMainNetConfig = Config{core.RpcUrl,  big.NewInt(configs.ChainId)}
	DposMainNetParams = &DposNetworkParameters{
		DefaultMainNetConfig,
		vm.StakingContractAddr.String(),
	}
)

type StakingAmountType uint16

const (
	FREE_AMOUNT_TYPE        StakingAmountType = 0
	RESTRICTING_AMOUNT_TYPE StakingAmountType = 1
)

func (sat StakingAmountType) GetValue() uint16 {
	return uint16(sat)
}


type ProgramVersion struct {
	Version *big.Int `json:"Version"`
	Sign    string   `json:"Sign"`
}
