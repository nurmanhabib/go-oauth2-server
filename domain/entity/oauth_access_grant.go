package entity

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type OAuthAccessGrant struct {
	ID            string         `json:"id" gorm:"size:36;primaryKey"`
	UserID        string         `json:"user_id" gorm:"size:36"`
	OAuthClientID string         `json:"oauth_client_id" gorm:"size:36;column:oauth_client_id"`
	Code          string         `json:"code"`
	RedirectURI   string         `json:"redirect_uri"`
	Scopes        string         `json:"scopes"`
	ExpiresIn     time.Duration  `json:"expires_in"`
	RevokedAt     sql.NullTime   `json:"revoked_at"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (o *OAuthAccessGrant) TableName() string {
	return "oauth_access_grants"
}

func (o *OAuthAccessGrant) IsExpired() bool {
	if o.ExpiresIn == 0 {
		return false
	}

	return o.CreatedAt.Add(o.ExpiresIn).Before(time.Now())
}

func (o *OAuthAccessGrant) IsRevoked() bool {
	return o.RevokedAt.Valid && o.RevokedAt.Time.Before(time.Now())
}
