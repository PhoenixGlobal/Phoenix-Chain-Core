package downloader

import (
	"fmt"

	"Phoenix-Chain-Core/ethereum/core/types"
)

// peerDropFn is a callback type for dropping a peer detected as malicious.
type peerDropFn func(id string)

// dataPack is a data message returned by a peer for some query.
type dataPack interface {
	PeerId() string
	Items() int
	Stats() string
}

// headerPack is a batch of block headers returned by a peer.
type headerPack struct {
	peerID  string
	headers []*types.Header
}

func (p *headerPack) PeerId() string { return p.peerID }
func (p *headerPack) Items() int     { return len(p.headers) }
func (p *headerPack) Stats() string  { return fmt.Sprintf("%d", len(p.headers)) }

// bodyPack is a batch of block bodies returned by a peer.
type bodyPack struct {
	peerID       string
	transactions [][]*types.Transaction
	extraData    [][]byte
}

func (p *bodyPack) PeerId() string { return p.peerID }
func (p *bodyPack) Items() int {
	return len(p.transactions)
}
func (p *bodyPack) Stats() string { return fmt.Sprintf("%d", len(p.transactions)) }

// receiptPack is a batch of receipts returned by a peer.
type receiptPack struct {
	peerID   string
	receipts [][]*types.Receipt
}

func (p *receiptPack) PeerId() string { return p.peerID }
func (p *receiptPack) Items() int     { return len(p.receipts) }
func (p *receiptPack) Stats() string  { return fmt.Sprintf("%d", len(p.receipts)) }

// statePack is a batch of states returned by a peer.
type statePack struct {
	peerID string
	states [][]byte
}

func (p *statePack) PeerId() string { return p.peerID }
func (p *statePack) Items() int     { return len(p.states) }
func (p *statePack) Stats() string  { return fmt.Sprintf("%d", len(p.states)) }

// dposStoragePack is a batch of dpos storage returned by a peer.
type dposStoragePack struct {
	peerID string
	kvs    []DPOSStorageKV
	last   bool
	kvNum  uint64
}

type DPOSStorageKV [2][]byte

func (p *dposStoragePack) PeerId() string { return p.peerID }
func (p *dposStoragePack) Items() int     { return len(p.kvs) }
func (p *dposStoragePack) Stats() string  { return fmt.Sprintf("%d", len(p.kvs)) }
func (p *dposStoragePack) KVs() [][2][]byte {
	var kv [][2][]byte
	for _, value := range p.kvs {
		kv = append(kv, value)
	}
	return kv
}

// dposStoragePack is a batch of dpos storage returned by a peer.
type dposInfoPack struct {
	peerID string
	latest *types.Header
	pivot  *types.Header
}

func (p *dposInfoPack) PeerId() string { return p.peerID }
func (p *dposInfoPack) Items() int     { return 1 }
func (p *dposInfoPack) Stats() string  { return fmt.Sprint(1) }
