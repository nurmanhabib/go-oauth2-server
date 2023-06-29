package persistence

import (
	"context"

	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
	"github.com/nurmanhabib/go-oauth2-server/domain/repository"
	"gorm.io/gorm"
)

type OAuthAccessTokenRepo struct {
	db *gorm.DB
}

func NewOAuthAccessTokenRepo(db *gorm.DB) repository.OAuthAccessTokenRepo {
	return &OAuthAccessTokenRepo{db: db}
}

func (O *OAuthAccessTokenRepo) FindByID(ctx context.Context, id string) (*entity.OAuthAccessToken, error) {
	// TODO implement me
	panic("implement me")
}

func (O *OAuthAccessTokenRepo) FindByToken(ctx context.Context, token string) (*entity.OAuthAccessToken, error) {
	// TODO implement me
	panic("implement me")
}

func (O *OAuthAccessTokenRepo) FindByRefreshToken(ctx context.Context, token string) (*entity.OAuthAccessToken, error) {
	// TODO implement me
	panic("implement me")
}

func (O *OAuthAccessTokenRepo) Save(ctx context.Context, token *entity.OAuthAccessToken) error {
	// TODO implement me
	panic("implement me")
}

func (O *OAuthAccessTokenRepo) Update(ctx context.Context, token *entity.OAuthAccessToken, id string) error {
	// TODO implement me
	panic("implement me")
}

func (O *OAuthAccessTokenRepo) Delete(ctx context.Context, id string) error {
	// TODO implement me
	panic("implement me")
}
