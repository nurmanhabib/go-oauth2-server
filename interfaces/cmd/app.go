package cmd

import "github.com/urfave/cli/v2"

type App struct {
	app *cli.App
}

func NewApp() *App {
	return &App{app: cli.NewApp()}
}

func (a *App) Init() {
	a.app.Action = func(context *cli.Context) error {
		return nil
	}
}

type Command func(app *cli.App)
