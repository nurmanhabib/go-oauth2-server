package persistence

import (
	"context"

	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
	"github.com/nurmanhabib/go-oauth2-server/domain/repository"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repository.UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) FindByID(ctx context.Context, id string) (*entity.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u *UserRepo) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u *UserRepo) Save(ctx context.Context, user *entity.User) error {
	// TODO implement me
	panic("implement me")
}

func (u *UserRepo) Update(ctx context.Context, user *entity.User, id string) error {
	// TODO implement me
	panic("implement me")
}

func (u *UserRepo) Delete(ctx context.Context, id string) error {
	// TODO implement me
	panic("implement me")
}
