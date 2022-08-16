package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/nurmanhabib/go-oauth2-server/domain/repository"
	"github.com/nurmanhabib/go-oauth2-server/interfaces/handler/oauth"
)

func oauthRoute(e *gin.Engine, srv *server.Server, clientRepo repository.OauthClientRepository, middleware ...gin.HandlerFunc) {
	h := oauth.NewHandler(srv, clientRepo)
	g := e.Group("oauth", middleware...)

	g.GET("/authorize", h.Authorize)
	g.POST("/token", h.Token)
	g.POST("/approve", h.Approve)
	g.POST("/decline", h.Decline)
}
