package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/nurmanhabib/go-oauth2-server/config"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/dao"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/provider/auth/userprovider"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/provider/connection"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/provider/hashing"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/provider/oauth"
	"github.com/nurmanhabib/go-oauth2-server/interfaces/cmd"
	"github.com/nurmanhabib/go-oauth2-server/interfaces/routes"
	"github.com/nurmanhabib/go-oauth2-server/util"
	"github.com/urfave/cli/v2"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Println("no .env file provided")
	}

	conf := config.New(
		config.WithDatabase(),
	)

	db, errDB := connection.NewDBConnection(conf)
	if errDB != nil {
		log.Fatalln(errDB)
	}

	repos := dao.New(db)

	app := cmd.New()
	app.Action = func(ctx *cli.Context) error {
		oauthDep := &oauth.Dependency{
			Repo: oauth.Repo{
				AccessGrantRepo: repos.OauthAccessGrant,
				AccessTokenRepo: repos.OauthAccessToken,
				ClientRepo:      repos.OauthClient,
				UserRepo:        repos.User,
			},
		}

		oauthConfig := oauth.NewConfig(oauthDep)
		oauthSrv := oauth.NewServer(oauthConfig)

		hasher := &hashing.Bcrypt{}
		userProvider := userprovider.NewDatabaseUserProvider(repos.User, hasher)

		r := routes.New(
			routes.WithOAuthServer(oauthSrv),
			routes.WithRepositories(repos),
			routes.WithAuthUserProvider(userProvider),
		)

		// shutdownTimeout is probably better stored in a config file.
		// time.ParseDuration can easily parse duration in strings formatted (such as "5s", "10ms")
		// into time.Duration
		shutdownTimeout := 10 * time.Second

		addr := fmt.Sprintf(":%d", conf.Port)
		err := util.RunHTTPServerWithGracefulShutdown(r, addr, shutdownTimeout)
		if err != nil {
			return err
		}

		return nil
	}

	app.Commands = cmd.Commands(conf, db)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Unable to run CLI command, err: %v", err)
	}
}
