package dtos

type AuthorizeDTO struct {
	TelegramID int64
	Name       string
	PhotoURL   *string
}

func (dto AuthorizeDTO) GetTelegramID() int64 {
	return dto.TelegramID
}

func (dto AuthorizeDTO) GetName() string {
	return dto.Name
}

func (dto AuthorizeDTO) GetPhotoURL() *string {
	return dto.PhotoURL
}
