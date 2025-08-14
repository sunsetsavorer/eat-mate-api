package entities

import "github.com/google/uuid"

type VoteEntity struct {
	GroupID  uuid.UUID     `json:"group_id"`
	UserID   int64         `json:"user_id"`
	BranchID uuid.NullUUID `json:"branch_id"`

	User   UserEntity   `json:"user"`
	Branch BranchEntity `json:"branch"`
}

func (e VoteEntity) GetGroupID() uuid.UUID {
	return e.GroupID
}

func (e *VoteEntity) SetGroupID(v uuid.UUID) {
	e.GroupID = v
}

func (e VoteEntity) GetUserID() int64 {
	return e.UserID
}

func (e *VoteEntity) SetUserID(v int64) {
	e.UserID = v
}

func (e VoteEntity) GetBranchID() uuid.NullUUID {
	return e.BranchID
}

func (e *VoteEntity) SetBranchID(v uuid.NullUUID) {
	e.BranchID = v
}

func (e VoteEntity) GetUser() UserEntity {
	return e.User
}

func (e *VoteEntity) SetUser(v UserEntity) {
	e.User = v
}

func (e VoteEntity) GetBranch() BranchEntity {
	return e.Branch
}

func (e *VoteEntity) SetBranch(v BranchEntity) {
	e.Branch = v
}
