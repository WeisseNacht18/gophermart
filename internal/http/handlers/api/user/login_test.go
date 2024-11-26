package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	api "github.com/WeisseNacht18/gophermart/internal/http/handlers/api/user"
	"github.com/WeisseNacht18/gophermart/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestLoginHandler(t *testing.T) {
	storage.NewMockStorage()
	storage.AddUser("userLogin", "1234")

	t.Run("correct logining", func(t *testing.T) {
		body := api.Authorization{
			Login:    "userLogin",
			Password: "1234",
		}
		byteData, _ := json.Marshal(body)
		request := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(byteData))
		request.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		api.LoginHandler(w, request)

		res := w.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		assert.NotEqual(t, 0, len(res.Cookies()), "cookies was not found")
	})

	t.Run("logining with incorrect json", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader("incorrect json"))
		request.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		api.LoginHandler(w, request)

		res := w.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("logining with incorrect Content-Type", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader("incorrect json"))
		request.Header.Set("Content-Type", "text/plain")

		w := httptest.NewRecorder()
		api.LoginHandler(w, request)

		res := w.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("logining with incorrect login", func(t *testing.T) {
		body := api.Authorization{
			Login:    "userUnexistLogin",
			Password: "1234",
		}
		byteData, _ := json.Marshal(body)
		request := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(byteData))
		request.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		api.LoginHandler(w, request)

		res := w.Result()
		assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
	})

	t.Run("logining with incorrect password", func(t *testing.T) {
		body := api.Authorization{
			Login:    "userLogin",
			Password: "12345678",
		}
		byteData, _ := json.Marshal(body)
		request := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(byteData))
		request.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		api.LoginHandler(w, request)

		res := w.Result()
		assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
	})
}
