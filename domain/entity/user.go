package entity

import (
	"database/sql"
	"time"
)

// User is a user entity container.
type User struct {
	ID              string
	Name            string
	Email           string
	EmailVerifiedAt sql.NullTime
	Password        string
	RememberToken   sql.NullString
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// GetAuthIdentifier is a function to take as auth identifier.
func (u *User) GetAuthIdentifier() string {
	if u == nil {
		return ""
	}

	return u.ID
}

// GetAuthPassword is a function to take as auth password.
func (u *User) GetAuthPassword() string {
	if u == nil {
		return ""
	}

	return u.Password
}
