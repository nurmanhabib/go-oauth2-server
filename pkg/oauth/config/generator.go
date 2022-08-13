package config

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/generates"
)

// Generator is a container for configuration to generate a token.
type Generator struct {
	AccessTokenGenerator       oauth2.AccessGenerate
	AuthorizationCodeGenerator oauth2.AuthorizeGenerate
}

// WithJWTAccessTokenGenerator create access token with JWT format.
func WithJWTAccessTokenGenerator(kid string, key []byte, method jwt.SigningMethod) Option {
	return WithAccessTokenGenerator(generates.NewJWTAccessGenerate(kid, key, method))
}

// WithAccessTokenGenerator is a function for the custom generator for the access token.
func WithAccessTokenGenerator(gen oauth2.AccessGenerate) Option {
	return func(config *Config) {
		config.Generator.AccessTokenGenerator = gen
	}
}

// WithAuthorizationCodeGenerator is a function for the custom generator for the access grant.
func WithAuthorizationCodeGenerator(gen oauth2.AuthorizeGenerate) Option {
	return func(config *Config) {
		config.Generator.AuthorizationCodeGenerator = gen
	}
}
