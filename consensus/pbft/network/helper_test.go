package network

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"testing"
	"time"

	"Phoenix-Chain-Core/consensus/pbft/protocols"
	types2 "Phoenix-Chain-Core/consensus/pbft/types"
	"Phoenix-Chain-Core/libs/log"

	"Phoenix-Chain-Core/consensus/pbft/utils"
	"Phoenix-Chain-Core/ethereum/core/types"
	"Phoenix-Chain-Core/ethereum/p2p/discover"
	"Phoenix-Chain-Core/libs/common"

	"Phoenix-Chain-Core/ethereum/p2p"
)

var loglevel = flag.Int("loglevel", 4, "verbosity of logs")

// Create a new PrepareBlock for testing.
func newFakePrepareBlock() *protocols.PrepareBlock {
	block := types.NewBlockWithHeader(&types.Header{
		GasLimit: uint64(3141592),
		GasUsed:  uint64(21000),
		Coinbase: common.MustStringToAddress("0x3zy0ruv447se9nlwscrfskzvqv85e8d35gau40"),
		Root:     common.HexToHash("ef1552a40b7165c3cd773806b9e0c165b75356e0314bf0706f279c729f51e017"),
		Nonce:    types.EncodeNonce(utils.RandBytes(81)),
		Time:     1426516743,
		Extra:    make([]byte, 100),
	})
	return &protocols.PrepareBlock{
		Epoch:        1,
		ViewNumber:   1,
		Block:        block,
		BlockIndex:   1,
		PrepareQC:    newQuorumCert(),
		ViewChangeQC: &types2.ViewChangeQC{},
		Signature:    newSignature(),
	}
}

// Create a new PrepareVote for testing.
func newFakePrepareVote() *protocols.PrepareVote {
	return &protocols.PrepareVote{
		Epoch:       1,
		ViewNumber:  1,
		BlockHash:   common.BytesToHash([]byte("I'm hash")),
		BlockNumber: 1,
		BlockIndex:  1,
		ParentQC:    newQuorumCert(),
		Signature:   newSignature(),
	}
}

func newQuorumCert() *types2.QuorumCert {
	return &types2.QuorumCert{
		Epoch:        1,
		ViewNumber:   1,
		BlockHash:    common.Hash{},
		BlockNumber:  1,
		BlockIndex:   1,
		Signature:    newSignature(),
		ValidatorSet: utils.NewBitArray(32),
	}
}

func newSignature() types2.Signature {
	return types2.Signature{}
}

// Create a new ViewChange for testing.
func newFakeViewChange() *protocols.ViewChange {
	return &protocols.ViewChange{
		Epoch:       1,
		ViewNumber:  1,
		BlockHash:   common.BytesToHash([]byte("I'm hash of viewChange")),
		BlockNumber: 1,
		PrepareQC:   newQuorumCert(),
		Signature:   newSignature(),
	}
}

// Create a new GetPrepareBlock for testing.
func newFakeGetPrepareBlock() *protocols.GetPrepareBlock {
	return &protocols.GetPrepareBlock{
		Epoch:      1,
		ViewNumber: 1,
		BlockIndex: 1,
	}
}

// Create a new GetBlockQuorumCert for testing.
func newFakeGetBlockQuorumCert() *protocols.GetBlockQuorumCert {
	return &protocols.GetBlockQuorumCert{
		BlockHash:   common.BytesToHash([]byte("GetBlockQuorumCert")),
		BlockNumber: 1,
	}
}

// Create a new BlockQuorumCert for testing.
func newFakeBlockQuorumCert() *protocols.BlockQuorumCert {
	return &protocols.BlockQuorumCert{
		BlockQC: newQuorumCert(),
	}
}

// Create a new GetQCBlockList for testing.
func newFakeGetQCBlockList() *protocols.GetQCBlockList {
	return &protocols.GetQCBlockList{
		BlockHash:   common.BytesToHash([]byte("GetQCBlockList")),
		BlockNumber: 1,
	}
}

// Create a new GetPrepareVote for testing.
func newFakeGetPrepareVote() *protocols.GetPrepareVote {
	return &protocols.GetPrepareVote{
		Epoch:      1,
		ViewNumber: 1,
		BlockIndex: 1,
		UnKnownSet: utils.NewBitArray(32),
	}
}

// Create a new PrepareVotes for testing.
func newFakePrepareVotes() *protocols.PrepareVotes {
	return &protocols.PrepareVotes{
		Epoch:      1,
		ViewNumber: 1,
		BlockIndex: 1,
		Votes:      []*protocols.PrepareVote{newFakePrepareVote()},
	}
}

// Create a new PrepareBlockHash for testing.
func newFakePrepareBlockHash() *protocols.PrepareBlockHash {
	return &protocols.PrepareBlockHash{
		BlockHash:   common.BytesToHash([]byte("PrepareBlockHash")),
		BlockNumber: 1,
	}
}

// Create a new QCBlockList for testing.
func newFakeQCBlockList() *protocols.QCBlockList {
	block := types.NewBlockWithHeader(&types.Header{
		GasLimit: uint64(3141592),
		GasUsed:  uint64(21000),
		Coinbase: common.MustStringToAddress("0x3zy0ruv447se9nlwscrfskzvqv85e8d35gau40"),
		Root:     common.HexToHash("ef1552a40b7165c3cd773806b9e0c165b75356e0314bf0706f279c729f51e017"),
		Nonce:    types.EncodeNonce(utils.RandBytes(81)),
		Time:     1426516743,
		Extra:    make([]byte, 100),
	})
	return &protocols.QCBlockList{
		QC:     []*types2.QuorumCert{newQuorumCert()},
		Blocks: []*types.Block{block},
	}
}

// Create a new GetLatestStatus for testing.
func newFakeGetLatestStatus() *protocols.GetLatestStatus {
	return &protocols.GetLatestStatus{
		BlockNumber: 1,
		LogicType:   TypeForQCBn,
	}
}

// Create a new LatestStatus for testing.
func newFakeLatestStatus() *protocols.LatestStatus {
	return &protocols.LatestStatus{
		BlockNumber:  1,
		BlockHash:    common.BytesToHash([]byte("qc")),
		LBlockNumber: 0,
		LBlockHash:   common.BytesToHash([]byte("lock")),
		LogicType:    TypeForQCBn,
	}
}

// Create a new PbftStatusData for testing.
func newFakePbftStatusData() *protocols.PbftStatusData {
	return &protocols.PbftStatusData{
		ProtocolVersion: 1,
		QCBn:            big.NewInt(1),
		QCBlock:         common.Hash{},
		LockBn:          big.NewInt(2),
		LockBlock:       common.Hash{},
		CmtBn:           big.NewInt(3),
		CmtBlock:        common.Hash{},
	}
}

// Create a new Ping for testing.
func newFakePing(pingTime string) *protocols.Ping {
	return &protocols.Ping{
		pingTime,
	}
}

// Create a new Pong for testing.
func newFakePong(pingTime string) *protocols.Pong {
	//pingTime := strconv.FormatInt(time.Now().UnixNano(), 10)
	return &protocols.Pong{
		pingTime,
	}
}

// fakePeer is a simulated peer to allow testing direct network calls.
type fakePeer struct {
	net   p2p.MsgReadWriter // Network layer reader/writer to simulate remote messaging.
	app   *p2p.MsgPipeRW    // Application layer reader/writer to simulate the local side.
	*peer                   // The peer belonging to PBFT layer.
}

// newFakePeer creates a new peer registered at the given protocol manager.
func newFakePeer(name string, version int, pm *EngineManager, shake bool) (*fakePeer, <-chan error) {
	// Create a message pipe to communicate through.
	app, net := p2p.MsgPipe()

	// Generate a random id and create the peer.
	var id discover.NodeID
	rand.Read(id[:])

	// Create a peer that belonging to pbft.
	peer := newPeer(version, p2p.NewPeer(id, name, nil), net)

	// Start the peer on a new thread
	errc := make(chan error, 1)
	go func() {
		//
		errc <- pm.handler(peer.Peer, peer.rw)
	}()
	tp := &fakePeer{app: app, net: net, peer: peer}
	return tp, errc
}

// Create a new peer for testing, return peer and ID.
func newTestPeer(version int, name string) (*peer, discover.NodeID) {
	_, net := p2p.MsgPipe()

	// Generate a random id and create the peer.
	var id discover.NodeID
	rand.Read(id[:])

	// Create a peer that belonging to pbft.
	peer := newPeer(version, p2p.NewPeer(id, name, nil), net)
	go peer.sendLoop()
	return peer, id
}

func newLinkedPeer(rw p2p.MsgReadWriter, version int, name string) (*peer, discover.NodeID) {
	// Generate a random id and create the peer.
	var id discover.NodeID
	rand.Read(id[:])

	// Create a peer that belonging to pbft.
	peer := newPeer(version, p2p.NewPeer(id, name, nil), rw)
	go peer.sendLoop()
	return peer, id
}

func Test_InitializePeers(t *testing.T) {

	// Randomly generated ID.
	nodeIds := RandomID()

	// init pbft
	pbft1 := &mockPbft{nodeIds, nodeIds[0]}
	pbft2 := &mockPbft{nodeIds, nodeIds[1]}
	pbft3 := &mockPbft{nodeIds, nodeIds[2]}
	pbft4 := &mockPbft{nodeIds, nodeIds[3]}

	// init handler
	h1 := NewEngineManger(pbft1)
	h2 := NewEngineManger(pbft2)
	h3 := NewEngineManger(pbft3)
	h4 := NewEngineManger(pbft4)

	// register
	//initializeHandler(peers, []*EngineManager{h1, h2, h3, h4})
	EnhanceEngineManager(nodeIds, []*EngineManager{h1, h2, h3, h4})

	// start handler.
	h1.Start()
	h2.Start()
	h4.Start()
	h3.Start()

	h1.Testing()
	h2.Testing()
	h3.Testing()
	h4.Testing()

	// Pretend to send data to p1.p2.p3
	time.Sleep(1 * time.Second)
	h1.Broadcast(newFakePrepareBlockHash())
	time.Sleep(1 * time.Second)
	h2.Broadcast(newFakeViewChange())
	time.Sleep(1 * time.Second)
	h3.Broadcast(newFakeGetPrepareVote())
	time.Sleep(3 * time.Second)
}

type mockPbft struct {
	consensusNodes []discover.NodeID
	peerID         discover.NodeID
}

func (s *mockPbft) NodeID() discover.NodeID {
	return s.peerID
}

func (s *mockPbft) ConsensusNodes() ([]discover.NodeID, error) {
	return s.consensusNodes, nil
}

func (s *mockPbft) Config() *types2.Config {
	return nil
}

func (s *mockPbft) ReceiveMessage(msg *types2.MsgInfo) error {
	log.Debug("ReceiveMessage", "from", msg.PeerID, "type", fmt.Sprintf("%T", msg.Msg))
	return nil
}

func (s *mockPbft) ReceiveSyncMsg(msg *types2.MsgInfo) error {
	log.Debug("ReceiveSyncMsg", "from", msg.PeerID, "type", fmt.Sprintf("%T", msg.Msg))
	return nil
}

func (s *mockPbft) HighestQCBlockBn() (uint64, common.Hash) {
	return 0, common.Hash{}
}

func (s *mockPbft) HighestLockBlockBn() (uint64, common.Hash) {
	return 0, common.Hash{}
}

func (s *mockPbft) HighestCommitBlockBn() (uint64, common.Hash) {
	return 0, common.Hash{}
}

func (s *mockPbft) MissingViewChangeNodes() (*protocols.GetViewChange, error) {
	return &protocols.GetViewChange{
		Epoch:      1,
		ViewNumber: 1,
	}, nil
}

func (s *mockPbft) MissingPrepareVote() (*protocols.GetPrepareVote, error) {
	return &protocols.GetPrepareVote{
		Epoch:      1,
		ViewNumber: 1,
	}, nil
}

func (s *mockPbft) MissingPreCommit() (*protocols.GetPreCommit, error) {
	return &protocols.GetPreCommit{
		Epoch:      1,
		ViewNumber: 1,
	}, nil
}

func (s *mockPbft) LatestStatus() *protocols.GetLatestStatus {
	qcNumber, qcHash := s.HighestQCBlockBn()
	lockNumber, lockHash := s.HighestLockBlockBn()
	return &protocols.GetLatestStatus{
		BlockNumber:  qcNumber,
		BlockHash:    qcHash,
		QuorumCert:   nil,
		LBlockNumber: lockNumber,
		LBlockHash:   lockHash,
		LQuorumCert:  nil,
		LogicType:    TypeForQCBn,
	}
}

func (s *mockPbft) OnPong(nodeID string, netLatency int64) error {
	return nil
}

func (s *mockPbft) BlockExists(blockNumber uint64, blockHash common.Hash) error {
	return nil
}
