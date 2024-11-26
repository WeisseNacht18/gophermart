package storage_test

import (
	"testing"

	"github.com/WeisseNacht18/gophermart/internal/builder"
	"github.com/WeisseNacht18/gophermart/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestJWTStorage(t *testing.T) {
	storage.NewJWTStorage()
	var token string
	t.Run("correct adding token", func(t *testing.T) {
		addedToken, err := storage.AddToken("admin")
		assert.NoError(t, err, "user was not added")
		token = addedToken
	})

	t.Run("find added token", func(t *testing.T) {
		login, err := storage.FindToken(token)
		assert.NoError(t, err, "login was not found")
		assert.Equal(t, "admin", login, "login != 'admin'")
	})

	t.Run("find unexist token", func(t *testing.T) {
		someToken, err := builder.BuildJWTStringWithLogin("user1")
		assert.NoError(t, err, "jwt string didn't create")
		_, err = storage.FindToken(someToken)
		assert.NotEqual(t, token, someToken)
		assert.Error(t, err, "login was found")
	})
}
