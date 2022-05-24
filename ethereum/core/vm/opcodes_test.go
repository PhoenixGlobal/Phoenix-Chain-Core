package vm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToOp(t *testing.T) {
	testCases := []struct {
		strOpCode  string
		wantOpCode OpCode
	}{
		{
			strOpCode:  "CALLCODE",
			wantOpCode: CALLCODE,
		}, {
			strOpCode:  "PUSH1",
			wantOpCode: PUSH1,
		}, {
			strOpCode:  "MLOAD",
			wantOpCode: MLOAD,
		},
	}
	for _, v := range testCases {
		opCode := StringToOp(v.strOpCode)
		assert.Equal(t, v.wantOpCode, opCode)
		assert.Equal(t, v.strOpCode, v.wantOpCode.String())
	}

	// test -> string
	str := OpCode(0x00).String()
	t.Log(str)

}

func TestOpCode_IsPush(t *testing.T) {
	testCases := []struct {
		opCode OpCode
		want   bool
	}{
		{opCode: CALLCODE, want: false},
		{opCode: CALLDATALOAD, want: false},
		{opCode: MLOAD, want: false},
		{opCode: SUB, want: false},
		{opCode: PUSH, want: false},
		{opCode: PUSH1, want: true},
	}
	for _, v := range testCases {
		assert.Equal(t, v.want, v.opCode.IsPush())
		assert.Equal(t, false, v.opCode.IsStaticJump())
	}
}
