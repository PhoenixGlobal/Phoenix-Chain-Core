// Package les implements the Light Ethereum Subprotocol.
package les

import (
	eth2 "github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/eth"
	downloader2 "github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/eth/downloader"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/eth/filters"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/eth/gasprice"
	ethapi2 "github.com/PhoenixGlobal/Phoenix-Chain-Core/internal/ethapi"
	"sync"
	"time"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core/db/snapshotdb"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/configs"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/consensus"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core/bloombits"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core/db/rawdb"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core/types"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/accounts"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/light"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/node"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/p2p"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/p2p/discv5"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/event"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/log"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/rpc"
)

type LightEthereum struct {
	lesCommons

	odr         *LesOdr
	relay       *LesTxRelay
	chainConfig *configs.ChainConfig
	// Channel for shutting down the service
	shutdownChan chan bool

	// Handlers
	peers      *peerSet
	txPool     *light.TxPool
	blockchain *light.LightChain
	serverPool *serverPool
	reqDist    *requestDistributor
	retriever  *retrieveManager

	bloomRequests chan chan *bloombits.Retrieval // Channel receiving bloom data retrieval requests
	bloomIndexer  *core.ChainIndexer

	ApiBackend *LesApiBackend

	eventMux       *event.TypeMux
	engine         consensus.Engine
	accountManager *accounts.Manager

	networkId     uint64
	netRPCService *ethapi2.PublicNetAPI

	wg sync.WaitGroup
}

func New(ctx *node.ServiceContext, config *eth2.Config) (*LightEthereum, error) {
	chainDb, err := ctx.OpenDatabase("lightchaindata", config.DatabaseCache, config.DatabaseHandles, "eth/db/chaindata/")
	if err != nil {
		return nil, err
	}
	basedb, err := snapshotdb.Open(ctx.ResolvePath(snapshotdb.DBPath), 0, 0, true)
	if err != nil {
		return nil, err
	}
	chainConfig, genesisHash, genesisErr := core.SetupGenesisBlock(chainDb, basedb, config.Genesis)
	if _, isCompat := genesisErr.(*configs.ConfigCompatError); genesisErr != nil && !isCompat {
		return nil, genesisErr
	}
	if err := basedb.Close(); err != nil {
		return nil, err
	}
	log.Info("Initialised chain configuration", "config", chainConfig)

	peers := newPeerSet()
	quitSync := make(chan struct{})

	leth := &LightEthereum{
		lesCommons: lesCommons{
			chainDb: chainDb,
			config:  config,
			iConfig: light.DefaultClientIndexerConfig,
		},
		chainConfig:    chainConfig,
		eventMux:       ctx.EventMux,
		peers:          peers,
		reqDist:        newRequestDistributor(peers, quitSync),
		accountManager: ctx.AccountManager,
		engine:         eth2.CreateConsensusEngine(ctx, chainConfig, false, chainDb, &config.PbftConfig, ctx.EventMux),
		shutdownChan:   make(chan bool),
		networkId:      config.NetworkId,
		bloomRequests:  make(chan chan *bloombits.Retrieval),
		bloomIndexer:   eth2.NewBloomIndexer(chainDb, configs.BloomBitsBlocksClient, configs.HelperTrieConfirmations),
	}

	leth.relay = NewLesTxRelay(peers, leth.reqDist)
	leth.serverPool = newServerPool(chainDb, quitSync, &leth.wg)
	leth.retriever = newRetrieveManager(peers, leth.reqDist, leth.serverPool)

	leth.odr = NewLesOdr(chainDb, light.DefaultClientIndexerConfig, leth.retriever)
	leth.chtIndexer = light.NewChtIndexer(chainDb, leth.odr, configs.CHTFrequency, configs.HelperTrieConfirmations)
	leth.bloomTrieIndexer = light.NewBloomTrieIndexer(chainDb, leth.odr, configs.BloomBitsBlocksClient, configs.BloomTrieFrequency)
	leth.odr.SetIndexers(leth.chtIndexer, leth.bloomTrieIndexer, leth.bloomIndexer)

	// Note: NewLightChain adds the trusted checkpoint so it needs an ODR with
	// indexers already set but not started yet
	if leth.blockchain, err = light.NewLightChain(leth.odr, leth.chainConfig, leth.engine); err != nil {
		return nil, err
	}
	// Note: AddChildIndexer starts the update process for the child
	leth.bloomIndexer.AddChildIndexer(leth.bloomTrieIndexer)
	leth.chtIndexer.Start(leth.blockchain)
	leth.bloomIndexer.Start(leth.blockchain)

	// Rewind the chain in case of an incompatible config upgrade.
	if compat, ok := genesisErr.(*configs.ConfigCompatError); ok {
		log.Warn("Rewinding chain to upgrade configuration", "err", compat)
		leth.blockchain.SetHead(compat.RewindTo)
		rawdb.WriteChainConfig(chainDb, genesisHash, chainConfig)
	}

	leth.txPool = light.NewTxPool(leth.chainConfig, leth.blockchain, leth.relay)
	if leth.protocolManager, err = NewProtocolManager(leth.chainConfig, light.DefaultClientIndexerConfig, true, config.NetworkId, leth.eventMux, leth.engine, leth.peers, leth.blockchain, nil, chainDb, leth.odr, leth.relay, leth.serverPool, quitSync, &leth.wg); err != nil {
		return nil, err
	}
	leth.ApiBackend = &LesApiBackend{ctx.ExtRPCEnabled(), leth, nil}
	gpoParams := config.GPO
	if gpoParams.Default == nil {
		gpoParams.Default = config.Miner.GasPrice
	}
	leth.ApiBackend.gpo = gasprice.NewOracle(leth.ApiBackend, gpoParams)
	return leth, nil
}

func lesTopic(genesisHash common.Hash, protocolVersion uint) discv5.Topic {
	var name string
	switch protocolVersion {
	case lpv1:
		name = "LES"
	case lpv2:
		name = "LES2"
	default:
		panic(nil)
	}
	return discv5.Topic(name + "@" + common.Bytes2Hex(genesisHash.Bytes()[0:8]))
}

// APIs returns the collection of RPC services the ethereum package offers.
// NOTE, some of these services probably need to be moved to somewhere else.
func (s *LightEthereum) APIs() []rpc.API {
	return append(ethapi2.GetAPIs(s.ApiBackend), []rpc.API{
		{
			Namespace: "phoenixchain",
			Version:   "1.0",
			Service:   downloader2.NewPublicDownloaderAPI(s.protocolManager.downloader, s.eventMux),
			Public:    true,
		}, {
			Namespace: "phoenixchain",
			Version:   "1.0",
			Service:   filters.NewPublicFilterAPI(s.ApiBackend, true),
			Public:    true,
		}, {
			Namespace: "net",
			Version:   "1.0",
			Service:   s.netRPCService,
			Public:    true,
		},
	}...)
}

func (s *LightEthereum) ResetWithGenesisBlock(gb *types.Block) {
	s.blockchain.ResetWithGenesisBlock(gb)
}

func (s *LightEthereum) BlockChain() *light.LightChain      { return s.blockchain }
func (s *LightEthereum) TxPool() *light.TxPool              { return s.txPool }
func (s *LightEthereum) Engine() consensus.Engine           { return s.engine }
func (s *LightEthereum) LesVersion() int                     { return int(ClientProtocolVersions[0]) }
func (s *LightEthereum) Downloader() *downloader2.Downloader { return s.protocolManager.downloader }
func (s *LightEthereum) EventMux() *event.TypeMux            { return s.eventMux }

// Protocols implements node.Service, returning all the currently configured
// network protocols to start.
func (s *LightEthereum) Protocols() []p2p.Protocol {
	return s.makeProtocols(ClientProtocolVersions)
}

// Start implements node.Service, starting all internal goroutines needed by the
// Ethereum protocol implementation.
func (s *LightEthereum) Start(srvr *p2p.Server) error {
	log.Warn("Light client mode is an experimental feature")
	s.startBloomHandlers(configs.BloomBitsBlocksClient)
	s.netRPCService = ethapi2.NewPublicNetAPI(srvr, s.networkId)
	// clients are searching for the first advertised protocol in the list
	protocolVersion := AdvertiseProtocolVersions[0]
	s.serverPool.start(srvr, lesTopic(s.blockchain.Genesis().Hash(), protocolVersion))
	s.protocolManager.Start(s.config.LightPeers)
	return nil
}

// Stop implements node.Service, terminating all internal goroutines used by the
// Ethereum protocol.
func (s *LightEthereum) Stop() error {
	s.odr.Stop()
	s.bloomIndexer.Close()
	s.chtIndexer.Close()
	s.blockchain.Stop()
	s.protocolManager.Stop()
	s.txPool.Stop()
	s.engine.Close()

	s.eventMux.Stop()

	time.Sleep(time.Millisecond * 200)
	s.chainDb.Close()
	close(s.shutdownChan)

	return nil
}
