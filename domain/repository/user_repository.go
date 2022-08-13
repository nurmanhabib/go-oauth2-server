package repository

import (
	"context"

	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
	"github.com/nurmanhabib/go-oauth2-server/pkg/auth"
)

// UserRepository is a repository interface for user entities.
type UserRepository interface {
	Find(ctx context.Context, user *entity.User) (*entity.User, error)
	FindByCredentials(ctx context.Context, credentials auth.Credentials) (*entity.User, error)
	Save(ctx context.Context, user *entity.User) error
}
