package config

import (
	"github.com/go-oauth2/oauth2/v4"
)

// AllowResponseTypes sets the type of response type that will be allowed.
func AllowResponseTypes(responseTypes ...oauth2.ResponseType) Option {
	return func(config *Config) {
		config.AllowResponseTypes = responseTypes
	}
}
