package commands

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/fairfieldfootball/league/backend/common"
	"github.com/fairfieldfootball/league/backend/core"
	"github.com/fairfieldfootball/league/backend/core/mongodb"
	"github.com/fairfieldfootball/league/backend/core/yahoo"

	"github.com/drone/envsubst"
	"github.com/joho/godotenv"
	cli "github.com/urfave/cli/v2"
)

type services struct {
	config        common.Config
	mongoProvider mongodb.Provider

	userStore   core.UserStore
	authService core.AuthService

	yahooClient  yahoo.Client
	yahooService *yahoo.Service
}

func newServices(cliCtx *cli.Context) (services, error) {
	if err := godotenv.Load(); err != nil {
		return services{}, err
	}

	var config common.Config

	dataRaw, err := ioutil.ReadFile(cliCtx.String("config"))
	if err != nil {
		return services{}, err
	}

	data, err := envsubst.EvalEnv(string(dataRaw))
	if err != nil {
		return services{}, err
	}

	if err := json.Unmarshal([]byte(data), &config); err != nil {
		return services{}, err
	}

	if err := config.Validate(); err != nil {
		return services{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), common.TimeoutServerOp)
	defer cancel()

	mongoProvider := mongodb.NewProvider(config.DB.URI)
	if err := mongoProvider.Setup(ctx); err != nil {
		return services{}, err
	}

	userStore, err := core.NewUserStore(mongoProvider.Client())
	if err != nil {
		return services{}, err
	}

	passwordStore, err := core.NewPasswordStore(mongoProvider.Client())
	if err != nil {
		return services{}, err
	}

	refreshTokenStore, err := core.NewRefreshTokenStore(mongoProvider.Client())
	if err != nil {
		return services{}, err
	}

	authService := core.NewAuthService(
		config,
		userStore,
		passwordStore,
		refreshTokenStore,
	)

	yahooClient := yahoo.NewClient(config)
	yahooRefreshTokenStore, err := yahoo.NewRefreshTokenStore(mongoProvider.Client())
	if err != nil {
		return services{}, err
	}

	yahooService := yahoo.NewService(yahooClient, yahooRefreshTokenStore)

	return services{
		config:        config,
		mongoProvider: mongoProvider,

		userStore:   userStore,
		authService: authService,

		yahooClient:  yahooClient,
		yahooService: &yahooService,
	}, nil
}
