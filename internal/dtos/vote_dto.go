package dtos

import "github.com/google/uuid"

type VoteDTO struct {
	GroupID  uuid.UUID
	UserID   int64
	BranchID uuid.UUID
}

func (dto VoteDTO) GetGroupID() uuid.UUID {
	return dto.GroupID
}

func (dto VoteDTO) GetUserID() int64 {
	return dto.UserID
}

func (dto VoteDTO) GetBranchID() uuid.UUID {
	return dto.BranchID
}
