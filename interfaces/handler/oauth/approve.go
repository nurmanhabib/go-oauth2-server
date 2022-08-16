package oauth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Approve is a handler to approve authorization request.
func (h *Handler) Approve(c *gin.Context) {
	r := c.Request
	w := c.Writer

	err := h.srv.HandleAuthorizeRequest(w, r)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
	}
}
