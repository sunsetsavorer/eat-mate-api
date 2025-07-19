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

type UserHdlr struct {
	*BaseHdlr
}

func NewUserHandler(baseHdlr *BaseHdlr) *UserHdlr {

	return &UserHdlr{baseHdlr}
}

func (hdlr UserHdlr) RegisterRoutes(router *gin.RouterGroup) {

	jwtService := services.NewJWTService(
		hdlr.config.JWT.Secret,
		hdlr.config.JWT.LifetimeInMinutes,
	)

	authMiddleware := middlewares.NewAuthMiddleware(
		hdlr.logger,
		jwtService,
	)

	userProtected := router.Group("users", authMiddleware.Check)
	{
		userProtected.PUT("/me/", hdlr.updateAction)
	}
}

func (hdlr UserHdlr) updateAction(c *gin.Context) {

	userID, exists := hdlr.GetUserID(c)
	if !exists {
		hdlr.logger.Errorf("failed to get `user id` from context")
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
		hdlr.logger.Errorf("failed to bind `update user` request: %v", err)
		c.JSON(
			httpresp.GetError(
				exceptions.NewBadRequestError(fmt.Errorf("failed to bind request")),
			),
		)
		return
	}

	if invalid := hdlr.validator.Struct(req); invalid != nil {
		hdlr.logger.Errorf("`update user` request validation error: %v", err)
		c.JSON(
			httpresp.GetError(invalid),
		)
		return
	}

	dto := dtos.UpdateUserDTO{
		UserID:   userID,
		Name:     req.Name,
		PhotoURL: req.PhotoURL,
	}

	userRepository := repositories.NewUserRepository(hdlr.db)

	uc := user.NewUpdateUserUseCase(
		hdlr.logger,
		userRepository,
	)

	if err := uc.Exec(dto); err != nil {
		hdlr.logger.Errorf("get error from `update user` usecase: %v", err)
		c.JSON(
			httpresp.GetError(err),
		)
		return
	}

	c.JSON(http.StatusOK, httpresp.SuccessDataResp{})
}
