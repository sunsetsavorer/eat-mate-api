package http

import (
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/config"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/db"
)

type BaseHdlr struct {
	db     *db.Db
	config *config.Config
}

func NewBaseHdlr(
	db *db.Db,
	config *config.Config,
) *BaseHdlr {

	return &BaseHdlr{
		db:     db,
		config: config,
	}
}
