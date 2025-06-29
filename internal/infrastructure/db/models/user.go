package models

import "database/sql"

type User struct {
	ID       int64          `gorm:"column:id;primaryKey"`
	Name     string         `gorm:"column:name;not null"`
	PhotoURL sql.NullString `gorm:"column:photo_url"`
}

func (User) TableName() string {
	return "users"
}
