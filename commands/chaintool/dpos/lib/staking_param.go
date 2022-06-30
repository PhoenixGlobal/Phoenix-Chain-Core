package lib

import (
	"Phoenix-Chain-Core/libs/common"
	"math/big"
)

type StakingParam struct {
	NodeId string
	Amount *big.Int
	StakingAmountType StakingAmountType
	BenefitAddress string
	ExternalId string
	NodeName string
	WebSite string
	Details string
	ProcessVersion ProgramVersion
	BlsPubKey string
	BlsProof string
	RewardPer *big.Int
}

func (sp StakingParam) SubmitInputParameters() []interface{} {
	return []interface{}{
		UInt16{ValueInner: sp.StakingAmountType.GetValue()},
		common.MustStringToAddress(sp.BenefitAddress),
		NodeId{HexStringId: sp.NodeId},
		Utf8String{ValueInner: sp.ExternalId},
		Utf8String{ValueInner: sp.NodeName},
		Utf8String{ValueInner: sp.WebSite},
		Utf8String{ValueInner: sp.Details},
		UInt256{ValueInner: sp.Amount},
		UInt32{ValueInner: sp.RewardPer},
		UInt32{ValueInner: sp.ProcessVersion.Version},
		HexStringParam{HexStringValue: sp.ProcessVersion.Sign},
		HexStringParam{HexStringValue: sp.BlsPubKey},
		HexStringParam{HexStringValue: sp.BlsProof},
	}
}
