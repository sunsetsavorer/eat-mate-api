package models

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/sunsetsavorer/eat-mate-api/internal/entities"
	"github.com/sunsetsavorer/eat-mate-api/pkg/nullable"
)

type BranchModel struct {
	ID           uuid.UUID      `gorm:"column:id;primaryKey"`
	BrandID      uuid.UUID      `gorm:"column:brand_id;not null"`
	Brand        BrandModel     `gorm:"foreignKey:BrandID"`
	Address      sql.NullString `gorm:"column:address"`
	ContactPhone sql.NullString `gorm:"column:contact_phone"`
}

func (BranchModel) TableName() string {
	return "branches"
}

func (m BranchModel) ToEntity() entities.BranchEntity {

	return entities.BranchEntity{
		ID:           m.ID,
		BrandID:      m.BrandID,
		Brand:        m.Brand.ToEntity(),
		Address:      nullable.NullStringToPtr(m.Address),
		ContactPhone: nullable.NullStringToPtr(m.ContactPhone),
	}
}

func (m *BranchModel) FromEntity(e entities.BranchEntity) {

	m.ID = e.GetID()
	m.BrandID = e.GetBrandID()
	m.Address = nullable.PtrToNullString(e.GetAddress())
	m.ContactPhone = nullable.PtrToNullString(e.GetContactPhone())
}
