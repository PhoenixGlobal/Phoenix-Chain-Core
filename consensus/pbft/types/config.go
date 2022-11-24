package types

import (
	"crypto/ecdsa"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/configs"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/p2p/discover"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/crypto/bls"
)

type OptionsConfig struct {
	NodePriKey *ecdsa.PrivateKey `json:"-"`
	NodeID     discover.NodeID   `json:"nodeID"`
	BlsPriKey  *bls.SecretKey    `json:"-"`
	WalMode    bool              `json:"walMode"`

	PeerMsgQueueSize  uint64 `json:"peerMsgQueueSize"`
	EvidenceDir       string `json:"evidenceDir"`
	MaxPingLatency    int64  `json:"maxPingLatency"`    // maxPingLatency is the time in milliseconds between Ping and Pong
	MaxQueuesLimit    int64  `json:"maxQueuesLimit"`    // The maximum value that a single node can send a message.
	BlacklistDeadline int64  `json:"blacklistDeadline"` // Blacklist expiration time. unit: minute.

	Period uint64 `json:"period"`
	Amount uint32 `json:"amount"`
}

type Config struct {
	Sys    *configs.PbftConfig `json:"sys"`
	Option *OptionsConfig      `json:"option"`
}
