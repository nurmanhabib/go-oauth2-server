package oauth

import (
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/nurmanhabib/go-oauth2-server/pkg/oauth"
	"github.com/nurmanhabib/go-oauth2-server/pkg/oauth/config"
)

// NewServer build new OAuth server.
func NewServer(config *config.Config) *server.Server {
	return oauth.NewServer(config)
}
