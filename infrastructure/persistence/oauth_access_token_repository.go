package persistence

import (
	"context"
	"time"

	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
	"gorm.io/gorm"
)

// OauthAccessTokenRepository is a container for implementing repositories in a persistent way to the database.
type OauthAccessTokenRepository struct {
	db *gorm.DB
}

// NewOauthAccessTokenRepository is a concrete repository constructor.
func NewOauthAccessTokenRepository(db *gorm.DB) *OauthAccessTokenRepository {
	return &OauthAccessTokenRepository{db: db}
}

// Find is a function to get access grant from within the database.
func (a *OauthAccessTokenRepository) Find(ctx context.Context, grant *entity.OauthAccessToken) (*entity.OauthAccessToken, error) {
	var result entity.OauthAccessToken

	q := a.db.WithContext(ctx).Take(&result, grant)
	if q.Error != nil {
		return nil, q.Error
	}

	if result.RevokedAt.Valid {
		return &result, entity.ErrAccessTokenRevoked
	}

	return &result, nil
}

// Save is a function to store access grant into the database.
func (a *OauthAccessTokenRepository) Save(ctx context.Context, grant *entity.OauthAccessToken) error {
	q := a.db.WithContext(ctx).Create(&grant)
	if q.Error != nil {
		return q.Error
	}

	return nil
}

// Revoke is a function to revoke access grant.
func (a *OauthAccessTokenRepository) Revoke(ctx context.Context, grant *entity.OauthAccessToken, time time.Time) error {
	q := a.db.WithContext(ctx).Model(&entity.OauthAccessToken{}).Where(grant).Update("revoked_at", time)
	if q.Error != nil {
		return q.Error
	}

	return nil
}
