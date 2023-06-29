package passport

import (
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/server"
)

type Config struct {
	AllowGrantTypes          []oauth2.GrantType
	ClientStore              oauth2.ClientStore
	TokenStore               oauth2.TokenStore
	AccessGenerator          oauth2.AccessGenerate
	AuthorizeGenerate        oauth2.AuthorizeGenerate
	ClientInfoHandler        server.ClientInfoHandler
	UserAuthorizationHandler server.UserAuthorizationHandler
}

type Option func(config *Config)

func NewConfig(options ...Option) *Config {
	c := &Config{
		AllowGrantTypes: []oauth2.GrantType{oauth2.AuthorizationCode, oauth2.Refreshing},
	}

	for _, apply := range options {
		apply(c)
	}

	return c
}

func WithGrantTypes(grants ...oauth2.GrantType) Option {
	return func(config *Config) {
		config.AllowGrantTypes = grants
	}
}
