package models

import (
	"database/sql"

	"github.com/sunsetsavorer/eat-mate-api/internal/entities"
	"github.com/sunsetsavorer/eat-mate-api/pkg/nullable"
)

type UserModel struct {
	ID       int64          `gorm:"column:id;primaryKey"`
	Name     string         `gorm:"column:name;not null"`
	PhotoURL sql.NullString `gorm:"column:photo_url"`
}

func (UserModel) TableName() string {
	return "users"
}

func (m UserModel) ToEntity() entities.UserEntity {

	return entities.UserEntity{
		ID:       m.ID,
		Name:     m.Name,
		PhotoURL: nullable.NullStringToPtr(m.PhotoURL),
	}
}

func (m *UserModel) FromEntity(e entities.UserEntity) {

	m.ID = e.GetID()
	m.Name = e.GetName()
	m.PhotoURL = nullable.PtrToNullString(e.GetPhotoURL())
}
