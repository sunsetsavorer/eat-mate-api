package dtos

import "github.com/google/uuid"

type CreateGroupDTO struct {
	Name               string
	SelectionMode      string
	IsPublic           bool
	PlaceBranchID      uuid.UUID
	PlaceBranchOptions []uuid.UUID
	OwnerID            int64
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

func (dto CreateGroupDTO) GetPlaceBranchID() uuid.UUID {
	return dto.PlaceBranchID
}

func (dto CreateGroupDTO) GetPlaceBranchOptions() []uuid.UUID {
	return dto.PlaceBranchOptions
}

func (dto CreateGroupDTO) GetOwnerID() int64 {
	return dto.OwnerID
}
