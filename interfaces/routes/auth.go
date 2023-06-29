package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nurmanhabib/go-oauth2-server/interfaces/handler"
)

func authRoute(e *gin.Engine) {
	h := handler.NewAuthHandler()

	e.GET("/login", h.ShowLoginForm)
	e.POST("/login", h.Authenticate)
}
