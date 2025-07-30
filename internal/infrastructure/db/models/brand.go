package models

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/sunsetsavorer/eat-mate-api/internal/entities"
	"github.com/sunsetsavorer/eat-mate-api/pkg/nullable"
)

type BrandModel struct {
	ID       uuid.UUID      `gorm:"column:id;primaryKey"`
	Name     string         `gorm:"column:name;not null"`
	IconPath sql.NullString `gorm:"column:icon_path"`
}

func (BrandModel) TableName() string {
	return "brands"
}

func (m BrandModel) ToEntity() entities.BrandEntity {

	return entities.BrandEntity{
		ID:       m.ID,
		Name:     m.Name,
		IconPath: nullable.NullStringToPtr(m.IconPath),
	}
}

func (m *BrandModel) FromEntity(e entities.BrandEntity) {

	m.ID = e.GetID()
	m.Name = e.GetName()
	m.IconPath = nullable.PtrToNullString(e.GetIconPath())
}
