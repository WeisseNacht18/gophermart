package api

import (
	"encoding/json"
	"net/http"

	"github.com/WeisseNacht18/gophermart/internal/entities"
	"github.com/WeisseNacht18/gophermart/internal/storage"
)

func GetWithdrawalsHandler(w http.ResponseWriter, r *http.Request) {
	chUserId := make(chan int)
	go func(login string, chOut chan int) {
		userId, err := storage.GetUserId(login)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		chOut <- userId
		close(chOut)
	}(r.Header.Get("login"), chUserId)

	chOut := make(chan entities.Withdraw)
	go func(chUserId <-chan int, chOut chan entities.Withdraw) {
		withdraws := storage.GetAllWithdraws(<-chUserId)

		for _, withdraw := range withdraws {
			chOut <- withdraw
		}

		close(chOut)
	}(chUserId, chOut)

	result := []entities.Withdraw{}
	for order := range chOut {
		result = append(result, order)
	}

	if len(result) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	bytes, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(bytes)
}
