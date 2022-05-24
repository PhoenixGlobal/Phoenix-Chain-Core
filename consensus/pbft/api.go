package pbft

import (
	"encoding/json"

	"Phoenix-Chain-Core/consensus/pbft/state"
	"Phoenix-Chain-Core/consensus/pbft/types"
	"Phoenix-Chain-Core/libs/crypto/bls"
)

type Status struct {
	Tree      *types.BlockTree `json:"blockTree"`
	State     *state.ViewState `json:"state"`
	Validator bool             `json:"validator"`
}

// API defines an exposed API function interface.
type API interface {
	Status() []byte
	Evidences() string
	GetPrepareQC(number uint64) *types.QuorumCert
	GetSchnorrNIZKProve() (*bls.SchnorrProof, error)
}

// PublicDebugConsensusAPI provides an API to access the PhoenixChain blockchain.
// It offers only methods that operate on public data that
// is freely available to anyone.
type PublicDebugConsensusAPI struct {
	engine API
}

// NewDebugConsensusAPI creates a new PhoenixChain blockchain API.
func NewDebugConsensusAPI(engine API) *PublicDebugConsensusAPI {
	return &PublicDebugConsensusAPI{engine: engine}
}

// ConsensusStatus returns the status data of the consensus engine.
func (s *PublicDebugConsensusAPI) ConsensusStatus() *Status {
	b := s.engine.Status()
	var status Status
	err := json.Unmarshal(b, &status)
	if err == nil {
		return &status
	}
	return nil
}

// GetPrepareQC returns the QC certificate corresponding to the blockNumber.
func (s *PublicDebugConsensusAPI) GetPrepareQC(number uint64) *types.QuorumCert {
	return s.engine.GetPrepareQC(number)
}

// PublicPhoenixchainConsensusAPI provides an API to access the PhoenixChain blockchain.
// It offers only methods that operate on public data that
// is freely available to anyone.
type PublicPhoenixchainConsensusAPI struct {
	engine API
}

// NewPublicPhoenixchainConsensusAPI creates a new PhoenixChain blockchain API.
func NewPublicPhoenixchainConsensusAPI(engine API) *PublicPhoenixchainConsensusAPI {
	return &PublicPhoenixchainConsensusAPI{engine: engine}
}

// Evidences returns the relevant data of the verification.
func (s *PublicPhoenixchainConsensusAPI) Evidences() string {
	return s.engine.Evidences()
}

// PublicAdminConsensusAPI provides an API to access the PhoenixChain blockchain.
// It offers only methods that operate on public data that
// is freely available to anyone.
type PublicAdminConsensusAPI struct {
	engine API
}

// NewPublicAdminConsensusAPI creates a new PhoenixChain blockchain API.
func NewPublicAdminConsensusAPI(engine API) *PublicAdminConsensusAPI {
	return &PublicAdminConsensusAPI{engine: engine}
}

func (s *PublicAdminConsensusAPI) GetSchnorrNIZKProve() string {
	proof, err := s.engine.GetSchnorrNIZKProve()
	if nil != err {
		return err.Error()
	}
	proofByte, err := proof.MarshalText()
	if nil != err {
		return err.Error()
	}
	return string(proofByte)
}
