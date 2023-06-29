package auth

import (
	"context"

	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
)

type ctxKey string

const (
	authenticatedKey = ctxKey("authenticated_user")
)

func WithAuthenticatedUser(ctx context.Context, user *entity.User) context.Context {
	return context.WithValue(ctx, authenticatedKey, user)
}

func GetUser(ctx context.Context) *entity.User {
	if user, ok := ctx.Value(authenticatedKey).(*entity.User); ok {
		return user
	}

	return nil
}
