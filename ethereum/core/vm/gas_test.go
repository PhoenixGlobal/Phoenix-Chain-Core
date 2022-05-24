package vm

import (
	"github.com/holiman/uint256"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCallGas(t *testing.T) {

	_, err := callGas(100, 2, uint256.NewInt().SetUint64(10))
	assert.Nil(t, err)

	_, err = callGas(100, 2, uint256.NewInt().SetUint64(1000000000000000))
	assert.Nil(t, err)
}
