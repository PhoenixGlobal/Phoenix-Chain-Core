## Compile：

Run in this directory： go build main.go generate or update chaintool.exe file.

## Command:

##### 1.Deploy contract：
```
./chaintool deploy
-abi        abi json file path (must)
-code       wasm file path (must)
-config     config path  (optional)

eg： ./chaintool deploy -abi "D:\\resource\\temp\\contractc.cpp.abi.json" -code "D:\\resource\\temp\\contractc.wasm"
```
##### 2.Contract call
```
./chaintool invoke
-addr     contract address (must)
-func     functon name and param (must)
-abi      abi json file path (must)
-type     transaction type ,default 2 (optional)

eg: ./chaintool invoke -addr "0xl3p70ayph8flwhx0ljxj86k9yt5kmetsusy5z0" -func "transfer("a",b,c) " -abi "D:\\resource\\temp\\contractc.cpp.abi.json" -type
```
##### 3.Send transaction
```
./chaintool sendTransaction
-from       msg sender (must)
-to         msg acceptor (must)
-value      transfer value (must)
-config     config path (optional)

```
##### 4.Send raw transaction
```
./chaintool sendRawTransaction
-pk         private key file (must)
-from       msg sender (must)
-to         msg acceptor (must)
-value      transfer value (must)
-config     config path (optional)
```
##### 5.Query transactionReceipt
```
./chaintool getTxReceipt
-hash       txhash (must)
-config     config path (optional)
```
##### 6.Prepare transaction stability test account
```
./chaintool prepare
-pkfile      account private key file path,defalut "./test/privatekeys.txt" (optional)
-size        the number of accounts,default 10 (optional)
-value       transfer value (must)
-config      config path (optional)

eg: ./chaintool.exe pre -size 10 -pkfile "./test/privateKeys.txt" -value 0xDE0B6B3A7640000
```

##### 7.Make Stability test
```
./chaintool stab
-pkfile      account private key file path, default "./test/privateKeys.txt"(optional)
-times       send transaction times,default 1000 (optional)
-interval    transaction send interval,if input 10 ,the interval will be 10*Millisecond ,default 10(option)
-config      config path (optional)

eg:  ./chaintool.exe stab -pkfile "./test/privateKeys.txt" -times 10000 -interval 10
```

note: If the command exits normally,the next time you can continue to run with the generated accounts and the command exits abnormally, you need to re-use the pre command to generate the test accounts.

##### Config Description： The config parameter is not passed in the command, and the `config.json` file in the current directory is read by default.

The config.json file is as follows：

```
{
  "url":"http://192.168.9.73:6789",
  "gas": "0x76c0",
  "gasPrice": "0x9184e72a000",
  "from":"0xlwxzlfr7snaaus7f0gy9j4t6x6jlk2zmj70fmq"
}
```

##### 8.dpos staking api 
```
./chaintool  staking

USAGE:
   chaintool staking  [command options] [arguments...]

COMMANDS:
     getVerifierList          1100,query the validator queue of the current settlement epoch
     getValidatorList         1101,query the list of validators in the current consensus round
     getCandidateList         1102,Query the list of all real-time candidates
     getRelatedListByDelAddr  1103,Query the NodeID and staking Id of the node entrusted by the current account address,parameter:add
     getDelegateInfo          1104,Query the delegation information of the current single node,parameter:stakingBlock,address,nodeid
     getCandidateInfo         1105,Query the staking information of the current node,parameter:nodeid
     getPackageReward         1200,query the block reward of the current settlement epoch
     getStakingReward         1201,query the staking reward of the current settlement epoch
     getAvgPackTime           1202,average time to query packaged blocks


eg:  ./chaintool.exe staking  getVerifierList  --rpcurl 'http://127.0.0.1:6771' -testnet
```

##### 9.dpos gov api 
```
./chaintool gov 
NAME:
   chaintool gov - use for gov func

USAGE:
   chaintool gov [command options] [arguments...]

COMMANDS:
     getProposal            2100,get proposal,parameter:proposalID
     getTallyResult         2101,get tally result,parameter:proposalID
     listProposal           2102,list proposal
     getActiveVersion       2103,query the effective version of the  chain
     getGovernParamValue    2104,query the governance parameter value of the current block height,parameter:module,name
     getAccuVerifiersCount  2105,query the cumulative number of votes available for a proposal,parameter:proposalID,blockHash
     listGovernParam        2106,query the list of governance parameters,parameter:module

eg:  ./chaintool.exe gov  getProposal  --rpcurl 'http://127.0.0.1:6771' -testnet --proposalID '0x41'
```

##### 10.dpos restricting api 
```
./chaintool  restricting getRestrictingInfo 
NAME:
   chaintool restricting getRestrictingInfo - 4100,get restricting info,parameter:address

USAGE:
   chaintool restricting getRestrictingInfo  [arguments...]

OPTIONS:
   --rpcurl value   the rpc url
   --testnet        use for testnet
   --address value  account address
   --json           print raw transaction

eg:  ./chaintool.exe restricting  getRestrictingInfo  --rpcurl 'http://127.0.0.1:6771' -testnet --address '0x7tfkaghs4vded6mz6k53xyv5cvqsl63h8c2v5t'
```


##### 11.dpos reward api 
```
NAME:
   chaintool reward getDelegateReward - 5100,query account not withdrawn commission rewards at each node,parameter:nodeList(can empty)

USAGE:
   chaintool reward getDelegateReward  [arguments...]

OPTIONS:
   --rpcurl value    the rpc url
   --testnet         use for testnet
   --nodeList value  node list,may empty
   --json            print raw transaction

eg:  ./chaintool.exe reward  getDelegateReward  --rpcurl 'http://127.0.0.1:6771' -testnet 
```

##### 12.dpos slashing api 
```
NAME:
   chaintool slashing - use for slashing

USAGE:
   chaintool slashing  [command options] [arguments...]

COMMANDS:
     checkDuplicateSign   3001,query whether the node has been reported for too many signatures,parameter:duplicateSignType,nodeid,blockNum
     zeroProduceNodeList  3002,query the list of nodes with zero block

OPTIONS:
   --help, -h  show help

eg:  ./chaintool.exe slashing  zeroProduceNodeList  --rpcurl 'http://127.0.0.1:6771' -testnet 
```

