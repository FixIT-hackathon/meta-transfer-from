package handlers

import (
	"context"
	"github.com/FixIT-hackathon/meta-transfer-from/internal/data"
	"github.com/ethereum/go-ethereum/ethclient"
	"net/http"
)

type ctxKey int

const (
	transfersCtxKey = 0
	ethClientCtxKey
)

func TransfersQ(r *http.Request) data.Transfers {
	return r.Context().Value(transfersCtxKey).(data.Transfers)
}

func CtxTransfersQ(entry data.Transfers) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, transfersCtxKey, entry)
	}
}

func EthClient(r *http.Request) *ethclient.Client {
	return r.Context().Value(ethClientCtxKey).(*ethclient.Client)
}

func CtxEthClient(entry *ethclient.Client) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ethClientCtxKey, entry)
	}
}
