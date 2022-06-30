package lib

import (
	"Phoenix-Chain-Core/ethereum/core/types"
	"Phoenix-Chain-Core/libs/common"
	"Phoenix-Chain-Core/libs/crypto"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
)

type Credentials struct {
	privateKey *ecdsa.PrivateKey
}

func NewCredential(hexPrivateKeyStr string) (*Credentials, error) {
	if len(hexPrivateKeyStr) < 64 {
		return nil, fmt.Errorf("length of private key is not fullfilled")
	}

	if strings.HasPrefix(hexPrivateKeyStr, "0x") || strings.HasPrefix(hexPrivateKeyStr, "0X") {
		hexPrivateKeyStr = hexPrivateKeyStr[2:]
	}

	key, err := crypto.HexToECDSA(hexPrivateKeyStr)
	if err != nil {
		return nil, err
	}

	return &Credentials{key}, nil
}

func (c *Credentials) Address() common.Address {
	return crypto.PubkeyToAddress(c.privateKey.PublicKey)
}

func (c *Credentials) SignTx(tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	signer := types.NewEIP155Signer(chainID)
	signedTx, err := types.SignTx(tx, signer, c.privateKey)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func (c *Credentials) MustStringToAddress(addr string) common.Address {
	return common.MustStringToAddress(addr)
}

func (c *Credentials) PrivateKey() *ecdsa.PrivateKey {
	return c.privateKey
}
