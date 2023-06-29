package repository

import (
	"context"

	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
)

type OAuthAccessGrantRepo interface {
	FindByID(ctx context.Context, id string) (*entity.OAuthAccessGrant, error)
	FindByCode(ctx context.Context, code string) (*entity.OAuthAccessGrant, error)
	Save(ctx context.Context, grant *entity.OAuthAccessGrant) error
	Update(ctx context.Context, grant *entity.OAuthAccessGrant, id string) error
	Delete(ctx context.Context, id string) error
}
