package repository

import (
	"context"

	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
)

type UserRepo interface {
	FindByID(ctx context.Context, id string) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	Save(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User, id string) error
	Delete(ctx context.Context, id string) error
}
