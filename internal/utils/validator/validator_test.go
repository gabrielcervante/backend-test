package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	t.Run("Test Validator True", func(t *testing.T) {
		assert.False(t, IsEmpty("ok"))
	})

	t.Run("Test Validator False", func(t *testing.T) {
		assert.True(t, IsEmpty(""))
	})
}
