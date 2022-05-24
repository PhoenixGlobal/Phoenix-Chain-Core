package tests

import (
	"math/big"
	"testing"

	"Phoenix-Chain-Core/configs"
)

func TestTransaction(t *testing.T) {
	t.Parallel()

	txt := new(testMatcher)
	txt.config(`^EIP155/`, configs.ChainConfig{
		EIP155Block:    big.NewInt(0),
		ChainID:        big.NewInt(1),
	})

	txt.walk(t, transactionTestDir, func(t *testing.T, name string, test *TransactionTest) {
		cfg := txt.findConfig(name)
		if err := txt.checkFailure(t, name, test.Run(cfg)); err != nil {
			t.Error(err)
		}
	})
}
