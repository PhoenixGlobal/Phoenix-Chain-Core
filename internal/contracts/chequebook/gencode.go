// +build none

// This program generates contract/code.go, which contains the chequebook code
// after deployment.
package main

import (
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/internal/contracts/chequebook/contract"
	"fmt"
	"io/ioutil"
	"math/big"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/accounts/abi/bind"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/accounts/abi/bind/backends"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/crypto"
)

var (
	testKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	testAlloc  = core.GenesisAlloc{
		crypto.PubkeyToAddress(testKey.PublicKey): {Balance: big.NewInt(500000000000)},
	}
)

func main() {
	backend := backends.NewSimulatedBackend(testAlloc, uint64(100000000))
	auth := bind.NewKeyedTransactor(testKey)

	// Deploy the contract, get the code.
	addr, _, _, err := contract.DeployChequebook(auth, backend)
	if err != nil {
		panic(err)
	}
	backend.Commit()
	code, err := backend.CodeAt(nil, addr, nil)
	if err != nil {
		panic(err)
	}
	if len(code) == 0 {
		panic("empty code")
	}

	// Write the output file.
	content := fmt.Sprintf(`package contract

// ContractDeployedCode is used to detect suicides. This constant needs to be
// updated when the contract code is changed.
const ContractDeployedCode = "%#x"
`, code)
	if err := ioutil.WriteFile("contract/code.go", []byte(content), 0644); err != nil {
		panic(err)
	}
}
