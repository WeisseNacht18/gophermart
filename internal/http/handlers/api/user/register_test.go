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

func TestRegisterHandler(t *testing.T) {
	storage.NewMockStorage()

	t.Run("correct register", func(t *testing.T) {
		body := api.Registration{
			Login:    "userRegister",
			Password: "1234",
		}
		byteData, _ := json.Marshal(body)
		request := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(byteData))
		request.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		api.RegisterHandler(w, request)

		res := w.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		assert.NotEqual(t, 0, len(res.Cookies()), "cookies was not found")
	})

	t.Run("incorrect rigister with duplicate login", func(t *testing.T) {
		body := api.Registration{
			Login:    "userRegister",
			Password: "1234",
		}
		byteData, _ := json.Marshal(body)
		request := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(byteData))
		request.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		api.RegisterHandler(w, request)

		res := w.Result()
		assert.Equal(t, http.StatusConflict, res.StatusCode)
	})

	t.Run("logining with incorrect json", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader("incorrect json"))
		request.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		api.RegisterHandler(w, request)

		res := w.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("logining with incorrect Content-Type", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader("incorrect json"))
		request.Header.Set("Content-Type", "text/plain")

		w := httptest.NewRecorder()
		api.RegisterHandler(w, request)

		res := w.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
}
