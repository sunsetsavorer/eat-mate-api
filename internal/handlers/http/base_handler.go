package http

import (
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/config"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/db"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/logger"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/validator"
)

type BaseHdlr struct {
	db        *db.Db
	config    *config.Config
	logger    *logger.Logger
	validator *validator.Validator
}

func NewBaseHdlr(
	db *db.Db,
	config *config.Config,
	logger *logger.Logger,
	validator *validator.Validator,
) *BaseHdlr {

	return &BaseHdlr{
		db:        db,
		config:    config,
		logger:    logger,
		validator: validator,
	}
}
