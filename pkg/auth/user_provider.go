package auth

import "context"

// UserProvider is an interface for finding and validating users.
type UserProvider interface {
	FindByIdentifier(ctx context.Context, identifier string) (User, error)
	FindByCredentials(ctx context.Context, credentials Credentials) (User, error)
	ValidateCredentials(ctx context.Context, user User, credentials Credentials) error
}
