package api

import (
	"net/http"
)

func PathVar(path string) string {
	return "{" + path + "}"
}

// https://go.dev/play/p/ze3l4tDCCQK
type RouteNeeds uint8

const (
	RouteNeedsUser RouteNeeds = 1 << iota
	RouteNeedsAccessToken
	RouteNeedsRefreshToken
	RouteNeedsYahooToken

	RouteNeedsSession = RouteNeedsAccessToken | RouteNeedsUser

	RouteNeedsNothing RouteNeeds = 0
)

type RouteRegistration struct {
	Handler  http.HandlerFunc
	Endpoint RouteEndpoint
	Needs    RouteNeeds
}

type RouteEndpoint struct {
	Method  string
	Path    string
	UseCORS bool
}
