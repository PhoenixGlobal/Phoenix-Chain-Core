package lib

import (
	phoenixchain "Phoenix-Chain-Core"
	"Phoenix-Chain-Core/commands/chaintool/core"
	"Phoenix-Chain-Core/ethereum/core/types"
	"Phoenix-Chain-Core/ethereum/ethclient"
	"Phoenix-Chain-Core/libs/common"
	"context"
	"encoding/json"
	"math/big"
)

type FunctionExecutor struct {
	httpEntry    string
	chainId      *big.Int
	contractAddr string
	credentials  *Credentials
}

type CallResponse struct {
	Code uint64          `json:"Code"`
	Ret  json.RawMessage `json:"Ret"`
}

func (fe *FunctionExecutor) SendWithRaw(f *Function) (string, error) {
	to := fe.credentials.MustStringToAddress(fe.contractAddr)
	data := f.ToBytes()

	gasPrice := fe.getDefaultGasPrice(f)
	gasLimit := fe.getDefaultGasLimit(f)
	chainId := fe.chainId

	r, err := fe.doSendRawTx(chainId, to, data, nil, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	//fmt.Println("[SendWithRaw] Http Response: " + string(r))
	return r, nil
}

func (fe *FunctionExecutor) SendWithResult(f *Function, result *string) error {
	raw, err := fe.SendWithRaw(f)
	if err != nil {
		return err
	}
	*result=raw
	return err
}

func (fe *FunctionExecutor) getDefaultGasPrice(f *Function) *big.Int {
	price := new(big.Int)
	switch f.Type {
	default:
		price = big.NewInt(0)
	}
	return price
}

func (fe *FunctionExecutor) getDefaultGasLimit(f *Function) uint64 {
	if IsLocalSupportFunction(f.Type) {
		return GetGasLimit(f)
	} else {
		return 0
	}
}

func (fe *FunctionExecutor) doSendRawTx(chainId *big.Int, to common.Address, data []byte, value *big.Int, gasPrice *big.Int, gasLimit uint64) (string, error) {
	client, err := ethclient.Dial(fe.httpEntry)
	if err != nil {
		return "", err
	}
	ctx := context.Background()

	if gasPrice.Cmp(big.NewInt(0)) == 0 {
		gasPrice, err = client.SuggestGasPrice(ctx)
		if err != nil {
			return "", err
		}
	}

	fromAddr := fe.credentials.Address()
	nonce:=core.GetNonce(fromAddr.String())
	//nonce, err := client.NonceAt(ctx, fromAddr, "pending")
	//if err != nil {
	//	return nil, err
	//}

	if gasLimit == 0 {
		msg := phoenixchain.CallMsg{
			From:     fromAddr,
			To:       &to,
			Gas:      0,
			GasPrice: gasPrice,
			Value:    value,
			Data:     data,
		}
		gasLimit, err = client.EstimateGas(ctx, msg)
		if err != nil {
			return "", err
		}
	}

	tx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, data)
	signedTx, err := fe.credentials.SignTx(tx, chainId)
	if err != nil {
		return "", err
	}
	err=client.SendTransaction(ctx, signedTx)
	hash:=signedTx.Hash().String()
	return hash,err
}

