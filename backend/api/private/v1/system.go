package v1

import (
	"net/http"
	"time"

	"github.com/fairfieldfootball/league/backend/api"
	"github.com/fairfieldfootball/league/backend/common"
)

func GetHealth(w http.ResponseWriter, r *http.Request) {
	api.Response(w, r, http.StatusNoContent)
}

func GetVersion(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	env, gitHash, buildTime := common.ServerVersion()

	api.JSONResponse(w, r, 0, map[string]interface{}{
		"env":         env,
		"last_commit": gitHash,
		"build_time":  buildTime,
		"time":        now,
	})
}
