package dtos

type UpdateUserDTO struct {
	UserID   int64
	Name     string
	PhotoURL *string
}

func (dto UpdateUserDTO) GetUserID() int64 {
	return dto.UserID
}

func (dto UpdateUserDTO) GetName() string {
	return dto.Name
}

func (dto UpdateUserDTO) GetPhotoURL() *string {
	return dto.PhotoURL
}
