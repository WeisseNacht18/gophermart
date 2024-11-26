package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/EClaesson/go-luhn"
	"github.com/WeisseNacht18/gophermart/internal/entities"
	"github.com/WeisseNacht18/gophermart/internal/http/handlers/api"
	"github.com/WeisseNacht18/gophermart/internal/storage"
)

func GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	//запросить все данные из бд
	//Запросить все данные из сервиса

	//сопоставить данные запросов
	//сформировать структуру на отправку
	//отправить получивщийся массив

	w.WriteHeader(http.StatusNotImplemented)
}

func AddOrderHandler(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.Header.Get("Content-Type"), "text/plain") {
		body, err := io.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		orderId := string(body)

		valid, err := luhn.IsValid(orderId)
		if err != nil || !valid {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

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

		chOrder := make(chan entities.Order)
		go func(orderId string, chOut chan entities.Order) {
			var order entities.Order

			response, err := http.Get(api.AccrualSystemAddress + "/api/orders/" + orderId)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			defer response.Body.Close()
			err = json.NewDecoder(response.Body).Decode(&order)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			chOut <- order
			close(chOut)
		}(orderId, chOrder)

		chExistLinkOrderAndUser := make(chan bool)
		chOrderUserId := make(chan int)
		go func(orderID string, chExist chan bool, chResult chan int) {
			orderUserId, err := storage.GetOrder(orderID)
			if err == nil {
				chExist <- true
			} else {
				chExist <- false
			}
			chResult <- orderUserId
		}(orderId, chExistLinkOrderAndUser, chOrderUserId)

		if !<-chExistLinkOrderAndUser {
			storage.AddOrder(<-chOrder, <-chUserId)
			w.WriteHeader(http.StatusAccepted)
			return
		}

		if <-chUserId == <-chOrderUserId {
			w.WriteHeader(http.StatusOK)
			return
		} else {
			w.WriteHeader(http.StatusConflict)
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
