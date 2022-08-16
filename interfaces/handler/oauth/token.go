package oauth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Token is handler to token request.
func (h *Handler) Token(c *gin.Context) {
	err := h.srv.HandleTokenRequest(c.Writer, c.Request)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}
