package routes

import (
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/dao"
)

type Provider struct {
	Repo   *dao.Repositories
	Server *server.Server
}
