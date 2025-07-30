package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sunsetsavorer/eat-mate-api/internal/dtos"
	"github.com/sunsetsavorer/eat-mate-api/internal/exceptions"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/httpresp"
	"github.com/sunsetsavorer/eat-mate-api/internal/repositories"
	"github.com/sunsetsavorer/eat-mate-api/internal/services"
	"github.com/sunsetsavorer/eat-mate-api/internal/usecases/user"
)

type AuthHandler struct {
	*BaseHandler
}

func NewAuthHandler(baseHandler *BaseHandler) *AuthHandler {

	return &AuthHandler{baseHandler}
}

func (h AuthHandler) RegisterRoutes(router *gin.RouterGroup) {

	auth := router.Group("auth")
	{
		auth.POST("/signin/", h.authorizeAction)
	}
}

func (h AuthHandler) authorizeAction(c *gin.Context) {

	var req AuthorizeRequest

	err := c.ShouldBind(&req)
	if err != nil {
		h.logger.Errorf("failed to bind `authorize` request: %v", err)
		c.JSON(
			httpresp.GetError(
				exceptions.NewBadRequestError(fmt.Errorf("failed to bind request")),
			),
		)
		return
	}

	if invalid := h.validator.Struct(req); invalid != nil {
		h.logger.Errorf("`authorize` request validation error: %v", err)
		c.JSON(
			httpresp.GetError(invalid),
		)
		return
	}

	dto := dtos.AuthorizeDTO{
		TelegramID: req.TelegramID,
		Name:       req.Name,
		PhotoURL:   req.PhotoURL,
	}

	userRepository := repositories.NewUserRepository(h.db)
	jwtService := services.NewJWTService(
		h.config.JWT.Secret,
		time.Minute*h.config.JWT.LifetimeInMinutes,
	)

	uc := user.NewAuthorizeUseCase(
		h.logger,
		userRepository,
		jwtService,
	)

	token, err := uc.Exec(dto)
	if err != nil {
		h.logger.Errorf("get error from `authorize` usecase: %v", err)
		c.JSON(httpresp.GetError(err))
		return
	}

	c.JSON(http.StatusOK, httpresp.SuccessDataResp{Data: token})
}
