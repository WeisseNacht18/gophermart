package handlers

import (
	api "github.com/WeisseNacht18/gophermart/internal/http/handlers/api/user"
	api_user_balance "github.com/WeisseNacht18/gophermart/internal/http/handlers/api/user/balance"
	"github.com/go-chi/chi/v5"
)

func AddHandlersToRouter(router *chi.Mux) {
	router.Post("/api/user/register", api.RegisterHandler)
	router.Post("/api/user/login", api.LoginHandler)
	router.Post("/api/user/orders", api.AddOrderHandler)
	router.Get("/api/user/orders", api.GetOrdersHandler)
	router.Get("/api/user/balance", api_user_balance.GetBalanceHandler)
	router.Post("/api/user/balance/withdraw", api_user_balance.PostBalanceWithdrawHandler)
	router.Get("/api/user/withdrawals", api.GetWithdrawalsHandler)
}
