package runtime

import (
	"Phoenix-Chain-Core/ethereum/core"
	"Phoenix-Chain-Core/ethereum/core/db/snapshotdb"
	"Phoenix-Chain-Core/ethereum/core/vm"
	"Phoenix-Chain-Core/libs/common"
)

func NewEnv(cfg *Config) *vm.EVM {
	context := vm.Context{
		CanTransfer: core.CanTransfer,
		Transfer:    core.Transfer,
		GetHash:     func(uint64) common.Hash { return common.Hash{} },

		Origin:      cfg.Origin,
		Coinbase:    cfg.Coinbase,
		BlockNumber: cfg.BlockNumber,
		Time:        cfg.Time,
		Difficulty:  cfg.Difficulty,
		GasLimit:    cfg.GasLimit,
		GasPrice:    cfg.GasPrice,
	}

	return vm.NewEVM(context, snapshotdb.Instance(), cfg.State, cfg.ChainConfig, cfg.EVMConfig)
}
