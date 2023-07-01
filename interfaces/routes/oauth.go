package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nurmanhabib/go-oauth2-server/interfaces/handler"
	"github.com/nurmanhabib/go-oauth2-server/interfaces/provider"
	"github.com/nurmanhabib/go-oauth2-server/interfaces/routes/middleware"
)

func oauthRoute(e *gin.Engine, provider *provider.Provider) {
	h := handler.NewOAuthHandler(provider.Server)

	g := e.Group("/oauth", middleware.Authenticated(provider.Repo))

	g.GET("/authorize", h.ShowUserConsent)
	g.POST("/approve", h.Authorize)
	g.POST("/token", h.TokenExchange)
}
