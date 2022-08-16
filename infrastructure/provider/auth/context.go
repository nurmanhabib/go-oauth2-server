package auth

import (
	"context"

	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
)

type ctxKey string

const (
	userCtx = ctxKey("AuthUserCtx")
)

// FromContext is a function to retrieve the current user via context.
func FromContext(ctx context.Context) *entity.User {
	u := ctx.Value(userCtx)

	if user, ok := u.(*entity.User); ok {
		return user
	}

	return nil
}

// WithContext is a function to store the current user in a context.
func WithContext(ctx context.Context, user *entity.User) context.Context {
	return context.WithValue(ctx, userCtx, user)
}
