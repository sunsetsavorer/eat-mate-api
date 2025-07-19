package repositories

import (
	"fmt"

	"github.com/sunsetsavorer/eat-mate-api/internal/entities"
	"github.com/sunsetsavorer/eat-mate-api/internal/exceptions"
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
		return entities.UserEntity{}, exceptions.NewRepositoryError(err)
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
		return exceptions.NewRepositoryError(err)
	}

	return nil
}

func (r UserRepository) Update(entity entities.UserEntity) error {

	var user models.UserModel

	err := r.db.Client.
		First(&user, entity.GetID()).
		Error

	if err != nil {
		return exceptions.NewNotFoundError(fmt.Errorf("user with specified id was not found"))
	}

	user.FromEntity(entity)

	err = r.db.Client.
		Save(&user).
		Error

	if err != nil {
		return exceptions.NewRepositoryError(fmt.Errorf("failed to update user"))
	}

	return nil
}
