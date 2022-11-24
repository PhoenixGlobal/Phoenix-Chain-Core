package dpos

import (
	phoenixchain "github.com/PhoenixGlobal/Phoenix-Chain-Core"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/commands/utils"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/accounts/keystore"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/ethclient"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common/hexutil"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common/vm"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/crypto/bls"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/rlp"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/web"
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"io"
	"io/ioutil"
	"os"
)

func CallDPosContract(client *ethclient.Client, funcType uint16, params ...interface{}) ([]byte, error) {
	send, to := EncodeDPOS(funcType, params...)
	var msg phoenixchain.CallMsg
	msg.Data = send
	msg.To = &to
	return client.CallContract(context.Background(), msg, nil)
}

func CallDPosStakingContract(client *ethclient.Client, funcType uint16, params *Dpos_1000) ([]byte, error) {
	send, to := EncodeDPOSStaking(funcType, params)
	var msg phoenixchain.CallMsg
	msg.Data = send
	msg.To = &to
	return client.CallContract(context.Background(), msg, nil)
}

// CallMsg contains parameters for contract calls.
type CallMsg struct {
	To   *common.Address // the destination contract (nil for contract creation)
	Data hexutil.Bytes   // input data, usually an ABI-encoded contract method invocation
}

func BuildDPosContract(funcType uint16, params ...interface{}) ([]byte, error) {
	send, to := EncodeDPOS(funcType, params...)
	var msg CallMsg
	msg.Data = send
	msg.To = &to
	return json.Marshal(msg)
}

func EncodeDPOS(funcType uint16, params ...interface{}) ([]byte, common.Address) {
	par := buildParams(funcType, params...)
	buf := new(bytes.Buffer)
	err := rlp.Encode(buf, par)
	if err != nil {
		panic(fmt.Errorf("encode rlp data fail: %v", err))
	}
	return buf.Bytes(), funcTypeToContractAddress(funcType)
}

func EncodeDPOSStaking(funcType uint16, params *Dpos_1000) ([]byte, common.Address) {
	par := buildStakingParams(funcType, params)
	buf := new(bytes.Buffer)
	err := rlp.Encode(buf, par)
	if err != nil {
		panic(fmt.Errorf("encode rlp data fail: %v", err))
	}
	data:=buf.Bytes()
	fmt.Printf("funcType:%d rlp data = %s\n", funcType, hexutil.Encode(data))
	return data, funcTypeToContractAddress(funcType)
}
func EncodeDPOSGov(funcType uint16, params *Dpos_2000) ([]byte, common.Address) {
	par := buildGovParams(funcType, params)
	buf := new(bytes.Buffer)
	err := rlp.Encode(buf, par)
	if err != nil {
		panic(fmt.Errorf("encode rlp data fail: %v", err))
	}
	data:=buf.Bytes()
	fmt.Printf("funcType:%d rlp data = %s\n", funcType, hexutil.Encode(data))
	return data, funcTypeToContractAddress(funcType)
}


func buildParams(funcType uint16, params ...interface{}) [][]byte {
	var res [][]byte
	res = make([][]byte, 0)
	fnType, _ := rlp.EncodeToBytes(funcType)
	res = append(res, fnType)
	for _, param := range params {
		val, err := rlp.EncodeToBytes(param)
		if err != nil {
			panic(err)
		}
		res = append(res, val)
	}
	return res
}

func buildStakingParams(funcType uint16, params *Dpos_1000) [][]byte {
	var res [][]byte
	res = make([][]byte, 0)
	fnType, _ := rlp.EncodeToBytes(funcType)
	res = append(res, fnType)

	typ, _ := rlp.EncodeToBytes(params.Typ)
	benefitAddress, _ := rlp.EncodeToBytes(params.BenefitAddress.Bytes())
	nodeId, _ := rlp.EncodeToBytes(params.NodeId)
	externalId, _ := rlp.EncodeToBytes(params.ExternalId)
	nodeName, _ := rlp.EncodeToBytes(params.NodeName)
	website, _ := rlp.EncodeToBytes(params.Website)
	details, _ := rlp.EncodeToBytes(params.Details)
	amount, _ := rlp.EncodeToBytes(params.Amount)
	rewardPer, _ := rlp.EncodeToBytes(params.RewardPer)
	programVersion, _ := rlp.EncodeToBytes(params.ProgramVersion)
	programVersionSign, _ := rlp.EncodeToBytes(params.ProgramVersionSign)
	blsPubKey, _ := rlp.EncodeToBytes(params.BlsPubKey)
	blsProof, _ := rlp.EncodeToBytes(params.BlsProof)

	res = append(res, typ)
	res = append(res, benefitAddress)
	res = append(res, nodeId)
	res = append(res, externalId)
	res = append(res, nodeName)
	res = append(res, website)
	res = append(res, details)
	res = append(res, amount)
	res = append(res, rewardPer)
	res = append(res, programVersion)
	res = append(res, programVersionSign)
	res = append(res, blsPubKey)
	res = append(res, blsProof)

	return res
}

func buildGovParams(funcType uint16, params *Dpos_2000) [][]byte {
	var res [][]byte
	res = make([][]byte, 0)
	fnType, _ := rlp.EncodeToBytes(funcType)
	res = append(res, fnType)

	verifier, _ := rlp.EncodeToBytes(params.Verifier)
	pipID, _ := rlp.EncodeToBytes(params.PIPID)
	res = append(res, verifier)
	res = append(res, pipID)
	return res
}

func funcTypeToContractAddress(funcType uint16) common.Address {
	toadd := common.ZeroAddr
	switch {
	case 0 < funcType && funcType < 2000:
		toadd = vm.StakingContractAddr
	case funcType >= 2000 && funcType < 3000:
		toadd = vm.GovContractAddr
	case funcType >= 3000 && funcType < 4000:
		toadd = vm.SlashingContractAddr
	case funcType >= 4000 && funcType < 5000:
		toadd = vm.RestrictingContractAddr
	case funcType >= 5000 && funcType < 6000:
		toadd = vm.DelegateRewardPoolAddr
	}
	return toadd
}

func netCheck(context *cli.Context) error {
	//hrp := context.String(addressHRPFlag.Name)
	//if err := common.SetAddressHRP(hrp); err != nil {
	//	return err
	//}
	return nil
}

func query(c *cli.Context, funcType uint16, params ...interface{}) error {
	url := c.String(rpcUrlFlag.Name)
	if url == "" {
		return errors.New("rpc url not set")
	}
	if c.Bool(jsonFlag.Name) {
		res, err := BuildDPosContract(funcType, params...)
		if err != nil {
			return err
		}
		fmt.Println(string(res))
		return nil
	} else {
		client, err := ethclient.Dial(url)
		if err != nil {
			return err
		}
		res, err := CallDPosContract(client, funcType, params...)
		if err != nil {
			return err
		}
		fmt.Println(string(res))
		return nil
	}
}

func getBlsProof(keyfilepath string) (bls.SchnorrProofHex,error) {
	var proofHex bls.SchnorrProofHex
	blsKey,err:=bls.LoadBLS(keyfilepath)
	if err != nil {
		return proofHex,fmt.Errorf("bls.LoadBLS error,%s", err.Error())
	}
	proof, _ := blsKey.MakeSchnorrNIZKP()
	proofByte, _ := proof.MarshalText()
	proofHex.UnmarshalText(proofByte)
	return proofHex,nil
}

func GetNodeKey(file string) (string, error)  {
	buf := make([]byte, 64)
	fd, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer fd.Close()
	if _, err := io.ReadFull(fd, buf); err != nil {
		return "", err
	}
	return string(buf),nil
}

// promptPassphrase prompts the user for a passphrase.  Set confirmation to true
// to require the user to confirm the passphrase.
func promptPassphrase(confirmation bool) string {
	passphrase, err := web.Stdin.PromptPassword("Passphrase: ")
	if err != nil {
		utils.Fatalf("Failed to read passphrase: %v", err)
	}

	if confirmation {
		confirm, err := web.Stdin.PromptPassword("Repeat passphrase: ")
		if err != nil {
			utils.Fatalf("Failed to read passphrase confirmation: %v", err)
		}
		if passphrase != confirm {
			utils.Fatalf("Passphrases do not match")
		}
	}

	return passphrase
}

// getPassphrase obtains a passphrase given by the user.
func getPassphrase(confirmation bool) string {

	// Otherwise prompt the user for the passphrase.
	return promptPassphrase(confirmation)
}

func getPrivateKey(keystorePath string) (*ecdsa.PrivateKey,string){
	// Read key from file.
	keyjson, err := ioutil.ReadFile(keystorePath)
	if err != nil {
		utils.Fatalf("Failed to read the keyfile at '%s': %v", keystorePath, err)
	}

	// Decrypt key with passphrase.
	passphrase := getPassphrase(false)
	key, err := keystore.DecryptKey(keyjson, passphrase)
	if err != nil {
		utils.Fatalf("Error decrypting key: %v", err)
	}
	return key.PrivateKey,key.Address.String()
}