package pbft

import (
	"crypto/ecdsa"
	"math/big"
	"time"

	"Phoenix-Chain-Core/libs/ethdb"

	"Phoenix-Chain-Core/ethereum/core/db/rawdb"

	"Phoenix-Chain-Core/pos/xcom"

	cvm "Phoenix-Chain-Core/libs/common/vm"

	"Phoenix-Chain-Core/consensus/pbft/network"

	"Phoenix-Chain-Core/configs"
	"Phoenix-Chain-Core/consensus"
	ctypes "Phoenix-Chain-Core/consensus/pbft/types"
	"Phoenix-Chain-Core/consensus/pbft/validator"
	"Phoenix-Chain-Core/ethereum/core"
	"Phoenix-Chain-Core/ethereum/core/types"
	"Phoenix-Chain-Core/ethereum/core/vm"
	"Phoenix-Chain-Core/ethereum/node"
	"Phoenix-Chain-Core/ethereum/p2p/discover"
	"Phoenix-Chain-Core/libs/common"
	"Phoenix-Chain-Core/libs/common/hexutil"
	"Phoenix-Chain-Core/libs/crypto"
	"Phoenix-Chain-Core/libs/crypto/bls"
	"Phoenix-Chain-Core/libs/event"
)

var (
	testKey, _  = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	testAddress = crypto.PubkeyToAddress(testKey.PublicKey)

	chainConfig      = configs.TestnetChainConfig
	testTxPoolConfig = core.DefaultTxPoolConfig

	// twenty billion von
	//twoentyBillion, _ = new(big.Int).SetString("200000000000000000000000000000", 10)
	// two billion von
	twoBillion, _ = new(big.Int).SetString("20000000000000000000000000000", 10)
)

// NewBlock returns a new block for testing.
func NewBlock(parent common.Hash, number uint64) *types.Block {
	header := &types.Header{
		Number:      big.NewInt(int64(number)),
		ParentHash:  parent,
		Time:        uint64(time.Now().UnixNano() / 1e6),
		Extra:       make([]byte, 97),
		ReceiptHash: common.BytesToHash(hexutil.MustDecode("0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421")),
		Root:        common.BytesToHash(hexutil.MustDecode("0x3a3fed92c98ca419be7d4a125c30a4e8b3424e46c607b9cc2b6dfa5a81cbef0d")),
		Coinbase:    common.Address{},
		GasLimit:    10000000000,
	}

	block := types.NewBlockWithHeader(header)
	return block
}

// NewBlock returns a new block for testing.
func NewBlockWithSign(parent common.Hash, number uint64, node *TestPBFT) *types.Block {
	header := &types.Header{
		Number:      big.NewInt(int64(number)),
		ParentHash:  parent,
		Time:        uint64(time.Now().UnixNano() / 1e6),
		Extra:       make([]byte, 97),
		ReceiptHash: common.BytesToHash(hexutil.MustDecode("0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421")),
		Root:        common.BytesToHash(hexutil.MustDecode("0x3a3fed92c98ca419be7d4a125c30a4e8b3424e46c607b9cc2b6dfa5a81cbef0d")),
		Coinbase:    common.Address{},
		GasLimit:    10000000000,
	}

	sign, _ := node.engine.signFn(header.SealHash().Bytes())
	copy(header.Extra[len(header.Extra)-consensus.ExtraSeal:], sign[:])

	block := types.NewBlockWithHeader(header)
	return block
}

// GenerateKeys returns the public and private key pair for testing.
func GenerateKeys(num int) ([]*ecdsa.PrivateKey, []*bls.SecretKey) {
	pk := make([]*ecdsa.PrivateKey, 0)
	sk := make([]*bls.SecretKey, 0)

	for i := 0; i < num; i++ {
		var blsKey bls.SecretKey
		blsKey.SetByCSPRNG()
		ecdsaKey, _ := crypto.GenerateKey()
		pk = append(pk, ecdsaKey)
		sk = append(sk, &blsKey)
	}
	return pk, sk
}

// GeneratePbftNode returns a list of PbftNode for testing.
func GeneratePbftNode(num int) ([]*ecdsa.PrivateKey, []*bls.SecretKey, []configs.PbftNode) {
	pk, sk := GenerateKeys(num)
	nodes := make([]configs.PbftNode, num)
	for i := 0; i < num; i++ {

		nodes[i].Node = *discover.NewNode(discover.PubkeyID(&pk[i].PublicKey), nil, 0, 0)
		nodes[i].BlsPubKey = *sk[i].GetPublicKey()

	}
	return pk, sk, nodes
}

// CreatePBFT returns a new PBFT for testing.
func CreatePBFT(pk *ecdsa.PrivateKey, sk *bls.SecretKey, period uint64, amount uint32) *Pbft {

	sysConfig := &configs.PbftConfig{
		Period:       period,
		Amount:       amount,
		InitialNodes: []configs.PbftNode{},
	}

	optConfig := &ctypes.OptionsConfig{
		NodePriKey:        pk,
		NodeID:            discover.PubkeyID(&pk.PublicKey),
		BlsPriKey:         sk,
		MaxQueuesLimit:    1000,
		BlacklistDeadline: 1,
	}

	ctx := node.NewServiceContext(&node.Config{DataDir: ""}, nil, new(event.TypeMux), nil)

	return New(sysConfig, optConfig, ctx.EventMux, ctx)
}

func CreateGenesis(db ethdb.Database) (core.Genesis, *types.Block) {
	var (
		gspec = core.Genesis{
			Config: chainConfig,
			Alloc:  core.GenesisAlloc{},
		}
	)
	xcom.GetEc(xcom.DefaultUnitTestNet)
	gspec.Alloc[xcom.PhoenixChainFundAccount()] = core.GenesisAccount{
		Balance: xcom.PhoenixChainFundBalance(),
	}
	gspec.Alloc[cvm.RewardManagerPoolAddr] = core.GenesisAccount{
		Balance: twoBillion,
	}
	block := gspec.MustCommit(db)
	return gspec, block
}

// CreateBackend returns a new Backend for testing.
func CreateBackend(engine *Pbft, nodes []configs.PbftNode) (*core.BlockChain, *core.BlockChainCache, *core.TxPool, consensus.Agency) {

	var db = rawdb.NewMemoryDatabase()
	gspec, _ := CreateGenesis(db)

	chain, _ := core.NewBlockChain(db, nil, gspec.Config, engine, vm.Config{}, nil)
	cache := core.NewBlockChainCache(chain)
	txpool := core.NewTxPool(testTxPoolConfig, chainConfig, cache)

	return chain, cache, txpool, validator.NewStaticAgency(nodes)
}

// CreateValidatorBackend returns a new ValidatorBackend for testing.
func CreateValidatorBackend(engine *Pbft, nodes []configs.PbftNode) (*core.BlockChain, *core.BlockChainCache, *core.TxPool, consensus.Agency) {
	var (
		db    = rawdb.NewMemoryDatabase()
		gspec = core.Genesis{
			Config: chainConfig,
			Alloc:  core.GenesisAlloc{},
		}
	)
	balanceBytes, _ := hexutil.Decode("0x2000000000000000000000000000000000000000000000000000000000000")
	balance := big.NewInt(0)
	gspec.Alloc[testAddress] = core.GenesisAccount{
		Code:    nil,
		Storage: nil,
		Balance: balance.SetBytes(balanceBytes),
		Nonce:   0,
	}
	gspec.MustCommit(db)

	chain, _ := core.NewBlockChain(db, nil, gspec.Config, engine, vm.Config{}, nil)
	cache := core.NewBlockChainCache(chain)
	txpool := core.NewTxPool(testTxPoolConfig, chainConfig, cache)

	return chain, cache, txpool, validator.NewInnerAgency(nodes, chain, int(engine.config.Sys.Amount), int(engine.config.Sys.Amount)*2)
}

// TestPBFT for testing.
type TestPBFT struct {
	engine *Pbft
	chain  *core.BlockChain
	cache  *core.BlockChainCache
	txpool *core.TxPool
	agency consensus.Agency
}

// Start turns on the pbft for testing.
func (t *TestPBFT) Start() error {
	return t.engine.Start(t.chain, t.cache, t.txpool, t.agency)
}

// MockNode returns a new TestPBFT for testing.
func MockNode(pk *ecdsa.PrivateKey, sk *bls.SecretKey, nodes []configs.PbftNode, period uint64, amount uint32) *TestPBFT {
	engine := CreatePBFT(pk, sk, period, amount)

	chain, cache, txpool, agency := CreateBackend(engine, nodes)
	return &TestPBFT{
		engine: engine,
		chain:  chain,
		cache:  cache,
		txpool: txpool,
		agency: agency,
	}
}

// MockValidator returns a new TestPBFT for testing.
func MockValidator(pk *ecdsa.PrivateKey, sk *bls.SecretKey, nodes []configs.PbftNode, period uint64, amount uint32) *TestPBFT {
	engine := CreatePBFT(pk, sk, period, amount)

	chain, cache, txpool, agency := CreateValidatorBackend(engine, nodes)
	return &TestPBFT{
		engine: engine,
		chain:  chain,
		cache:  cache,
		txpool: txpool,
		agency: agency,
	}
}

// NewEngineManager returns a list of EngineManager and NodeID.
func NewEngineManager(pbfts []*TestPBFT) ([]*network.EngineManager, []discover.NodeID) {
	nodeids := make([]discover.NodeID, 0)
	engines := make([]*network.EngineManager, 0)
	for _, c := range pbfts {
		engines = append(engines, c.engine.network)
		nodeids = append(nodeids, c.engine.config.Option.NodeID)
	}
	return engines, nodeids
}

// Mock4NodePipe returns a list of TestPBFT for testing.
func Mock4NodePipe(start bool) []*TestPBFT {
	pk, sk, pbftnodes := GeneratePbftNode(4)
	nodes := make([]*TestPBFT, 0)
	for i := 0; i < 4; i++ {
		node := MockNode(pk[i], sk[i], pbftnodes, 20000, 1)

		nodes = append(nodes, node)
		//fmt.Println(i, node.engine.config.Option.NodeID.TerminalString())
		nodes[i].Start()
	}

	netHandler, nodeids := NewEngineManager(nodes)

	network.EnhanceEngineManager(nodeids, netHandler)
	if start {
		for i := 0; i < 4; i++ {
			netHandler[i].Testing()
		}
	}
	return nodes
}
