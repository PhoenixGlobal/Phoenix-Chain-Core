package dpos

import (
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/commands/chaintool/dpos/lib"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/node"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/crypto"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/crypto/bls"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
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

var (
	StakingCmd = cli.Command{
		Name:  "staking",
		Usage: "use for staking",
		Subcommands: []cli.Command{
			CreateStakingCmd,
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
