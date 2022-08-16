package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func defaultRoute(e *gin.Engine) {
	e.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/login")
	})
}
