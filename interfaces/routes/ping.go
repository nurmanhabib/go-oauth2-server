package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nurmanhabib/go-oauth2-server/interfaces/handler/ping"
)

func pingRoute(e *gin.Engine) {
	e.GET("/ping", ping.Pong)
}
