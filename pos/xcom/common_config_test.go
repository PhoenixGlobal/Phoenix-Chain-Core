package xcom

import (
	"Phoenix-Chain-Core/libs/common"
	"Phoenix-Chain-Core/libs/rlp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultEMConfig(t *testing.T) {
	t.Run("DefaultMainNet", func(t *testing.T) {
		if getDefaultEMConfig(DefaultMainNet) == nil {
			t.Error("DefaultMainNet can't be nil config")
		}
		if err := CheckEconomicModel(); nil != err {
			t.Error(err)
		}
	})
	t.Run("DefaultTestNet", func(t *testing.T) {
		if getDefaultEMConfig(DefaultTestNet) == nil {
			t.Error("DefaultTestNet can't be nil config")
		}
		if err := CheckEconomicModel(); nil != err {
			t.Error(err)
		}
	})
	t.Run("DefaultUnitTestNet", func(t *testing.T) {
		if getDefaultEMConfig(DefaultUnitTestNet) == nil {
			t.Error("DefaultUnitTestNet can't be nil config")
		}
		if err := CheckEconomicModel(); nil != err {
			t.Error(err)
		}
	})
	if getDefaultEMConfig(10) != nil {
		t.Error("the chain config not support")
	}
}

func TestMainNetHash(t *testing.T) {
	tempEc := getDefaultEMConfig(DefaultMainNet)
	bytes, err := rlp.EncodeToBytes(tempEc)
	if err != nil {
		t.Error(err)
	}
	assert.True(t, common.RlpHash(bytes).Hex() == MainNetECHash)
}
