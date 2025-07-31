package entities

type UserEntity struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	PhotoURL *string `json:"photo_url"`
}

func (e UserEntity) GetID() int64 {
	return e.ID
}

func (e *UserEntity) SetID(v int64) {
	e.ID = v
}

func (e UserEntity) GetName() string {
	return e.Name
}

func (e *UserEntity) SetName(v string) {
	e.Name = v
}

func (e UserEntity) GetPhotoURL() *string {
	return e.PhotoURL
}

func (e *UserEntity) SetPhotoURL(v *string) {
	e.PhotoURL = v
}
