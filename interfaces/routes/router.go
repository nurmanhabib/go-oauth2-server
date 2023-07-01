package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/nurmanhabib/go-oauth2-server/interfaces/provider"
)

type Router struct {
	provider *provider.Provider
}

func NewRouter(provider *provider.Provider) *Router {
	return &Router{provider: provider}
}

func (r *Router) Init() *gin.Engine {
	e := gin.Default()

	e.LoadHTMLGlob("interfaces/views/*.html")

	store := cookie.NewStore([]byte("secret"))
	e.Use(sessions.Sessions("session_app", store))

	authRoute(e, r.provider)
	oauthRoute(e, r.provider)

	return e
}
