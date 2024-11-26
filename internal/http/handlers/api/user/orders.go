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

	chIn := make(chan entities.Order)
	go func(chUserId <-chan int, chOut chan entities.Order) {
		orders := storage.GetAllOrders(<-chUserId)

		for _, order := range orders {
			chOut <- order
		}

		close(chOut)
	}(chUserId, chIn)

	chOut := make(chan entities.Order)
	chOrdersForUpdate := make(chan entities.Order)

	go func(chIn <-chan entities.Order, chOut chan entities.Order, chOrdersForUpdate chan entities.Order) {
		for order := range chIn {
			if order.Status == "REGISTRED" || order.Status == "PROCESSED" {
				var order entities.Order

				response, err := http.Get(api.AccrualSystemAddress + "/api/orders/" + order.ID)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				defer response.Body.Close()

				var queryOrder entities.Order
				err = json.NewDecoder(response.Body).Decode(&queryOrder)

				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				order.Status = queryOrder.Status
				order.Accrual = queryOrder.Accrual

				chOrdersForUpdate <- order
			}

			chOut <- order
		}

		close(chOrdersForUpdate)
		close(chOut)
	}(chIn, chOut, chOrdersForUpdate)

	go func(chIn <-chan entities.Order) {
		for order := range chIn {
			err := storage.UpdateOrder(order)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}(chOrdersForUpdate)

	result := []entities.Order{}
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
