package entities

import "github.com/google/uuid"

type PlaceBranchEntity struct {
	ID           uuid.UUID   `json:"id"`
	PlaceID      uuid.UUID   `json:"place_id"`
	Place        PlaceEntity `json:"place"`
	Address      *string     `json:"address"`
	ContactPhone *string     `json:"contact_phone"`
}

func (e PlaceBranchEntity) GetID() uuid.UUID {
	return e.ID
}

func (e *PlaceBranchEntity) SetID(v uuid.UUID) {
	e.ID = v
}

func (e PlaceBranchEntity) GetPlaceID() uuid.UUID {
	return e.PlaceID
}

func (e *PlaceBranchEntity) SetPlaceID(v uuid.UUID) {
	e.PlaceID = v
}

func (e PlaceBranchEntity) GetPlace() PlaceEntity {
	return e.Place
}

func (e *PlaceBranchEntity) SetPlace(v PlaceEntity) {
	e.Place = v
}

func (e PlaceBranchEntity) GetAddress() *string {
	return e.Address
}

func (e *PlaceBranchEntity) SetAddress(v *string) {
	e.Address = v
}

func (e PlaceBranchEntity) GetContactPhone() *string {
	return e.ContactPhone
}

func (e *PlaceBranchEntity) SetContactPhone(v *string) {
	e.ContactPhone = v
}
