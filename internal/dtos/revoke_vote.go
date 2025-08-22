package dtos

import "github.com/google/uuid"

type RevokeVoteDTO struct {
	GroupID uuid.UUID
	UserID  int64
}

func (dto RevokeVoteDTO) GetGroupID() uuid.UUID {
	return dto.GroupID
}

func (dto RevokeVoteDTO) GetUserID() int64 {
	return dto.UserID
}
