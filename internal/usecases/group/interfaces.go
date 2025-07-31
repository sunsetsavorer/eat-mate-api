package group

import "github.com/sunsetsavorer/eat-mate-api/internal/entities"

type GroupRepositoryInterface interface {
	Create(entity entities.GroupEntity) (entities.GroupEntity, error)
	GetList(filter GroupsFilter) ([]entities.GroupEntity, int64, error)
}

type UserRepositoryInterface interface {
	IsInAnyGroup(ID int64) (bool, error)
}
