package configs

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main PhoenixChain network.
var MainnetBootnodes = []string{
	"enode://0c88323c5dab6e2b5bac92b54fe9e632e7b8cc5f0345433ac32b188980142d691320d0215c665b61e9de94982a0ddfb64f0407d5a9b4d93e97eae3e2f3704960@39.104.61.131:16888",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the test network.
var TestnetBootnodes = []string{
	"enode://33761aca567c3ce253635bfea65cb48d5518eaabfefd92e3d443414ea31a1abfff5dfda1f0c47f6b4cab0efd55a535a268acd82b3a9d078b40e328b890f49291@127.0.0.1:16789",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
