package dtos

import "github.com/google/uuid"

type KickMemberDTO struct {
	UserID   int64
	MemberID int64
	GroupID  uuid.UUID
}

func (dto KickMemberDTO) GetUserID() int64 {
	return dto.UserID
}

func (dto KickMemberDTO) GetMemberID() int64 {
	return dto.MemberID
}

func (dto KickMemberDTO) GetGroupID() uuid.UUID {
	return dto.GroupID
}
