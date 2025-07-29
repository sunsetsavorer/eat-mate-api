package models

import (
	"github.com/google/uuid"
	"github.com/sunsetsavorer/eat-mate-api/internal/entities"
)

type GroupMemberModel struct {
	GroupID uuid.UUID `gorm:"column:group_id"`
	UserID  int64     `gorm:"column:user_id"`
	User    UserModel `gorm:"foreignKey:UserID"`
	Role    string    `gorm:"column:role"`
}

func (GroupMemberModel) TableName() string {
	return "group_members"
}

func (m GroupMemberModel) ToEntity() entities.GroupMemberEntity {

	return entities.GroupMemberEntity{
		GroupID: m.GroupID,
		UserID:  m.UserID,
		User:    m.User.ToEntity(),
		Role:    m.Role,
	}
}

func (m *GroupMemberModel) FromEntity(e entities.GroupMemberEntity) {

	m.GroupID = e.GetGroupID()
	m.UserID = e.GetUserID()
	m.Role = e.GetRole()
}
