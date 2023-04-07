package handlers

import (
	"context"
	"net/http"

	"gitlab.com/distributed_lab/acs/role-svc/internal/config"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	linksParamCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxLinksParams(entry *config.LinksCfg) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, linksParamCtxKey, entry)
	}
}

func LinksParams(r *http.Request) *config.LinksCfg {
	return r.Context().Value(linksParamCtxKey).(*config.LinksCfg)
}
