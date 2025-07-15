package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sunsetsavorer/eat-mate-api/internal/dtos"
	"github.com/sunsetsavorer/eat-mate-api/internal/repositories"
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
		c.JSON(http.StatusBadRequest, "failed to bind request")
		return
	}

	dto := dtos.AuthorizeDTO{
		UserID:        req.UserID,
		UserName:      req.UserName,
		UserPhotoURL:  req.UserPhotoURL,
		TokenSecret:   hdlr.config.JWTSecret,
		TokenLifetime: time.Hour * 24,
	}

	userRepository := repositories.NewUserRepository(hdlr.db)

	uc := user.NewAuthorizeUseCase(
		hdlr.logger,
		userRepository,
	)

	token, err := uc.Exec(dto)
	if err != nil {
		hdlr.logger.Errorf("get error from `authorize` usecase: %v", err)
		c.JSON(http.StatusBadRequest, "failed to authorize")
		return
	}

	c.JSON(http.StatusOK, SuccessDataResp{token})
}
