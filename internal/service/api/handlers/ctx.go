package handlers

import (
	"context"
	"github.com/FixIT-hackathon/meta-transfer-from/internal/data"
	"net/http"
)

type ctxKey int

const (
	transfersCtxKey = 0
)

func TransfersQ(r *http.Request) data.Transfers {
	return r.Context().Value(transfersCtxKey).(data.Transfers)
}

func CtxTransfersQ(entry data.Transfers) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, transfersCtxKey, entry)
	}
}
