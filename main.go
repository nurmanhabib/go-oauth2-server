package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/dao"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/oauth"
	"github.com/nurmanhabib/go-oauth2-server/interfaces/provider"
	"github.com/nurmanhabib/go-oauth2-server/interfaces/routes"
	"github.com/nurmanhabib/go-oauth2-server/pkg/passport"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/go_oauth2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if errMigrate := db.AutoMigrate(
		&entity.User{},
		&entity.OAuthClient{},
		&entity.OAuthAccessGrant{},
		&entity.OAuthAccessToken{},
	); errMigrate != nil {
		panic(errMigrate)
	}

	repo := dao.NewRepo(db)
	builder := oauth.NewServerBuilder(repo)
	director := passport.NewDirector(builder)
	srv := director.Build()

	prov := &provider.Provider{
		Repo:   repo,
		Server: srv,
	}

	app := cli.NewApp()
	app.Action = func(ctx *cli.Context) error {
		router := routes.NewRouter(prov).Init()

		appPort := os.Getenv("APP_PORT")
		if appPort == "" {
			appPort = "8080"
		}

		return router.Run(":" + appPort)
	}

	app.Commands = cli.Commands{
		{
			Name: "client:create",
			Action: func(c *cli.Context) error {
				client := &entity.OAuthClient{
					ID:          uuid.New().String(),
					Secret:      "PGqJBBNqJbDf",
					Name:        "Example App",
					SuperApp:    false,
					RedirectURI: "https://oauthdebugger.com/debug",
					Scopes:      "public",
				}

				if err := repo.OAuthClientRepo.Save(c.Context, client); err != nil {
					return err
				}

				fmt.Printf("Client created\n")
				fmt.Printf("Client ID: %s\n", client.ID)
				fmt.Printf("Client Secret: %s\n", client.Secret)

				return nil
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatalf("Error executing command, err: %v", err)
	}
}
