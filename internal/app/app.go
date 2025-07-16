package app

import (
	"context"

	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/config"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/db"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/logger"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/server"
)

type App struct {
	db     *db.Db
	config *config.Config
	logger *logger.Logger
}

func NewApp() *App {

	return &App{}
}

func (app *App) InitInfrastructure() {

	app.logger = logger.NewLogger()

	config := config.NewConfig()
	err := config.LoadConfig()

	if err != nil {
		app.logger.Errorf("an error occured while load config: %v", err)
	}

	app.config = config

	db, err := db.NewDB(app.config.App.DbConn)

	if err != nil {
		app.logger.Errorf("db connection error: %v", err)
	}

	app.db = db
}

func (app App) ExecAndLoop() {

	router := app.InitRouter()

	srv := server.NewServer(app.config.App.Port, router)

	err := srv.Start()

	if err != nil {
		app.logger.Errorf("server start error: %v", err)
		return
	}

	defer srv.Stop(context.TODO())
}
