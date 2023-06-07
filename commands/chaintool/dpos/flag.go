package dpos

import "gopkg.in/urfave/cli.v1"

var (
	rpcUrlFlag = cli.StringFlag{
		Name:  "rpcurl",
		Usage: "the rpc url",
	}

	jsonFlag = cli.BoolFlag{
		Name:  "json",
		Usage: "print raw transaction",
	}

	addressHRPFlag = cli.StringFlag{
		Name:  "addressHRP",
		Usage: "set address hrp",
	}

	configPathFlag = cli.StringFlag{
		Name:  "config",
		Usage: "config path",
	}

	keystoreFlag = cli.StringFlag{
		Name:  "keystore",
		Usage: "keystore file path",
	}

	nodeKeyFlag = cli.StringFlag{
		Name:  "nodeKey",
		Usage: "nodeKey file path",
	}

	blsKeyfileFlag = cli.StringFlag{
		Name:  "blsKey",
		Usage: "file containing the blsKey",
	}

	stakingParamsFlag = cli.StringFlag{
		Name:  "stakingParams",
		Usage: "stakingParams file path",
	}

	addStakingParamsFlag = cli.StringFlag{
		Name:  "addStakingParams",
		Usage: "addStakingParams file path",
	}

	unStakingParamsFlag = cli.StringFlag{
		Name:  "unStakingParams",
		Usage: "unStakingParams file path",
	}

	updateStakingParamsFlag = cli.StringFlag{
		Name:  "updateStakingParams",
		Usage: "updateStakingParams file path",
	}

	delegateParamsFlag = cli.StringFlag{
		Name:  "delegateParams",
		Usage: "delegateParams file path",
	}

	unDelegateParamsFlag = cli.StringFlag{
		Name:  "unDelegateParams",
		Usage: "unDelegateParams file path",
	}

	govParamsFlag = cli.StringFlag{
		Name:  "govParams",
		Usage: "gov file path",
	}
)
