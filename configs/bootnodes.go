package configs

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main PhoenixChain network.
var MainnetBootnodes = []string{
	"enode://3f900d32151e8e1e6126124e0f3408ea70aaebfd7e4c1fd9f8e905cbd4573b45ab6d0dbef624e1fef955b68954e73fcce81747225242056e3bcfd59a2eb977f8@39.104.68.32:16888",
	"enode://0ed45e0e1b1c0db14e0dbe592e57c74bee7baecd64a1abffd66b499e5122d382c2de5ddeb2b57fa7c6b05950ba3c69551992ebcaaa03bfe7433b3610549a7efb@39.104.62.41:16888",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the test network.
var TestnetBootnodes = []string{
	"enode://33761aca567c3ce253635bfea65cb48d5518eaabfefd92e3d443414ea31a1abfff5dfda1f0c47f6b4cab0efd55a535a268acd82b3a9d078b40e328b890f49291@127.0.0.1:16789",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
