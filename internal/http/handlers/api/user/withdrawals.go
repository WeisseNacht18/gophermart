package api

import "net/http"

func GetWithdrawalsHandler(w http.ResponseWriter, r *http.Request) {
	//Запросить из базы данных список заказов, в которых были списания
	//если данных в ответе нет, то сказать что заказов не было и выдать пустой список
	//если списания были, то выдать этот список списаний

	w.WriteHeader(http.StatusNotImplemented)
}
