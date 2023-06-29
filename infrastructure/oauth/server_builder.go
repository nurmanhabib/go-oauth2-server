package oauth

import (
	"net/http"
	"strings"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/auth"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/dao"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/oauth/store"
	"github.com/nurmanhabib/go-oauth2-server/pkg/passport"
)

type ServerBuilder struct {
	repo    *dao.Repositories
	manager *manage.Manager
	server  *server.Server
}

func NewServerBuilder(repo *dao.Repositories) passport.Builder {
	return &ServerBuilder{
		repo: repo,
	}
}

func (b *ServerBuilder) Reset() {
	b.manager = manage.NewDefaultManager()
	b.server = server.NewDefaultServer(b.manager)
}

func (b *ServerBuilder) SetManager() {
	b.manager.MapClientStorage(store.NewClientStore(b.repo))
	b.manager.MapTokenStorage(store.NewTokenStore(b.repo))
	b.manager.SetValidateURIHandler(func(baseURI, redirectURI string) error {
		allowedURIs := strings.Fields(baseURI)

		if !sliceContains(allowedURIs, redirectURI) {
			return errors.ErrInvalidRedirectURI
		}

		return nil
	})
}

func (b *ServerBuilder) SetServer() {
	b.server.SetAllowGetAccessRequest(false)
	b.server.SetAllowedGrantType(oauth2.AuthorizationCode, oauth2.Refreshing)
	b.server.SetAllowedResponseType(oauth2.Code)
	b.server.SetClientInfoHandler(server.ClientFormHandler)
	b.server.SetUserAuthorizationHandler(func(w http.ResponseWriter, r *http.Request) (userID string, err error) {
		user := auth.GetUser(r.Context())
		if user != nil {
			return user.ID, nil
		}

		return "", errors.ErrAccessDenied
	})
}

func (b *ServerBuilder) GetServer() *server.Server {
	return b.server
}

func sliceContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
