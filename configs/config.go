package configs

import (
	"fmt"
	"math/big"

	"Phoenix-Chain-Core/ethereum/p2p/discover"
	"Phoenix-Chain-Core/libs/common"
	"Phoenix-Chain-Core/libs/crypto/bls"
)

// Genesis hashes to enforce below configs on.
var (
	MainnetGenesisHash = common.HexToHash("0x3baa806f49c3e07c63f90271f21e0b3866e97dea5f33bf4eb258738884b96e08")
	TestnetGenesisHash = common.HexToHash("0xb03366a596c4c5e26b2a4b49c8f4710c3e89983f6abafc59eb0e7e214c29db38")
)

var TrustedCheckpoints = map[common.Hash]*TrustedCheckpoint{
	MainnetGenesisHash: MainnetTrustedCheckpoint,
	TestnetGenesisHash: TestnetTrustedCheckpoint,
}

var (
	initialMainNetConsensusNodes = []initNode{
		{
			"enode://f4d3989cfa7c30f09ebd24ab733ff97d5e10006fdba93af7c64e02e4cc9becec28d888badfeb9515a23a9c17497161e6bec03ea28113ecaa7adab8cb1b5990a0@www.phoenix.global:16789",
			"ca0c71e76b2ba6147467258238714ffe55ce5e57f3b50adb80f8d58aadc9f4eb7e338bb46d4ac8cdd4801625257d4d05c3aa0f446048b00ebb36eb94c3bc3dd18baa0e164dd150cc11bd708baef1e9b08c0496d6b498944b7eaa8cd685238c93",
		},
		{
			"enode://0116873bb440fb2aeadbb402e6093f5b97f0574cec02c53db6ff99cc5c2dc6dc4ee578961640ed80f9f5b69605493c3a88c8137eb06497596fe960b4c4a9053e@www.phoenix.global:16789",
			"e4f3774f6e043bc106503c291935047eed3f6b109bd3c29762fab6d14a00c775a315dac8ee2a82832476232efed4ff126e006ae4bccc9c708a615da105a1b3a0998dc2ea8a5f910e3280225b1bbc19c1132dfc661e4227fbf29bbb391706a589",
		},
		{
			"enode://8c816c9441f3b6d64efe2382110d17cd1ccd10a4a60db742383c36e6cabaeef2a2f87815ecd0f285e74ad6f9911a2db1dde7fe33a249e9c625ce225b0dbe25f0@www.phoenix.global:16789",
			"5c8e0280a5b72cb756437db1d62cf0803d4789a8cc3ff1989b26175cb63880d953c43352f6aa3e2c16c7cc933a669c1862a8011506262625a2a1ba7e79b141ce9ce54fa64c2d98073ecb8644accf63129229581eb318cc14d4f5e529171c2f95",
		},
		{
			"enode://33761aca567c3ce253635bfea65cb48d5518eaabfefd92e3d443414ea31a1abfff5dfda1f0c47f6b4cab0efd55a535a268acd82b3a9d078b40e328b890f49291@www.phoenix.global:16789",
			"660067d3c46fc7e41cd9b38e81937dbcb5a124524daefa3c9ade3330d84f92f19146e53c2c60abeb4a2044351077571866947234b4516e5200f925fc6f3d69b732bfc6a103f86cc2c25a9925b170adba04dd194864b90e9eb22c21196907830f",
		},
		{
			"enode://be5409650380b505bbb754663b0cc04f105fd792dcab394a78cd62be641f8ba976ff11bc814becfa23174b0852cc5f6db6d631a89d411c140d104c0927fb3822@www.phoenix.global:16789",
			"305a676d0b03d3f2c2018791bb4a7a855d4b2d226a09b9ed029dedec1a7774364dc8a8f55ba4fb8d39a40093d6bacb17732465002baf959913737c5f6e062704bb29b815ab5d8684ffa7de79a3e518fde7c7df34a18b92660a64f37eddfddb0d",
		},
		{
			"enode://39f455da5a67144c6f453bd3c375db85071c51c400c6b408a1361493f4ca1b169ac3e1ab76cfad8ba6612118c484e374516f0747facd9b3c22acd4fdd4c347b5@www.phoenix.global:16789",
			"ed62192c500c0fc4b1d7ad204b4eb88edb602a19e6767722e2e8aed44bd0eee4e8998e2407c654c2684cb975ab7c25019c09d5a519e32232aa532433af09753f8c3c3ee6d8ade73eca72fe5194f917a990bef42d11e9f417989e0c4d9db1a412",
		},
		{
			"enode://fb334dd1f0803291aebc242eeac3deb4f0132cbfa7c1d9c6acd1fb8277f9c83d3ed12cf9cca0a0f7d224d7761f0f93dc1c82c5540db8640bf0f21316795e0ccd@www.phoenix.global:16789",
			"4375213672319c43b731639e4ad70ed7b4ec8ae46289fdba02a5cc5b2813731f7f71fcb7e07263e58fc9e2badb9b780e01c90e10918620c01734429ef30a70df138edddb253f0efcde885339c7e16f9115bff7d24e33a7e283f45d7b4227d997",
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
		ChainID:     big.NewInt(100),
		//AddressHRP:  "0x",
		EmptyBlock:  "on",
		EIP155Block: big.NewInt(1),
		Pbft: &PbftConfig{
			InitialNodes:  ConvertNodeUrl(initialMainNetConsensusNodes),
			Amount:        10,
			ValidatorMode: "dpos",
			Period:        20000,
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
