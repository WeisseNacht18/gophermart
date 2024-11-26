package api

import "net/http"

func PostBalanceWithdrawHandler(w http.ResponseWriter, r *http.Request) {
	//проверить заказ алгоритмом Луна
	//проверить баланс
	//если не хватает то выдать определенный статус
	//если все ок, то выдать что все ок
	w.WriteHeader(http.StatusNotImplemented)
}
