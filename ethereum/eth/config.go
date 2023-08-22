package eth

import (
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/eth/downloader"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/eth/gasprice"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/miner"
	"math/big"
	"time"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/configs"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/consensus/pbft/types"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core"
)

// DefaultConfig contains default settings for use on the Ethereum main net.
var DefaultConfig = Config{
	SyncMode: downloader.FullSync,
	PbftConfig: types.OptionsConfig{
		WalMode:           true,
		PeerMsgQueueSize:  1024,
		EvidenceDir:       "evidence",
		MaxPingLatency:    5000,
		MaxQueuesLimit:    4096,
		BlacklistDeadline: 10,
		Period:            2000,
		Amount:            1,
	},
	NetworkId:     1,
	LightPeers:    100,
	DatabaseCache: 768,
	TrieCache:     32,
	TrieTimeout:   60 * time.Minute,
	TrieDBCache:   512,
	DBDisabledGC:      false,
	DBGCInterval:      86400,
	DBGCTimeout:       time.Minute,
	DBGCMpt:           true,
	DBGCBlock:         10,
	VMWasmType:        "wagon",
	VmTimeoutDuration: 0, // default 0 ms for vm exec timeout
	Miner: miner.Config{
		GasFloor: configs.GenesisGasLimit,
		GasPrice: big.NewInt(configs.GVon),
		Recommit: 3 * time.Second,
	},

	MiningLogAtDepth:       7,
	TxChanSize:             4096,
	ChainHeadChanSize:      10,
	ChainSideChanSize:      10,
	ResultQueueSize:        10,
	ResubmitAdjustChanSize: 10,
	MinRecommitInterval:    1 * time.Second,
	MaxRecommitInterval:    15 * time.Second,
	IntervalAdjustRatio:    0.1,
	IntervalAdjustBias:     200 * 1000.0 * 1000.0,
	StaleThreshold:         7,
	DefaultCommitRatio:     0.95,

	BodyCacheLimit:    256,
	BlockCacheLimit:   256,
	MaxFutureBlocks:   256,
	BadBlockLimit:     10,
	TriesInMemory:     128,
	BlockChainVersion: 3,

	TxPool: core.DefaultTxPoolConfig,
	GPO: gasprice.Config{
		Blocks:     20,
		Percentile: 60,
	},

	//MPCPool: core.DefaultMPCPoolConfig,
	//VCPool:  core.DefaultVCPoolConfig,
}

//go:generate gencodec -type Config -formats toml -out gen_config.go

type Config struct {
	// The genesis block, which is inserted if the database is empty.
	// If nil, the Ethereum main net block is used.
	Genesis *core.Genesis `toml:",omitempty"`

	PbftConfig types.OptionsConfig `toml:",omitempty"`

	// Protocol options
	NetworkId uint64 // Network ID to use for selecting peers to connect to
	SyncMode  downloader.SyncMode
	NoPruning bool

	// Light client options
	LightServ  int `toml:",omitempty"` // Maximum percentage of time allowed for serving LES requests
	LightPeers int `toml:",omitempty"` // Maximum number of LES client peers

	// Database options
	SkipBcVersionCheck bool `toml:"-"`
	DatabaseHandles    int  `toml:"-"`
	DatabaseCache      int
	DatabaseFreezer    string

	TrieCache    int
	TrieTimeout  time.Duration
	TrieDBCache  int
	DBDisabledGC bool
	DBGCInterval uint64
	DBGCTimeout  time.Duration
	DBGCMpt      bool
	DBGCBlock    int

	// VM options
	VMWasmType        string
	VmTimeoutDuration uint64

	// Mining options
	Miner	miner.Config
	// minning conig
	MiningLogAtDepth       uint          // miningLogAtDepth is the number of confirmations before logging successful mining.
	TxChanSize             int           // txChanSize is the size of channel listening to NewTxsEvent.The number is referenced from the size of tx pool.
	ChainHeadChanSize      int           // chainHeadChanSize is the size of channel listening to ChainHeadEvent.
	ChainSideChanSize      int           // chainSideChanSize is the size of channel listening to ChainSideEvent.
	ResultQueueSize        int           // resultQueueSize is the size of channel listening to sealing result.
	ResubmitAdjustChanSize int           // resubmitAdjustChanSize is the size of resubmitting interval adjustment channel.
	MinRecommitInterval    time.Duration // minRecommitInterval is the minimal time interval to recreate the mining block with any newly arrived transactions.
	MaxRecommitInterval    time.Duration // maxRecommitInterval is the maximum time interval to recreate the mining block with any newly arrived transactions.
	IntervalAdjustRatio    float64       // intervalAdjustRatio is the impact a single interval adjustment has on sealing work resubmitting interval.
	IntervalAdjustBias     float64       // intervalAdjustBias is applied during the new resubmit interval calculation in favor of increasing upper limit or decreasing lower limit so that the limit can be reachable.
	StaleThreshold         uint64        // staleThreshold is the maximum depth of the acceptable stale block.
	DefaultCommitRatio     float64

	// block config
	BodyCacheLimit           int
	BlockCacheLimit          int
	MaxFutureBlocks          int
	BadBlockLimit            int
	TriesInMemory            int
	BlockChainVersion        int // BlockChainVersion ensures that an incompatible database forces a resync from scratch.
	DefaultTxsCacheSize      int
	DefaultBroadcastInterval time.Duration

	// Transaction pool options
	TxPool core.TxPoolConfig

	// Gas Price Oracle options
	GPO gasprice.Config

	// Miscellaneous options
	DocRoot string `toml:"-"`

	// MPC pool options
	//MPCPool core.MPCPoolConfig
	//VCPool  core.VCPoolConfig
	Debug bool

	// RPCGasCap is the global gas cap for eth-call variants.
	RPCGasCap *big.Int `toml:",omitempty"`
}
