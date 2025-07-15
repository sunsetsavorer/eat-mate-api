package user

import "github.com/sunsetsavorer/eat-mate-api/internal/entities"

type (
	UserRepositoryInterface interface {
		GetByID(ID int64) (entities.UserEntity, error)
		Create(entity entities.UserEntity) error
	}
)
