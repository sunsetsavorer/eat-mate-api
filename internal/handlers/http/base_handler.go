package http

import (
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
