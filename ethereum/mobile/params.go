// Contains all the wrappers from the params package.

package phoenixchain

import (
	"encoding/json"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/configs"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/p2p/discv5"
)

// MainnetGenesis returns the JSON spec to use for the main Ethereum network. It
// is actually empty since that defaults to the hard coded binary genesis block.
func MainnetGenesis() string {
	return ""
}

// TestnetGenesis returns the JSON spec to use for the Alpha test network.
func TestnetGenesis() string {
	enc, err := json.Marshal(core.DefaultTestnetGenesisBlock())
	if err != nil {
		panic(err)
	}
	return string(enc)
}

// FoundationBootnodes returns the enode URLs of the P2P bootstrap nodes operated
// by the foundation running the V5 discovery protocol.
func FoundationBootnodes() *Enodes {
	nodes := &Enodes{nodes: make([]*discv5.Node, len(configs.DiscoveryV5Bootnodes))}
	for i, url := range configs.DiscoveryV5Bootnodes {
		nodes.nodes[i] = discv5.MustParseNode(url)
	}
	return nodes
}
