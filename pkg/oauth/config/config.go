package config

import (
	"github.com/go-oauth2/oauth2/v4"
)

// Config is a set of configurations for building an OAuth 2.0 server.
type Config struct {
	Repo
	Handler
	Generator

	ClientStorage oauth2.ClientStore
	TokenStorage  oauth2.TokenStore

	AllowGrantTypes []oauth2.GrantType

	ManagerAdjustments []ManagerOption
	ServerAdjustments  []ServerOption
}

// New is a function of building a new configuration with some defaults.
func New(options ...Option) *Config {
	conf := &Config{}

	defaultMemoryTokenStorage()(conf)
	defaultClientInfoHandler()(conf)

	for _, apply := range options {
		apply(conf)
	}

	return conf
}

// ManagerAdjustment is a function for the adjustment of the base manager.
func (c *Config) ManagerAdjustment(option ManagerOption) {
	if c.ManagerAdjustments == nil {
		c.ManagerAdjustments = []ManagerOption{}
	}

	c.ManagerAdjustments = append(c.ManagerAdjustments, option)
}

// ServerAdjustment is a function for the adjustment of the base server.
func (c *Config) ServerAdjustment(option ServerOption) {
	if c.ServerAdjustments == nil {
		c.ServerAdjustments = []ServerOption{}
	}

	c.ServerAdjustments = append(c.ServerAdjustments, option)
}
