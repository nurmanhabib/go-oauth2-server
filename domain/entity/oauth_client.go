package entity

import "time"

// OauthClient is an OAuth Client entity container.
type OauthClient struct {
	ID          string
	Name        string
	Secret      string
	SuperApp    bool
	RedirectURI string
	Scopes      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
