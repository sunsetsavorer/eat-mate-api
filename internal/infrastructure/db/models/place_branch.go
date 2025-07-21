package models

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/sunsetsavorer/eat-mate-api/internal/entities"
	"github.com/sunsetsavorer/eat-mate-api/pkg/nullable"
)

type PlaceBranchModel struct {
	ID           uuid.UUID      `gorm:"column:id;primaryKey"`
	PlaceID      uuid.UUID      `gorm:"column:place_id;not null"`
	Place        PlaceModel     `gorm:"foreignKey:PlaceID"`
	Address      sql.NullString `gorm:"column:address"`
	ContactPhone sql.NullString `gorm:"column:contact_phone"`
}

func (PlaceBranchModel) TableName() string {
	return "place_branches"
}

func (m PlaceBranchModel) ToEntity() entities.PlaceBranchEntity {

	return entities.PlaceBranchEntity{
		ID:           m.ID,
		PlaceID:      m.PlaceID,
		Place:        m.Place.ToEntity(),
		Address:      nullable.NullStringToPtr(m.Address),
		ContactPhone: nullable.NullStringToPtr(m.ContactPhone),
	}
}

func (m *PlaceBranchModel) FromEntity(e entities.PlaceBranchEntity) {

	var place PlaceModel

	place.FromEntity(e.GetPlace())

	m.ID = e.GetID()
	m.PlaceID = e.GetPlaceID()
	m.Place = place
	m.Address = nullable.PtrToNullString(e.GetAddress())
	m.ContactPhone = nullable.PtrToNullString(e.GetContactPhone())
}
