package handler

import (
	"net/http"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/provider/auth"
	"github.com/nurmanhabib/go-oauth2-server/pkg/oauth/config"
)

// WithAuthenticatedUserHandler is a function to set how to get the current user.
func WithAuthenticatedUserHandler() config.Option {
	return func(c *config.Config) {
		c.Handler.UserAuthorizationHandler = authenticatedUserHandler
	}
}

func authenticatedUserHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	u := auth.FromContext(r.Context())
	if u == nil {
		return "", errors.New("unauthenticated")
	}

	return u.ID, nil
}
