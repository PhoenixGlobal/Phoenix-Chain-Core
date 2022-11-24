package vm

import (
	"github.com/holiman/uint256"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common"
)

func TestValidJumpdest(t *testing.T) {
	code := []byte{0x00, 0x5b}
	contract := &Contract{
		Code:      code,
		CodeHash:  common.BytesToHash(code),
		jumpdests: make(map[common.Hash]bitvec),
	}
	r := contract.validJumpdest(uint256.NewInt().SetUint64(3))
	if r {
		t.Errorf("Expected false, got true")
	}
	r = contract.validJumpdest(uint256.NewInt().SetUint64(1))
	if !r {
		t.Errorf("Expected true, got false")
	}
	r = contract.validJumpdest(uint256.NewInt().SetUint64(2))
	if r {
		t.Errorf("Expected false, got true")
	}
}

func TestGetOp(t *testing.T) {
	code := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06}
	testCases := []struct {
		n    uint64
		want OpCode
	}{
		{n: 0, want: STOP},
		{n: 1, want: ADD},
		{n: 2, want: MUL},
		{n: 3, want: SUB},
		{n: 4, want: DIV},
		{n: 5, want: SDIV},
		{n: 6, want: MOD},
	}
	c := &Contract{
		Code: code,
	}
	// iterate and verify.
	for _, v := range testCases {
		opCode := c.GetOp(v.n)
		assert.Equal(t, v.want, opCode)
	}
}

func TestGetByte(t *testing.T) {
	code := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06}
	testCases := []struct {
		n    uint64
		want byte
	}{
		{n: 0, want: byte(STOP)},
		{n: 1, want: byte(ADD)},
		{n: 2, want: byte(MUL)},
		{n: 3, want: byte(SUB)},
		{n: 4, want: byte(DIV)},
		{n: 5, want: byte(SDIV)},
		{n: 6, want: byte(MOD)},
		{n: 100, want: byte(0x00)},
	}
	c := &Contract{
		Code: code,
	}
	// iterate and verify.
	for _, v := range testCases {
		r := c.GetByte(v.n)
		assert.Equal(t, v.want, r)
	}
}

func TestCaller(t *testing.T) {
	addr := common.BytesToAddress([]byte("aaa"))
	contract := &Contract{
		CallerAddress: addr,
	}
	cr := contract.Caller()
	if cr != addr {
		t.Errorf("Not equal, expect: %s, actual: %s", addr, cr)
	}
}

func TestUseGas(t *testing.T) {
	contract := &Contract{
		Gas: 1000,
	}
	cr := contract.UseGas(100)
	if !cr {
		t.Errorf("Expected: true, got false")
	}
	laveGas := contract.Gas - 100
	if laveGas != 800 {
		t.Errorf("Expected: 800, actual: %d", laveGas)
	}

	// Simulation does not hold.
	cr = contract.UseGas(1000)
	if cr {
		t.Errorf("Expected: false, got true")
	}
}

func TestValue(t *testing.T) {
	contract := &Contract{
		value: buildBigInt(100),
	}
	cr := contract.Value()
	if cr.Cmp(new(big.Int).SetUint64(100)) != 0 {
		t.Errorf("Expected: 100, got: %d", cr)
	}
}

func TestSetting(t *testing.T) {
	// test SetCallCode of method.
	contract := &Contract{
		value: buildBigInt(100),
	}
	addr := common.BytesToAddress([]byte("I'm address"))
	code := []byte{0x00, 0x10}
	hash := common.BytesToHash(code)
	contract.SetCallCode(&addr, hash, code)

	if *contract.CodeAddr != addr {
		t.Errorf("Expected: %s, got %s", addr, contract.CodeAddr)
	}
	assert.Equal(t, code, contract.Code)
	if contract.CodeHash != hash {
		t.Errorf("Expected: %s, got %s", hash, contract.CodeHash)
	}

	// test SetCodeOptionalHash
	optional := codeAndHash{
		code: code,
		hash: hash,
	}
	contract.SetCodeOptionalHash(&addr, &optional)
	if *contract.CodeAddr != addr {
		t.Errorf("Expected: %s, got %s", addr, contract.CodeAddr)
	}
	assert.Equal(t, code, contract.Code)
	if contract.CodeHash != hash {
		t.Errorf("Expected: %s, got %s", hash, contract.CodeHash)
	}

	// test SetCallAbi.
	contract.SetCallAbi(&addr, hash, []byte{})

}
