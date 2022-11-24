package evidence

import (
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core/types/pbfttypes"
	"crypto/ecdsa"
	"math/big"
	"testing"
	"time"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/crypto/bls"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/crypto"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/p2p/discover"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core/types"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/consensus/pbft/protocols"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/consensus/pbft/utils"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common"
)

func newBlock(blockNumber int64) *types.Block {
	header := &types.Header{
		Number:      big.NewInt(blockNumber),
		ParentHash:  common.BytesToHash(utils.Rand32Bytes(32)),
		Time:        uint64(time.Now().UnixNano()),
		Extra:       make([]byte, 97),
		ReceiptHash: common.BytesToHash(utils.Rand32Bytes(32)),
		Root:        common.BytesToHash(utils.Rand32Bytes(32)),
	}
	block := types.NewBlockWithHeader(header)
	return block
}

func GenerateKeys(num int) ([]*ecdsa.PrivateKey, []*bls.SecretKey) {
	pk := make([]*ecdsa.PrivateKey, 0)
	sk := make([]*bls.SecretKey, 0)

	for i := 0; i < num; i++ {
		var blsKey bls.SecretKey
		blsKey.SetByCSPRNG()
		ecdsaKey, _ := crypto.GenerateKey()
		pk = append(pk, ecdsaKey)
		sk = append(sk, &blsKey)
	}
	return pk, sk
}

func createValidateNode(num int) ([]*pbfttypes.ValidateNode, []*bls.SecretKey) {
	pk, sk := GenerateKeys(num)
	nodes := make([]*pbfttypes.ValidateNode, num)
	for i := 0; i < num; i++ {

		nodes[i] = &pbfttypes.ValidateNode{
			Index:   uint32(i),
			Address: crypto.PubkeyToNodeAddress(pk[i].PublicKey),
			PubKey:  &pk[i].PublicKey,
			NodeID:  discover.PubkeyID(&pk[i].PublicKey),
		}
		nodes[i].BlsPubKey = sk[i].GetPublicKey()

	}
	return nodes, sk
}

func makePrepareBlock(epoch, viewNumber uint64, block *types.Block, blockIndex uint32, ProposalIndex uint32, t *testing.T, secretKeys *bls.SecretKey) *protocols.PrepareBlock {
	p := &protocols.PrepareBlock{
		Epoch:         epoch,
		ViewNumber:    viewNumber,
		Block:         block,
		BlockIndex:    blockIndex,
		ProposalIndex: ProposalIndex,
	}

	// bls sign
	buf, err := p.CannibalizeBytes()
	if err != nil {
		t.Fatalf("%s", "prepareBlock cannibalizeBytes error")
	}
	p.Signature.SetBytes(secretKeys.Sign(string(buf)).Serialize())
	//t.Logf("prepareBlock signature:%s", p.Signature.String())

	return p
}

func makePrepareVote(epoch, viewNumber uint64, blockHash common.Hash, blockNumber uint64, blockIndex uint32, validatorIndex uint32, t *testing.T, secretKeys *bls.SecretKey) *protocols.PrepareVote {
	p := &protocols.PrepareVote{
		Epoch:          epoch,
		ViewNumber:     viewNumber,
		BlockHash:      blockHash,
		BlockNumber:    blockNumber,
		BlockIndex:     blockIndex,
		ValidatorIndex: validatorIndex,
	}

	// bls sign
	buf, err := p.CannibalizeBytes()
	if err != nil {
		t.Fatalf("%s", "prepareVote cannibalizeBytes error")
	}
	p.Signature.SetBytes(secretKeys.Sign(string(buf)).Serialize())
	//t.Logf("prepareVote signature:%s", p.Signature.String())

	return p
}

func makePreCommit(epoch, viewNumber uint64, blockHash common.Hash, blockNumber uint64, blockIndex uint32, validatorIndex uint32, t *testing.T, secretKeys *bls.SecretKey) *protocols.PreCommit {
	p := &protocols.PreCommit{
		Epoch:          epoch,
		ViewNumber:     viewNumber,
		BlockHash:      blockHash,
		BlockNumber:    blockNumber,
		BlockIndex:     blockIndex,
		ValidatorIndex: validatorIndex,
	}

	// bls sign
	buf, err := p.CannibalizeBytes()
	if err != nil {
		t.Fatalf("%s", "prepareVote cannibalizeBytes error")
	}
	p.Signature.SetBytes(secretKeys.Sign(string(buf)).Serialize())
	//t.Logf("prepareVote signature:%s", p.Signature.String())

	return p
}

func makeViewChange(epoch, viewNumber uint64, blockHash common.Hash, blockNumber uint64, validatorIndex uint32, t *testing.T, secretKeys *bls.SecretKey) *protocols.ViewChange {
	p := &protocols.ViewChange{
		Epoch:          epoch,
		ViewNumber:     viewNumber,
		BlockHash:      blockHash,
		BlockNumber:    blockNumber,
		ValidatorIndex: validatorIndex,
	}

	// bls sign
	buf, err := p.CannibalizeBytes()
	if err != nil {
		t.Fatalf("%s", "viewChange cannibalizeBytes error")
	}
	p.Signature.SetBytes(secretKeys.Sign(string(buf)).Serialize())
	t.Logf("viewChange signature:%s", p.Signature.String())

	return p
}
