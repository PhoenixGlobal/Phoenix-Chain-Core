package configs

import (
	"fmt"
	"math/big"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/p2p/discover"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/crypto/bls"
)

// Genesis hashes to enforce below configs on.
var (
	MainnetGenesisHash = common.HexToHash("0x54656f18949b491f293a29584603e64318101bae6c5fb6da70c4258d0b08c55a")
	TestnetGenesisHash = common.HexToHash("0xb4d9b7badc9e28a3461f8b0786b233b23f165e35aa84c4a363998816f6d042b1")
)

var TrustedCheckpoints = map[common.Hash]*TrustedCheckpoint{
	MainnetGenesisHash: MainnetTrustedCheckpoint,
	TestnetGenesisHash: TestnetTrustedCheckpoint,
}

const ChainId = 908

var (
	initialMainNetConsensusNodes = []initNode{
		{
			"enode://6c719bccb655c76748c2e9a31402fa5112de24cf6ad4de3e03ff66a324c25913e28bc68418e1e88615eb696bb9088bc3b8aa14dda41dbe6f794919e645982691@39.104.68.32:16789",
			"0c257c93b61b51f7d45b77c50ed8bac0b95736622318aeb7634a1f77be82c5eb61eb3305d396e1739d1d576ba504ba1559f02202642803d772307262d56eabed12c785902c3f2c5d0badb59e98edc8638d4cb45421ceacd3c30aa5e14b5a1790",
		},
		{
			"enode://797e7c01fb7462d0a3eaf8dd8b982d7cb6363503521627d3f31a49b6389c357839e218d154388b11d9e10011f81ae3bbcc93e07d6f2063257d8015018494f3e1@39.104.68.32:16790",
			"d02a0dcc882879d3ee2ae82af8be442b234e9f4541f0d7a96c9199d2223599af77dd3125af0caf318cd1b0d11eb10f18f5b8fff79fedb680ecbdffd20840204103d4f1905414191ed0cc414ca0bc681cd9b178de3b37c9e9f2e29cbe29d7ce18",
		},
		{
			"enode://5db732679631792e4098f0c3cc32653acc394867dc4fbed0855f478adddfe758823571ec9b647a157957af84c90254e8c8d8a9a0848bea1e24cd4006d584cd3e@39.104.62.41:16791",
			"bb18a4dd5041b94d2367ad9e532c4c4da964bfc435a539a94180c8b7d9241cdbf2dc94fbb58b536560e0a3e30713c70873bb96aeab9a085047e4d6de5987f0790c95d9bc1abf7a6452cfaa6ca19564a511c99899678fcde1b84f5b2798099688",
		},
		{
			"enode://f6acf06029e09eba5c222fcaf9cbb55b178457b93d01b2bf92065c7dfdfa844f477fd4fdf5291372cb1ac8af89095bc68b51567d5eaed5d6fd5782b2ce0e64b6@39.104.62.41:16792",
			"e5182404de9160f49afb9ae7305220360e286981e572b7ce94934fb316caae0e414bfd7aff40ea419c63d40e0cb4e20da3119c77426798ec189b7e803cb4d9b6086b737ad468a4dbf33facb8ca382c28c60363923a30e2021d333ec8af8e6807",
		},
	}

	initialTestnetConsensusNodes = []initNode{
		{
			"enode://1ae3a2bfe9d93a4c96d44cabdda7d1ba2b5c653ab13d4593020add6238de1074b159f7eb072ca1a0ede3252fd44023aabda61a7f4fef16001bc9f83e1877296f@127.0.0.1:16789",
			"964211deb768dce2782600e0f9bbc55a1c5c6813b4ca262a0ba59fe389070950f2a7d25cb73e1e25dcc2262a12451714070e81b9fddd43ba1c0c456584338d1a2e5ec274a4cd4f8ef808174b323cadb2d2e06b64b08d590186b2efebc3c10b05",
		},
	}

	// MainnetChainConfig is the chain parameters to run a node on the main network.
	MainnetChainConfig = &ChainConfig{
		ChainID:     big.NewInt(ChainId),
		//AddressHRP:  "0x",
		EmptyBlock:  "on",
		EIP155Block: big.NewInt(1),
		Pbft: &PbftConfig{
			InitialNodes:  ConvertNodeUrl(initialMainNetConsensusNodes),
			Amount:        1,
			ValidatorMode: "dpos",
			Period:        2000,
		},
		GenesisVersion: GenesisVersion,
	}

	// MainnetTrustedCheckpoint contains the light client trusted checkpoint for the main network.
	MainnetTrustedCheckpoint = &TrustedCheckpoint{
		Name:         "mainnet",
		SectionIndex: 193,
		SectionHead:  common.HexToHash("0xc2d574295ecedc4d58530ae24c31a5a98be7d2b3327fba0dd0f4ed3913828a55"),
		CHTRoot:      common.HexToHash("0x5d1027dfae688c77376e842679ceada87fd94738feb9b32ef165473bfbbb317b"),
		BloomRoot:    common.HexToHash("0xd38be1a06aabd568e10957fee4fcc523bc64996bcf31bae3f55f86e0a583919f"),
	}

	// TestnetChainConfig is the chain parameters to run a node on the test network.
	TestnetChainConfig = &ChainConfig{
		ChainID:     big.NewInt(104),
		//AddressHRP:  "0x",
		EmptyBlock:  "on",
		EIP155Block: big.NewInt(1),
		Pbft: &PbftConfig{
			InitialNodes:  ConvertNodeUrl(initialTestnetConsensusNodes),
			Amount:        1,
			ValidatorMode: "dpos",
			Period:        20000,
		},
		GenesisVersion: GenesisVersion,
	}

	// TestnetTrustedCheckpoint contains the light client trusted checkpoint for the test network.
	TestnetTrustedCheckpoint = &TrustedCheckpoint{
		Name:         "testnet",
		SectionIndex: 123,
		SectionHead:  common.HexToHash("0xa372a53decb68ce453da12bea1c8ee7b568b276aa2aab94d9060aa7c81fc3dee"),
		CHTRoot:      common.HexToHash("0x6b02e7fada79cd2a80d4b3623df9c44384d6647fc127462e1c188ccd09ece87b"),
		BloomRoot:    common.HexToHash("0xf2d27490914968279d6377d42868928632573e823b5d1d4a944cba6009e16259"),
	}

	GrapeChainConfig = &ChainConfig{
		//AddressHRP:  "0x",
		ChainID:     big.NewInt(304),
		EmptyBlock:  "on",
		EIP155Block: big.NewInt(3),
		Pbft: &PbftConfig{
			Period: 3,
		},
		GenesisVersion: GenesisVersion,
	}

	// AllEthashProtocolChanges contains every protocol change (EIPs) introduced
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllEthashProtocolChanges = &ChainConfig{big.NewInt(1337), "", big.NewInt(0), big.NewInt(0), nil, nil, GenesisVersion}

	TestChainConfig = &ChainConfig{big.NewInt(1),  "", big.NewInt(0), big.NewInt(0), nil, new(PbftConfig), GenesisVersion}
)

// TrustedCheckpoint represents a set of post-processed trie roots (CHT and
// BloomTrie) associated with the appropriate section index and head hash. It is
// used to start light syncing from this checkpoint and avoid downloading the
// entire header chain while still being able to securely access old headers/logs.
type TrustedCheckpoint struct {
	Name         string      `json:"-"`
	SectionIndex uint64      `json:"sectionIndex"`
	SectionHead  common.Hash `json:"sectionHead"`
	CHTRoot      common.Hash `json:"chtRoot"`
	BloomRoot    common.Hash `json:"bloomRoot"`
}

// ChainConfig is the core config which determines the blockchain settings.
//
// ChainConfig is stored in the database on a per block basis. This means
// that any network, identified by its genesis block, can have its own
// set of configuration options.
type ChainConfig struct {
	ChainID     *big.Int `json:"chainId"` // chainId identifies the current chain and is used for replay protection
	//AddressHRP  string   `json:"addressHRP"`
	EmptyBlock  string   `json:"emptyBlock"`
	EIP155Block *big.Int `json:"eip155Block,omitempty"` // EIP155 HF block
	EWASMBlock  *big.Int `json:"ewasmBlock,omitempty"`  // EWASM switch block (nil = no fork, 0 = already activated)
	// Various consensus engines
	Clique *CliqueConfig `json:"clique,omitempty"`
	Pbft   *PbftConfig   `json:"pbft,omitempty"`

	GenesisVersion uint32 `json:"genesisVersion"`
}

type PbftNode struct {
	Node      discover.Node `json:"node"`
	BlsPubKey bls.PublicKey `json:"blsPubKey"`
}

type initNode struct {
	Enode     string
	BlsPubkey string
}

type Validator struct {
	NodeId          discover.NodeID
	BlsPubKey       bls.PublicKeyHex
}

type PbftConfig struct {
	Period        uint64     `json:"period,omitempty"`        // Number of seconds between blocks to enforce
	Amount        uint32     `json:"amount,omitempty"`        //The maximum number of blocks generated per cycle
	InitialNodes  []PbftNode `json:"initialNodes,omitempty"`  //Genesis consensus node
	ValidatorMode string     `json:"validatorMode,omitempty"` //Validator mode for easy testing
}

// CliqueConfig is the consensus engine configs for proof-of-authority based sealing.
type CliqueConfig struct {
	Period uint64 `json:"period"` // Number of seconds between blocks to enforce
	Epoch  uint64 `json:"epoch"`  // Epoch length to reset votes and checkpoint
}

// String implements the stringer interface, returning the consensus engine details.
func (c *CliqueConfig) String() string {
	return "clique"
}

// String implements the fmt.Stringer interface.
func (c *ChainConfig) String() string {
	var engine interface{}
	switch {
	case c.Clique != nil:
		engine = c.Clique
	case c.Pbft != nil:
		engine = c.Pbft
	default:
		engine = "unknown"
	}
	return fmt.Sprintf("{ChainID: %v EIP155: %v Engine: %v}",
		c.ChainID,
		c.EIP155Block,
		engine,
	)
}

// IsEIP155 returns whether num is either equal to the EIP155 fork block or greater.
func (c *ChainConfig) IsEIP155(num *big.Int) bool {
	//	return isForked(c.EIP155Block, num)
	return true
}

// IsEWASM returns whether num represents a block number after the EWASM fork
func (c *ChainConfig) IsEWASM(num *big.Int) bool {
	return isForked(c.EWASMBlock, num)
}

// GasTable returns the gas table corresponding to the current phase (homestead or homestead reprice).
//
// The returned GasTable's fields shouldn't, under any circumstances, be changed.
func (c *ChainConfig) GasTable(num *big.Int) GasTable {
	return GasTableConstantinople
}

// CheckCompatible checks whether scheduled fork transitions have been imported
// with a mismatching chain configuration.
func (c *ChainConfig) CheckCompatible(newcfg *ChainConfig, height uint64) *ConfigCompatError {
	bhead := new(big.Int).SetUint64(height)

	// Iterate checkCompatible to find the lowest conflict.
	var lasterr *ConfigCompatError
	for {
		err := c.checkCompatible(newcfg, bhead)
		if err == nil || (lasterr != nil && err.RewindTo == lasterr.RewindTo) {
			break
		}
		lasterr = err
		bhead.SetUint64(err.RewindTo)
	}
	return lasterr
}

func (c *ChainConfig) checkCompatible(newcfg *ChainConfig, head *big.Int) *ConfigCompatError {
	if isForkIncompatible(c.EIP155Block, newcfg.EIP155Block, head) {
		return newCompatError("EIP155 fork block", c.EIP155Block, newcfg.EIP155Block)
	}
	if isForkIncompatible(c.EWASMBlock, newcfg.EWASMBlock, head) {
		return newCompatError("ewasm fork block", c.EWASMBlock, newcfg.EWASMBlock)
	}
	return nil
}

// isForkIncompatible returns true if a fork scheduled at s1 cannot be rescheduled to
// block s2 because head is already past the fork.
func isForkIncompatible(s1, s2, head *big.Int) bool {
	return (isForked(s1, head) || isForked(s2, head)) && !configNumEqual(s1, s2)
}

// isForked returns whether a fork scheduled at block s is active at the given head block.
func isForked(s, head *big.Int) bool {
	if s == nil || head == nil {
		return false
	}
	return s.Cmp(head) <= 0
}

func configNumEqual(x, y *big.Int) bool {
	if x == nil {
		return y == nil
	}
	if y == nil {
		return x == nil
	}
	return x.Cmp(y) == 0
}

// ConfigCompatError is raised if the locally-stored blockchain is initialised with a
// ChainConfig that would alter the past.
type ConfigCompatError struct {
	What string
	// block numbers of the stored and new configurations
	StoredConfig, NewConfig *big.Int
	// the block number to which the local chain must be rewound to correct the error
	RewindTo uint64
}

func newCompatError(what string, storedblock, newblock *big.Int) *ConfigCompatError {
	var rew *big.Int
	switch {
	case storedblock == nil:
		rew = newblock
	case newblock == nil || storedblock.Cmp(newblock) < 0:
		rew = storedblock
	default:
		rew = newblock
	}
	err := &ConfigCompatError{what, storedblock, newblock, 0}
	if rew != nil && rew.Sign() > 0 {
		err.RewindTo = rew.Uint64() - 1
	}
	return err
}

func (err *ConfigCompatError) Error() string {
	return fmt.Sprintf("mismatching %s in database (have %d, want %d, rewindto %d)", err.What, err.StoredConfig, err.NewConfig, err.RewindTo)
}

func ConvertNodeUrl(initialNodes []initNode) []PbftNode {
	bls.Init(bls.BLS12_381)
	NodeList := make([]PbftNode, 0, len(initialNodes))
	for _, n := range initialNodes {

		pbftNode := new(PbftNode)

		if node, err := discover.ParseNode(n.Enode); nil == err {
			pbftNode.Node = *node
		}

		if n.BlsPubkey != "" {
			var blsPk bls.PublicKey
			if err := blsPk.UnmarshalText([]byte(n.BlsPubkey)); nil == err {
				pbftNode.BlsPubKey = blsPk
			}
		}

		NodeList = append(NodeList, *pbftNode)
	}
	return NodeList
}
