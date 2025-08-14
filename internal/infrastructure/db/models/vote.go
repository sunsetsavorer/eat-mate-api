package models

import (
	"github.com/google/uuid"
	"github.com/sunsetsavorer/eat-mate-api/internal/entities"
)

type VoteModel struct {
	GroupID  uuid.UUID `gorm:"column:group_id"`
	UserID   int64     `gorm:"column:user_id"`
	BranchID uuid.UUID `gorm:"column:branch_id"`

	User   UserModel   `gorm:"foreignKey:UserID"`
	Branch BranchModel `gorm:"foreignKey:BranchID"`
}

func (VoteModel) TableName() string {
	return "votes"
}

func (m VoteModel) ToEntity() entities.VoteEntity {

	return entities.VoteEntity{
		GroupID:  m.GroupID,
		UserID:   m.UserID,
		BranchID: m.BranchID,

		User:   m.User.ToEntity(),
		Branch: m.Branch.ToEntity(),
	}
}

func (m *VoteModel) FromEntity(e entities.VoteEntity) {

	m.GroupID = e.GetGroupID()
	m.UserID = e.GetUserID()
	m.BranchID = e.GetBranchID()
}
