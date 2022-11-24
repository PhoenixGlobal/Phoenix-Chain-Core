package core

import (
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/internal/ethapi"
	"context"

	"encoding/json"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/accounts"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common/hexutil"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/log"
)

type AuditLogger struct {
	log log.Logger
	api ExternalAPI
}

func (l *AuditLogger) List(ctx context.Context) (Accounts, error) {
	l.log.Info("List", "type", "request", "metadata", MetadataFromContext(ctx).String())
	res, e := l.api.List(ctx)

	l.log.Info("List", "type", "response", "data", res.String())

	return res, e
}

func (l *AuditLogger) New(ctx context.Context) (accounts.Account, error) {
	return l.api.New(ctx)
}

func (l *AuditLogger) SignTransaction(ctx context.Context, args SendTxArgs, methodSelector *string) (*ethapi.SignTransactionResult, error) {
	sel := "<nil>"
	if methodSelector != nil {
		sel = *methodSelector
	}
	l.log.Info("SignTransaction", "type", "request", "metadata", MetadataFromContext(ctx).String(),
		"tx", args.String(),
		"methodSelector", sel)

	res, e := l.api.SignTransaction(ctx, args, methodSelector)
	if res != nil {
		l.log.Info("SignTransaction", "type", "response", "data", common.Bytes2Hex(res.Raw), "error", e)
	} else {
		l.log.Info("SignTransaction", "type", "response", "data", res, "error", e)
	}
	return res, e
}

func (l *AuditLogger) Sign(ctx context.Context, addr common.MixedcaseAddress, data hexutil.Bytes) (hexutil.Bytes, error) {
	l.log.Info("Sign", "type", "request", "metadata", MetadataFromContext(ctx).String(),
		"addr", addr.String(), "data", common.Bytes2Hex(data))
	b, e := l.api.Sign(ctx, addr, data)
	l.log.Info("Sign", "type", "response", "data", common.Bytes2Hex(b), "error", e)
	return b, e
}

func (l *AuditLogger) EcRecover(ctx context.Context, data, sig hexutil.Bytes) (common.Address, error) {
	l.log.Info("EcRecover", "type", "request", "metadata", MetadataFromContext(ctx).String(),
		"data", common.Bytes2Hex(data))
	a, e := l.api.EcRecover(ctx, data, sig)
	l.log.Info("EcRecover", "type", "response", "addr", a.String(), "error", e)
	return a, e
}

func (l *AuditLogger) Export(ctx context.Context, addr common.Address) (json.RawMessage, error) {
	l.log.Info("Export", "type", "request", "metadata", MetadataFromContext(ctx).String(),
		"addr", addr)
	j, e := l.api.Export(ctx, addr)
	// In this case, we don't actually log the json-response, which may be extra sensitive
	l.log.Info("Export", "type", "response", "json response size", len(j), "error", e)
	return j, e
}

func (l *AuditLogger) Import(ctx context.Context, keyJSON json.RawMessage) (Account, error) {
	// Don't actually log the json contents
	l.log.Info("Import", "type", "request", "metadata", MetadataFromContext(ctx).String(),
		"keyJSON size", len(keyJSON))
	a, e := l.api.Import(ctx, keyJSON)
	l.log.Info("Import", "type", "response", "addr", a.String(), "error", e)
	return a, e
}

func NewAuditLogger(path string, api ExternalAPI) (*AuditLogger, error) {
	l := log.New("api", "signer")
	handler, err := log.FileHandler(path, log.LogfmtFormat())
	if err != nil {
		return nil, err
	}
	l.SetHandler(handler)
	l.Info("Configured", "audit log", path)
	return &AuditLogger{l, api}, nil
}
