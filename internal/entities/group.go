package entities

import "github.com/google/uuid"

type GroupEntity struct {
	ID            uuid.UUID           `json:"id"`
	Name          string              `json:"name"`
	IsPublic      bool                `json:"is_public"`
	IsActive      bool                `json:"is_active"`
	SelectionMode string              `json:"selection_mode"`
	BranchID      uuid.NullUUID       `json:"branch_id"`
	Branch        BranchEntity        `json:"branch"`
	BranchOptions []BranchEntity      `json:"branch_options"`
	Members       []GroupMemberEntity `json:"members"`
	Votes         []VoteEntity        `json:"votes"`
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

func (e GroupEntity) GetBranchID() uuid.NullUUID {
	return e.BranchID
}

func (e *GroupEntity) SetBranchID(v uuid.NullUUID) {
	e.BranchID = v
}

func (e GroupEntity) GetBranch() BranchEntity {
	return e.Branch
}

func (e *GroupEntity) SetBranch(v BranchEntity) {
	e.Branch = v
}

func (e GroupEntity) GetBranchOptions() []BranchEntity {
	return e.BranchOptions
}

func (e *GroupEntity) SetBranchOptions(v []BranchEntity) {
	e.BranchOptions = v
}

func (e GroupEntity) GetMembers() []GroupMemberEntity {
	return e.Members
}

func (e *GroupEntity) SetMembers(v []GroupMemberEntity) {
	e.Members = v
}

func (e GroupEntity) GetVotes() []VoteEntity {
	return e.Votes
}

func (e *GroupEntity) SetVotes(v []VoteEntity) {
	e.Votes = v
}
