package config

import "github.com/nurmanhabib/go-oauth2-server/domain/repository"

// Repo is a container to hold the repository needed.
type Repo struct {
	OauthClientRepo repository.OauthClientRepository
}
