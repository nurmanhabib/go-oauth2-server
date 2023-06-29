package handler

import (
	"html/template"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/server"
)

type OAuthHandler struct {
	srv *server.Server
}

func NewOAuthHandler(srv *server.Server) *OAuthHandler {
	return &OAuthHandler{srv: srv}
}

func (h *OAuthHandler) ShowUserConsent(c *gin.Context) {
	authorizeReq, err := h.srv.ValidationAuthorizeRequest(c.Request)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusUnauthorized)
		return
	}

	client, err := h.srv.Manager.GetClient(c, authorizeReq.ClientID)
	if err != nil {
		h.showError(c, err)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("interfaces", "views", "oauth_authorize.html"))
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"name":  client.GetID(),
		"query": template.URL(c.Request.URL.RawQuery),
	}

	err = tmpl.Execute(c.Writer, data)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (h *OAuthHandler) Authorize(c *gin.Context) {
	err := h.srv.HandleAuthorizeRequest(c.Writer, c.Request)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	}
}

func (h *OAuthHandler) showError(c *gin.Context, err error) {
	c.Status(http.StatusBadRequest)

	tmpl, errTmpl := template.ParseFiles(path.Join("interfaces", "views", "oauth_error.html"))
	if errTmpl != nil {
		http.Error(c.Writer, errTmpl.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"message": err.Error(),
	}

	errTmpl = tmpl.Execute(c.Writer, data)
	if errTmpl != nil {
		http.Error(c.Writer, errTmpl.Error(), http.StatusInternalServerError)
	}
}
