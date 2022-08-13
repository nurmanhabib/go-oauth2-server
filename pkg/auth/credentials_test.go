package auth_test

import (
	"testing"

	"github.com/nurmanhabib/go-oauth2-server/pkg/auth"
	"github.com/stretchr/testify/assert"
)

func TestCredentials_HasPassword(t *testing.T) {
	t.Run("if nil credentials given", func(t *testing.T) {
		var credentials auth.Credentials

		assert.Nil(t, credentials)
		assert.False(t, credentials.HasPassword())
	})

	t.Run("if password credentials given", func(t *testing.T) {
		credentials := auth.Credentials{
			"username": "john",
			"password": "secret",
		}

		assert.True(t, credentials.HasPassword())
	})

	t.Run("if no password credentials given", func(t *testing.T) {
		credentials := auth.Credentials{
			"username": "john",
		}

		assert.False(t, credentials.HasPassword())
	})
}

func TestCredentials_GetPassword(t *testing.T) {
	t.Run("if nil credentials given", func(t *testing.T) {
		var credentials auth.Credentials

		assert.Nil(t, credentials)
		assert.Empty(t, credentials.GetPassword())
	})

	t.Run("if password credentials given", func(t *testing.T) {
		credentials := auth.Credentials{
			"username": "john",
			"password": "secret",
		}

		assert.Equal(t, "secret", credentials.GetPassword())
	})

	t.Run("if no password credentials given", func(t *testing.T) {
		credentials := auth.Credentials{
			"username": "john",
		}

		assert.Equal(t, "", credentials.GetPassword())
		assert.Empty(t, credentials.GetPassword())
	})
}

func TestCredentials_WithoutPassword(t *testing.T) {
	t.Run("if nil credentials given", func(t *testing.T) {
		var credentials auth.Credentials

		assert.Nil(t, credentials)
		assert.Empty(t, credentials.WithoutPassword())
	})

	t.Run("if filled credentials given", func(t *testing.T) {
		credentials := auth.Credentials{
			"username": "john",
			"status":   1,
			"password": "secret",
		}

		expectedCredentials := auth.Credentials{
			"username": "john",
			"status":   1,
		}

		assert.Equal(t, expectedCredentials, credentials.WithoutPassword())
	})
}
