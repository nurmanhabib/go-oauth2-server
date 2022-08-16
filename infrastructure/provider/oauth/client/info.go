package client

import "github.com/nurmanhabib/go-oauth2-server/domain/entity"

// Info is a transformer for the ClientInfo implementation.
type Info struct {
	*entity.OauthClient
}

// GetID is a function to get client_id.
func (i *Info) GetID() string {
	if i == nil || i.OauthClient == nil {
		return ""
	}

	return i.ID
}

// GetSecret is a function to get client_secret.
func (i *Info) GetSecret() string {
	if i == nil || i.OauthClient == nil {
		return ""
	}

	return i.Secret
}

// GetDomain is a function to get redirect_uri.
func (i *Info) GetDomain() string {
	if i == nil || i.OauthClient == nil {
		return ""
	}

	return i.RedirectURI
}

// GetUserID is a function to get the client owner (in this case, we don't support it).
func (i *Info) GetUserID() string {
	return ""
}
