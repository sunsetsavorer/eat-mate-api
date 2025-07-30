package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/config"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/db"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/logger"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/validator"
)

type BaseHandler struct {
	db        *db.Db
	config    *config.Config
	logger    *logger.Logger
	validator *validator.Validator
}

func NewBaseHandler(
	db *db.Db,
	config *config.Config,
	logger *logger.Logger,
	validator *validator.Validator,
) *BaseHandler {

	return &BaseHandler{
		db:        db,
		config:    config,
		logger:    logger,
		validator: validator,
	}
}

func (h BaseHandler) GetUserID(c *gin.Context) (int64, bool) {

	if userID, ok := c.Get("user_id"); ok {
		return userID.(int64), true
	}

	return 0, false
}
