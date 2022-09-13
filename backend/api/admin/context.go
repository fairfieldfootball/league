package admin

import (
	"context"

	"github.com/fairfieldfootball/league/backend/api"
	"github.com/fairfieldfootball/league/backend/common"
	"github.com/fairfieldfootball/league/backend/core"
	"github.com/fairfieldfootball/league/backend/core/yahoo"
)

type ctxKey int

const (
	ctxKeyServerContext ctxKey = iota
)

type ServerContext struct {
	Config common.Config

	AuthService  *core.AuthService
	YahooService *yahoo.Service

	// RefreshTokenStore *core.RefreshTokenStore
	// PasswordStore *core.PasswordStore
	// UserStore     *core.UserStore
}

func AttachServerContext(ctx context.Context, srvCtx ServerContext) context.Context {
	return context.WithValue(ctx, ctxKeyServerContext, srvCtx)
}

func MustHaveServerContext(r api.Contexter) ServerContext {
	srvCtx, ok := r.Context().Value(ctxKeyServerContext).(ServerContext)
	if !ok {
		panic("must have admin server context")
	}
	return srvCtx
}
