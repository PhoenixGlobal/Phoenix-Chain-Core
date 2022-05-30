package configs

import (
	"Phoenix-Chain-Core/ethereum/p2p/discover"
	"Phoenix-Chain-Core/libs/crypto/bls"
	"fmt"
	"math/big"
	"reflect"
	"testing"
)

func TestCheckCompatible(t *testing.T) {
	type test struct {
		stored, new *ChainConfig
		head        uint64
		wantErr     *ConfigCompatError
	}
	tests := []test{
		{stored: AllEthashProtocolChanges, new: AllEthashProtocolChanges, head: 0, wantErr: nil},
		{stored: AllEthashProtocolChanges, new: AllEthashProtocolChanges, head: 100, wantErr: nil},
		{
			stored:  &ChainConfig{EIP155Block: big.NewInt(10)},
			new:     &ChainConfig{EIP155Block: big.NewInt(20)},
			head:    9,
			wantErr: nil,
		},
		{
			stored: AllEthashProtocolChanges,
			new:    &ChainConfig{EIP155Block: nil},
			head:   3,
			wantErr: &ConfigCompatError{
				What:         "EIP155 fork block",
				StoredConfig: big.NewInt(0),
				NewConfig:    nil,
				RewindTo:     0,
			},
		},
		{
			stored: AllEthashProtocolChanges,
			new:    &ChainConfig{EIP155Block: big.NewInt(1)},
			head:   3,
			wantErr: &ConfigCompatError{
				What:         "EIP155 fork block",
				StoredConfig: big.NewInt(0),
				NewConfig:    big.NewInt(1),
				RewindTo:     0,
			},
		},
	}

	for _, test := range tests {
		err := test.stored.CheckCompatible(test.new, test.head)
		if !reflect.DeepEqual(err, test.wantErr) {
			t.Errorf("error mismatch:\nstored: %v\nnew: %v\nhead: %v\nerr: %v\nwant: %v", test.stored, test.new, test.head, err, test.wantErr)
		}
	}
}

func TestInit(t *testing.T) {
	for _, n := range initialChosenTestnetDPosNodes {
		validator := new(Validator)
		if node, err := discover.ParseNode(n.Enode); nil == err {
			validator.NodeId = node.ID
		}
		fmt.Println(11111,validator.NodeId)
		if n.BlsPubkey != "" {
			var blsPk bls.PublicKeyHex
			if err := blsPk.UnmarshalText([]byte(n.BlsPubkey)); nil == err {
				validator.BlsPubKey = blsPk
			}
		}
		fmt.Println(22222,validator.BlsPubKey)
		InitialChosenTestnetValidators = append(InitialChosenTestnetValidators, *validator)
	}
}