package models

import (
	"github.com/google/uuid"
	"github.com/sunsetsavorer/eat-mate-api/internal/entities"
)

type GroupModel struct {
	ID            uuid.UUID          `gorm:"column:id;default:gen_random_uuid();primaryKey"`
	Name          string             `gorm:"column:name;not null"`
	IsPublic      bool               `gorm:"column:is_public"`
	IsActive      bool               `gorm:"column:is_active"`
	SelectionMode string             `gorm:"column:selection_mode"`
	BranchID      uuid.NullUUID      `gorm:"column:branch_id"`
	Branch        BranchModel        `gorm:"foreignKey:BranchID"`
	BranchOptions []BranchModel      `gorm:"many2many:group_branch_options;joinForeignKey:group_id;joinReferences:branch_id"`
	Members       []GroupMemberModel `gorm:"foreignKey:GroupID"`
}

func (GroupModel) TableName() string {
	return "groups"
}

func (m GroupModel) ToEntity() entities.GroupEntity {

	options := make([]entities.BranchEntity, len(m.BranchOptions))

	for i, e := range m.BranchOptions {
		options[i] = e.ToEntity()
	}

	members := make([]entities.GroupMemberEntity, len(m.Members))

	for i, e := range m.Members {
		members[i] = e.ToEntity()
	}

	return entities.GroupEntity{
		ID:            m.ID,
		Name:          m.Name,
		IsPublic:      m.IsPublic,
		IsActive:      m.IsActive,
		SelectionMode: m.SelectionMode,
		BranchID:      m.BranchID,
		Branch:        m.Branch.ToEntity(),
		BranchOptions: options,
		Members:       members,
	}
}

func (m *GroupModel) FromEntity(e entities.GroupEntity) {

	m.ID = e.GetID()
	m.Name = e.GetName()
	m.IsPublic = e.GetIsPublic()
	m.IsActive = e.GetIsActive()
	m.SelectionMode = e.GetSelectionMode()
	m.BranchID = e.GetBranchID()
}
