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
	var user entity.User

	err := u.db.WithContext(ctx).Take(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User

	err := u.db.WithContext(ctx).Take(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) Save(ctx context.Context, user *entity.User) error {
	return u.db.WithContext(ctx).Create(&user).Error
}

func (u *UserRepo) Update(ctx context.Context, user *entity.User, id string) error {
	q := u.db.WithContext(ctx).Where("id = ?", id).Updates(&user)
	if q.Error != nil {
		return q.Error
	}

	if q.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (u *UserRepo) Delete(ctx context.Context, id string) error {
	q := u.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.User{})
	if q.Error != nil {
		return q.Error
	}

	if q.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
