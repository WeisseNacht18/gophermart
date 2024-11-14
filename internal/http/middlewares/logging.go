package middlewares

import (
	"net/http"
)

func WithLogging(h http.Handler) http.Handler {
	logFn := func(w http.ResponseWriter, r *http.Request) {

	}

	return http.HandlerFunc(logFn)
}
