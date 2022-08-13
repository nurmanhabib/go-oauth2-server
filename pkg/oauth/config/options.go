package config

import (
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
)

// Option is option pattern.
type Option func(*Config)

// ManagerOption is a type for manager customization.
type ManagerOption func(*manage.Manager)

// ServerOption is a type for server customization.
type ServerOption func(*server.Server)
