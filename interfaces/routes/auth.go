package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nurmanhabib/go-oauth2-server/interfaces/handler"
	"github.com/nurmanhabib/go-oauth2-server/interfaces/provider"
)

func authRoute(e *gin.Engine, provider *provider.Provider) {
	h := handler.NewAuthHandler(provider)

	e.GET("/login", h.ShowLoginForm)
	e.POST("/login", h.Authenticate)
}
