package entity

import (
	"database/sql"
	"time"
)

// OauthAccessToken is an OAuth Access Token entity container.
type OauthAccessToken struct {
	ID               string
	UserID           string
	OauthClientID    string
	Token            string
	ExpiresIn        int
	RefreshToken     string
	RefreshExpiresIn int
	Scopes           string
	RevokedAt        sql.NullTime
	CreatedAt        time.Time
	UpdatedAt        time.Time
}