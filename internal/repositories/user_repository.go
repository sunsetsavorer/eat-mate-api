package repositories

import (
	"fmt"

	"github.com/sunsetsavorer/eat-mate-api/internal/entities"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/db"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/db/models"
)

type UserRepository struct {
	db *db.Db
}

func NewUserRepository(db *db.Db) *UserRepository {

	return &UserRepository{db}
}

func (r UserRepository) GetByID(ID int64) (entities.UserEntity, error) {

	var user models.UserModel

	err := r.db.Client.
		First(&user, ID).
		Error

	if err != nil {
		return entities.UserEntity{}, fmt.Errorf("user finding error: %v", err)
	}

	entity := user.ToEntity()

	return entity, nil
}

func (r UserRepository) Create(entity entities.UserEntity) error {

	var user models.UserModel

	user.FromEntity(entity)

	err := r.db.Client.
		Create(&user).
		Error

	if err != nil {
		return fmt.Errorf("user creating error: %v", err)
	}

	return nil
}
