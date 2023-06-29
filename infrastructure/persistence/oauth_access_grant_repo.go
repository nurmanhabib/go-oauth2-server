package persistence

import (
	"context"

	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
	"github.com/nurmanhabib/go-oauth2-server/domain/repository"
	"gorm.io/gorm"
)

type OAuthAccessGrantRepo struct {
	db *gorm.DB
}

func NewOAuthAccessGrantRepo(db *gorm.DB) repository.OAuthAccessGrantRepo {
	return &OAuthAccessGrantRepo{db: db}
}

func (O *OAuthAccessGrantRepo) FindByID(ctx context.Context, id string) (*entity.OAuthAccessGrant, error) {
	// TODO implement me
	panic("implement me")
}

func (O *OAuthAccessGrantRepo) FindByCode(ctx context.Context, code string) (*entity.OAuthAccessGrant, error) {
	// TODO implement me
	panic("implement me")
}

func (O *OAuthAccessGrantRepo) Save(ctx context.Context, grant *entity.OAuthAccessGrant) error {
	// TODO implement me
	panic("implement me")
}

func (O *OAuthAccessGrantRepo) Update(ctx context.Context, grant *entity.OAuthAccessGrant, id string) error {
	// TODO implement me
	panic("implement me")
}

func (O *OAuthAccessGrantRepo) Delete(ctx context.Context, id string) error {
	// TODO implement me
	panic("implement me")
}
