package routes

import (
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/dao"
	"github.com/nurmanhabib/go-oauth2-server/pkg/auth"
)

// RouterOption is option pattern to set Router.
type RouterOption func(*Router)

// WithOAuthServer is an option to set OAuthServer.
func WithOAuthServer(srv *server.Server) RouterOption {
	return func(router *Router) {
		router.oauthSrv = srv
	}
}

// WithRepositories is an option to set repositories.
func WithRepositories(repos *dao.Repositories) RouterOption {
	return func(router *Router) {
		router.repos = repos
	}
}

// WithAuthUserProvider is an option to set auth user provider.
func WithAuthUserProvider(provider auth.UserProvider) RouterOption {
	return func(router *Router) {
		router.authUserProvider = provider
	}
}
