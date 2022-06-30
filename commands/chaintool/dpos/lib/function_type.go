package lib

type FunctionType int

const (
	STAKING_FUNC_TYPE FunctionType = 1000
	UPDATE_STAKING_INFO_FUNC_TYPE FunctionType = 1001
	ADD_STAKING_FUNC_TYPE FunctionType = 1002
	WITHDREW_STAKING_FUNC_TYPE FunctionType = 1003
)

func (ft FunctionType) GetType() int {
	return int(ft)
}