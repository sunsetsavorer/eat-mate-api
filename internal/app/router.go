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

	baseHandler := http.NewBaseHandler(
		app.db,
		app.config,
		app.logger,
		app.validator,
	)

	v1 := router.Group("v1")
	{
		authHandler := http.NewAuthHandler(baseHandler)
		authHandler.RegisterRoutes(v1)

		userHandler := http.NewUserHandler(baseHandler)
		userHandler.RegisterRoutes(v1)

		branchHandler := http.NewBranchHandler(baseHandler)
		branchHandler.RegisterRoutes(v1)

		groupHandler := http.NewGroupHandler(baseHandler)
		groupHandler.RegisterRoutes(v1)
	}

	return router
}
