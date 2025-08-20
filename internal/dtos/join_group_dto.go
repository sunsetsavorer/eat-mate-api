package dtos

import "github.com/google/uuid"

type JoinGroupDTO struct {
	UserID  int64
	GroupID uuid.UUID
}

func (dto JoinGroupDTO) GetUserID() int64 {
	return dto.UserID
}

func (dto JoinGroupDTO) GetGroupID() uuid.UUID {
	return dto.GroupID
}
