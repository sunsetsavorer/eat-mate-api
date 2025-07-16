package dtos

type AuthorizeDTO struct {
	UserID       int64
	UserName     string
	UserPhotoURL *string
}

func (dto AuthorizeDTO) GetUserID() int64 {
	return dto.UserID
}

func (dto AuthorizeDTO) GetUserName() string {
	return dto.UserName
}

func (dto AuthorizeDTO) GetUserPhotoURL() *string {
	return dto.UserPhotoURL
}
