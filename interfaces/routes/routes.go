package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/dao"
	"github.com/nurmanhabib/go-oauth2-server/interfaces/middleware"
	"github.com/nurmanhabib/go-oauth2-server/pkg/auth"
)

// Router is container router configuration needs.
type Router struct {
	oauthSrv         *server.Server
	repos            *dao.Repositories
	authUserProvider auth.UserProvider
}

// New build routes with gin.Engine.
func New(options ...RouterOption) *gin.Engine {
	r := &Router{}

	for _, apply := range options {
		apply(r)
	}

	// Middleware
	authWeb := middleware.NewAuthWeb(r.repos.User)

	e := gin.Default()

	defaultRoute(e)
	pingRoute(e)
	loginRoute(e, r.authUserProvider)

	oauthRoute(e, r.oauthSrv, r.repos.OauthClient, authWeb.Authenticate())

	return e
}
