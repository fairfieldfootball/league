package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fairfieldfootball/league/backend/core/yahoo"

	"github.com/pkg/browser"
	cli "github.com/urfave/cli/v2"
)

var (
	Auth = &cli.Command{
		Name:  "auth",
		Usage: "authenticate with yahoo",
		Subcommands: []*cli.Command{
			{
				Name:   "getcode",
				Usage:  "get the oauth code",
				Action: getCode,
			},
			{
				Name:  "login",
				Usage: "login with the oauth code",
				Flags: []cli.Flag{&cli.StringFlag{
					Name:     "code",
					Usage:    "oauth code",
					Required: true,
				}},
				Action: login,
			},
			{
				Name:   "refresh",
				Usage:  "refresh the access token",
				Action: refresh,
			},
		},
	}
)

const (
	fileAuth = "auth.json"
)

func loadAuth() (yahoo.Auth, error) {
	data, err := os.ReadFile(fileAuth)
	if err != nil {
		return yahoo.Auth{}, err
	}
	var auth yahoo.Auth
	if err := json.Unmarshal(data, &auth); err != nil {
		return yahoo.Auth{}, err
	}
	return auth, nil
}

func writeAuth(auth yahoo.Auth) error {
	authData, err := json.Marshal(auth)
	if err != nil {
		return err
	}
	if err := os.WriteFile(fileAuth, authData, 0666); err != nil {
		return err
	}
	return nil
}

func getCode(cliCtx *cli.Context) error {
	services, err := newServices(cliCtx)
	if err != nil {
		return err
	}

	url := fmt.Sprintf(
		"https://api.login.yahoo.com/oauth2/request_auth?client_id=%s&redirect_uri=oob&response_type=code",
		services.config.Yahoo.ClientID,
	)
	return browser.OpenURL(url)
}

func login(cliCtx *cli.Context) error {
	services, err := newServices(cliCtx)
	if err != nil {
		return err
	}

	code := cliCtx.String("code")

	auth, err := services.yahooClient.Auth(yahoo.TokensActionGet, code)
	if err != nil {
		return err
	}

	return writeAuth(auth)
}

func refresh(cliCtx *cli.Context) error {
	auth, err := loadAuth()
	if err != nil {
		return err
	}

	services, err := newServices(cliCtx)
	if err != nil {
		return err
	}

	newAuth, err := services.yahooClient.Auth(yahoo.TokensActionRefresh, auth.RefreshToken)
	if err != nil {
		return err
	}
	return writeAuth(newAuth)
}
