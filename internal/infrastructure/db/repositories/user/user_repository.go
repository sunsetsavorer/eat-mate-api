package user

import (
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/db"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/db/models"
)

type UserRepository struct {
	Db *db.Db
}

func NewUserRepository(db *db.Db) *UserRepository {

	return &UserRepository{db}
}

func (repo UserRepository) IsExistsByTelegramID(telegramID int64) bool {

	var user models.UserModel

	repo.Db.firs

	return true
}
