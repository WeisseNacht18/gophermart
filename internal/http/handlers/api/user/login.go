package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/WeisseNacht18/gophermart/internal/storage"
)

type Authorization struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		var content Authorization

		err := json.NewDecoder(r.Body).Decode(&content)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		password, err := storage.GetUserPassword(content.Login)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if password != content.Password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := storage.AddToken(content.Login)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		authCookie := &http.Cookie{
			Name:  "auth",
			Value: token,
		}

		http.SetCookie(w, authCookie)

		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
