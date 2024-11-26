package main

import (
	"github.com/WeisseNacht18/gophermart/internal/app"
	"github.com/WeisseNacht18/gophermart/internal/config"
)

func main() {
	config := config.NewConfig()

	app := app.NewApp(config)
	app.Run()
}
