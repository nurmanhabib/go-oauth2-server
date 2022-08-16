package oauth

import (
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/nurmanhabib/go-oauth2-server/pkg/oauth/config"
)

// NewServer is a wrapper function building a new OAuth 2.0 server with our configuration.
func NewServer(config *config.Config) *server.Server {
	mgr := NewManager(config)
	srv := server.NewDefaultServer(mgr)

	if config.Handler.ClientInfoHandler != nil {
		srv.SetClientInfoHandler(config.Handler.ClientInfoHandler)
	}

	if config.Handler.ClientScopeHandler != nil {
		srv.SetClientScopeHandler(config.Handler.ClientScopeHandler)
	}

	if config.Handler.UserAuthorizationHandler != nil {
		srv.SetUserAuthorizationHandler(config.Handler.UserAuthorizationHandler)
	}

	if config.Handler.InternalErrorHandler != nil {
		srv.SetInternalErrorHandler(config.Handler.InternalErrorHandler)
	}

	if config.Handler.ResponseErrorHandler != nil {
		srv.SetResponseErrorHandler(config.Handler.ResponseErrorHandler)
	}

	if len(config.AllowGrantTypes) > 0 {
		srv.SetAllowedGrantType(config.AllowGrantTypes...)
	}

	if config.ServerAdjustments != nil {
		for _, apply := range config.ServerAdjustments {
			apply(srv)
		}
	}

	return srv
}
