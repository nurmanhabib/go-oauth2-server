package ping

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Pong is handler to response ping request.
func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":  "pong",
		"datetime": time.Now(),
	})
}
