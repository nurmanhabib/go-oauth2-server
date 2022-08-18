package oauth

import (
	"github.com/go-oauth2/oauth2/v4"
	"github.com/nurmanhabib/go-oauth2-server/domain/repository"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/provider/oauth/handler"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/provider/oauth/store"
	"github.com/nurmanhabib/go-oauth2-server/pkg/oauth/config"
)

// Dependency is a container for OAuth configuration needs.
type Dependency struct {
	Repo
}

// Repo is a container for the needs of the OAuth configuration repository.
type Repo struct {
	AccessGrantRepo repository.OauthAccessGrantRepository
	AccessTokenRepo repository.OauthAccessTokenRepository
	ClientRepo      repository.OauthClientRepository
	UserRepo        repository.UserRepository
}

// NewConfig is an OAuth configuration constructor.
func NewConfig(dependency *Dependency) *config.Config {
	return config.New(
		handler.WithAuthenticatedUserHandler(),
		handler.WithClientScopeHandler(dependency.ClientRepo),

		// Storage
		store.WithClientStorage(dependency.ClientRepo),
		store.WithTokenStorage(
			dependency.AccessGrantRepo,
			dependency.AccessTokenRepo,
			dependency.ClientRepo,
			dependency.UserRepo,
		),

		// Allow Grant Types
		config.AllowGrantTypes(
			oauth2.AuthorizationCode,
			oauth2.Refreshing,
			// oauth2.ClientCredentials,
			// oauth2.Implicit,

			// Deprecated https://datatracker.ietf.org/doc/html/draft-ietf-oauth-security-topics-16#section-2.4
			// oauth2.PasswordCredentials,
		),

		// Allow Response Types
		config.AllowResponseTypes(
			oauth2.Code,
			// oauth2.Token,
		),

		// config.WithJWTAccessTokenGenerator("", []byte("00000000"), jwt.SigningMethodHS512),
	)
}
