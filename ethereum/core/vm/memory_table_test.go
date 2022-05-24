package vm

import (
	"github.com/holiman/uint256"
	"testing"
)

func TestMemorySha3(t *testing.T) {
	stack := newstack()
	stack.push(uint256.NewInt().SetBytes([]byte{0x01}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x02}))
	r, _ := memorySha3(stack)
	if r != 3 {
		t.Errorf("Expected: 3, got %d", r)
	}
}

func TestMemoryCallDataCopy(t *testing.T) {
	stack := newstack()
	stack.push(uint256.NewInt().SetBytes([]byte{0x01}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x02}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x03}))
	r, _ := memoryCallDataCopy(stack)
	if r != 4 {
		t.Errorf("Expected: 4, got %d", r)
	}
}

func TestMemoryReturnDataCopy(t *testing.T) {
	stack := newstack()
	stack.push(uint256.NewInt().SetBytes([]byte{0x01}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x02}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x03}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x05}))
	r, _ := memoryReturnDataCopy(stack)
	if r != 7 {
		t.Errorf("Expected: 7, got %d", r)
	}
}

func TestMemoryCodeCopy(t *testing.T) {
	stack := newstack()
	stack.push(uint256.NewInt().SetBytes([]byte{0x01}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x02}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x03}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x05}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x06}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x07}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x08}))
	r, _ := memoryCodeCopy(stack)
	if r != 14 {
		t.Errorf("Expected: 14, got %d", r)
	}
}

func TestMemoryExtCodeCopy(t *testing.T) {
	stack := newstack()
	stack.push(uint256.NewInt().SetBytes([]byte{0x01}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x02}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x03}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x05}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x06}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x07}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x08}))
	r, _ := memoryExtCodeCopy(stack)
	if r != 12 {
		t.Errorf("Expected: 12, got %d", r)
	}
}

func TestMemoryMLoad(t *testing.T) {
	stack := newstack()
	stack.push(uint256.NewInt().SetBytes([]byte{0x08}))
	r, _ := memoryMLoad(stack)
	if r != 40 {
		t.Errorf("Expected: 40, got %d", r)
	}
}

func TestMemoryMStore8(t *testing.T) {
	stack := newstack()
	stack.push(uint256.NewInt().SetBytes([]byte{0x08}))
	r, _ := memoryMStore8(stack)
	if r != 9 {
		t.Errorf("Expected: 9, got %d", r)
	}
}

func TestMemoryMStore(t *testing.T) {
	stack := newstack()
	stack.push(uint256.NewInt().SetBytes([]byte{0x08}))
	r, _ := memoryMStore(stack)
	if r != 40 {
		t.Errorf("Expected: 40, got %d", r)
	}
}

func TestMemoryCreate(t *testing.T) {
	stack := newstack()
	stack.push(uint256.NewInt().SetBytes([]byte{0x08}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x03}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x08}))
	r, _ := memoryCreate(stack)
	if r != 11 {
		t.Errorf("Expected: 11, got %d", r)
	}
	r, _ = memoryCreate(stack)
	if r != 11 {
		t.Errorf("Expected: 11, got %d", r)
	}
}

func TestMemoryCall(t *testing.T) {
	stack := newstack()
	stack.push(uint256.NewInt().SetBytes([]byte{0x08}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x08}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x08}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x08}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x08}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x06}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x08}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x03}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x08}))
	r, _ := memoryCall(stack)
	if r != 16 {
		t.Errorf("Expected: 16, got %d", r)
	}

	// memoryDelegateCall verify.
	r, _ = memoryDelegateCall(stack)
	if r != 16 {
		t.Errorf("Expected: 16, got %d", r)
	}

	// memoryStaticCall verify.
	r, _ = memoryDelegateCall(stack)
	if r != 16 {
		t.Errorf("Expected: 16, got %d", r)
	}
}

func TestMemoryReturn(t *testing.T) {
	stack := newstack()
	stack.push(uint256.NewInt().SetBytes([]byte{0x08}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x03}))
	stack.push(uint256.NewInt().SetBytes([]byte{0x08}))
	r, _ := memoryReturn(stack)
	if r != 11 {
		t.Errorf("Expected: 11, got %d", r)
	}

	// for memoryRevert.
	r, _ = memoryRevert(stack)
	if r != 11 {
		t.Errorf("Expected: 11, got %d", r)
	}

	// for memoryLog.
	r, _ = memoryLog(stack)
	if r != 11 {
		t.Errorf("Expected: 11, got %d", r)
	}
}
