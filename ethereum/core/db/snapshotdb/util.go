package snapshotdb

import (
	"bytes"
	"golang.org/x/crypto/sha3"
	"io"
	"math/big"
	"math/rand"
	"sort"
	"time"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core/types"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/rlp"
)

func generateKVHash(k, v []byte, hash common.Hash) common.Hash {
	var buf bytes.Buffer
	buf.Write(k)
	buf.Write(v)
	buf.Write(hash.Bytes())
	return rlpHash(buf.Bytes())
}

func rlpHash(x interface{}) (h common.Hash) {
	hw := sha3.NewLegacyKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}

func encode(x interface{}) ([]byte, error) {
	return rlp.EncodeToBytes(x)
}

func decode(r io.Reader, val interface{}) error {
	return rlp.Decode(r, val)
}

func generateHeader(num *big.Int, parentHash common.Hash) *types.Header {
	h := new(types.Header)
	h.Number = num
	h.ParentHash = parentHash
	return h
}

func generateHash(n string) common.Hash {
	var buf bytes.Buffer
	buf.Write([]byte(n))
	return rlpHash(buf.Bytes())
}

func randomString2(s string) []byte {
	b := new(bytes.Buffer)
	if s != "" {
		b.Write([]byte(s))
	}
	for i := 0; i < 8; i++ {
		b.WriteByte(' ' + byte(rand.Uint64()))
	}
	return b.Bytes()
}

func generatekv(n int) kvs {
	rand.Seed(time.Now().UnixNano())
	kvs := make(kvs, n)
	for i := 0; i < n; i++ {
		kvs[i] = kv{
			key:   randomString2(""),
			value: randomString2(""),
		}
	}
	sort.Sort(kvs)
	return kvs
}

func generatekvWithPrefix(n int, p string) kvs {
	rand.Seed(time.Now().UnixNano())
	kvs := make(kvs, n)
	for i := 0; i < n; i++ {
		kvs[i] = kv{
			key:   randomString2(p),
			value: randomString2(p),
		}
	}
	sort.Sort(kvs)
	return kvs
}

type kv struct {
	key   []byte
	value []byte
}

type kvs []kv

func (k kvs) Len() int {
	return len(k)
}

func (k kvs) Less(i, j int) bool {
	n := bytes.Compare(k[i].key, k[j].key)
	if n == -1 {
		return true
	}
	return false
}

func (k kvs) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

type kvsMaxToMin []kv

func (k kvsMaxToMin) Len() int {
	return len(k)
}

func (k kvsMaxToMin) Less(i, j int) bool {
	if bytes.Compare(k[i].key, k[j].key) >= 0 {
		return true
	}
	return false
}

func (k kvsMaxToMin) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

func (k *kvsMaxToMin) Push(x interface{}) {
	*k = append(*k, x.(kv))
}

func (k *kvsMaxToMin) Pop() interface{} {
	n := len(*k)
	x := (*k)[n-1]
	*k = (*k)[:n-1]
	return x
}
