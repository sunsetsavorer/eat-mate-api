package dtos

import "github.com/google/uuid"

type CreateGroupDTO struct {
	Name          string
	SelectionMode string
	IsPublic      bool
	BranchID      uuid.NullUUID
	BranchOptions []uuid.UUID
	OwnerID       int64
}

func (dto CreateGroupDTO) GetName() string {
	return dto.Name
}

func (dto CreateGroupDTO) GetSelectionMode() string {
	return dto.SelectionMode
}

func (dto CreateGroupDTO) GetIsPublic() bool {
	return dto.IsPublic
}

func (dto CreateGroupDTO) GetBranchID() uuid.NullUUID {
	return dto.BranchID
}

func (dto CreateGroupDTO) GetBranchOptions() []uuid.UUID {
	return dto.BranchOptions
}

func (dto CreateGroupDTO) GetOwnerID() int64 {
	return dto.OwnerID
}
