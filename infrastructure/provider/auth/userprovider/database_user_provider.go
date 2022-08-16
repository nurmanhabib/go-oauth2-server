package userprovider

import (
	"context"
	"fmt"

	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
	"github.com/nurmanhabib/go-oauth2-server/domain/repository"
	"github.com/nurmanhabib/go-oauth2-server/pkg/auth"
	"github.com/nurmanhabib/go-oauth2-server/pkg/hashing"
	"gorm.io/gorm"
)

// DatabaseUserProvider is a container for the needs of the user provider in the database.
type DatabaseUserProvider struct {
	repo   repository.UserRepository
	hasher hashing.Hasher
}

// NewDatabaseUserProvider is constructor.
func NewDatabaseUserProvider(repo repository.UserRepository, hasher hashing.Hasher) auth.UserProvider {
	return &DatabaseUserProvider{
		repo:   repo,
		hasher: hasher,
	}
}

// FindByIdentifier is a function to search for users by identifier.
// The identifier here is not the username, but the user id as the primary key.
func (p *DatabaseUserProvider) FindByIdentifier(ctx context.Context, identifier string) (auth.User, error) {
	userEntity, err := p.repo.Find(ctx, &entity.User{ID: identifier})
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, auth.ErrUserNotFound
		default:
			return nil, fmt.Errorf("error find by credentials: %w", err)
		}
	}

	return userEntity, nil
}

// FindByCredentials is a function to search for users based on credentials.
// Usually the user is required to enter a pair of credentials in the form of a username and password,
// the application needs to check whether the username is available or not in the application.
//
// To retrieve credentials without password credentials.WithoutPassword().
func (p *DatabaseUserProvider) FindByCredentials(ctx context.Context, credentials auth.Credentials) (auth.User, error) {
	userEntity, err := p.repo.FindByCredentials(ctx, credentials.WithoutPassword())
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, auth.ErrUserNotFound
		default:
			return nil, fmt.Errorf("error find by credentials: %w", err)
		}
	}

	return userEntity, nil
}

// ValidateCredentials is a function for authentication by comparing the password credentials.
func (p *DatabaseUserProvider) ValidateCredentials(ctx context.Context, user auth.User, credentials auth.Credentials) error {
	ok := p.hasher.Check(ctx, credentials.GetPassword(), user.GetAuthPassword())
	if !ok {
		return auth.ErrInvalidCredentials
	}

	return nil
}
