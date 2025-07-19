package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sunsetsavorer/eat-mate-api/internal/exceptions"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/httpresp"
	"github.com/sunsetsavorer/eat-mate-api/internal/usecases"
)

type AuthMiddleware struct {
	log        usecases.LoggerInterface
	jwtService usecases.JWTServiceInterface
}

func NewAuthMiddleware(
	log usecases.LoggerInterface,
	jwtService usecases.JWTServiceInterface,
) *AuthMiddleware {

	return &AuthMiddleware{
		log,
		jwtService,
	}
}

func (mv AuthMiddleware) Check(c *gin.Context) {

	token := c.GetHeader("Authorization")

	_, err := mv.jwtService.ParseToken(token)
	if err != nil {
		mv.log.Errorf("get error while validating token: %v, %v", token, err)
		c.JSON(
			httpresp.GetError(exceptions.NewUnauthorizedError(
				fmt.Errorf("unauthorized"),
			)),
		)
		return
	}

	c.Next()
}
