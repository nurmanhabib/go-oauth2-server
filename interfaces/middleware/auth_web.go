package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-session/session"
	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
	"github.com/nurmanhabib/go-oauth2-server/domain/repository"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/provider/auth"
)

// AuthWeb is a container for the needs of auth web middleware.
type AuthWeb struct {
	userRepo repository.UserRepository
}

// NewAuthWeb returns AuthWeb middleware.
func NewAuthWeb(userRepo repository.UserRepository) *AuthWeb {
	return &AuthWeb{userRepo: userRepo}
}

// Authenticate is a handler to check current user in session.
func (aw AuthWeb) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		w := c.Writer
		ctx := r.Context()

		sessionStore, err := session.Start(ctx, w, r)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		userID, ok := sessionStore.Get("login_user_id")
		if !ok {
			sessionStore.Set("intended_uri", r.URL.String())
			_ = sessionStore.Save()

			c.Redirect(http.StatusFound, "/login")
			return
		}

		user, errUser := aw.userRepo.Find(ctx, &entity.User{ID: userID.(string)})
		if errUser != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, errUser)
			return
		}

		c.Request = r.WithContext(auth.WithContext(ctx, user))

		c.Next()
	}
}
