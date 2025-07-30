package entities

import "github.com/google/uuid"

type BranchEntity struct {
	ID           uuid.UUID   `json:"id"`
	BrandID      uuid.UUID   `json:"brand_id"`
	Brand        BrandEntity `json:"brand"`
	Address      *string     `json:"address"`
	ContactPhone *string     `json:"contact_phone"`
}

func (e BranchEntity) GetID() uuid.UUID {
	return e.ID
}

func (e *BranchEntity) SetID(v uuid.UUID) {
	e.ID = v
}

func (e BranchEntity) GetBrandID() uuid.UUID {
	return e.BrandID
}

func (e *BranchEntity) SetBrandID(v uuid.UUID) {
	e.BrandID = v
}

func (e BranchEntity) GetBrand() BrandEntity {
	return e.Brand
}

func (e *BranchEntity) SetBrand(v BrandEntity) {
	e.Brand = v
}

func (e BranchEntity) GetAddress() *string {
	return e.Address
}

func (e *BranchEntity) SetAddress(v *string) {
	e.Address = v
}

func (e BranchEntity) GetContactPhone() *string {
	return e.ContactPhone
}

func (e *BranchEntity) SetContactPhone(v *string) {
	e.ContactPhone = v
}
