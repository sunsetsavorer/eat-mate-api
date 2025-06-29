package main

import "github.com/sunsetsavorer/eat-mate-api/internal/app"

func main() {

	app := app.NewApp()

	app.InitInfrastructure()
	app.ExecAndLoop()
}
