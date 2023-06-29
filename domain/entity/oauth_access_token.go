package entity

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type OAuthAccessToken struct {
	ID               string         `json:"id" gorm:"size:36;primaryKey"`
	UserID           string         `json:"user_id" gorm:"size:36"`
	OAuthClientID    string         `json:"oauth_client_id" gorm:"size:36;column:oauth_client_id"`
	Token            string         `json:"code" gorm:"uniqueIndex"`
	Scopes           string         `json:"scopes"`
	ExpiresIn        time.Duration  `json:"expires_in"`
	RefreshToken     string         `json:"refresh_token" gorm:"uniqueIndex"`
	RefreshExpiresIn time.Duration  `json:"refresh_expires_in"`
	RevokedAt        sql.NullTime   `json:"revoked_at"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (o *OAuthAccessToken) TableName() string {
	return "oauth_access_tokens"
}

func (o *OAuthAccessToken) IsExpired() bool {
	if o.ExpiresIn == 0 {
		return false
	}

	return o.CreatedAt.Add(o.ExpiresIn).Before(time.Now())
}

func (o *OAuthAccessToken) IsRefreshExpired() bool {
	if o.RefreshExpiresIn == 0 {
		return false
	}

	return o.CreatedAt.Add(o.RefreshExpiresIn).Before(time.Now())
}

func (o *OAuthAccessToken) IsRevoked() bool {
	return o.RevokedAt.Valid && o.RevokedAt.Time.Before(time.Now())
}
