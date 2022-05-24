package vm

import (
	"Phoenix-Chain-Core/configs"
	"github.com/holiman/uint256"
)

// enable1884 applies EIP-1884 to the given jump table:
// - Increase cost of BALANCE to 700
// - Increase cost of EXTCODEHASH to 700
// - Define SELFBALANCE, with cost GasFastStep (5)
func enable1884(jt *JumpTable) {
	// Gas cost changes
	jt[BALANCE].constantGas = configs.BalanceGasEIP1884
	jt[EXTCODEHASH].constantGas = configs.ExtcodeHashGasEIP1884

	// New opcode
	jt[SELFBALANCE] = &operation{
		execute:     opSelfBalance,
		constantGas: GasFastStep,
		minStack:    minStack(0, 1),
		maxStack:    maxStack(0, 1),
	}
}

func opSelfBalance(pc *uint64, interpreter *EVMInterpreter, callContext *callCtx) ([]byte, error) {
	balance, _ := uint256.FromBig(interpreter.evm.StateDB.GetBalance(callContext.contract.Address()))
	callContext.stack.push(balance)
	return nil, nil
}

// enable1344 applies EIP-1344 (ChainID Opcode)
// - Adds an opcode that returns the current chain’s EIP-155 unique identifier
func enable1344(jt *JumpTable) {
	// New opcode
	jt[CHAINID] = &operation{
		execute:     opChainID,
		constantGas: GasQuickStep,
		minStack:    minStack(0, 1),
		maxStack:    maxStack(0, 1),
	}
}

// opChainID implements CHAINID opcode
func opChainID(pc *uint64, interpreter *EVMInterpreter, callContext *callCtx) ([]byte, error) {
	chainId, _ := uint256.FromBig(interpreter.evm.chainConfig.ChainID)
	callContext.stack.push(chainId)
	return nil, nil
}

// enable2315 applies EIP-2315 (Simple Subroutines)
// - Adds opcodes that jump to and return from subroutines
func enable2315(jt *JumpTable) {
	// New opcode
	jt[BEGINSUB] = &operation{
		execute:     opBeginSub,
		constantGas: GasQuickStep,
		minStack:    minStack(0, 0),
		maxStack:    maxStack(0, 0),
	}
	// New opcode
	jt[JUMPSUB] = &operation{
		execute:     opJumpSub,
		constantGas: GasSlowStep,
		minStack:    minStack(1, 0),
		maxStack:    maxStack(1, 0),
		jumps:       true,
	}
	// New opcode
	jt[RETURNSUB] = &operation{
		execute:     opReturnSub,
		constantGas: GasFastStep,
		minStack:    minStack(0, 0),
		maxStack:    maxStack(0, 0),
		jumps:       true,
	}
}
