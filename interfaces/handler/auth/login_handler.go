package auth

import (
	"html/template"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/go-session/session"
	"github.com/nurmanhabib/go-oauth2-server/pkg/auth"
)

// LoginHandler is login handler.
type LoginHandler struct {
	UserProvider auth.UserProvider
}

// Index is login form handler.
func (*LoginHandler) Index(c *gin.Context) {
	tmpl, err := template.ParseFiles(path.Join("resources", "views", "oauth_login.html"))

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(c.Writer, gin.H{})

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

// Submit is a function to submit login request.
func (h *LoginHandler) Submit(c *gin.Context) {
	r := c.Request
	w := c.Writer
	ctx := r.Context()

	credentials := auth.Credentials{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	}

	user, err := auth.Attempt(ctx, credentials, h.UserProvider)
	if err != nil {
		c.Redirect(http.StatusFound, "/login?error=invalid_credentials")
		return
	}

	sessionStore, err := session.Start(r.Context(), w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sessionStore.Set("login_user_id", user.GetAuthIdentifier())
	errStore := sessionStore.Save()
	if errStore != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, errStore)
		return
	}

	redirectURI := "/oauth/authorize"
	intendedURI, ok := sessionStore.Get("intended_uri")

	if ok {
		redirectURI = intendedURI.(string)
		sessionStore.Delete("intended_uri")
		_ = sessionStore.Save()
	}

	c.Redirect(http.StatusFound, redirectURI)
}
