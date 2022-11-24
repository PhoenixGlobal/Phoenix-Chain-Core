package dpos

import (
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/commands/chaintool/dpos/lib"
	"fmt"
	"math/big"
	"testing"
)

func TestStakingContract_Staking(t *testing.T) {
	config := lib.DposMainNetParams
	sc := lib.NewStakingContract(config, lib.KeyCredentials)

	sp := lib.StakingParam{
		NodeId:            "0xe0fbf44d916ee125098327e378a3b5f23b421c9636b7662e8fec1d0c8ab4c7addcbeae865c005734f8451873eef11f64f361bdac3c57ad2143e892af24127769",
		Amount:            big.NewInt(10),
		StakingAmountType: lib.FREE_AMOUNT_TYPE,
		BenefitAddress:    lib.MainFanAccount,
		ExternalId:        "121412312",
		NodeName:          "chendai-node3",
		WebSite:           "www.baidu.com",
		Details:           "chendai-node3-details",
		ProcessVersion: lib.ProgramVersion{
			Version: big.NewInt(4096),
			Sign:    "0x0dca7024507a5d94c84b9c9deb417d56bf58f6fe5e37ecee86e64a62d1f518b67ddeeed7ba59a619b7f30ecd881164e96f9781b30309c07ea8985929401692de00",
		},
		BlsPubKey: "0x18fb11d0264dca2650b6e1a99208042d5a4552477f1899f04223b601b68371723a86da2325e566dab293b1f2d11b7e0610fef25bb1a6e44a3e1cfa580c0d4420dcaa480e7d01485ea8cb0f0225b6ac0a4693f0fc7600ab995acf09b3a025e790",
		BlsProof:  "0xe2f3e141679767187d3a82359c142a62a48281f33694d9d5f3bdcc771af3123d9f67f82001b4db46a601b14cbeb755efad9de350ce1e5c38966b1fa2747ef86f",
		RewardPer: big.NewInt(1000),
	}

	result, err := sc.Staking(sp)
	if err != nil {
		t.Errorf("StakingContract.Staking failed: %s", err)
	}
	fmt.Println("staking result is ",result)
}