package middlewares

import (
	"net/http"

	"github.com/WeisseNacht18/gophermart/internal/storage"
)

func WithAuthentification(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/api/user/register" || r.RequestURI == "/api/user/login" {
			next.ServeHTTP(w, r)
			return
		}

		token, err := r.Cookie("auth")

		if err == nil {
			login, err := storage.FindToken(token.Value)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
			}

			r.Header.Set("login", login)

			next.ServeHTTP(w, r)
			return
		}

		w.WriteHeader(http.StatusUnauthorized)
	})
}
