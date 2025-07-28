package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sunsetsavorer/eat-mate-api/internal/middlewares"
	"github.com/sunsetsavorer/eat-mate-api/internal/services"
)

type GroupHdlr struct {
	*BaseHdlr
}

func NewGroupHandler(baseHdlr *BaseHdlr) *GroupHdlr {

	return &GroupHdlr{baseHdlr}
}

func (hdlr GroupHdlr) RegisterRoutes(router *gin.RouterGroup) {

	jwtService := services.NewJWTService(
		hdlr.config.JWT.Secret,
		hdlr.config.JWT.LifetimeInMinutes,
	)

	authMiddleware := middlewares.NewAuthMiddleware(
		hdlr.logger,
		jwtService,
	)

	groupProtected := router.Group("groups", authMiddleware.Check)
	{
		groupProtected.POST("/", hdlr.createAction)
	}
}

func (hdlr GroupHdlr) createAction(c *gin.Context) {

}
