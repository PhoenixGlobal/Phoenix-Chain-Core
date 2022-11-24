package testing

import (
	"fmt"
	"sync"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/log"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/p2p/discover"
)

type TestPeer interface {
	ID() discover.NodeID
	Drop(error)
}

// TestPeerPool is an example peerPool to demonstrate registration of peer connections
type TestPeerPool struct {
	lock  sync.Mutex
	peers map[discover.NodeID]TestPeer
}

func NewTestPeerPool() *TestPeerPool {
	return &TestPeerPool{peers: make(map[discover.NodeID]TestPeer)}
}

func (p *TestPeerPool) Add(peer TestPeer) {
	p.lock.Lock()
	defer p.lock.Unlock()
	log.Trace(fmt.Sprintf("pp add peer  %v", peer.ID()))
	p.peers[peer.ID()] = peer

}

func (p *TestPeerPool) Remove(peer TestPeer) {
	p.lock.Lock()
	defer p.lock.Unlock()
	delete(p.peers, peer.ID())
}

func (p *TestPeerPool) Has(id discover.NodeID) bool {
	p.lock.Lock()
	defer p.lock.Unlock()
	_, ok := p.peers[id]
	return ok
}

func (p *TestPeerPool) Get(id discover.NodeID) TestPeer {
	p.lock.Lock()
	defer p.lock.Unlock()
	return p.peers[id]
}
