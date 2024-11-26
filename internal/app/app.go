package app

import (
	"log"
	"net/http"

	"github.com/WeisseNacht18/gophermart/internal/config"
	"github.com/WeisseNacht18/gophermart/internal/http/handlers"
	"github.com/WeisseNacht18/gophermart/internal/http/handlers/api"
	"github.com/WeisseNacht18/gophermart/internal/http/middlewares"
	"github.com/WeisseNacht18/gophermart/internal/storage"
	"github.com/go-chi/chi/v5"
)

type App struct {
	name   string
	Config config.Config
}

func NewApp(config config.Config) *App {
	return &App{
		name:   "gophermart",
		Config: config,
	}
}

func (app *App) Run() {
	router := chi.NewRouter()

	api.NewApi(app.Config.AccrualSystemAddress)

	storage.NewStorage(app.Config)

	middlewares.AddMiddlewaresToRouter(router)

	handlers.AddHandlersToRouter(router)

	err := http.ListenAndServe(app.Config.RunAddress, router)

	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
