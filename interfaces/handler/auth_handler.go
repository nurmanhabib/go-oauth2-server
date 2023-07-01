package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/nurmanhabib/go-oauth2-server/interfaces/provider"
	"gorm.io/gorm"
)

type AuthHandler struct {
	provider *provider.Provider
}

func NewAuthHandler(provider *provider.Provider) *AuthHandler {
	return &AuthHandler{provider: provider}
}

func (h *AuthHandler) ShowLoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "oauth_login.html", gin.H{})
}

func (h *AuthHandler) Authenticate(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	user, err := h.provider.Repo.UserRepo.FindByEmail(c, email)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.HTML(http.StatusOK, "oauth_login.html", gin.H{"error": "Email or password do not match."})
			return

		default:
			c.HTML(http.StatusInternalServerError, "oauth_login.html", gin.H{"error": fmt.Sprintf("error query: %v", err)})
			return
		}
	}

	if user.Password != password {
		c.HTML(http.StatusOK, "oauth_login.html", gin.H{"error": "Email or password do not match."})
		return
	}

	session := sessions.Default(c)
	session.Set("LoggedInUserID", user.ID)
	session.Save()

	if intendedURL := session.Get("intended_url"); intendedURL != nil {
		session.Delete("intended_url")
		session.Save()

		c.Redirect(http.StatusFound, intendedURL.(string))
		return
	}

	c.Redirect(http.StatusFound, "/home")
}
