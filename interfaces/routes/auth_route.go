package routes

import (
	"github.com/gin-gonic/gin"
	authHandler "github.com/nurmanhabib/go-oauth2-server/interfaces/handler/auth"
	"github.com/nurmanhabib/go-oauth2-server/pkg/auth"
)

func loginRoute(e *gin.Engine, userProvider auth.UserProvider) {
	h := authHandler.LoginHandler{
		UserProvider: userProvider,
	}

	e.GET("/login", h.Index)
	e.POST("/login", h.Submit)
}
