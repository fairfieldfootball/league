package main

import (
	"fmt"
	"os"

	"github.com/fairfieldfootball/league/backend/cmd/yahoo/commands"

	cli "github.com/urfave/cli/v2"
)

func main() {
	cmd := &cli.App{
		Name:  "yahoo",
		Usage: "interact with the yahoo api",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config",
				Usage: "the config file to load",
			},
		},
		Commands: []*cli.Command{
			commands.Auth,
			commands.FFC,
		},
	}

	if err := cmd.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
