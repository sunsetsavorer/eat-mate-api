package entities

import "github.com/google/uuid"

type GroupMemberEntity struct {
	GroupID uuid.UUID  `json:"group_id"`
	UserID  int64      `json:"user_id"`
	User    UserEntity `json:"user"`
	Role    string     `json:"role"`
}

func (e GroupMemberEntity) GetGroupID() uuid.UUID {
	return e.GroupID
}

func (e *GroupMemberEntity) SetGroupID(v uuid.UUID) {
	e.GroupID = v
}

func (e GroupMemberEntity) GetUserID() int64 {
	return e.UserID
}

func (e *GroupMemberEntity) SetUserID(v int64) {
	e.UserID = v
}

func (e GroupMemberEntity) GetUser() UserEntity {
	return e.User
}

func (e *GroupMemberEntity) SetUser(v UserEntity) {
	e.User = v
}

func (e GroupMemberEntity) GetRole() string {
	return e.Role
}

func (e *GroupMemberEntity) SetRole(v string) {
	e.Role = v
}
