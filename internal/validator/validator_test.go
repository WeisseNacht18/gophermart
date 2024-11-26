package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator(t *testing.T) {
	t.Run("correct run address", func(t *testing.T) {
		assert.NoError(t, IsValidRunAddress("127.0.0.1:8080"), "incorrect run address")
	})

	t.Run("incorrect run address", func(t *testing.T) {
		assert.Error(t, IsValidRunAddress("355.12.33.0:9999999"), "correct run address")
	})
}
