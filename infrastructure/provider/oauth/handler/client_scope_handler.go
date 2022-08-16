package handler

import (
	"context"
	"strings"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
	"github.com/nurmanhabib/go-oauth2-server/domain/repository"
	"github.com/nurmanhabib/go-oauth2-server/pkg/oauth/config"
	"github.com/nurmanhabib/go-oauth2-server/util"
)

// ClientScopeHandler is a handler for scope validation.
type ClientScopeHandler struct {
	clientRepo repository.OauthClientRepository
}

// WithClientScopeHandler is a function for the validation set scope.
func WithClientScopeHandler(clientRepo repository.OauthClientRepository) config.Option {
	return func(c *config.Config) {
		h := &ClientScopeHandler{clientRepo: clientRepo}
		c.ClientScopeHandler = h.Handler
	}
}

// Handler is a function for scope validation.
func (c *ClientScopeHandler) Handler(tgr *oauth2.TokenGenerateRequest) (bool, error) {
	oauthClient, err := c.clientRepo.Find(context.Background(), &entity.OauthClient{ID: tgr.ClientID})
	if err != nil {
		return false, err
	}

	allowedScopes := strings.Fields(oauthClient.Scopes)
	requestScopes := strings.Fields(tgr.Scope)

	declinedScopes := util.SliceSubtract(requestScopes, allowedScopes)

	if len(declinedScopes) > 0 {
		return false, nil
	}

	return true, nil
}
