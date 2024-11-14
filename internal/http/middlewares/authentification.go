package middlewares

import (
	"net/http"
)

func WithAuthentification(h http.Handler) http.Handler {
	authFn := func(w http.ResponseWriter, r *http.Request) {

	}

	return http.HandlerFunc(authFn)
}
