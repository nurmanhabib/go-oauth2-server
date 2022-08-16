package persistence

import (
	"context"

	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
	"gorm.io/gorm"
)

// OauthClientRepository is a container for implementing repositories in a persistent way to the database.
type OauthClientRepository struct {
	db *gorm.DB
}

// NewOauthClientRepository is a concrete repository constructor.
func NewOauthClientRepository(db *gorm.DB) *OauthClientRepository {
	return &OauthClientRepository{db: db}
}

// Find is a function to get client from within the database.
func (a *OauthClientRepository) Find(ctx context.Context, client *entity.OauthClient) (*entity.OauthClient, error) {
	var result entity.OauthClient

	q := a.db.WithContext(ctx).Take(&result, client)
	if q.Error != nil {
		return nil, q.Error
	}

	return &result, nil
}

// Save is a function to store client into the database.
func (a *OauthClientRepository) Save(ctx context.Context, client *entity.OauthClient) error {
	q := a.db.WithContext(ctx).Create(&client)
	if q.Error != nil {
		return q.Error
	}

	return nil
}
