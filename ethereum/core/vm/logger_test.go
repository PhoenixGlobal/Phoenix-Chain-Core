package vm

import (
	"bytes"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common/mock"
	"github.com/holiman/uint256"
	"math/big"
	"strings"
	"testing"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/log"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/configs"
)

type dummyContractRef struct {
	calledForEach bool
}

func (dummyContractRef) ReturnGas(*big.Int)          {}
func (dummyContractRef) Address() common.Address     { return common.Address{} }
func (dummyContractRef) Value() *big.Int             { return new(big.Int) }
func (dummyContractRef) SetCode(common.Hash, []byte) {}
func (d *dummyContractRef) ForEachStorage(callback func(key common.Hash, value []byte) bool) {
	d.calledForEach = true
}
func (d *dummyContractRef) SubBalance(amount *big.Int) {}
func (d *dummyContractRef) AddBalance(amount *big.Int) {}
func (d *dummyContractRef) SetBalance(*big.Int)        {}
func (d *dummyContractRef) SetNonce(uint64)            {}
func (d *dummyContractRef) Balance() *big.Int          { return new(big.Int) }

type dummyStatedb struct {
	mock.MockStateDB
}

func (*dummyStatedb) GetRefund() uint64 { return 1337 }

func TestStoreCapture(t *testing.T) {
	var (
		env      = NewEVM(Context{}, nil, &dummyStatedb{}, configs.TestChainConfig, Config{})
		logger   = NewStructLogger(nil)
		mem      = NewMemory()
		stack    = newstack()
		rstack   = newReturnStack()
		contract = NewContract(&dummyContractRef{}, &dummyContractRef{}, new(big.Int), 0)
	)
	stack.push(uint256.NewInt().SetUint64(1))
	stack.push(uint256.NewInt())
	var index common.Hash
	logger.CaptureState(env, 0, SSTORE, 0, 0, mem, stack, rstack, nil, contract, 0, nil)
	if len(logger.storage[contract.Address()]) == 0 {
		t.Fatalf("expected exactly 1 changed value on address %x, got %d", contract.Address(), len(logger.storage[contract.Address()]))
	}
	exp := common.BigToHash(big.NewInt(1))
	if logger.storage[contract.Address()][index] != exp {
		t.Errorf("expected %x, got %x", exp, logger.storage[contract.Address()][index])
	}
}

func TestNewWasmLogger(t *testing.T) {
	logger := log.New("test.wasm")

	buf := new(bytes.Buffer)

	logger.SetHandler(log.LvlFilterHandler(log.LvlDebug, log.StreamHandler(buf, log.FormatFunc(func(r *log.Record) []byte {
		return []byte(r.Msg)
	}))))

	wasmLog := NewWasmLogger(Config{Debug: true}, logger)

	wasmLog.Info("hello")
	wasmLog.Flush()
	if !strings.Contains(buf.String(), "hello") {
		t.Fatal("output log error", buf.String())
	}
}
