package store

import (
	"context"
	"fmt"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/dao"
	"gorm.io/gorm"
)

type ClientStore struct {
	repo *dao.Repositories
}

func NewClientStore(repo *dao.Repositories) oauth2.ClientStore {
	return &ClientStore{repo: repo}
}

func (c *ClientStore) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {
	client, err := c.repo.OAuthClientRepo.FindByID(ctx, id)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, errors.ErrInvalidClient

		default:
			return nil, fmt.Errorf("failed query find client [%s]: %w", id, err)
		}
	}

	return client, nil
}
