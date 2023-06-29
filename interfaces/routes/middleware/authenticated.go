package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/auth"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/dao"
)

func Authenticated(repo *dao.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID, _ := session.Get("LoggedInUserID").(string)

		if userID == "" {
			c.Redirect(http.StatusFound, "/login")
			return
		}

		user, err := repo.UserRepo.FindByID(c, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		c.Set("AuthUser", auth.WithAuthenticatedUser(c, user))
	}
}
