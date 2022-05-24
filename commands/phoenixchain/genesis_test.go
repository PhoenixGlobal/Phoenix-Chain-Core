package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var customGenesisTests = []struct {
	genesis string
	query   string
	result  string
}{
	// Plain genesis file without anything extra
	{
		genesis: `{
    "alloc":{
        "0xzqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqp7pn3ep":{
            "balance":"0"
        },
        "0xzqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqp7pn3ep":{
            "balance":"0"
        },
        "0xzqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqrdyjj2v":{
            "balance":"200000000000000000000000000"
        },
        "0xzqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqyva9ztf":{
            "balance":"0"
        },
        "0xzqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq93t3hkm":{
            "balance":"0"
        },
        "0xvr8v48qjjrh9dwvdfctqauz98a7yp5se3mm5yk":{
            "balance":"8050000000000000000000000000"
        },
        "0x2klaf9rjl4qjz929kqt382wr49a00zc9ng995w":{
            "balance":"2000000000000000000000000000"
        }
    },
    "economicModel":{
        "common":{
            "maxEpochMinutes":4,
            "maxConsensusVals":4,
            "additionalCycleTime":28
        },
        "staking":{
            "stakeThreshold": 1000000000000000000000000,
            "operatingThreshold": 10000000000000000000,
            "maxValidators": 30,
            "unStakeFreezeDuration": 2
        },
        "slashing":{
           "slashFractionDuplicateSign": 100,
           "duplicateSignReportReward": 50,
           "maxEvidenceAge":1,
           "slashBlocksReward":20
        },
         "gov": {
            "versionProposalVoteDurationSeconds": 160,
            "versionProposalSupportRate": 6670,
            "textProposalVoteDurationSeconds": 160,
            "textProposalVoteRate": 5000,
            "textProposalSupportRate": 6670,          
            "cancelProposalVoteRate": 5000,
            "cancelProposalSupportRate": 6670,
            "paramProposalVoteDurationSeconds": 160,
            "paramProposalVoteRate": 5000,
            "paramProposalSupportRate": 6670      
        },
        "reward":{
            "newBlockRate": 50,
            "phoenixchainFoundationYear": 10 
        },
        "innerAcc":{
            "phoenixchainFundAccount": "0xfyeszufxwxk62p46djncj86rd553skppy4qgz4",
            "phoenixchainFundBalance": 0,
            "cdfAccount": "0xc8enpvs5v6974shxgxxav5dsn36e5jl4r0hwhh",
            "cdfBalance": 331811981000000000000000000
        }
    },
    "coinbase":"0xqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq542u6a",
    "extraData":"",
    "gasLimit":"0x2fefd8",
    "nonce":"0x0376e56dffd12ab53bb149bda4e0cbce2b6aabe4cccc0df0b5a39e12977a2fcd23",
    "parentHash":"0x0000000000000000000000000000000000000000000000000000000000000000",
    "timestamp":"0x00",
    "config":{
        "pbft":{
            "initialNodes":[
                {
                    "node":"enode://4fcc251cf6bf3ea53a748971a223f5676225ee4380b65c7889a2b491e1551d45fe9fcc19c6af54dcf0d5323b5aa8ee1d919791695082bae1f86dd282dba4150f@0.0.0.0:16789",
                    "blsPubKey":"d341a0c485c9ec00cecf7ea16323c547900f6a1bacb9daacb00c2b8bacee631f75d5d31b75814b7f1ae3a4e18b71c617bc2f230daa0c893746ed87b08b2df93ca4ddde2816b3ac410b9980bcc048521562a3b2d00e900fd777d3cf88ce678719"
                }
            ],
            "amount":10,
			"period":10000,
            "validatorMode":"dpos"
        }
    }
}`,
		query:  "phoenixchain.getBlock(0).nonce",
		result: "0x0376e56dffd12ab53bb149bda4e0cbce2b6aabe4cccc0df0b5a39e12977a2fcd23000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
	},
	//Genesis file with only pbft config
	{
		genesis: `{
    "alloc":{
        "0xzqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqp7pn3ep":{
            "balance":"0"
        },
        "0xzqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqp7pn3ep":{
            "balance":"0"
        },
        "0xzqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqrdyjj2v":{
            "balance":"200000000000000000000000000"
        },
        "0xzqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqyva9ztf":{
            "balance":"0"
        },
        "0xzqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq93t3hkm":{
            "balance":"0"
        },
        "0xvr8v48qjjrh9dwvdfctqauz98a7yp5se3mm5yk":{
            "balance":"8050000000000000000000000000"
        },
        "0x2klaf9rjl4qjz929kqt382wr49a00zc9ng995w":{
            "balance":"2000000000000000000000000000"
        }
    },
    "economicModel":{
        "common":{
            "maxEpochMinutes":4,
            "maxConsensusVals":4,
            "additionalCycleTime":28
        },
        "staking":{
            "stakeThreshold": 1000000000000000000000000,
            "operatingThreshold": 10000000000000000000,
            "maxValidators": 30,
            "unStakeFreezeDuration": 2
        },
        "slashing":{
           "slashFractionDuplicateSign": 100,
           "duplicateSignReportReward": 50,
           "maxEvidenceAge":1,
           "slashBlocksReward":20
        },
         "gov": {
            "versionProposalVoteDurationSeconds": 160,
            "versionProposalSupportRate": 6670,
            "textProposalVoteDurationSeconds": 160,
            "textProposalVoteRate": 5000,
            "textProposalSupportRate": 6670,          
            "cancelProposalVoteRate": 5000,
            "cancelProposalSupportRate": 6670,
            "paramProposalVoteDurationSeconds": 160,
            "paramProposalVoteRate": 5000,
            "paramProposalSupportRate": 6670      
        },
        "reward":{
            "newBlockRate": 50,
            "phoenixchainFoundationYear": 10 
        },
        "innerAcc":{
            "phoenixchainFundAccount": "0xfyeszufxwxk62p46djncj86rd553skppy4qgz4",
            "phoenixchainFundBalance": 0,
            "cdfAccount": "0xc8enpvs5v6974shxgxxav5dsn36e5jl4r0hwhh",
            "cdfBalance": 331811981000000000000000000
        }
    },
    "coinbase":"0xqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq542u6a",
    "extraData":"",
    "gasLimit":"0x2fefd8",
    "nonce":"0x0376e56dffd12ab53bb149bda4e0cbce2b6aabe4cccc0df0b5a39e12977a2fcd23",
    "parentHash":"0x0000000000000000000000000000000000000000000000000000000000000000",
    "timestamp":"0x00",
    "config":{
        "pbft":{
            "initialNodes":[
                {
                    "node":"enode://4fcc251cf6bf3ea53a748971a223f5676225ee4380b65c7889a2b491e1551d45fe9fcc19c6af54dcf0d5323b5aa8ee1d919791695082bae1f86dd282dba4150f@0.0.0.0:16789",
                    "blsPubKey":"d341a0c485c9ec00cecf7ea16323c547900f6a1bacb9daacb00c2b8bacee631f75d5d31b75814b7f1ae3a4e18b71c617bc2f230daa0c893746ed87b08b2df93ca4ddde2816b3ac410b9980bcc048521562a3b2d00e900fd777d3cf88ce678719"
                }
            ],
            "amount":10,
			"period":10000,
            "validatorMode":"dpos"
        }
    }
}`,
		query:  "phoenixchain.getBlock(0).nonce",
		result: "0x0376e56dffd12ab53bb149bda4e0cbce2b6aabe4cccc0df0b5a39e12977a2fcd23000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
	},
	//Genesis file with specific chain configurations
	{
		genesis: `{
    "alloc":{
        "0xzqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqp7pn3ep":{
            "balance":"0"
        },
        "0xzqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqzsjx8h7":{
            "balance":"0"
        },
        "0xzqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqrdyjj2v":{
            "balance":"200000000000000000000000000"
        },
        "0xzqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqyva9ztf":{
            "balance":"0"
        },
        "0xzqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq93t3hkm":{
            "balance":"0"
        },
        "0xvr8v48qjjrh9dwvdfctqauz98a7yp5se3mm5yk":{
            "balance":"8050000000000000000000000000"
        },
        "0x2klaf9rjl4qjz929kqt382wr49a00zc9ng995w":{
            "balance":"2000000000000000000000000000"
        }
    },
    "economicModel":{
        "common":{
            "maxEpochMinutes":4,
            "maxConsensusVals":4,
            "additionalCycleTime":28
        },
        "staking":{
            "stakeThreshold": 1000000000000000000000000,
            "operatingThreshold": 10000000000000000000,
            "maxValidators": 30,
            "unStakeFreezeDuration": 2
        },
        "slashing":{
           "slashFractionDuplicateSign": 100,
           "duplicateSignReportReward": 50,
           "maxEvidenceAge":1,
           "slashBlocksReward":20
        },
         "gov": {
            "versionProposalVoteDurationSeconds": 160,
            "versionProposalSupportRate": 6670,
            "textProposalVoteDurationSeconds": 160,
            "textProposalVoteRate": 5000,
            "textProposalSupportRate": 6670,          
            "cancelProposalVoteRate": 5000,
            "cancelProposalSupportRate": 6670,
            "paramProposalVoteDurationSeconds": 160,
            "paramProposalVoteRate": 5000,
            "paramProposalSupportRate": 6670      
        },
        "reward":{
            "newBlockRate": 50,
            "phoenixchainFoundationYear": 10 
        },
        "innerAcc":{
            "phoenixchainFundAccount": "0xfyeszufxwxk62p46djncj86rd553skppy4qgz4",
            "phoenixchainFundBalance": 0,
            "cdfAccount": "0xc8enpvs5v6974shxgxxav5dsn36e5jl4r0hwhh",
            "cdfBalance": 331811981000000000000000000
        }
    },
    "coinbase":"0xqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq542u6a",
    "extraData":"",
    "gasLimit":"0x2fefd8",
    "nonce":"0x0376e56dffd12ab53bb149bda4e0cbce2b6aabe4cccc0df0b5a39e12977a2fcd23",
    "parentHash":"0x0000000000000000000000000000000000000000000000000000000000000000",
    "timestamp":"0x00",
    "config":{
        "chainId":101,
        "eip155Block":0,
        "interpreter":"wasm",
        "pbft":{
            "initialNodes":[
                {
                    "node":"enode://4fcc251cf6bf3ea53a748971a223f5676225ee4380b65c7889a2b491e1551d45fe9fcc19c6af54dcf0d5323b5aa8ee1d919791695082bae1f86dd282dba4150f@0.0.0.0:16789",
                    "blsPubKey":"d341a0c485c9ec00cecf7ea16323c547900f6a1bacb9daacb00c2b8bacee631f75d5d31b75814b7f1ae3a4e18b71c617bc2f230daa0c893746ed87b08b2df93ca4ddde2816b3ac410b9980bcc048521562a3b2d00e900fd777d3cf88ce678719"
                }
            ],
            "amount":10,
			"period":10000,
            "validatorMode":"dpos"
        }
    }
}`,
		query:  "phoenixchain.getBlock(0).nonce",
		result: "0x0376e56dffd12ab53bb149bda4e0cbce2b6aabe4cccc0df0b5a39e12977a2fcd23000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
	},
}

// Tests that initializing PhoenixChain with a custom genesis block and chain definitions
// work properly.
func TestCustomGenesis(t *testing.T) {
	for i, tt := range customGenesisTests {
		// Create a temporary data directory to use and inspect later
		datadir := tmpdir(t)
		defer os.RemoveAll(datadir)

		// Initialize the data directory with the custom genesis block
		json := filepath.Join(datadir, "genesis.json")
		if err := ioutil.WriteFile(json, []byte(tt.genesis), 0600); err != nil {
			t.Fatalf("test %d: failed to write genesis file: %v", i, err)
		}
		runPhoenixChain(t, "--datadir", datadir, "init", json).WaitExit()

		// Query the custom genesis block
		phoenixchain := runPhoenixChain(t,
			"--datadir", datadir, "--maxpeers", "60", "--port", "0",
			"--nodiscover", "--nat", "none", "--ipcdisable", "--testnet",
			"--exec", tt.query, "console")
		t.Log("testi", i)
		phoenixchain.ExpectRegexp(tt.result)
		phoenixchain.ExpectExit()
	}
}
