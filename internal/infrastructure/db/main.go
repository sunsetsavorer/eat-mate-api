package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	Client *gorm.DB
}

func NewDB(dsn string) (*Db, error) {

	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return &Db{}, fmt.Errorf("failed to create db connection: %v", err)
	}

	return &Db{client}, nil
}
