package entities

import "github.com/google/uuid"

type GroupEntity struct {
	ID                 uuid.UUID           `json:"id"`
	Name               string              `json:"name"`
	IsPublic           bool                `json:"is_public"`
	IsActive           bool                `json:"is_active"`
	SelectionMode      string              `json:"selection_mode"`
	PlaceBranchID      uuid.NullUUID       `json:"place_branch_id"`
	PlaceBranch        PlaceBranchEntity   `json:"place_branch"`
	PlaceBranchOptions []PlaceBranchEntity `json:"place_branch_options"`
}

func (e GroupEntity) GetID() uuid.UUID {
	return e.ID
}

func (e *GroupEntity) SetID(v uuid.UUID) {
	e.ID = v
}

func (e GroupEntity) GetName() string {
	return e.Name
}

func (e *GroupEntity) SetName(v string) {
	e.Name = v
}

func (e GroupEntity) GetIsPublic() bool {
	return e.IsPublic
}

func (e *GroupEntity) SetIsPublic(v bool) {
	e.IsPublic = v
}

func (e GroupEntity) GetIsActive() bool {
	return e.IsActive
}

func (e *GroupEntity) SetIsActive(v bool) {
	e.IsActive = v
}

func (e GroupEntity) GetSelectionMode() string {
	return e.SelectionMode
}

func (e *GroupEntity) SetSelectionMode(v string) {
	e.SelectionMode = v
}

func (e GroupEntity) GetPlaceBranchID() uuid.NullUUID {
	return e.PlaceBranchID
}

func (e *GroupEntity) SetPlaceBranchID(v uuid.NullUUID) {
	e.PlaceBranchID = v
}

func (e GroupEntity) GetPlaceBranch() PlaceBranchEntity {
	return e.PlaceBranch
}

func (e *GroupEntity) SetPlaceBranch(v PlaceBranchEntity) {
	e.PlaceBranch = v
}

func (e GroupEntity) GetPlaceBranchOptions() []PlaceBranchEntity {
	return e.PlaceBranchOptions
}

func (e *GroupEntity) SetPlaceBranchOptions(v []PlaceBranchEntity) {
	e.PlaceBranchOptions = v
}
