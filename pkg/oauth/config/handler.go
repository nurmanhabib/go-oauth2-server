package config

import "github.com/go-oauth2/oauth2/v4/server"

// Handler is a container for configuration for multiple handlers.
type Handler struct {
	ClientInfoHandler        server.ClientInfoHandler
	UserAuthorizationHandler server.UserAuthorizationHandler
	InternalErrorHandler     server.InternalErrorHandler
	ResponseErrorHandler     server.ResponseErrorHandler
}

func defaultClientInfoHandler() Option {
	return func(config *Config) {
		config.Handler.ClientInfoHandler = server.ClientFormHandler
	}
}
