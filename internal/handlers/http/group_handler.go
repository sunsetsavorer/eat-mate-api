package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sunsetsavorer/eat-mate-api/internal/dtos"
	"github.com/sunsetsavorer/eat-mate-api/internal/exceptions"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/httpresp"
	"github.com/sunsetsavorer/eat-mate-api/internal/middlewares"
	"github.com/sunsetsavorer/eat-mate-api/internal/repositories"
	"github.com/sunsetsavorer/eat-mate-api/internal/services"
	"github.com/sunsetsavorer/eat-mate-api/internal/usecases/group"
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

	userID, exists := hdlr.GetUserID(c)
	if !exists {
		hdlr.logger.Errorf("failed to get `user id` from context")
		c.JSON(
			httpresp.GetError(
				exceptions.NewUnauthorizedError(fmt.Errorf("unauthorized")),
			),
		)
		return
	}

	var req CreateGroupRequest

	err := c.ShouldBind(&req)
	if err != nil {
		hdlr.logger.Errorf("failed to bind `create group` request: %v", err)
		c.JSON(
			httpresp.GetError(
				exceptions.NewBadRequestError(fmt.Errorf("failed to bind request")),
			),
		)
		return
	}

	if invalid := hdlr.validator.Struct(&req); invalid != nil {
		hdlr.logger.Errorf("`create group` request validation error: %v", invalid)
		c.JSON(httpresp.GetError(invalid))
		return
	}

	dto := dtos.CreateGroupDTO{
		Name:          req.Name,
		SelectionMode: req.SelectionMode,
		IsPublic:      req.IsPublic,
		BranchID:      req.BranchID,
		BranchOptions: req.BranchOptions,
		OwnerID:       userID,
	}

	groupRepo := repositories.NewGroupRepository(hdlr.db)
	userRepo := repositories.NewUserRepository(hdlr.db)

	uc := group.NewCreateGroupUseCase(
		hdlr.logger,
		groupRepo,
		userRepo,
	)

	res, err := uc.Exec(dto)
	if err != nil {
		hdlr.logger.Errorf("get error from `create group` usecase: %v", err)
		c.JSON(httpresp.GetError(err))
		return
	}

	c.JSON(http.StatusOK, httpresp.SuccessDataResp{Data: res})
}
