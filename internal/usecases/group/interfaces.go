package group

import "github.com/sunsetsavorer/eat-mate-api/internal/entities"

type GroupRepositoryInterface interface {
	Create(entity entities.GroupEntity) (entities.GroupEntity, error)
}

type UserRepositoryInterface interface {
	IsInAnyGroup(ID int64) (bool, error)
}
