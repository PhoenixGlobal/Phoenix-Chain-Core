package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/commands/chaintool/core"

	"gopkg.in/urfave/cli.v1"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/commands/chaintool/dpos"
)

var (
	app *cli.App
)

func init() {
	app = cli.NewApp()

	// Initialize the CLI app
	app.Commands = []cli.Command{
		core.DeployCmd,
		core.InvokeCmd,
		core.SendTransactionCmd,
		core.SendRawTransactionCmd,
		core.GetTxReceiptCmd,
		core.StabilityCmd,
		core.StabPrepareCmd,
		core.AnalyzeStressTestCmd,
		dpos.GovCmd,
		dpos.SlashingCmd,
		dpos.StakingCmd,
		dpos.RestrictingCmd,
		dpos.RewardCmd,
	}

	app.Name = "chaintool"
	app.Version = "1.0.0"

	sort.Sort(cli.CommandsByName(app.Commands))
	app.After = func(ctx *cli.Context) error {
		return nil
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
