package repository

import (
	"context"

	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
)

type OAuthAccessTokenRepo interface {
	FindByID(ctx context.Context, id string) (*entity.OAuthAccessToken, error)
	FindByToken(ctx context.Context, token string) (*entity.OAuthAccessToken, error)
	FindByRefreshToken(ctx context.Context, token string) (*entity.OAuthAccessToken, error)
	Save(ctx context.Context, token *entity.OAuthAccessToken) error
	Update(ctx context.Context, token *entity.OAuthAccessToken, id string) error
	Delete(ctx context.Context, id string) error
}
