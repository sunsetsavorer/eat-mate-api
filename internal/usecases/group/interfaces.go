package group

import (
	"github.com/google/uuid"
	"github.com/sunsetsavorer/eat-mate-api/internal/entities"
)

type GroupRepositoryInterface interface {
	Create(entity entities.GroupEntity) (entities.GroupEntity, error)
	GetList(filter GroupsFilter) ([]entities.GroupEntity, int64, error)
	GetByID(ID uuid.UUID) (entities.GroupEntity, error)
	AddMember(entity entities.GroupMemberEntity) error
	RemoveMember(entity entities.GroupMemberEntity) error
	GetMemberByID(userID int64, groupID uuid.UUID) (entities.GroupMemberEntity, error)
	DeactivateByID(groupID uuid.UUID) error
	SaveVote(entity entities.VoteEntity) error
}

type UserRepositoryInterface interface {
	IsInAnyGroup(ID int64) (bool, error)
}
