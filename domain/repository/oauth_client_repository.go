package repository

import (
	"context"

	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
)

// OauthClientRepository is a repository interface for application entities.
type OauthClientRepository interface {
	Find(ctx context.Context, client *entity.OauthClient) (*entity.OauthClient, error)
	Save(ctx context.Context, client *entity.OauthClient) error
}
