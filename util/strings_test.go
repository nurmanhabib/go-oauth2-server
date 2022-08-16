package util_test

import (
	"testing"

	"github.com/nurmanhabib/go-oauth2-server/util"
	"github.com/stretchr/testify/assert"
)

func TestSliceSubtract(t *testing.T) {
	t.Run("if difference given", func(t *testing.T) {
		a := []string{"view"}
		b := []string{"view", "read", "write"}

		c := util.SliceSubtract(a, b)

		var expected []string

		assert.Equal(t, 0, len(c))
		assert.Equal(t, expected, c)
		assert.Nil(t, c)
	})

	t.Run("if difference given", func(t *testing.T) {
		a := []string{"view", "update"}
		b := []string{"view", "read", "write"}

		c := util.SliceSubtract(a, b)

		expected := []string{"update"}

		assert.Equal(t, len(expected), len(c))
		assert.Equal(t, expected, c)
		assert.NotNil(t, c)
	})

	t.Run("if duplicate given", func(t *testing.T) {
		a := []string{"view", "view", "update"}
		b := []string{"view", "read", "write"}

		c := util.SliceSubtract(a, b)

		expected := []string{"update"}

		assert.Equal(t, len(expected), len(c))
		assert.Equal(t, expected, c)
		assert.NotNil(t, c)
	})
}
