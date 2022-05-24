package main

import (
	"Phoenix-Chain-Core/ethereum/eth"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"unicode"

	cli "gopkg.in/urfave/cli.v1"

	"Phoenix-Chain-Core/ethereum/core/db/snapshotdb"

	"Phoenix-Chain-Core/commands/utils"
	"Phoenix-Chain-Core/configs"
	"Phoenix-Chain-Core/ethereum/node"
	"github.com/naoina/toml"
)

var (
	dumpConfigCommand = cli.Command{
		Action:    utils.MigrateFlags(dumpConfig),
		Name:      "dumpconfig",
		Usage:     "Show configuration values",
		ArgsUsage: "",
		//Flags:       append(append(nodeFlags, rpcFlags...), whisperFlags...),
		Flags:       append(append(nodeFlags, rpcFlags...)),
		Category:    "MISCELLANEOUS COMMANDS",
		Description: `The dumpconfig command shows configuration values.`,
	}

	configFileFlag = cli.StringFlag{
		Name:  "config",
		Usage: "TOML configuration file",
	}
)

// These settings ensure that TOML keys use the same names as Go struct fields.
var tomlSettings = toml.Config{
	NormFieldName: func(rt reflect.Type, key string) string {
		return key
	},
	FieldToKey: func(rt reflect.Type, field string) string {
		return field
	},
	MissingField: func(rt reflect.Type, field string) error {
		link := ""
		if unicode.IsUpper(rune(rt.Name()[0])) && rt.PkgPath() != "main" {
			link = fmt.Sprintf(", see https://godoc.org/%s#%s for available fields", rt.PkgPath(), rt.Name())
		}
		return fmt.Errorf("field '%s' is not defined in %s%s", field, rt.String(), link)
	},
}

type ethstatsConfig struct {
	URL string `toml:",omitempty"`
}

type phoenixchainConfig struct {
	Eth  eth.Config
	Node node.Config
	Ethstats ethstatsConfig
}

func loadConfig(file string, cfg *phoenixchainConfig) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tomlSettings.NewDecoder(bufio.NewReader(f)).Decode(cfg)
	// Add file name to errors that have a line number.
	if _, ok := err.(*toml.LineError); ok {
		err = errors.New(file + ", " + err.Error())
	}
	return err
}

func loadConfigFile(filePath string, cfg *phoenixchainConfig) error {
	file, err := os.Open(filePath)
	if err != nil {
		utils.Fatalf("Failed to read config file: %v", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(cfg)
	if err != nil {
		utils.Fatalf("invalid config file: %v", err)
	}
	return err
}

func defaultNodeConfig() node.Config {
	cfg := node.DefaultConfig
	cfg.Name = clientIdentifier
	cfg.Version = configs.VersionWithCommit(gitCommit, gitDate)
	cfg.HTTPModules = append(cfg.HTTPModules, "phoenixchain")
	cfg.WSModules = append(cfg.WSModules, "phoenixchain")
	cfg.IPCPath = "phoenixchain.ipc"
	return cfg
}

func makeConfigNode(ctx *cli.Context) (*node.Node, phoenixchainConfig) {

	// Load defaults.
	cfg := phoenixchainConfig{
		Eth:  eth.DefaultConfig,
		Node: defaultNodeConfig(),
	}

	//

	// Load config file.
	if file := ctx.GlobalString(configFileFlag.Name); file != "" {
		/*	if err := loadConfig(file, &cfg); err != nil {
			utils.Fatalf("%v", err)
		}*/
		if err := loadConfigFile(file, &cfg); err != nil {
			utils.Fatalf("%v", err)
		}
	}

	// Current version only supports full syncmode
	// ctx.GlobalSet(utils.SyncModeFlag.Name, cfg.Eth.SyncMode.String())

	// Apply flags.
	utils.SetNodeConfig(ctx, &cfg.Node)
	utils.SetPbft(ctx, &cfg.Eth.PbftConfig, &cfg.Node)
	stack, err := node.New(&cfg.Node)
	if err != nil {
		utils.Fatalf("Failed to create the protocol stack: %v", err)
	}

	utils.SetEthConfig(ctx, stack, &cfg.Eth)

	// pass on the rpc port to mpc pool conf.
	//cfg.Eth.MPCPool.LocalRpcPort = cfg.Node.HTTPPort

	// pass on the rpc port to vc pool conf.
	//cfg.Eth.VCPool.LocalRpcPort = cfg.Node.HTTPPort

	// load pbft config file.
	//if pbftConfig := cfg.Eth.LoadPbftConfig(cfg.Node); pbftConfig != nil {
	//	cfg.Eth.PbftConfig = *pbftConfig
	//}

	//if ctx.GlobalIsSet(utils.EthStatsURLFlag.Name) {
	//	cfg.Ethstats.URL = ctx.GlobalString(utils.EthStatsURLFlag.Name)
	//}

	//utils.SetShhConfig(ctx, stack, &cfg.Shh)

	return stack, cfg
}

func makeFullNode(ctx *cli.Context) *node.Node {

	stack, cfg := makeConfigNode(ctx)

	snapshotdb.SetDBPathWithNode(stack.ResolvePath(snapshotdb.DBPath))

	utils.RegisterEthService(stack, &cfg.Eth)

	// Add the Ethereum Stats daemon if requested.
	if cfg.Ethstats.URL != "" {
		utils.RegisterEthStatsService(stack, cfg.Ethstats.URL)
	}
	return stack
}

func makeFullNodeForPBFT(ctx *cli.Context) (*node.Node, phoenixchainConfig) {
	stack, cfg := makeConfigNode(ctx)
	snapshotdb.SetDBPathWithNode(stack.ResolvePath(snapshotdb.DBPath))

	utils.RegisterEthService(stack, &cfg.Eth)

	// Add the Ethereum Stats daemon if requested.
	if cfg.Ethstats.URL != "" {
		utils.RegisterEthStatsService(stack, cfg.Ethstats.URL)
	}
	return stack, cfg
}

// dumpConfig is the dumpconfig command.
func dumpConfig(ctx *cli.Context) error {
	_, cfg := makeConfigNode(ctx)
	comment := ""

	if cfg.Eth.Genesis != nil {
		cfg.Eth.Genesis = nil
		comment += "# Note: this config doesn't contain the genesis block.\n\n"
	}

	out, err := tomlSettings.Marshal(&cfg)
	if err != nil {
		return err
	}
	io.WriteString(os.Stdout, comment)
	os.Stdout.Write(out)
	return nil
}
