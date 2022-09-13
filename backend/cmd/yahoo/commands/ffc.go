package commands

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fairfieldfootball/league/backend/core/yahoo"
	cli "github.com/urfave/cli/v2"
)

var (
	FFC = &cli.Command{
		Name:   "ffc",
		Usage:  "get ffc history",
		Action: getFFC,
	}
)

func getFFC(cliCtx *cli.Context) error {
	auth, err := loadAuth()
	if err != nil {
		return err
	}

	services, err := newServices(cliCtx)
	if err != nil {
		return err
	}

	yahooService := services.yahooService.WithSession(yahoo.Session{
		AccessToken: auth.AccessToken,
	})

	ffc, err := yahooService.FFC(context.Background())
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(ffc, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(data))
	return nil
}
