// Contains all the wrappers from the node package to support client side node
// management on mobile platforms.

package phoenixchain

import (
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/eth"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/eth/downloader"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/ethclient"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/ethstats"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/internal/debug"
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/configs"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/les"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/node"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/p2p"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/p2p/nat"
	//	whisper "github.com/PhoenixGlobal/Phoenix-Chain-Core/whisper/whisperv6"
	//	whisper "github.com/PhoenixGlobal/Phoenix-Chain-Core/whisper/whisperv6"
)

// NodeConfig represents the collection of configuration values to fine tune the Geth
// node embedded into a mobile process. The available values are a subset of the
// entire API provided by go-ethereum to reduce the maintenance surface and dev
// complexity.
type NodeConfig struct {
	// Bootstrap nodes used to establish connectivity with the rest of the network.
	BootstrapNodes *Enodes

	// MaxPeers is the maximum number of peers that can be connected. If this is
	// set to zero, then only the configured static and trusted peers can connect.
	MaxPeers int

	// EthereumEnabled specifies whether the node should run the Ethereum protocol.
	EthereumEnabled bool

	// EthereumNetworkID is the network identifier used by the Ethereum protocol to
	// decide if remote peers should be accepted or not.
	EthereumNetworkID int64 // uint64 in truth, but Java can't handle that...

	// EthereumGenesis is the genesis JSON to use to seed the blockchain with. An
	// empty genesis state is equivalent to using the mainnet's state.
	EthereumGenesis string

	// EthereumDatabaseCache is the system memory in MB to allocate for database caching.
	// A minimum of 16MB is always reserved.
	EthereumDatabaseCache int

	// EthereumNetStats is a netstats connection string to use to report various
	// chain, transaction and node stats to a monitoring server.
	//
	// It has the form "nodename:secret@host:port"
	EthereumNetStats string

	// WhisperEnabled specifies whether the node should run the Whisper protocol.
	//	WhisperEnabled bool

	// Listening address of pprof server.
	PprofAddress string
}

// defaultNodeConfig contains the default node configuration values to use if all
// or some fields are missing from the user's specified list.
var defaultNodeConfig = &NodeConfig{
	BootstrapNodes:        FoundationBootnodes(),
	MaxPeers:              25,
	EthereumEnabled:       true,
	EthereumNetworkID:     1,
	EthereumDatabaseCache: 16,
}

// NewNodeConfig creates a new node option set, initialized to the default values.
func NewNodeConfig() *NodeConfig {
	config := *defaultNodeConfig
	return &config
}

// Node represents a Geth Ethereum node instance.
type Node struct {
	node *node.Node
}

// NewNode creates and configures a new Geth node.
func NewNode(datadir string, config *NodeConfig) (stack *Node, _ error) {
	// If no or partial configurations were specified, use defaults
	if config == nil {
		config = NewNodeConfig()
	}
	if config.MaxPeers == 0 {
		config.MaxPeers = defaultNodeConfig.MaxPeers
	}
	if config.BootstrapNodes == nil || config.BootstrapNodes.Size() == 0 {
		config.BootstrapNodes = defaultNodeConfig.BootstrapNodes
	}

	if config.PprofAddress != "" {
		debug.StartPProf(config.PprofAddress)
	}

	// Create the empty networking stack
	nodeConf := &node.Config{
		Name:        clientIdentifier,
		Version:     configs.VersionWithMeta,
		DataDir:     datadir,
		KeyStoreDir: filepath.Join(datadir, "keystore"), // Mobile should never use internal keystores!
		P2P: p2p.Config{
			NoDiscovery:      true,
			DiscoveryV5:      true,
			BootstrapNodesV5: config.BootstrapNodes.nodes,
			ListenAddr:       ":0",
			NAT:              nat.Any(),
			MaxPeers:         config.MaxPeers,
		},
	}
	rawStack, err := node.New(nodeConf)
	if err != nil {
		return nil, err
	}

	debug.Memsize.Add("node", rawStack)

	var genesis *core.Genesis
	if config.EthereumGenesis != "" {
		// Parse the user supplied genesis spec if not mainnet
		genesis = new(core.Genesis)
		if err := json.Unmarshal([]byte(config.EthereumGenesis), genesis); err != nil {
			return nil, fmt.Errorf("invalid genesis spec: %v", err)
		}
		// If we have the testnet, hard code the chain configs too
		if config.EthereumGenesis == TestnetGenesis() {
			genesis.Config = configs.TestnetChainConfig
			if config.EthereumNetworkID == 1 {
				config.EthereumNetworkID = 3
			}
		}
	}
	// Register the Ethereum protocol if requested
	if config.EthereumEnabled {
		ethConf := eth.DefaultConfig
		ethConf.Genesis = genesis
		ethConf.SyncMode = downloader.LightSync
		ethConf.NetworkId = uint64(config.EthereumNetworkID)
		ethConf.DatabaseCache = config.EthereumDatabaseCache
		if err := rawStack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
			return les.New(ctx, &ethConf)
		}); err != nil {
			return nil, fmt.Errorf("ethereum init: %v", err)
		}
		// If netstats reporting is requested, do it
		if config.EthereumNetStats != "" {
			if err := rawStack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
				var lesServ *les.LightEthereum
				ctx.Service(&lesServ)

				return ethstats.New(config.EthereumNetStats, nil, lesServ)
			}); err != nil {
				return nil, fmt.Errorf("netstats init: %v", err)
			}
		}
	}
	return &Node{rawStack}, nil
}

// Close terminates a running node along with all it's services, tearing internal
// state doen too. It's not possible to restart a closed node.
func (n *Node) Close() error {
	return n.node.Close()
}

// Start creates a live P2P node and starts running it.
func (n *Node) Start() error {
	return n.node.Start()
}

// Stop terminates a running node along with all it's services. If the node was
// not started, an error is returned.
func (n *Node) Stop() error {
	return n.node.Stop()
}

// GetEthereumClient retrieves a client to access the Ethereum subsystem.
func (n *Node) GetEthereumClient() (client *EthereumClient, _ error) {
	rpc, err := n.node.Attach()
	if err != nil {
		return nil, err
	}
	return &EthereumClient{ethclient.NewClient(rpc)}, nil
}

// GetNodeInfo gathers and returns a collection of metadata known about the host.
func (n *Node) GetNodeInfo() *NodeInfo {
	return &NodeInfo{n.node.Server().NodeInfo()}
}

// GetPeersInfo returns an array of metadata objects describing connected peers.
func (n *Node) GetPeersInfo() *PeerInfos {
	return &PeerInfos{n.node.Server().PeersInfo()}
}
