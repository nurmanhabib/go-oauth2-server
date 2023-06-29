package handler

import (
	"html/template"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) ShowLoginForm(c *gin.Context) {
	tmpl, err := template.ParseFiles(path.Join("interfaces", "views", "oauth_login.html"))
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(c.Writer, nil)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (h *AuthHandler) Authenticate(c *gin.Context) {

}
