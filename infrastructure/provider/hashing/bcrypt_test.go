package hashing

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBcrypt_Create(t *testing.T) {
	b := &Bcrypt{}
	passwordHashed, err := b.Create(context.Background(), "password")

	if assert.NoError(t, err) {
		assert.NotEmpty(t, passwordHashed)
	}
}
