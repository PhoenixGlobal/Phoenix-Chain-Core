package protocols

import (
	"fmt"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common"
	ctypes "github.com/PhoenixGlobal/Phoenix-Chain-Core/consensus/pbft/types"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/consensus/pbft/utils"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/crypto"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/rlp"
	"sync/atomic"
)

// Message used to PreCommit.
type PreCommit struct {
	Epoch          uint64             `json:"epoch"`
	ViewNumber     uint64             `json:"viewNumber"`
	BlockHash      common.Hash        `json:"blockHash"`
	BlockNumber    uint64             `json:"blockNumber"`
	BlockIndex     uint32             `json:"blockIndex"` // The block number of the current ViewNumber proposal, 0....10
	ValidatorIndex uint32             `json:"validatorIndex"`
	ParentQC       *ctypes.QuorumCert `json:"parentQC" rlp:"nil"`
	Signature      ctypes.Signature   `json:"signature"`
	messageHash    atomic.Value       `rlp:"-"`
}

func (pv *PreCommit) String() string {
	return fmt.Sprintf("{Epoch:%d,ViewNumber:%d,Hash:%s,Number:%d,BlockIndex:%d,ValidatorIndex:%d}",
		pv.Epoch, pv.ViewNumber, pv.BlockHash.TerminalString(), pv.BlockNumber, pv.BlockIndex, pv.ValidatorIndex)
}

func (pv *PreCommit) MsgHash() common.Hash {
	if mhash := pv.messageHash.Load(); mhash != nil {
		return mhash.(common.Hash)
	}
	v := utils.BuildHash(PreCommitsMsg,
		utils.MergeBytes(common.Uint64ToBytes(pv.ViewNumber), pv.BlockHash.Bytes(), common.Uint32ToBytes(pv.BlockIndex), pv.Signature.Bytes()))
	pv.messageHash.Store(v)
	return v
}

func (pv *PreCommit) BHash() common.Hash {
	return pv.BlockHash
}
func (pv *PreCommit) EpochNum() uint64 {
	return pv.Epoch
}
func (pv *PreCommit) ViewNum() uint64 {
	return pv.ViewNumber
}

func (pv *PreCommit) BlockNum() uint64 {
	return pv.BlockNumber
}

func (pv *PreCommit) NodeIndex() uint32 {
	return pv.ValidatorIndex
}

func (pv *PreCommit) CannibalizeBytes() ([]byte, error) {
	buf, err := rlp.EncodeToBytes([]interface{}{
		pv.Epoch,
		pv.ViewNumber,
		pv.BlockHash,
		pv.BlockNumber,
		pv.BlockIndex,
		ctypes.RoundStepPreCommit,
	})

	if err != nil {
		return nil, err
	}
	return crypto.Keccak256(buf), nil
}

func (pv *PreCommit) Sign() []byte {
	return pv.Signature.Bytes()
}

func (pv *PreCommit) SetSign(sign []byte) {
	pv.Signature.SetBytes(sign)
}

func (pv *PreCommit) EqualState(vote *PreCommit) bool {
	return pv.Epoch == vote.Epoch &&
		pv.BlockHash == vote.BlockHash &&
		pv.BlockNumber == vote.BlockNumber &&
		pv.BlockIndex == vote.BlockIndex
}

// Message used to get PreCommit.
type GetPreCommit struct {
	Epoch       uint64
	ViewNumber  uint64
	BlockNumber  uint64
	BlockIndex  uint32
	UnKnownSet  *utils.BitArray
	messageHash atomic.Value `json:"-" rlp:"-"`
}

func (s *GetPreCommit) String() string {
	return fmt.Sprintf("{Epoch:%d,ViewNumber:%d,BlockNumber:%d,BlockIndex:%d,UnKnownSet:%s}", s.Epoch, s.ViewNumber,s.BlockNumber, s.BlockIndex, s.UnKnownSet.String())
}

func (s *GetPreCommit) MsgHash() common.Hash {
	if mhash := s.messageHash.Load(); mhash != nil {
		return mhash.(common.Hash)
	}
	v := utils.BuildHash(GetPreCommitMsg, utils.MergeBytes(common.Uint64ToBytes(s.Epoch), common.Uint64ToBytes(s.ViewNumber),
		common.Uint32ToBytes(s.BlockIndex)))
	s.messageHash.Store(v)
	return v
}

func (s *GetPreCommit) BHash() common.Hash {
	return common.Hash{}
}