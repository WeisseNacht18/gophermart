package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/EClaesson/go-luhn"
	"github.com/WeisseNacht18/gophermart/internal/storage"
)

type QueryWithdraw struct {
	Order string  `json:"order"`
	Sum   float32 `json:"sum"`
}

func PostBalanceWithdrawHandler(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.Header.Get("Content-Type"), "application/json") {

		var query QueryWithdraw

		err := json.NewDecoder(r.Body).Decode(&query)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		valid, err := luhn.IsValid(query.Order)
		if err != nil || !valid {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		userId, err := storage.GetUserId(r.Header.Get("login"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		balance, err := storage.GetBalance(userId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if query.Sum > balance.Current {
			w.WriteHeader(http.StatusPaymentRequired)
			return
		}

		err = storage.UpdateBalance(userId, balance.Current-query.Sum, balance.Withdrawn+query.Sum)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = storage.AddWithdraw(userId, query.Order, query.Sum)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
