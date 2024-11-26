package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/WeisseNacht18/gophermart/internal/storage"
)

func GetBalanceHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := storage.GetUserId(r.Header.Get("login"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println(userId)

	balance, err := storage.GetBalance(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("some point")

	bytes, err := json.Marshal(balance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
