package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/nurmanhabib/go-oauth2-server/config"
	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/dao"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/provider/auth/userprovider"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/provider/connection"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/provider/hashing"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/provider/oauth"
	"github.com/nurmanhabib/go-oauth2-server/interfaces/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file provided")
	}

	conf := config.New(
		config.WithDatabase(),
	)

	db, errDB := connection.NewDBConnection(conf)
	if errDB != nil {
		log.Fatalln(errDB)
	}

	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.OauthClient{})
	db.AutoMigrate(&entity.OauthAccessGrant{})
	db.AutoMigrate(&entity.OauthAccessToken{})

	repos := dao.New(db)

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

	log.Fatal(r.Run(fmt.Sprintf(":%d", conf.Port)))
}
