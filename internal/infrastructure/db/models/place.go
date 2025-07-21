package models

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/sunsetsavorer/eat-mate-api/internal/entities"
	"github.com/sunsetsavorer/eat-mate-api/pkg/nullable"
)

type PlaceModel struct {
	ID       uuid.UUID      `gorm:"column:id;primaryKey"`
	Name     string         `gorm:"column:name;not null"`
	IconPath sql.NullString `gorm:"column:icon_path"`
}

func (m PlaceModel) TableName() string {
	return "places"
}

func (m PlaceModel) ToEntity() entities.PlaceEntity {

	return entities.PlaceEntity{
		ID:       m.ID,
		Name:     m.Name,
		IconPath: nullable.NullStringToPtr(m.IconPath),
	}
}

func (m *PlaceModel) FromEntity(e entities.PlaceEntity) {

	m.ID = e.GetID()
	m.Name = e.GetName()
	m.IconPath = nullable.PtrToNullString(e.GetIconPath())
}
