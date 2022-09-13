package api

import (
	"context"

	"github.com/fairfieldfootball/league/backend/auth"
	"github.com/fairfieldfootball/league/backend/common"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ctxKey int

const (
	ctxKeyLogger ctxKey = iota
	ctxKeyRequestID
	ctxKeyUserToken
	ctxKeyAccessToken
	ctxKeyRefreshToken
	ctxKeyYahooToken
)

type Contexter interface {
	Context() context.Context
}

func CtxAccessToken(r Contexter) (auth.AccessToken, bool) {
	accessToken, ok := r.Context().Value(ctxKeyAccessToken).(auth.AccessToken)
	return accessToken, ok
}

func MustHaveAccessToken(r Contexter) auth.AccessToken {
	accessToken, ok := CtxAccessToken(r)
	if !ok {
		panic("must have access token")
	}
	return accessToken
}

func CtxLogger(r Contexter) (common.Logger, bool) {
	logger, ok := r.Context().Value(ctxKeyLogger).(common.Logger)
	return logger, ok
}

func MustHaveLogger(r Contexter) common.Logger {
	logger, ok := CtxLogger(r)
	if !ok {
		panic("must have logger")
	}
	return logger
}

func CtxRefreshToken(r Contexter) (auth.RefreshToken, bool) {
	refreshToken, ok := r.Context().Value(ctxKeyRefreshToken).(auth.RefreshToken)
	return refreshToken, ok
}

func MustHaveRefreshToken(r Contexter) auth.RefreshToken {
	refreshToken, ok := CtxRefreshToken(r)
	if !ok {
		panic("must have refresh token")
	}
	return refreshToken
}

func CtxYahooToken(r Contexter) (string, bool) {
	yahooToken, ok := r.Context().Value(ctxKeyYahooToken).(string)
	return yahooToken, ok
}

func MustHaveYahooToken(r Contexter) string {
	yahooToken, ok := CtxYahooToken(r)
	if !ok {
		panic("must have yahoo token")
	}
	return yahooToken
}

func CtxRequestID(r Contexter) (string, bool) {
	requestID, ok := r.Context().Value(ctxKeyRequestID).(string)
	return requestID, ok
}

func MustHaveRequestID(r Contexter) string {
	requestID, ok := CtxRequestID(r)
	if !ok {
		panic("must have request id")
	}
	return requestID
}

func CtxUser(r Contexter) (auth.User, bool) {
	userToken, ok := r.Context().Value(ctxKeyUserToken).(auth.User)
	return userToken, ok
}

func MustHaveUser(r Contexter) auth.User {
	userToken, ok := CtxUser(r)
	if !ok {
		panic("must have user token")
	}
	return userToken
}

func CtxUserID(r Contexter) primitive.ObjectID {
	if user, ok := CtxUser(r); ok {
		return user.ID
	}
	if refreshToken, ok := CtxRefreshToken(r); ok {
		return refreshToken.UserID
	}
	if accessToken, ok := CtxAccessToken(r); ok {
		return accessToken.UserID
	}
	return primitive.NilObjectID
}

type ContextBuilder interface {
	Attach(key, val interface{}) ContextBuilder
	Context() context.Context

	AttachAccessToken(accessToken auth.AccessToken) ContextBuilder
	AttachLogger(logger common.Logger) ContextBuilder
	AttachRequestID(requestID string) ContextBuilder
	AttachRefreshToken(refreshToken auth.RefreshToken) ContextBuilder
	AttachUserToken(userToken auth.User) ContextBuilder
	AttachYahooToken(yahooToken string) ContextBuilder
}

type contextBuilder struct {
	ctx context.Context
}

func NewContextBuilder(ctx context.Context) ContextBuilder {
	return &contextBuilder{ctx}
}

func (b *contextBuilder) Attach(key, val interface{}) ContextBuilder {
	b.ctx = context.WithValue(b.ctx, key, val)
	return b
}

func (b *contextBuilder) Context() context.Context {
	return b.ctx
}

func (b *contextBuilder) AttachAccessToken(accessToken auth.AccessToken) ContextBuilder {
	return b.Attach(ctxKeyAccessToken, accessToken)
}

func (b *contextBuilder) AttachLogger(logger common.Logger) ContextBuilder {
	return b.Attach(ctxKeyLogger, logger)
}

func (b *contextBuilder) AttachRefreshToken(refreshToken auth.RefreshToken) ContextBuilder {
	return b.Attach(ctxKeyRefreshToken, refreshToken)
}

func (b *contextBuilder) AttachRequestID(requestID string) ContextBuilder {
	return b.Attach(ctxKeyRequestID, requestID)
}

func (b *contextBuilder) AttachUserToken(userToken auth.User) ContextBuilder {
	return b.Attach(ctxKeyUserToken, userToken)
}

func (b *contextBuilder) AttachYahooToken(yahooToken string) ContextBuilder {
	return b.Attach(ctxKeyYahooToken, yahooToken)
}
