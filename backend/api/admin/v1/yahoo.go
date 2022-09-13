package v1

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/fairfieldfootball/league/backend/api"
	"github.com/fairfieldfootball/league/backend/api/admin"
	"github.com/fairfieldfootball/league/backend/auth"
	"github.com/fairfieldfootball/league/backend/common"
	"github.com/fairfieldfootball/league/backend/core/yahoo"
)

type authenticateRequest struct {
	Code string `bson:"code"`
}

func mustHaveYahooSession(r *http.Request) yahoo.Session {
	return yahoo.Session{
		AccessToken: api.MustHaveYahooToken(r),
		UserID:      api.CtxUserID(r),
	}
}

func YahooGame(w http.ResponseWriter, r *http.Request) {
	srvCtx := admin.MustHaveServerContext(r)

	yahooService := srvCtx.YahooService.WithSession(mustHaveYahooSession(r))

	ffc, err := yahooService.FFC(r.Context())
	if err != nil {
		api.ErrorResponse(w, r, err)
		return
	}

	if sess, ok := yahooService.EndSession(); ok {
		http.SetCookie(w, &http.Cookie{
			Name:     auth.CookieYahooToken,
			Value:    sess.AccessToken,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
			Secure:   srvCtx.Config.Server.SSLEnabled,
			Path:     "/",
			Expires:  time.Now().Add(time.Duration(sess.AccessTokenExpiry) * time.Second),
		})
	}

	api.JSONResponse(w, r, 0, ffc)
}

func YahooSessionBegin(w http.ResponseWriter, r *http.Request) {
	srvCtx := admin.MustHaveServerContext(r)
	userID := api.CtxUserID(r)

	defer r.Body.Close()

	var req authenticateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		api.ErrorResponse(w, r, common.NewErr("failed to decode request", common.ErrCodeBadRequest))
		return
	}

	tokens, err := srvCtx.YahooService.Authenticate(r.Context(), userID, req.Code)
	if err != nil {
		api.ErrorResponse(w, r, err)
		return
	}

	writeAuth(w, tokens, srvCtx.Config.Server.SSLEnabled)

	api.Response(w, r, http.StatusCreated)
}

func writeAuth(w http.ResponseWriter, tokens yahoo.Auth, secure bool) {
	http.SetCookie(w, &http.Cookie{
		Name:     auth.CookieYahooToken,
		Value:    tokens.AccessToken,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   secure,
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(tokens.ExpiresIn) * time.Second),
	})
}

func YahooSessionEnd(w http.ResponseWriter, r *http.Request) {
	srvCtx := admin.MustHaveServerContext(r)
	userID := api.CtxUserID(r)

	if err := srvCtx.YahooService.ClearSession(r.Context(), userID); err != nil {
		api.ErrorResponse(w, r, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     auth.CookieYahooToken,
		Value:    "",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   srvCtx.Config.Server.SSLEnabled,
		Path:     "/",
		MaxAge:   -1,
	})

	api.Response(w, r, 0)
}
