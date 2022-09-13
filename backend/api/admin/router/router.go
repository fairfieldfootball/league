package router

import (
	"net/http"

	"github.com/fairfieldfootball/league/backend/api"
	"github.com/fairfieldfootball/league/backend/api/admin/v1"
)

const (
	pathV1 = "/v1"

	pathUser        = "/user"
	pathUserSession = "/user/session"

	pathYahooFFC     = "/yahoo/ffc"
	pathYahooSession = "/yahoo/session"
)

// var (
// 	pathYahooGame = "/yahoo/games/" + api.PathVar(v1.ParamGame)
// )

var (
	Versions = []string{pathV1}

	Registry = map[string][]api.RouteRegistration{
		pathV1: {
			// auth routes
			{
				v1.Whoami,
				api.RouteEndpoint{http.MethodGet, pathUser, false},
				api.RouteNeedsSession,
			},
			{
				v1.UpdateUserData,
				api.RouteEndpoint{http.MethodPut, pathUser, true},
				api.RouteNeedsSession,
			},
			{
				v1.Register,
				api.RouteEndpoint{http.MethodPost, pathUser, false},
				api.RouteNeedsNothing,
			},
			{
				v1.Login,
				api.RouteEndpoint{http.MethodPost, pathUserSession, true},
				api.RouteNeedsNothing,
			},
			{
				v1.RefreshAccess,
				api.RouteEndpoint{http.MethodPut, pathUserSession, false},
				api.RouteNeedsRefreshToken,
			},
			{
				v1.Logout,
				api.RouteEndpoint{http.MethodDelete, pathUserSession, false},
				api.RouteNeedsNothing,
			},

			// yahoo routes
			{
				v1.YahooGame,
				api.RouteEndpoint{http.MethodGet, pathYahooFFC, false},
				api.RouteNeedsYahooToken,
			},
			{
				v1.YahooSessionBegin,
				api.RouteEndpoint{http.MethodPost, pathYahooSession, true},
				api.RouteNeedsNothing,
			},
			{
				v1.YahooSessionEnd,
				api.RouteEndpoint{http.MethodDelete, pathYahooSession, false},
				api.RouteNeedsNothing,
			},
		},
	}
)
