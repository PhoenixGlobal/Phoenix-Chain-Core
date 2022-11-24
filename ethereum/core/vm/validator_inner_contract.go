package vm

import (
	"errors"
	"fmt"

	"bytes"

	"encoding/json"

	"encoding/binary"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/configs"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/p2p/discover"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/crypto"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/crypto/bls"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/log"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/rlp"
)

const (
	CurrentValidatorKey = "current_validator"
	NextValidatorKey    = "next_validator"

	txTypeUpdate  = 2000
	txTypeCurrent = 2001
	txTypeNext    = 2002
	txTypeSwitch  = 2003
)

type ValidateNode struct {
	Index     uint               `json:"index"`
	NodeID    discover.NodeID    `json:"nodeID"`
	Address   common.NodeAddress `json:"-"`
	BlsPubKey bls.PublicKey      `json:"blsPubKey"`
}

type NodeList []*ValidateNode

func (nl *NodeList) String() string {
	s := ""
	for _, v := range *nl {
		s = s + fmt.Sprintf("{Index: %d NodeID: %s Address: %s blsPubKey: %s},", v.Index, v.NodeID, v.Address.String(), fmt.Sprintf("%x", v.BlsPubKey.Serialize()))
	}
	return s
}

type Validators struct {
	ValidateNodes    NodeList `json:"validateNodes"`
	ValidBlockNumber uint64   `json:"-"`
}

func (vds *Validators) String() string {
	return fmt.Sprintf("{validateNodes: [%s] validBlockNumber: %d}", vds.ValidateNodes.String(), vds.ValidBlockNumber)
}

type ValidatorInnerContractBase interface {
	UpdateValidators(validators *Validators) error
	CurrentValidators() (*Validators, error)
	NextValidators() (*Validators, error)
	SwitchValidators(validBlockNumber uint64) error
}

type validatorInnerContract struct {
	ValidatorInnerContractBase

	Contract *Contract
	Evm      *EVM
}

func (vic *validatorInnerContract) RequiredGas(input []byte) uint64 {
	return configs.EcrecoverGas
}

func (vic *validatorInnerContract) Run(input []byte) ([]byte, error) {
	var cmd = map[string]interface{}{
		"UpdateValidators":  vic.UpdateValidators,
		"CurrentValidatros": vic.CurrentValidators,
		"NextValidators":    vic.NextValidators,
		"SwitchValidators":  vic.SwitchValidators,
	}
	return vic.execute(input, cmd)
}

func (vic *validatorInnerContract) UpdateValidators(validators *Validators) error {
	if len(validators.ValidateNodes) <= 0 {
		log.Error("Empty validator nodes")
		return errors.New("Empty validator nodes")
	}

	var newVds Validators
	for _, node := range validators.ValidateNodes {
		pubkey, err := node.NodeID.Pubkey()
		if err != nil {
			log.Error("Get pubkey from nodeID fail", "error", err)
			return err
		}
		node.Address = crypto.PubkeyToNodeAddress(*pubkey)
		newVds.ValidateNodes = append(newVds.ValidateNodes, node)
	}
	log.Debug("Update validators", "validators", newVds.String(), "address", vic.Contract.Address())

	vs, err := rlp.EncodeToBytes(newVds)
	if err != nil {
		log.Error("RLP encode error", "validators", newVds.String(), "error", err)
		return err
	}
	vic.Evm.StateDB.SetState(vic.Contract.Address(), []byte(NextValidatorKey), vs)
	return nil
}

func (vic *validatorInnerContract) CurrentValidators() (*Validators, error) {
	state := vic.Evm.StateDB
	b := state.GetState(vic.Contract.Address(), []byte(CurrentValidatorKey))

	var vds Validators
	err := rlp.DecodeBytes(b, &vds)
	return &vds, err
}

func (vic *validatorInnerContract) NextValidators() (*Validators, error) {
	state := vic.Evm.StateDB
	b := state.GetState(vic.Contract.Address(), []byte(NextValidatorKey))

	var vds Validators
	err := rlp.DecodeBytes(b, &vds)
	return &vds, err
}

func (vic *validatorInnerContract) SwitchValidators(validBlockNumber uint64) error {
	state := vic.Evm.StateDB
	b := state.GetState(vic.Contract.Address(), []byte(NextValidatorKey))
	var nvs Validators
	err := rlp.DecodeBytes(b, &nvs)
	if err == nil {
		nvs.ValidBlockNumber = validBlockNumber
		b, _ = rlp.EncodeToBytes(nvs)
		state.SetState(vic.Contract.Address(), []byte(CurrentValidatorKey), b)
		log.Debug("Switch validators success", "number", vic.Evm.BlockNumber, "validators", nvs.String())
		return nil
	}

	log.Warn("Get next validators fail, try to get current validators", "error", err, "validBlockNumber", validBlockNumber)
	// Try to get current validators.
	b = state.GetState(vic.Contract.Address(), []byte(CurrentValidatorKey))
	err = rlp.DecodeBytes(b, &nvs)
	if err != nil {
		log.Error("RLP decode current Validators fail", "validBlockNumber", validBlockNumber, "error", err)
		return err
	}

	// There not next validators, so update ValidBlockNumber and setting current as next.
	nvs.ValidBlockNumber = validBlockNumber
	b, _ = rlp.EncodeToBytes(nvs)
	state.SetState(vic.Contract.Address(), []byte(CurrentValidatorKey), b)
	log.Debug("Switch validators success", "number", vic.Evm.BlockNumber, "validators", nvs.String())
	return nil
}

func (vic *validatorInnerContract) execute(input []byte, cmd map[string]interface{}) (ret []byte, err error) {
	defer func() {
		if er := recover(); er != nil {
			ret, err = nil, fmt.Errorf("Validator inner contract execute fail: %v", er)
			log.Error("Validator inner contract execute fail", "error", err)
		}
	}()

	var source [][]byte
	if err = rlp.Decode(bytes.NewReader(input), &source); err != nil {
		log.Error("Validator inner contract execute fail", "error", err)
		return nil, errors.New("RLP decode fail")
	}

	if len(source) < 2 {
		log.Error("Params base length not match")
		return nil, errors.New("Params base length not match")
	}

	funcName := string(source[1])
	if _, ok := cmd[funcName]; !ok {
		log.Error("Function undefined", "function", funcName)
		return nil, errors.New("Function undefined")
	}

	txType := common.BytesToInt64(source[0])
	switch txType {
	case txTypeUpdate:
		var vds Validators
		err = json.Unmarshal(source[2], &vds)
		if err != nil {
			log.Error("Parse params fail", "params", string(source[2]), "error", err)
			return nil, err
		}
		err = vic.UpdateValidators(&vds)
		return nil, err

	case txTypeCurrent:
		var vds *Validators = nil
		vds, err = vic.CurrentValidators()
		if err != nil {
			log.Error("Get current validators fail", "error", err)
			return nil, err
		}
		b, _ := json.Marshal(&vds)
		return b, nil

	case txTypeNext:
		var vds *Validators = nil
		vds, err = vic.NextValidators()
		if err != nil {
			log.Error("Get next validators fail", "error", err)
			return nil, err
		}
		b, _ := json.Marshal(&vds)
		return b, nil
	case txTypeSwitch:
		log.Debug("Switch validators", "source", len(source))
		validBlockNumber := binary.BigEndian.Uint64(source[2])
		return nil, vic.SwitchValidators(validBlockNumber)
	default:
		log.Error("Unexpected transaction type", "txType", txType)
		return nil, errors.New("unexpected transaction type")
	}
}
