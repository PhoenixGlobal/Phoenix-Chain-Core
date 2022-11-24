package chequebook

import (
	"errors"
	"math/big"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common"
)

const Version = "1.0"

var errNoChequebook = errors.New("no chequebook")

type API struct {
	chequebookf func() *Chequebook
}

func NewAPI(ch func() *Chequebook) *API {
	return &API{ch}
}

func (a *API) Balance() (string, error) {
	ch := a.chequebookf()
	if ch == nil {
		return "", errNoChequebook
	}
	return ch.Balance().String(), nil
}

func (a *API) Issue(beneficiary common.Address, amount *big.Int) (cheque *Cheque, err error) {
	ch := a.chequebookf()
	if ch == nil {
		return nil, errNoChequebook
	}
	return ch.Issue(beneficiary, amount)
}

func (a *API) Cash(cheque *Cheque) (txhash string, err error) {
	ch := a.chequebookf()
	if ch == nil {
		return "", errNoChequebook
	}
	return ch.Cash(cheque)
}

func (a *API) Deposit(amount *big.Int) (txhash string, err error) {
	ch := a.chequebookf()
	if ch == nil {
		return "", errNoChequebook
	}
	return ch.Deposit(amount)
}
