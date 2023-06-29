package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Router struct {
	provider *Provider
}

func NewRouter(provider *Provider) *Router {
	return &Router{provider: provider}
}

func (r *Router) Init() *gin.Engine {
	e := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	e.Use(sessions.Sessions("session_app", store))

	authRoute(e)
	oauthRoute(e, r.provider)

	return e
}
