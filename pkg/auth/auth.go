package auth

import "context"

// Attempt is a function to attempt user by credentials.
func Attempt(ctx context.Context, credentials Credentials, provider UserProvider) (User, error) {
	u, err := provider.FindByCredentials(ctx, credentials)
	if err != nil {
		return nil, err
	}

	err = provider.ValidateCredentials(ctx, u, credentials)
	if err != nil {
		return nil, err
	}

	return u, nil
}
