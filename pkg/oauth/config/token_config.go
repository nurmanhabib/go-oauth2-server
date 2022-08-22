package config

import (
	"time"

	"github.com/go-oauth2/oauth2/v4/manage"
)

type TokenConfig struct {
	AuthorizeCodeGrant *manage.Config
	PasswordGrant      *manage.Config
	ClientGrant        *manage.Config
}

func defaultTokenConfig() Option {
	return func(config *Config) {
		config.TokenConfig.AuthorizeCodeGrant = manage.DefaultAuthorizeCodeTokenCfg
		config.TokenConfig.PasswordGrant = manage.DefaultPasswordTokenCfg
		config.TokenConfig.ClientGrant = manage.DefaultClientTokenCfg
	}
}

func WithAccessTokenAllGrantExpiry(exp time.Duration) Option {
	return func(config *Config) {
		config.TokenConfig.AuthorizeCodeGrant.AccessTokenExp = exp
		config.TokenConfig.PasswordGrant.AccessTokenExp = exp
		config.TokenConfig.ClientGrant.AccessTokenExp = exp
	}
}

func WithRefreshTokenAllGrantExpiry(exp time.Duration) Option {
	return func(config *Config) {
		config.TokenConfig.AuthorizeCodeGrant.RefreshTokenExp = exp
		config.TokenConfig.PasswordGrant.RefreshTokenExp = exp
		config.TokenConfig.ClientGrant.RefreshTokenExp = exp
	}
}
