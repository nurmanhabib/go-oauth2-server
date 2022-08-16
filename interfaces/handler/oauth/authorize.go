package oauth

import (
	"html/template"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
)

// Authorize is a handler to show user consent.
func (h *Handler) Authorize(c *gin.Context) {
	ctx := c.Request.Context()

	ar, err := h.srv.ValidationAuthorizeRequest(c.Request)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusUnauthorized)
		return
	}

	client, errClient := h.clientRepo.Find(ctx, &entity.OauthClient{ID: ar.ClientID})
	if errClient != nil {
		http.Error(c.Writer, errors.ErrInvalidClient.Error(), http.StatusBadRequest)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("resources", "views", "oauth_authorize.html"))
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"name":  client.Name,
		"query": template.URL(c.Request.URL.RawQuery),
	}

	err = tmpl.Execute(c.Writer, data)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}
