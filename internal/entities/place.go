package entities

import "github.com/google/uuid"

type PlaceEntity struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	IconPath *string   `json:"icon_path"`
}

func (e PlaceEntity) GetID() uuid.UUID {
	return e.ID
}

func (e *PlaceEntity) SetID(v uuid.UUID) {
	e.ID = v
}

func (e PlaceEntity) GetName() string {
	return e.Name
}

func (e *PlaceEntity) SetName(v string) {
	e.Name = v
}

func (e PlaceEntity) GetIconPath() *string {
	return e.IconPath
}

func (e *PlaceEntity) SetIconPath(v *string) {
	e.IconPath = v
}
