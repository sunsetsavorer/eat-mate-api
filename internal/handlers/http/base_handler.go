package http

import (
	"fmt"
	"net/http"

	"github.com/sunsetsavorer/eat-mate-api/internal/exceptions"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/config"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/db"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/logger"
)

type BaseHdlr struct {
	db     *db.Db
	config *config.Config
	logger *logger.Logger
}

func NewBaseHdlr(
	db *db.Db,
	config *config.Config,
	logger *logger.Logger,
) *BaseHdlr {

	return &BaseHdlr{
		db:     db,
		config: config,
		logger: logger,
	}
}

func (h BaseHdlr) getError(err error) (int, any) {

	if _, ok := err.(*exceptions.ManyRequestsError); ok {
		return http.StatusTooManyRequests, ErrorResp[OtherError]{
			Errors: OtherError{err.Error()},
		}
	}

	if _, ok := err.(*exceptions.NotFoundError); ok {
		return http.StatusNotFound, ErrorResp[OtherError]{
			Errors: OtherError{err.Error()},
		}
	}

	if _, ok := err.(*exceptions.BadRequestError); ok {
		return http.StatusBadRequest, ErrorResp[OtherError]{
			Errors: OtherError{err.Error()},
		}
	}

	return http.StatusBadRequest, fmt.Errorf("unknow response")
}
