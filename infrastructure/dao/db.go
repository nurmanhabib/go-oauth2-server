package dao

import (
	"github.com/nurmanhabib/go-oauth2-server/domain/repository"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/persistence"
	"gorm.io/gorm"
)

// Repositories is a container for collecting repositories.
type Repositories struct {
	DB *gorm.DB

	OauthAccessGrant repository.OauthAccessGrantRepository
	OauthAccessToken repository.OauthAccessTokenRepository
	OauthClient      repository.OauthClientRepository
	User             repository.UserRepository
}

// New is a constructor for building repository collections.
func New(db *gorm.DB) *Repositories {
	return &Repositories{
		DB:               db,
		OauthAccessGrant: persistence.NewOauthAccessGrantRepository(db),
		OauthAccessToken: persistence.NewOauthAccessTokenRepository(db),
		OauthClient:      persistence.NewOauthClientRepository(db),
		User:             persistence.NewUserRepository(db),
	}
}
