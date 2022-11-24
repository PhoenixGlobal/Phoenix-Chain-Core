package reward

import (
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/pos/xutil"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/p2p/discover"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common"
)

const DelegateRewardPerLength = 1000

var (
	HistoryIncreasePrefix    = []byte("RewardHistory")
	LastYearEndBalancePrefix = []byte("RewardBalance")
	YearStartBlockNumberKey  = []byte("YearStartBlockNumberKey")
	YearStartTimeKey         = []byte("YearStartTimeKey")
	RemainingRewardKey       = []byte("RemainingRewardKey")
	NewBlockRewardKey        = []byte("NewBlockRewardKey")
	StakingRewardKey         = []byte("StakingRewardKey")
	ChainYearNumberKey       = []byte("ChainYearNumberKey")
	delegateRewardPerKey     = []byte("DelegateRewardPerKey")
)

// GetHistoryIncreaseKey used for search the balance of reward pool at last year
func GetHistoryIncreaseKey(year uint32) []byte {
	return append(HistoryIncreasePrefix, common.Uint32ToBytes(year)...)
}

//
func HistoryBalancePrefix(year uint32) []byte {
	return append(LastYearEndBalancePrefix, common.Uint32ToBytes(year)...)
}

func DelegateRewardPerKey(nodeID discover.NodeID, stakingNum, epoch uint64) []byte {
	index := uint32(epoch / DelegateRewardPerLength)
	add, err := xutil.NodeId2Addr(nodeID)
	if err != nil {
		panic(err)
	}
	perKeyLength := len(delegateRewardPerKey)
	lengthUint32, lengthUint64 := 4, 8
	keyAdd := make([]byte, perKeyLength+common.AddressLength+lengthUint64+lengthUint32)
	n := copy(keyAdd[:perKeyLength], delegateRewardPerKey)
	n += copy(keyAdd[n:n+common.AddressLength], add.Bytes())
	n += copy(keyAdd[n:n+lengthUint64], common.Uint64ToBytes(stakingNum))
	copy(keyAdd[n:n+lengthUint32], common.Uint32ToBytes(index))

	return keyAdd
}

func DelegateRewardPerKeys(nodeID discover.NodeID, stakingNum, fromEpoch, toEpoch uint64) [][]byte {
	indexFrom := uint32(fromEpoch / DelegateRewardPerLength)
	indexTo := uint32(toEpoch / DelegateRewardPerLength)
	add, err := xutil.NodeId2Addr(nodeID)
	if err != nil {
		panic(err)
	}
	perKeyLength := len(delegateRewardPerKey)
	lengthUint64 := 8

	delegateRewardPerPrefix := make([]byte, perKeyLength+common.AddressLength+lengthUint64)
	n := copy(delegateRewardPerPrefix[:perKeyLength], delegateRewardPerKey)
	n += copy(delegateRewardPerPrefix[n:n+common.AddressLength], add.Bytes())
	n += copy(delegateRewardPerPrefix[n:n+lengthUint64], common.Uint64ToBytes(stakingNum))

	keys := make([][]byte, 0)
	for i := indexFrom; i <= indexTo; i++ {
		delegateRewardPerKey := append(delegateRewardPerPrefix[:], common.Uint32ToBytes(i)...)
		keys = append(keys, delegateRewardPerKey)
	}
	return keys
}
