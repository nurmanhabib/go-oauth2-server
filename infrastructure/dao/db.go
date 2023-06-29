package dao

import (
	"github.com/nurmanhabib/go-oauth2-server/domain/repository"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/persistence"
	"gorm.io/gorm"
)

type Repositories struct {
	OAuthAccessGrantRepo repository.OAuthAccessGrantRepo
	OAuthAccessTokenRepo repository.OAuthAccessTokenRepo
	OAuthClientRepo      repository.OAuthClientRepo
	UserRepo             repository.UserRepo
}

func NewRepo(db *gorm.DB) *Repositories {
	return &Repositories{
		OAuthAccessGrantRepo: persistence.NewOAuthAccessGrantRepo(db),
		OAuthAccessTokenRepo: persistence.NewOAuthAccessTokenRepo(db),
		OAuthClientRepo:      persistence.NewOAuthClientRepo(db),
		UserRepo:             persistence.NewUserRepo(db),
	}
}
