package middleware

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/auth"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/dao"
	"gorm.io/gorm"
)

func Authenticated(repo *dao.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("current: %s\n", c.Request.URL.String())
		session := sessions.Default(c)
		userID, _ := session.Get("LoggedInUserID").(string)

		if userID == "" {
			session.Set("intended_url", c.Request.URL.String())
			session.Save()

			c.Redirect(http.StatusFound, "/login")
			return
		}

		user, err := repo.UserRepo.FindByID(c, userID)
		if err != nil {
			switch err {
			case gorm.ErrRecordNotFound:
				session.Set("intended_url", c.Request.URL.String())
				session.Delete("LoggedInUserID")
				session.Save()

				c.Redirect(http.StatusFound, "/login")
				return

			default:
				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		c.Request = c.Request.Clone(auth.WithAuthenticatedUser(c.Request.Context(), user))
		c.Set("AuthUser", auth.WithAuthenticatedUser(c, user))

		c.Next()
	}
}
