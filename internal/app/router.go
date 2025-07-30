package app

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sunsetsavorer/eat-mate-api/internal/handlers/http"
)

func (app App) InitRouter() *gin.Engine {

	if !app.config.App.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(
		cors.New(
			cors.Config{
				AllowOrigins:     app.config.App.AllowOrigins,
				AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
				AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control"},
				ExposeHeaders:    []string{"Content-Length"},
				AllowCredentials: true,
				MaxAge:           12 * time.Hour,
			},
		),
	)

	baseHdlr := http.NewBaseHdlr(
		app.db,
		app.config,
		app.logger,
		app.validator,
	)

	v1 := router.Group("v1")
	{
		authHdlr := http.NewAuthHandler(baseHdlr)
		authHdlr.RegisterRoutes(v1)

		userHdlr := http.NewUserHandler(baseHdlr)
		userHdlr.RegisterRoutes(v1)

		branchHdlr := http.NewBranchHandler(baseHdlr)
		branchHdlr.RegisterRoutes(v1)

		groupHdlr := http.NewGroupHandler(baseHdlr)
		groupHdlr.RegisterRoutes(v1)
	}

	return router
}
