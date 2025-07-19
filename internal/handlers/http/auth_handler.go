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

type AuthHdlr struct {
	*BaseHdlr
}

func NewAuthHandler(baseHdlr *BaseHdlr) *AuthHdlr {

	return &AuthHdlr{baseHdlr}
}

func (hdlr AuthHdlr) RegisterRoutes(router *gin.RouterGroup) {

	auth := router.Group("auth")
	{
		auth.POST("/signin/", hdlr.authorizeAction)
	}
}

func (hdlr AuthHdlr) authorizeAction(c *gin.Context) {

	var req AuthorizeRequest

	err := c.ShouldBind(&req)
	if err != nil {
		hdlr.logger.Errorf("failed to bind `authorize` request: %v", err)
		c.JSON(
			httpresp.GetError(
				exceptions.NewBadRequestError(fmt.Errorf("failed to bind request")),
			),
		)
		return
	}

	if invalid := hdlr.validator.Struct(req); invalid != nil {
		hdlr.logger.Errorf("`authorize` request validation error: %v", err)
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

	userRepository := repositories.NewUserRepository(hdlr.db)
	jwtService := services.NewJWTService(
		hdlr.config.JWT.Secret,
		time.Minute*hdlr.config.JWT.LifetimeInMinutes,
	)

	uc := user.NewAuthorizeUseCase(
		hdlr.logger,
		userRepository,
		jwtService,
	)

	token, err := uc.Exec(dto)
	if err != nil {
		hdlr.logger.Errorf("get error from `authorize` usecase: %v", err)
		c.JSON(httpresp.GetError(err))
		return
	}

	c.JSON(http.StatusOK, httpresp.SuccessDataResp{Data: token})
}
