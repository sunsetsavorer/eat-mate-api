package dtos

import "github.com/google/uuid"

type LeaveGroupDTO struct {
	UserID  int64
	GroupID uuid.UUID
}

func (dto LeaveGroupDTO) GetUserID() int64 {
	return dto.UserID
}

func (dto LeaveGroupDTO) GetGroupID() uuid.UUID {
	return dto.GroupID
}
