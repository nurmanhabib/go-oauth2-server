package store

import (
	"context"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
	"github.com/nurmanhabib/go-oauth2-server/domain/repository"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/provider/oauth/client"
	"github.com/nurmanhabib/go-oauth2-server/pkg/oauth/config"
)

// ClientStorage is a container to accommodate client storage needs.
type ClientStorage struct {
	repo repository.OauthClientRepository
}

// WithClientStorage is a function to set to OAuth config.
func WithClientStorage(repo repository.OauthClientRepository) config.Option {
	return func(c *config.Config) {
		c.ClientStorage = NewClientStorage(repo)
	}
}

// NewClientStorage is constructor.
func NewClientStorage(repo repository.OauthClientRepository) *ClientStorage {
	return &ClientStorage{repo: repo}
}

// GetByID is a function to get client info by id.
func (c *ClientStorage) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {
	clientEntity, err := c.repo.Find(ctx, &entity.OauthClient{ID: id})
	if err != nil {
		return nil, err
	}

	return &client.Info{OauthClient: clientEntity}, nil
}
