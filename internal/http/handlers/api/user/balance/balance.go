package api

import "net/http"

func GetBalanceHandler(w http.ResponseWriter, r *http.Request) {
	// запросить баланс из бд и выдать пользоватлю

	w.WriteHeader(http.StatusNotImplemented)
}
