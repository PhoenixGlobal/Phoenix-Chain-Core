package types

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/consensus/pbft/utils"
)

func TestCodec(t *testing.T) {
	// EncodeExtra
	pbftVersion := 1
	qc := &QuorumCert{
		Epoch:        1,
		ViewNumber:   0,
		BlockHash:    common.BytesToHash(utils.Rand32Bytes(32)),
		BlockNumber:  1,
		BlockIndex:   0,
		Signature:    Signature{},
		ValidatorSet: utils.NewBitArray(25),
	}
	data, err := EncodeExtra(byte(pbftVersion), qc)
	assert.Nil(t, err)
	assert.True(t, len(data) > 0)

	// DecodeExtra
	version, cert, err := DecodeExtra(data)
	assert.Nil(t, err)
	assert.Equal(t, byte(pbftVersion), version)
	assert.Equal(t, qc.BlockHash, cert.BlockHash)
}
