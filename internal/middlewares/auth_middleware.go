package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sunsetsavorer/eat-mate-api/internal/exceptions"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/httpresp"
	"github.com/sunsetsavorer/eat-mate-api/internal/usecases"
)

type AuthMiddleware struct {
	logger     usecases.LoggerInterface
	jwtService usecases.JWTServiceInterface
}

func NewAuthMiddleware(
	logger usecases.LoggerInterface,
	jwtService usecases.JWTServiceInterface,
) *AuthMiddleware {

	return &AuthMiddleware{
		logger,
		jwtService,
	}
}

func (mv AuthMiddleware) Check(c *gin.Context) {

	token := c.GetHeader("Authorization")

	tokenStruct, err := mv.jwtService.ParseToken(token)
	if err != nil {
		mv.logger.Errorf("get error while validating token: %v, %v", token, err)
		c.AbortWithStatusJSON(
			httpresp.GetError(exceptions.NewUnauthorizedError(
				fmt.Errorf("unauthorized"),
			)),
		)
		return
	}

	c.Set("user_id", tokenStruct.Payload.UserID)
	c.Next()
}
