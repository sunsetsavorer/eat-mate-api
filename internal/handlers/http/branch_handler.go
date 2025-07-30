package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sunsetsavorer/eat-mate-api/internal/dtos"
	"github.com/sunsetsavorer/eat-mate-api/internal/exceptions"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/httpresp"
	"github.com/sunsetsavorer/eat-mate-api/internal/repositories"
	"github.com/sunsetsavorer/eat-mate-api/internal/usecases/branch"
)

type BranchHdlr struct {
	*BaseHdlr
}

func NewBranchHandler(baseHdlr *BaseHdlr) *BranchHdlr {

	return &BranchHdlr{baseHdlr}
}

func (hdlr BranchHdlr) RegisterRoutes(router *gin.RouterGroup) {

	branch := router.Group("branches")
	{
		branch.GET("/", hdlr.getListAction)
	}
}

func (hdlr BranchHdlr) getListAction(c *gin.Context) {

	var req GetBranchesRequest

	err := c.ShouldBindQuery(&req)
	if err != nil {
		hdlr.logger.Errorf("failed bind `get branches` request: %v", err)
		c.JSON(
			httpresp.GetError(
				exceptions.NewBadRequestError(fmt.Errorf("failed to bind request")),
			),
		)
		return
	}

	if invalid := hdlr.validator.Struct(&req); invalid != nil {
		hdlr.logger.Errorf("`get branches` request validation error: %v", err)
		c.JSON(httpresp.GetError(invalid))
		return
	}

	dto := dtos.GetBranchesDTO{
		Page:  req.Page,
		Limit: req.Limit,
		Query: req.Query,
	}

	repo := repositories.NewBranchRepository(hdlr.db)

	uc := branch.NewGetBranchesUseCase(
		hdlr.logger,
		repo,
	)

	res, err := uc.Exec(dto)
	if err != nil {
		hdlr.logger.Errorf("get error from `get branches` usecase: %v", err)
		c.JSON(
			httpresp.GetError(err),
		)
		return
	}

	c.JSON(http.StatusOK, httpresp.SuccessDataResp{Data: res})
}
