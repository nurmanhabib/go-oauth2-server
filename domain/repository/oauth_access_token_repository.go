package repository

import (
	"context"
	"time"

	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
)

// OauthAccessTokenRepository is a repository interface for access token entities.
type OauthAccessTokenRepository interface {
	Find(ctx context.Context, grant *entity.OauthAccessToken) (*entity.OauthAccessToken, error)
	Save(ctx context.Context, grant *entity.OauthAccessToken) error
	Revoke(ctx context.Context, token *entity.OauthAccessToken, time time.Time) error
}
