package consensus

import (
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/crypto/bls"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/p2p/discover"
)

type EvidenceType uint8

type Evidence interface {
	//Verify(ecdsa.PublicKey) error
	Equal(Evidence) bool
	//return lowest number
	BlockNumber() uint64
	Epoch() uint64
	ViewNumber() uint64
	Hash() []byte
	//Address() common.NodeAddress
	NodeID() discover.NodeID
	BlsPubKey() *bls.PublicKey
	Validate() error
	Type() EvidenceType
	ValidateMsg() bool
}

type Evidences []Evidence

func (e Evidences) Len() int {
	return len(e)
}

type EvidencePool interface {
	//Deserialization of evidence
	//UnmarshalEvidence(data string) (Evidences, error)
	//Get current evidences
	Evidences() Evidences
	//Clear all evidences
	Clear(epoch uint64, blockNumber uint64)
	Close()
}
