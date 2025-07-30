package entities

import "github.com/google/uuid"

type BrandEntity struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	IconPath *string   `json:"icon_path"`
}

func (e BrandEntity) GetID() uuid.UUID {
	return e.ID
}

func (e *BrandEntity) SetID(v uuid.UUID) {
	e.ID = v
}

func (e BrandEntity) GetName() string {
	return e.Name
}

func (e *BrandEntity) SetName(v string) {
	e.Name = v
}

func (e BrandEntity) GetIconPath() *string {
	return e.IconPath
}

func (e *BrandEntity) SetIconPath(v *string) {
	e.IconPath = v
}
