package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sunsetsavorer/eat-mate-api/internal/dtos"
	"github.com/sunsetsavorer/eat-mate-api/internal/exceptions"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/httpresp"
	"github.com/sunsetsavorer/eat-mate-api/internal/repositories"
	"github.com/sunsetsavorer/eat-mate-api/internal/usecases/place"
)

type PlaceBranchHdlr struct {
	*BaseHdlr
}

func NewPlaceBranchHandler(baseHdlr *BaseHdlr) *PlaceBranchHdlr {

	return &PlaceBranchHdlr{baseHdlr}
}

func (hdlr PlaceBranchHdlr) RegisterRoutes(router *gin.RouterGroup) {

	placeBranch := router.Group("place_branches")
	{
		placeBranch.GET("/", hdlr.getListAction)
	}
}

func (hdlr PlaceBranchHdlr) getListAction(c *gin.Context) {

	var req GetPlaceBranchesRequest

	err := c.ShouldBindQuery(&req)
	if err != nil {
		hdlr.logger.Errorf("failed bind `get place branches` request: %v", err)
		c.JSON(
			httpresp.GetError(
				exceptions.NewBadRequestError(fmt.Errorf("failed to bind request")),
			),
		)
		return
	}

	if invalid := hdlr.validator.Struct(&req); invalid != nil {
		hdlr.logger.Errorf("`get place branches` request validation error: %v", err)
		c.JSON(httpresp.GetError(invalid))
		return
	}

	dto := dtos.GetPlaceBranchesDTO{
		Page:  req.Page,
		Limit: req.Limit,
		Query: req.Query,
	}

	repo := repositories.NewPlaceBranchRepository(hdlr.db)

	uc := place.NewGetPlaceBranchesUseCase(
		hdlr.logger,
		repo,
	)

	res, err := uc.Exec(dto)
	if err != nil {
		hdlr.logger.Errorf("get error from `get place branches` usecase: %v", err)
		c.JSON(
			httpresp.GetError(err),
		)
		return
	}

	c.JSON(http.StatusOK, httpresp.SuccessDataResp{Data: res})
}
