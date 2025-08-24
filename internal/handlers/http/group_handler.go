package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		groupProtected.POST("/:group_id/members/", h.joinAction)
		groupProtected.DELETE("/:group_id/members/", h.leaveAction)
		groupProtected.DELETE("/:group_id/members/:member_id/", h.kickAction)
		groupProtected.POST("/:group_id/votes/", h.voteAction)
		groupProtected.DELETE("/:group_id/votes/", h.revokeVoteAction)
	}

	group := router.Group("groups")
	{
		group.GET("/", h.getListAction)
		group.GET("/:group_id/", h.getAction)
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

func (h GroupHandler) getAction(c *gin.Context) {

	groupIDStr := c.Param("group_id")
	if groupIDStr == "" {
		h.logger.Errorf("failed to get group id from ctx: %s", groupIDStr)
		c.JSON(
			httpresp.GetError(
				exceptions.NewBadRequestError(
					fmt.Errorf("empty group_id"),
				),
			),
		)
		return
	}

	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		h.logger.Errorf("failed to parse group id from ctx: %s", groupID)
		c.JSON(
			httpresp.GetError(
				exceptions.NewBadRequestError(
					fmt.Errorf("invalid group id"),
				),
			),
		)
		return
	}

	groupRepository := repositories.NewGroupRepository(h.db)

	uc := group.NewGetGroupUseCase(
		h.logger,
		groupRepository,
	)

	response, err := uc.Exec(groupID)
	if err != nil {
		h.logger.Errorf("get error from `get group` usecase: %v", err)
		c.JSON(httpresp.GetError(err))
		return
	}

	c.JSON(http.StatusOK, httpresp.SuccessDataResp{Data: response})
}

func (h GroupHandler) joinAction(c *gin.Context) {

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

	groupIDStr := c.Param("group_id")
	if groupIDStr == "" {
		h.logger.Errorf("failed to get group id from ctx: %s", groupIDStr)
		c.JSON(
			httpresp.GetError(
				exceptions.NewBadRequestError(
					fmt.Errorf("empty group_id"),
				),
			),
		)
		return
	}

	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		h.logger.Errorf("failed to parse group id from ctx: %s", groupID)
		c.JSON(
			httpresp.GetError(
				exceptions.NewBadRequestError(
					fmt.Errorf("invalid group id"),
				),
			),
		)
		return
	}

	groupRepository := repositories.NewGroupRepository(h.db)
	userRepository := repositories.NewUserRepository(h.db)

	dto := dtos.JoinGroupDTO{
		GroupID: groupID,
		UserID:  userID,
	}

	uc := group.NewJoinGroupUseCase(
		h.logger,
		groupRepository,
		userRepository,
	)

	err = uc.Exec(dto)
	if err != nil {
		h.logger.Errorf("get error from `join group usecase`: %v", err)
		c.JSON(httpresp.GetError(err))
		return
	}

	// TODO: add ws

	c.JSON(http.StatusOK, httpresp.SuccessDataResp{})
}

func (h GroupHandler) leaveAction(c *gin.Context) {

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

	groupIDStr := c.Param("group_id")
	if groupIDStr == "" {
		h.logger.Errorf("failed to get group id from ctx: %s", groupIDStr)
		c.JSON(
			httpresp.GetError(
				exceptions.NewBadRequestError(
					fmt.Errorf("empty group_id"),
				),
			),
		)
		return
	}

	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		h.logger.Errorf("failed to parse group id from ctx: %s", groupID)
		c.JSON(
			httpresp.GetError(
				exceptions.NewBadRequestError(
					fmt.Errorf("invalid group id"),
				),
			),
		)
		return
	}

	groupRepository := repositories.NewGroupRepository(h.db)
	userRepository := repositories.NewUserRepository(h.db)

	dto := dtos.LeaveGroupDTO{
		UserID:  userID,
		GroupID: groupID,
	}

	uc := group.NewLeaveGroupUseCase(
		h.logger,
		groupRepository,
		userRepository,
	)

	err = uc.Exec(dto)
	if err != nil {
		h.logger.Errorf("get error from `leave group` usecase: %v", err)
		c.JSON(httpresp.GetError(err))
		return
	}

	// TODO: add ws

	c.JSON(http.StatusOK, httpresp.SuccessDataResp{})
}

func (h GroupHandler) voteAction(c *gin.Context) {

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

	groupIDStr := c.Param("group_id")
	if groupIDStr == "" {
		h.logger.Errorf("failed to get group id from ctx: %s", groupIDStr)
		c.JSON(
			httpresp.GetError(
				exceptions.NewBadRequestError(
					fmt.Errorf("empty group_id"),
				),
			),
		)
		return
	}

	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		h.logger.Errorf("failed to parse group id from ctx: %s", groupID)
		c.JSON(
			httpresp.GetError(
				exceptions.NewBadRequestError(
					fmt.Errorf("invalid group id"),
				),
			),
		)
		return
	}

	var req VoteRequest

	err = c.ShouldBind(&req)
	if err != nil {
		h.logger.Errorf("failed to bind `create vote` request: %v", err)
		c.JSON(
			httpresp.GetError(
				exceptions.NewBadRequestError(fmt.Errorf("failed to bind request")),
			),
		)
		return
	}

	if invalid := h.validator.Struct(&req); invalid != nil {
		h.logger.Errorf("`create vote` request validation error: %v", invalid)
		c.JSON(httpresp.GetError(invalid))
		return
	}

	groupRepository := repositories.NewGroupRepository(h.db)

	dto := dtos.VoteDTO{
		GroupID:  groupID,
		UserID:   userID,
		BranchID: req.BranchID,
	}

	uc := group.NewVoteUseCase(
		h.logger,
		groupRepository,
	)

	err = uc.Exec(dto)
	if err != nil {
		h.logger.Errorf("get error from `vote` usecase: %v", err)
		c.JSON(httpresp.GetError(err))
		return
	}

	// TODO: add ws

	c.JSON(http.StatusOK, httpresp.SuccessDataResp{})
}

func (h GroupHandler) revokeVoteAction(c *gin.Context) {

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

	groupIDStr := c.Param("group_id")
	if groupIDStr == "" {
		h.logger.Errorf("failed to get group id from ctx: %s", groupIDStr)
		c.JSON(
			httpresp.GetError(
				exceptions.NewBadRequestError(
					fmt.Errorf("empty group_id"),
				),
			),
		)
		return
	}

	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		h.logger.Errorf("failed to parse group id from ctx: %s", groupID)
		c.JSON(
			httpresp.GetError(
				exceptions.NewBadRequestError(
					fmt.Errorf("invalid group id"),
				),
			),
		)
		return
	}

	groupRepository := repositories.NewGroupRepository(h.db)

	dto := dtos.RevokeVoteDTO{
		GroupID: groupID,
		UserID:  userID,
	}

	uc := group.NewRevokeVoteUseCase(
		h.logger,
		groupRepository,
	)

	err = uc.Exec(dto)
	if err != nil {
		h.logger.Errorf("get error from `revoke vote` usecase: %v", err)
		c.JSON(httpresp.GetError(err))
		return
	}

	// TODO: add ws

	c.JSON(http.StatusOK, httpresp.SuccessDataResp{})
}

func (h GroupHandler) kickAction(c *gin.Context) {

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

	groupIDStr := c.Param("group_id")
	if groupIDStr == "" {
		h.logger.Errorf("failed to get group id from ctx: %s", groupIDStr)
		c.JSON(
			httpresp.GetError(
				exceptions.NewBadRequestError(
					fmt.Errorf("empty group_id"),
				),
			),
		)
		return
	}

	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		h.logger.Errorf("failed to parse group id from ctx: %s", groupID)
		c.JSON(
			httpresp.GetError(
				exceptions.NewBadRequestError(
					fmt.Errorf("invalid group id"),
				),
			),
		)
		return
	}

	memberID, err := strconv.ParseInt(c.Param("member_id"), 10, 64)
	if err != nil {
		h.logger.Errorf("failed to parse member id from route: %v", err)
		c.JSON(httpresp.GetError(
			exceptions.NewBadRequestError(fmt.Errorf("invalid member id")),
		))
		return
	}

}
