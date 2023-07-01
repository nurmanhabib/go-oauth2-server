package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/server"
)

type OAuthHandler struct {
	srv *server.Server
}

func NewOAuthHandler(srv *server.Server) *OAuthHandler {
	return &OAuthHandler{srv: srv}
}

func (h *OAuthHandler) redirectToDefaultClient(c *gin.Context) {
	defaultClientID := os.Getenv("DEFAULT_CLIENT_ID")
	client, err := h.srv.Manager.GetClient(c, defaultClientID)
	if err != nil {
		h.showError(c, err)
		return
	}

	domains := strings.Fields(client.GetDomain())
	if len(domains) == 0 {
		h.showError(c, fmt.Errorf("missing default redirect uri for client_id [%s]", client.GetID()))
		return
	}

	redirectURI := domains[0]

	params := url.Values{
		"client_id":     {client.GetID()},
		"redirect_uri":  {redirectURI},
		"scope":         {"public"},
		"response_type": {oauth2.Code.String()},
	}

	c.Redirect(http.StatusFound, "/oauth/authorize?"+params.Encode())
}

func (h *OAuthHandler) ShowUserConsent(c *gin.Context) {
	if c.Request.FormValue("client_id") == "" {
		h.redirectToDefaultClient(c)
		return
	}

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

	// Ensures that if it is considered a first-party application,
	// it is immediately authorized without displaying user consent.
	if clientApp, ok := client.(interface {
		IsSuperApp() bool
	}); ok && clientApp.IsSuperApp() {
		h.Authorize(c)
		return
	}

	c.HTML(http.StatusOK, "oauth_authorize.html", gin.H{
		"name":  client.GetID(),
		"query": template.URL(c.Request.URL.RawQuery),
	})
}

func (h *OAuthHandler) Authorize(c *gin.Context) {
	err := h.srv.HandleAuthorizeRequest(c.Writer, c.Request)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	}
}

func (h *OAuthHandler) TokenExchange(c *gin.Context) {
	err := h.srv.HandleTokenRequest(c.Writer, c.Request)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	}
}

func (h *OAuthHandler) showError(c *gin.Context, err error) {
	var msg string

	if err != nil {
		msg = err.Error()
	}

	c.HTML(http.StatusBadRequest, "oauth_error.html", gin.H{"message": msg})
}
