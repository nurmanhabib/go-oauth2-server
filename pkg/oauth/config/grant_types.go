package config

import (
	"github.com/go-oauth2/oauth2/v4"
)

// AllowGrantTypes sets the type of grant type that will be allowed.
func AllowGrantTypes(grantTypes ...oauth2.GrantType) Option {
	return func(config *Config) {
		config.AllowGrantTypes = grantTypes
	}
}
