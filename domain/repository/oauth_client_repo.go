package repository

import (
	"context"

	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
)

type OAuthClientRepo interface {
	FindByID(ctx context.Context, id string) (*entity.OAuthClient, error)
	Save(ctx context.Context, client *entity.OAuthClient) error
	Update(ctx context.Context, client *entity.OAuthClient, id string) error
	Delete(ctx context.Context, id string) error
}
