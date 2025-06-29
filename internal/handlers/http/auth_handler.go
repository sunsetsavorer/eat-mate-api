package http

import "github.com/gin-gonic/gin"

type AuthHdlr struct {
	*BaseHdlr
}

func NewAuthHandler(baseHdlr *BaseHdlr) *AuthHdlr {

	return &AuthHdlr{baseHdlr}
}

func (hdlr AuthHdlr) RegisterRoutes(router *gin.RouterGroup) {

}
