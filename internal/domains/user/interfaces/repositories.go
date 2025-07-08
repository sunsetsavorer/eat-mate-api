package interfaces

import "github.com/sunsetsavorer/eat-mate-api/internal/domains/user/entities"

type UserRepositoryInterface interface {
	GetByID(ID int64) (entities.UserEntity, error)
}
