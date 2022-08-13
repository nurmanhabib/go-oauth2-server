package entity

import (
	"database/sql"
	"time"
)

// OauthAccessGrant is an OAuth Access Grant entity container.
type OauthAccessGrant struct {
	ID            string
	UserID        string
	OauthClientID string
	Code          string
	RedirectURI   string
	Scopes        string
	ExpiresIn     int
	RevokedAt     sql.NullTime
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
