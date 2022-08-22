package oauth

import (
	"strings"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/nurmanhabib/go-oauth2-server/pkg/oauth/config"
	"github.com/nurmanhabib/go-oauth2-server/util"
)

// NewManager is a wrapper function building a new OAuth 2.0 manager with our configuration.
func NewManager(config *config.Config) *manage.Manager {
	m := manage.NewDefaultManager()

	if config.ClientStorage != nil {
		m.MapClientStorage(config.ClientStorage)
	}

	if config.TokenStorage != nil {
		m.MapTokenStorage(config.TokenStorage)
	}

	if config.Generator.AuthorizationCodeGenerator != nil {
		m.MapAuthorizeGenerate(config.Generator.AuthorizationCodeGenerator)
	}

	if config.Generator.AccessTokenGenerator != nil {
		m.MapAccessGenerate(config.Generator.AccessTokenGenerator)
	}

	if cfg := config.TokenConfig.AuthorizeCodeGrant; cfg != nil {
		m.SetAuthorizeCodeTokenCfg(cfg)
	}

	if cfg := config.TokenConfig.PasswordGrant; cfg != nil {
		m.SetPasswordTokenCfg(cfg)
	}

	if cfg := config.TokenConfig.ClientGrant; cfg != nil {
		m.SetClientTokenCfg(cfg)
	}

	m.SetValidateURIHandler(validateRedirectURI)

	if config.ManagerAdjustments != nil {
		for _, apply := range config.ManagerAdjustments {
			apply(m)
		}
	}

	return m
}

// validateRedirectURI allows multiple RedirectURIs separated by spaces or newlines.
func validateRedirectURI(baseURI, redirectURI string) error {
	allowedURIs := strings.Fields(baseURI)

	if !util.SliceContains(allowedURIs, redirectURI) {
		return errors.ErrInvalidRedirectURI
	}

	return nil
}
