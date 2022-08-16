package oauth

import (
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/nurmanhabib/go-oauth2-server/domain/repository"
)

// Handler is container oauth handler needs.
type Handler struct {
	srv        *server.Server
	clientRepo repository.OauthClientRepository
}

// NewHandler is constructor to build new OAuth handler.
func NewHandler(srv *server.Server, clientRepo repository.OauthClientRepository) *Handler {
	return &Handler{
		srv:        srv,
		clientRepo: clientRepo,
	}
}
