package lib

const (
	BASE_DEFAULT_GAS_LIMIT = uint64(21000)
	BASE_NON_ZERO_GAS_LIMIT = uint64(68)
	BASE_ZERO_GAS_LIMIT = uint64(4)
)

func IsLocalSupportFunction(ftype FunctionType) bool {
	switch ftype {
	case
		STAKING_FUNC_TYPE,
		UPDATE_STAKING_INFO_FUNC_TYPE,
		ADD_STAKING_FUNC_TYPE,
		WITHDREW_STAKING_FUNC_TYPE:
		return true
	default:
		return false
	}
}

func GetGasLimit(f *Function) uint64 {
	if IsLocalSupportFunction(f.Type) {
		bytes := f.ToBytes()
		return BASE_DEFAULT_GAS_LIMIT +
			getContractGasLimit(f.Type) +
			getFunctionGasLimit(f.Type) +
			getInterfaceDynamicGasLimit(f.Type, f.InputParams) +
			getDataGasLimit(bytes)
	}
	return 0
}

func getContractGasLimit(ftype FunctionType) uint64 {
	switch ftype {
	case
		STAKING_FUNC_TYPE,
		UPDATE_STAKING_INFO_FUNC_TYPE,
		ADD_STAKING_FUNC_TYPE,
		WITHDREW_STAKING_FUNC_TYPE:
		return uint64(6000)
	default:
		return uint64(0)
	}
}

func getFunctionGasLimit(ftype FunctionType) uint64 {
	switch ftype {
	case STAKING_FUNC_TYPE:
		return uint64(32000)
	case UPDATE_STAKING_INFO_FUNC_TYPE:
		return uint64(12000)
	case ADD_STAKING_FUNC_TYPE,
		WITHDREW_STAKING_FUNC_TYPE:
		return uint64(20000)
	default:
		return uint64(0)
	}
}

func getInterfaceDynamicGasLimit(ftype FunctionType, inputParameters []interface{}) uint64 {
	return uint64(0)
}

func getDataGasLimit(rlpData []byte) uint64 {
	var nonZeroSize uint64 = 0
	var zeroSize uint64 = 0

	for _, b := range rlpData {
		if b != 0 {
			nonZeroSize++
		} else {
			zeroSize++
		}
	}
	return nonZeroSize*BASE_NON_ZERO_GAS_LIMIT + zeroSize*BASE_ZERO_GAS_LIMIT
}
