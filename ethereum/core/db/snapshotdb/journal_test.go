package snapshotdb

import (
	"math/big"
	"testing"
)

func TestEncodeJournalKey(t *testing.T) {
	num := new(big.Int).SetUint64(111)
	key := EncodeWalKey(num)
	num2 := DecodeWalKey(key)
	if num.Cmp(num2) != 0 {
		t.Error("num should same")
	}
}
