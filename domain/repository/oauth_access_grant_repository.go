package repository

import (
	"context"
	"time"

	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
)

// OauthAccessGrantRepository is a repository interface for access grant entities.
type OauthAccessGrantRepository interface {
	Find(ctx context.Context, grant *entity.OauthAccessGrant) (*entity.OauthAccessGrant, error)
	Save(ctx context.Context, grant *entity.OauthAccessGrant) error
	Revoke(ctx context.Context, grant *entity.OauthAccessGrant, time time.Time) error
}
