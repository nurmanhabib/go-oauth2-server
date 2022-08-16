package persistence

import (
	"context"
	"time"

	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
	"gorm.io/gorm"
)

// OauthAccessGrantRepository is a container for implementing repositories in a persistent way to the database.
type OauthAccessGrantRepository struct {
	db *gorm.DB
}

// NewOauthAccessGrantRepository is a concrete repository constructor.
func NewOauthAccessGrantRepository(db *gorm.DB) *OauthAccessGrantRepository {
	return &OauthAccessGrantRepository{db: db}
}

// Find is a function to get access grant from within the database.
func (a *OauthAccessGrantRepository) Find(ctx context.Context, grant *entity.OauthAccessGrant) (*entity.OauthAccessGrant, error) {
	var result entity.OauthAccessGrant

	q := a.db.WithContext(ctx).Take(&result, grant)
	if q.Error != nil {
		return nil, q.Error
	}

	if result.RevokedAt.Valid {
		return &result, entity.ErrAccessGrantRevoked
	}

	return &result, nil
}

// Save is a function to store access grant into the database.
func (a *OauthAccessGrantRepository) Save(ctx context.Context, grant *entity.OauthAccessGrant) error {
	q := a.db.WithContext(ctx).Create(&grant)
	if q.Error != nil {
		return q.Error
	}

	return nil
}

// Revoke is a function to revoke access grant.
func (a *OauthAccessGrantRepository) Revoke(ctx context.Context, grant *entity.OauthAccessGrant, time time.Time) error {
	q := a.db.WithContext(ctx).Model(&entity.OauthAccessGrant{}).Where(grant).Update("revoked_at", time)
	if q.Error != nil {
		return q.Error
	}

	return nil
}
