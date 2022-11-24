package tests

import (
	"fmt"
	"math/big"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/configs"
)

// Forks table defines supported forks and their chain config.
var Forks = map[string]*configs.ChainConfig{
	"Frontier": {
		ChainID: big.NewInt(1),
	},
	"EIP155": {
		ChainID:        big.NewInt(1),
		EIP155Block:    big.NewInt(0),
	},
}

// UnsupportedForkError is returned when a test requests a fork that isn't implemented.
type UnsupportedForkError struct {
	Name string
}

func (e UnsupportedForkError) Error() string {
	return fmt.Sprintf("unsupported fork %q", e.Name)
}
