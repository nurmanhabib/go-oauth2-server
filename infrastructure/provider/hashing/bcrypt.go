package hashing

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

// Bcrypt is a struct for hasher implementation.
type Bcrypt struct{}

// Check is a function to validate a plain text with a hashed value.
func (b *Bcrypt) Check(ctx context.Context, plain string, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	if err != nil {
		return false
	}

	return true
}

// Create is a function to create a hashed string.
func (b *Bcrypt) Create(ctx context.Context, plain string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(password), nil
}
