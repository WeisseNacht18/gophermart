package app

type App struct {
	name string
}

func New() (app *App) {
	app = &App{
		name: "gophermart",
	}

	return
}

func (app *App) Run() {

}
