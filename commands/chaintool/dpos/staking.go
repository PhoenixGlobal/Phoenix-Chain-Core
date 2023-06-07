package dpos

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/commands/chaintool/dpos/lib"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/node"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/crypto"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/crypto/bls"
	"io"
	"math/big"
	"os"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/p2p/discover"
	"gopkg.in/urfave/cli.v1"
)

// createStaking
type Dpos_1000 struct {
	Typ                uint16
	BenefitAddress     common.Address
	NodeId             discover.NodeID
	ExternalId         string
	NodeName           string
	Website            string
	Details            string
	Amount             *big.Int
	RewardPer          uint64
	ProgramVersion     uint32
	ProgramVersionSign common.VersionSign
	BlsPubKey          bls.PublicKeyHex
	BlsProof           bls.SchnorrProofHex
}

// editorCandidate
type Dpos_1001 struct {
	BenefitAddress common.Address
	NodeId         discover.NodeID
	RewardPer      uint16
	ExternalId     string
	NodeName       string
	Website        string
	Details        string
}

// increaseStaking
type Dpos_1002 struct {
	NodeId discover.NodeID
	Typ    uint16
	Amount *big.Int
}

// withdrawStaking
type Dpos_1003 struct {
	NodeId discover.NodeID
}

// delegate
type Dpos_1004 struct {
	Typ    uint16
	NodeId discover.NodeID
	Amount *big.Int
}

// withdrawDelegate
type Dpos_1005 struct {
	StakingBlockNum *big.Int
	NodeId          discover.NodeID
	Amount          *big.Int
}

var (
	StakingCmd = cli.Command{
		Name:  "staking",
		Usage: "use for staking",
		Subcommands: []cli.Command{
			CreateStakingCmd,
			AddStakingCmd,
			UnStakingCmd,
			UpdateStakingCmd,
			DelegateCmd,
			UnDelegateCmd,
			GetVerifierListCmd,
			getValidatorListCmd,
			getCandidateListCmd,
			getRelatedListByDelAddrCmd,
			getDelegateInfoCmd,
			getCandidateInfoCmd,
			getPackageRewardCmd,
			getStakingRewardCmd,
			getAvgPackTimeCmd,
		},
	}
	CreateStakingCmd = cli.Command{
		Name:   "createStaking",
		Usage:  "1000,create Staking",
		Before: netCheck,
		Action: createStaking,
		Flags:  []cli.Flag{configPathFlag,keystoreFlag,blsKeyfileFlag,nodeKeyFlag,stakingParamsFlag},
	}
	AddStakingCmd = cli.Command{
		Name:   "addStaking",
		Usage:  "1002,add Staking",
		Before: netCheck,
		Action: addStaking,
		Flags:  []cli.Flag{configPathFlag,keystoreFlag,addStakingParamsFlag},
	}
	UnStakingCmd = cli.Command{
		Name:   "unStaking",
		Usage:  "1003,withdraw Staking",
		Before: netCheck,
		Action: unStaking,
		Flags:  []cli.Flag{configPathFlag,keystoreFlag,unStakingParamsFlag},
	}
	UpdateStakingCmd = cli.Command{
		Name:   "UpdateStaking",
		Usage:  "1001,update Staking",
		Before: netCheck,
		Action: updateStaking,
		Flags:  []cli.Flag{configPathFlag,keystoreFlag,updateStakingParamsFlag},
	}
	DelegateCmd = cli.Command{
		Name:   "delegate",
		Usage:  "1004,delegate",
		Before: netCheck,
		Action: delegate,
		Flags:  []cli.Flag{configPathFlag,keystoreFlag,addStakingParamsFlag},
	}
	UnDelegateCmd = cli.Command{
		Name:   "unDelegate",
		Usage:  "1005,withdraw Delegate",
		Before: netCheck,
		Action: unDelegate,
		Flags:  []cli.Flag{configPathFlag,keystoreFlag,unDelegateParamsFlag},
	}
	GetVerifierListCmd = cli.Command{
		Name:   "getVerifierList",
		Usage:  "1100,query the validator queue of the current settlement epoch",
		Before: netCheck,
		Action: getVerifierList,
		Flags:  []cli.Flag{rpcUrlFlag, addressHRPFlag, jsonFlag},
	}
	getValidatorListCmd = cli.Command{
		Name:   "getValidatorList",
		Usage:  "1101,query the list of validators in the current consensus round",
		Before: netCheck,
		Action: getValidatorList,
		Flags:  []cli.Flag{rpcUrlFlag, addressHRPFlag, jsonFlag},
	}
	getCandidateListCmd = cli.Command{
		Name:   "getCandidateList",
		Usage:  "1102,Query the list of all real-time candidates",
		Before: netCheck,
		Action: getCandidateList,
		Flags:  []cli.Flag{rpcUrlFlag, addressHRPFlag, jsonFlag},
	}
	getRelatedListByDelAddrCmd = cli.Command{
		Name:   "getRelatedListByDelAddr",
		Usage:  "1103,Query the NodeID and staking Id of the node entrusted by the current account address,parameter:add",
		Before: netCheck,
		Action: getRelatedListByDelAddr,
		Flags:  []cli.Flag{rpcUrlFlag, addressHRPFlag, addFlag, jsonFlag},
	}
	getDelegateInfoCmd = cli.Command{
		Name:   "getDelegateInfo",
		Usage:  "1104,Query the delegation information of the current single node,parameter:stakingBlock,address,nodeid",
		Before: netCheck,
		Action: getDelegateInfo,
		Flags:  []cli.Flag{rpcUrlFlag, addressHRPFlag, stakingBlockNumFlag, addFlag, nodeIdFlag, jsonFlag},
	}
	getCandidateInfoCmd = cli.Command{
		Name:   "getCandidateInfo",
		Usage:  "1105,Query the staking information of the current node,parameter:nodeid",
		Before: netCheck,
		Action: getCandidateInfo,
		Flags:  []cli.Flag{rpcUrlFlag, addressHRPFlag, nodeIdFlag, jsonFlag},
	}
	getPackageRewardCmd = cli.Command{
		Name:   "getPackageReward",
		Usage:  "1200,query the block reward of the current settlement epoch",
		Before: netCheck,
		Action: getPackageReward,
		Flags:  []cli.Flag{rpcUrlFlag, addressHRPFlag, jsonFlag},
	}
	getStakingRewardCmd = cli.Command{
		Name:   "getStakingReward",
		Usage:  "1201,query the staking reward of the current settlement epoch",
		Before: netCheck,
		Action: getStakingReward,
		Flags:  []cli.Flag{rpcUrlFlag, addressHRPFlag, jsonFlag},
	}
	getAvgPackTimeCmd = cli.Command{
		Name:   "getAvgPackTime",
		Usage:  "1202,average time to query packaged blocks",
		Before: netCheck,
		Action: getAvgPackTime,
		Flags:  []cli.Flag{rpcUrlFlag, addressHRPFlag, jsonFlag},
	}
	addFlag = cli.StringFlag{
		Name:  "address",
		Usage: "account address",
	}
	stakingBlockNumFlag = cli.Uint64Flag{
		Name:  "stakingBlock",
		Usage: "block height when staking is initiated",
	}
	nodeIdFlag = cli.StringFlag{
		Name:  "nodeid",
		Usage: "node id",
	}
)

func createStaking(c *cli.Context) error {
	//url := c.String(rpcUrlFlag.Name)
	//if url == "" {
	//	return errors.New("rpc url not set")
	//}

	keystorePath:=c.String(keystoreFlag.Name)
	priKey,_:=getPrivateKey(keystorePath)

	//Load Dpos_1000 from json
	stakingParams:=c.String(stakingParamsFlag.Name)
	file, err := os.Open(stakingParams)
	if err != nil {
		return fmt.Errorf("Failed to read stakingParams file: %v", err)
	}
	defer file.Close()
	file.Seek(0, io.SeekStart)
	var dpos_1000 Dpos_1000
	if err := json.NewDecoder(file).Decode(&dpos_1000); err != nil {
		return fmt.Errorf("parse config to json error,%s", err.Error())
	}

	//set Dpos_1000.BlsProof
	blsKeyfilePath := c.String(blsKeyfileFlag.Name)
	blsProof, err := getBlsProof(blsKeyfilePath)
	if err != nil {
		return fmt.Errorf("getBlsProof error: %v", err)
	}
	dpos_1000.BlsProof=blsProof

	//set Dpos_1000.ProgramVersionSign,   ProgramVersion is GenesisVersion
	nodeKeyPath := c.String(nodeKeyFlag.Name)
	nodeKey, err := GetNodeKey(nodeKeyPath)
	if err != nil {
		return fmt.Errorf("getNodeKey error: %v", err)
	}
	node.GetCryptoHandler().SetPrivateKey(crypto.HexMustToECDSA(nodeKey))
	versionSign := common.VersionSign{}
	versionSign.SetBytes(node.GetCryptoHandler().MustSign(dpos_1000.ProgramVersion))
	fmt.Println("versionSign is ",versionSign.String())
	dpos_1000.ProgramVersionSign=versionSign
	fmt.Println("CreateStaking params dpos_1000 is ",dpos_1000.Amount,dpos_1000.ProgramVersion,dpos_1000)

	priKeyStr := hex.EncodeToString(crypto.FromECDSA(priKey))
	keyCredentials,err:=lib.NewCredential(priKeyStr)
	config := lib.DposMainNetParams
	sc := lib.NewStakingContract(config, keyCredentials)
	sp := lib.StakingParam{
		NodeId:            "0x"+dpos_1000.NodeId.String(),
		Amount:            dpos_1000.Amount,
		StakingAmountType: lib.StakingAmountType(dpos_1000.Typ),
		BenefitAddress:    dpos_1000.BenefitAddress.String(),
		ExternalId:        dpos_1000.ExternalId,
		NodeName:          dpos_1000.NodeName,
		WebSite:           dpos_1000.Website,
		Details:           dpos_1000.Details,
		ProcessVersion: lib.ProgramVersion{
			Version: big.NewInt(int64(dpos_1000.ProgramVersion)),
			Sign:    dpos_1000.ProgramVersionSign.String(),
		},
		BlsPubKey:  "0x"+dpos_1000.BlsPubKey.String(),
		BlsProof:  "0x"+dpos_1000.BlsProof.String(),
		RewardPer: big.NewInt(int64(dpos_1000.RewardPer)),
	}
	fmt.Println("Staking params sp is ",sp)
	result, err := sc.Staking(sp)
	if err != nil {
		return fmt.Errorf("StakingContract.Staking failed: %v", err)
	}
	fmt.Println("Staking success,res is ",result)

	////SendRawTransaction
	//configPath:=c.String(configPathFlag.Name)
	//data,to:= EncodeDPOSStaking(1000, &dpos_1000)
	//res, err :=core.SendRawTransactionWithData(configPath,fromAddress,to.String(),0,priKey,data)
	//if err != nil {
	//	return err
	//}
	//fmt.Println("Staking success,res is ",string(res))

	//client, err := ethclient.Dial(url)
	//if err != nil {
	//	return err
	//}
	//res, err := CallDPosStakingContract(client, 1000, &dpos_1000)
	//if err != nil {
	//	return err
	//}
	//fmt.Println("Staking success,res is ",string(res))
	return nil

}

func addStaking(c *cli.Context) error {
	keystorePath:=c.String(keystoreFlag.Name)
	priKey,_:=getPrivateKey(keystorePath)

	//Load Dpos_1002 from json
	addStakingParams:=c.String(addStakingParamsFlag.Name)
	file, err := os.Open(addStakingParams)
	if err != nil {
		return fmt.Errorf("Failed to read addStakingParams file: %v", err)
	}
	defer file.Close()
	file.Seek(0, io.SeekStart)
	var dpos_1002 Dpos_1002
	if err := json.NewDecoder(file).Decode(&dpos_1002); err != nil {
		return fmt.Errorf("parse config to json error,%s", err.Error())
	}

	fmt.Println("AddStaking params dpos_1002 is ",dpos_1002.NodeId,dpos_1002.Amount,dpos_1002.Typ)

	priKeyStr := hex.EncodeToString(crypto.FromECDSA(priKey))
	keyCredentials,err:=lib.NewCredential(priKeyStr)
	config := lib.DposMainNetParams
	sc := lib.NewStakingContract(config, keyCredentials)

	nodeId:= "0x"+dpos_1002.NodeId.String()

	result, err := sc.AddStaking(nodeId,lib.StakingAmountType(dpos_1002.Typ),dpos_1002.Amount)
	if err != nil {
		return fmt.Errorf("StakingContract.AddStaking failed: %v", err)
	}
	fmt.Println("AddStaking success,res is ",result)
	return nil
}

func unStaking(c *cli.Context) error {
	keystorePath:=c.String(keystoreFlag.Name)
	priKey,_:=getPrivateKey(keystorePath)

	//Load Dpos_1003 from json
	unStakingParams:=c.String(addStakingParamsFlag.Name)
	file, err := os.Open(unStakingParams)
	if err != nil {
		return fmt.Errorf("Failed to read unStakingParams file: %v", err)
	}
	defer file.Close()
	file.Seek(0, io.SeekStart)
	var dpos_1003 Dpos_1003
	if err := json.NewDecoder(file).Decode(&dpos_1003); err != nil {
		return fmt.Errorf("parse config to json error,%s", err.Error())
	}

	fmt.Println("UnStaking params dpos_1003 is ",dpos_1003.NodeId)

	priKeyStr := hex.EncodeToString(crypto.FromECDSA(priKey))
	keyCredentials,err:=lib.NewCredential(priKeyStr)
	config := lib.DposMainNetParams
	sc := lib.NewStakingContract(config, keyCredentials)

	nodeId:= "0x"+dpos_1003.NodeId.String()

	result, err := sc.UnStaking(nodeId)
	if err != nil {
		return fmt.Errorf("StakingContract.UnStaking failed: %v", err)
	}
	fmt.Println("UnStaking success,res is ",result)
	return nil
}

func updateStaking(c *cli.Context) error {
	keystorePath:=c.String(keystoreFlag.Name)
	priKey,_:=getPrivateKey(keystorePath)

	//Load Dpos_1001 from json
	updateStakingParams:=c.String(updateStakingParamsFlag.Name)
	file, err := os.Open(updateStakingParams)
	if err != nil {
		return fmt.Errorf("Failed to read updateStakingParams file: %v", err)
	}
	defer file.Close()
	file.Seek(0, io.SeekStart)
	var dpos_1001 Dpos_1001
	if err := json.NewDecoder(file).Decode(&dpos_1001); err != nil {
		return fmt.Errorf("parse config to json error,%s", err.Error())
	}

	fmt.Println("UpdateStaking params dpos_1001 is ",dpos_1001)

	priKeyStr := hex.EncodeToString(crypto.FromECDSA(priKey))
	keyCredentials,err:=lib.NewCredential(priKeyStr)
	config := lib.DposMainNetParams
	sc := lib.NewStakingContract(config, keyCredentials)
	sp := lib.UpdateStakingParam{
		NodeId:            "0x"+dpos_1001.NodeId.String(),
		BenefitAddress:    dpos_1001.BenefitAddress.String(),
		ExternalId:        dpos_1001.ExternalId,
		NodeName:          dpos_1001.NodeName,
		WebSite:           dpos_1001.Website,
		Details:           dpos_1001.Details,
		RewardPer: big.NewInt(int64(dpos_1001.RewardPer)),
	}
	fmt.Println("UpdateStaking params sp is ",sp)
	result, err := sc.UpdateStakingInfo(sp)
	if err != nil {
		return fmt.Errorf("StakingContract.UpdateStakingInfo failed: %v", err)
	}
	fmt.Println("UpdateStakingInfo success,res is ",result)
	return nil
}

func delegate(c *cli.Context) error {
	keystorePath:=c.String(keystoreFlag.Name)
	priKey,_:=getPrivateKey(keystorePath)

	//Load Dpos_1004 from json
	delegateParams:=c.String(delegateParamsFlag.Name)
	file, err := os.Open(delegateParams)
	if err != nil {
		return fmt.Errorf("Failed to read delegateParams file: %v", err)
	}
	defer file.Close()
	file.Seek(0, io.SeekStart)
	var dpos_1004 Dpos_1004
	if err := json.NewDecoder(file).Decode(&dpos_1004); err != nil {
		return fmt.Errorf("parse config to json error,%s", err.Error())
	}

	fmt.Println("Delegate params dpos_1004 is ",dpos_1004.NodeId,dpos_1004.Amount,dpos_1004.Typ)

	priKeyStr := hex.EncodeToString(crypto.FromECDSA(priKey))
	keyCredentials,err:=lib.NewCredential(priKeyStr)
	config := lib.DposMainNetParams
	sc := lib.NewDelegateContract(config, keyCredentials)

	nodeId:= "0x"+dpos_1004.NodeId.String()

	result, err := sc.Delegate(nodeId,lib.StakingAmountType(dpos_1004.Typ),dpos_1004.Amount)
	if err != nil {
		return fmt.Errorf("DelegateContract.Delegate failed: %v", err)
	}
	fmt.Println("Delegate success,res is ",result)
	return nil
}

func unDelegate(c *cli.Context) error {
	keystorePath:=c.String(keystoreFlag.Name)
	priKey,_:=getPrivateKey(keystorePath)

	//Load Dpos_1005 from json
	unDelegateParams:=c.String(unDelegateParamsFlag.Name)
	file, err := os.Open(unDelegateParams)
	if err != nil {
		return fmt.Errorf("Failed to read unDelegateParams file: %v", err)
	}
	defer file.Close()
	file.Seek(0, io.SeekStart)
	var dpos_1005 Dpos_1005
	if err := json.NewDecoder(file).Decode(&dpos_1005); err != nil {
		return fmt.Errorf("parse config to json error,%s", err.Error())
	}

	fmt.Println("unDelegate params dpos_1005 is ",dpos_1005)

	priKeyStr := hex.EncodeToString(crypto.FromECDSA(priKey))
	keyCredentials,err:=lib.NewCredential(priKeyStr)
	config := lib.DposMainNetParams
	sc := lib.NewDelegateContract(config, keyCredentials)

	nodeId:= "0x"+dpos_1005.NodeId.String()

	result, err := sc.UnDelegate(nodeId,dpos_1005.StakingBlockNum,dpos_1005.Amount)
	if err != nil {
		return fmt.Errorf("DelegateContract.UnDelegate failed: %v", err)
	}
	fmt.Println("UnDelegate success,res is ",result)
	return nil
}

func getVerifierList(c *cli.Context) error {
	return query(c, 1100)
}

func getValidatorList(c *cli.Context) error {
	return query(c, 1101)
}

func getCandidateList(c *cli.Context) error {
	return query(c, 1102)
}

func getRelatedListByDelAddr(c *cli.Context) error {
	addstring := c.String(addFlag.Name)
	if addstring == "" {
		return errors.New("The Del's account address is not set")
	}
	add, err := common.StringToAddress(addstring)
	if err != nil {
		return err
	}
	return query(c, 1103, add)
}

func getDelegateInfo(c *cli.Context) error {
	addstring := c.String(addFlag.Name)
	if addstring == "" {
		return errors.New("The Del's account address is not set")
	}
	add, err := common.StringToAddress(addstring)
	if err != nil {
		return err
	}
	nodeIDstring := c.String(nodeIdFlag.Name)
	if nodeIDstring == "" {
		return errors.New("The verifier's node ID is not set")
	}
	nodeid, err := discover.HexID(nodeIDstring)
	if err != nil {
		return err
	}
	stakingBlockNum := c.Uint64(stakingBlockNumFlag.Name)
	return query(c, 1104, stakingBlockNum, add, nodeid)
}

func getCandidateInfo(c *cli.Context) error {
	nodeIDstring := c.String(nodeIdFlag.Name)
	if nodeIDstring == "" {
		return errors.New("The verifier's node ID is not set")
	}
	nodeid, err := discover.HexID(nodeIDstring)
	if err != nil {
		return err
	}
	return query(c, 1105, nodeid)
}

func getPackageReward(c *cli.Context) error {
	return query(c, 1200)
}

func getStakingReward(c *cli.Context) error {
	return query(c, 1201)
}

func getAvgPackTime(c *cli.Context) error {
	return query(c, 1202)
}
