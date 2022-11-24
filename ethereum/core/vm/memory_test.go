package vm

import (
	"github.com/holiman/uint256"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common"
)

func TestSet(t *testing.T) {
	m := NewMemory()
	testCases := []struct {
		value  []byte
		offset uint64
		size   uint64
		want   string
	}{
		{[]byte{0x00}, 0, 1, "00"},
		{[]byte{0x01}, 0, 1, "01"},
		{[]byte{0x00, 0x01, 0x02}, 0, 3, "000102"},
	}
	for _, v := range testCases {
		m.Resize(v.size)
		m.Set(v.offset, v.size, v.value)
		actual := common.Bytes2Hex(m.store)
		if actual != v.want {
			t.Errorf("Expected: %s, got: %s", v.want, actual)
		}
	}
}

func TestSet32(t *testing.T) {
	m := NewMemory()
	testCases := []struct {
		val    common.Hash
		offset uint64
		want   common.Hash
	}{
		{common.BytesToHash([]byte{0x11}), 0, common.BytesToHash([]byte{0x11})},
		{common.BytesToHash([]byte{0x00}), 0, common.BytesToHash([]byte{0x00})},
		{common.BytesToHash([]byte{0x11, 0xab}), 0, common.BytesToHash([]byte{0x11, 0xab})},
	}
	for _, v := range testCases {
		m.Resize(32)
		m.Set32(v.offset, uint256.NewInt().SetBytes(v.val.Bytes()))
		actual := common.Bytes2Hex(m.Data())
		if actual != v.want.HexWithNoPrefix() {
			t.Errorf("Expected: %s, got: %s", v.want.Hex(), actual)
		}
	}
}

func TestResize(t *testing.T) {
	m := NewMemory()
	testCases := []struct {
		size int64
	}{
		{size: 10},
		{size: 1000},
		{size: 2000},
	}
	for _, v := range testCases {
		m.Resize(uint64(v.size))
		assert.Equal(t, int(v.size), m.Len())
	}
}

func TestGet(t *testing.T) {
	m := NewMemory()
	testCases := []struct {
		value  []byte
		offset uint64
		size   uint64
		want   string
	}{
		{[]byte{0x00}, 0, 0, ""},
		{[]byte{0x00}, 0, 1, "00"},
		{[]byte{0x01}, 0, 1, "01"},
		{[]byte{0x00, 0x01, 0x02}, 0, 3, "000102"},
	}
	for _, v := range testCases {
		m.Resize(v.size)
		m.Set(v.offset, v.size, v.value)
		actual := common.Bytes2Hex(m.GetCopy(int64(v.offset), int64(v.size)))
		if actual != v.want {
			t.Errorf("Expected: %s, got: %s", v.want, actual)
		}
	}
}

func TestGetPtr(t *testing.T) {
	m := NewMemory()
	testCases := []struct {
		value  []byte
		offset uint64
		size   uint64
		want   string
	}{
		{[]byte{0x00}, 0, 0, ""},
		{[]byte{0x00}, 0, 1, "00"},
		{[]byte{0x01}, 0, 1, "01"},
		{[]byte{0x00, 0x01, 0x02}, 0, 3, "000102"},
	}
	for _, v := range testCases {
		m.Resize(v.size)
		m.Set(v.offset, v.size, v.value)
		actual := common.Bytes2Hex(m.GetPtr(int64(v.offset), int64(v.size)))
		if actual != v.want {
			t.Errorf("Expected: %s, got: %s", v.want, actual)
		}
		m.Print()
	}
}
