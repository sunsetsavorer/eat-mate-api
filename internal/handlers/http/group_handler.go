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

type GroupHandler struct {
	*BaseHandler
}

func NewGroupHandler(baseHandler *BaseHandler) *GroupHandler {

	return &GroupHandler{baseHandler}
}

func (h GroupHandler) RegisterRoutes(router *gin.RouterGroup) {

	jwtService := services.NewJWTService(
		h.config.JWT.Secret,
		h.config.JWT.LifetimeInMinutes,
	)

	authMiddleware := middlewares.NewAuthMiddleware(
		h.logger,
		jwtService,
	)

	groupProtected := router.Group("groups", authMiddleware.Check)
	{
		groupProtected.POST("/", h.createAction)
	}

	group := router.Group("groups")
	{
		group.GET("/", h.getListAction)
	}
}

func (h GroupHandler) createAction(c *gin.Context) {

	userID, exists := h.GetUserID(c)
	if !exists {
		h.logger.Errorf("failed to get `user id` from context")
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
		h.logger.Errorf("failed to bind `create group` request: %v", err)
		c.JSON(
			httpresp.GetError(
				exceptions.NewBadRequestError(fmt.Errorf("failed to bind request")),
			),
		)
		return
	}

	if invalid := h.validator.Struct(&req); invalid != nil {
		h.logger.Errorf("`create group` request validation error: %v", invalid)
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

	groupRepo := repositories.NewGroupRepository(h.db)
	userRepo := repositories.NewUserRepository(h.db)

	uc := group.NewCreateGroupUseCase(
		h.logger,
		groupRepo,
		userRepo,
	)

	response, err := uc.Exec(dto)
	if err != nil {
		h.logger.Errorf("get error from `create group` usecase: %v", err)
		c.JSON(httpresp.GetError(err))
		return
	}

	c.JSON(http.StatusOK, httpresp.SuccessDataResp{Data: response})
}

func (h GroupHandler) getListAction(c *gin.Context) {

	var req GetGroupsRequest

	err := c.ShouldBindQuery(&req)
	if err != nil {
		h.logger.Errorf("failed to bind `get groups` request: %v", err)
		c.JSON(
			httpresp.GetError(
				exceptions.NewBadRequestError(fmt.Errorf("failed to bind request")),
			),
		)
		return
	}

	if invalid := h.validator.Struct(&req); invalid != nil {
		h.logger.Errorf("`get groups` request validation error: %v", invalid)
		c.JSON(httpresp.GetError(invalid))
		return
	}

	dto := dtos.GetGroupsDTO{
		Page:  req.Page,
		Limit: req.Limit,
	}

	groupRepo := repositories.NewGroupRepository(h.db)

	uc := group.NewGetGroupsUseCase(
		h.logger,
		groupRepo,
	)

	response, err := uc.Exec(dto)
	if err != nil {
		h.logger.Errorf("get error from `get groups` usecase: %v", err)
		c.JSON(httpresp.GetError(err))
		return
	}

	c.JSON(http.StatusOK, httpresp.SuccessDataResp{Data: response})
}
