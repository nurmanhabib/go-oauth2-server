package entity

import (
	"time"

	"gorm.io/gorm"
)

type OAuthClient struct {
	ID          string         `json:"id" gorm:"size:36;primaryKey"`
	Secret      string         `json:"secret"`
	Name        string         `json:"name"`
	SuperApp    bool           `json:"super_app"`
	RedirectURI string         `json:"redirect_uri" gorm:"type:text"`
	Scopes      string         `json:"scopes"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (o *OAuthClient) TableName() string {
	return "oauth_clients"
}

func (o *OAuthClient) GetID() string {
	return o.ID
}

func (o *OAuthClient) GetSecret() string {
	return o.Secret
}

func (o *OAuthClient) GetDomain() string {
	return o.RedirectURI
}

func (o *OAuthClient) IsPublic() bool {
	return !o.SuperApp
}

func (o *OAuthClient) GetUserID() string {
	return ""
}
