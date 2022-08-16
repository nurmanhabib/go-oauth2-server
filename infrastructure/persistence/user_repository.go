package persistence

import (
	"context"

	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
	"github.com/nurmanhabib/go-oauth2-server/pkg/auth"
	"gorm.io/gorm"
)

// UserRepository is a container for implementing repositories in a persistent way to the database.
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository is a concrete repository constructor.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Find is a function to get user from within the database.
func (u *UserRepository) Find(ctx context.Context, user *entity.User) (*entity.User, error) {
	var result entity.User

	q := u.db.WithContext(ctx).Take(&result, user)
	if q.Error != nil {
		return nil, q.Error
	}

	return &result, nil
}

// FindByCredentials is a function to get user by auth credentials from within the database.
func (u *UserRepository) FindByCredentials(ctx context.Context, credentials auth.Credentials) (*entity.User, error) {
	var result entity.User

	mCredentials := map[string]interface{}{}
	for k, v := range credentials {
		mCredentials[k] = v
	}

	q := u.db.WithContext(ctx).Take(&result, mCredentials)
	if q.Error != nil {
		return nil, q.Error
	}

	return &result, nil
}

// Save is a function to store user into the database.
func (u *UserRepository) Save(ctx context.Context, user *entity.User) error {
	q := u.db.WithContext(ctx).Create(&user)
	if q.Error != nil {
		return q.Error
	}

	return nil
}
