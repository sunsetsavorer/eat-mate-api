package app

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sunsetsavorer/eat-mate-api/internal/handlers/http"
)

func (app App) InitRouter() *gin.Engine {

	if !app.config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(
		cors.New(
			cors.Config{
				AllowOrigins:     app.config.AllowOrigins,
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
	)

	v1 := router.Group("v1")
	{
		userHdlr := http.NewAuthHandler(baseHdlr)
		userHdlr.RegisterRoutes(v1)
	}

	return router
}
