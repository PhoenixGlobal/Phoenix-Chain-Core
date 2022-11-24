package types

import (
	"errors"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/rlp"
)

// EncodeExtra encode pbft version and `QuorumCert` as extra data.
func EncodeExtra(pbftVersion byte, qc *QuorumCert) ([]byte, error) {
	extra := []byte{pbftVersion}
	bxBytes, err := rlp.EncodeToBytes(qc)
	if err != nil {
		return nil, err
	}
	extra = append(extra, bxBytes...)
	return extra, nil
}

// DecodeExtra decode extra data as pbft version and `QuorumCert`.
func DecodeExtra(extra []byte) (byte, *QuorumCert, error) {
	if len(extra) == 0 {
		return 0, nil, errors.New("empty extra")
	}
	version := extra[0]
	var qc QuorumCert
	err := rlp.DecodeBytes(extra[1:], &qc)
	if err != nil {
		return 0, nil, err
	}
	return version, &qc, nil
}
