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
	"github.com/sunsetsavorer/eat-mate-api/internal/usecases/user"
)

type UserHandler struct {
	*BaseHandler
}

func NewUserHandler(baseHandler *BaseHandler) *UserHandler {

	return &UserHandler{baseHandler}
}

func (h UserHandler) RegisterRoutes(router *gin.RouterGroup) {

	jwtService := services.NewJWTService(
		h.config.JWT.Secret,
		h.config.JWT.LifetimeInMinutes,
	)

	authMiddleware := middlewares.NewAuthMiddleware(
		h.logger,
		jwtService,
	)

	userProtected := router.Group("users", authMiddleware.Check)
	{
		userProtected.PUT("/me/", h.updateAction)
	}
}

func (h UserHandler) updateAction(c *gin.Context) {

	userID, exists := h.GetUserID(c)
	if !exists {
		h.logger.Errorf("failed to get `user id` from context")
		c.JSON(
			httpresp.GetError(
				exceptions.NewUnauthorizedError(
					fmt.Errorf("unauthorized"),
				),
			),
		)
		return
	}

	var req UpdateUserRequest

	err := c.ShouldBind(&req)
	if err != nil {
		h.logger.Errorf("failed to bind `update user` request: %v", err)
		c.JSON(
			httpresp.GetError(
				exceptions.NewBadRequestError(fmt.Errorf("failed to bind request")),
			),
		)
		return
	}

	if invalid := h.validator.Struct(req); invalid != nil {
		h.logger.Errorf("`update user` request validation error: %v", invalid)
		c.JSON(httpresp.GetError(invalid))
		return
	}

	dto := dtos.UpdateUserDTO{
		UserID:   userID,
		Name:     req.Name,
		PhotoURL: req.PhotoURL,
	}

	userRepository := repositories.NewUserRepository(h.db)

	uc := user.NewUpdateUserUseCase(
		h.logger,
		userRepository,
	)

	if err := uc.Exec(dto); err != nil {
		h.logger.Errorf("get error from `update user` usecase: %v", err)
		c.JSON(
			httpresp.GetError(err),
		)
		return
	}

	c.JSON(http.StatusOK, httpresp.SuccessDataResp{})
}
