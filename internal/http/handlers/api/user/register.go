package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/WeisseNacht18/gophermart/internal/storage"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Registration struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		var content Registration

		err := json.NewDecoder(r.Body).Decode(&content)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		isExist := storage.CheckUser(content.Login)

		if isExist {
			w.WriteHeader(http.StatusConflict)
			return
		}

		err = storage.AddUser(content.Login, content.Password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
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
