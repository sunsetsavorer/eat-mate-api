package dtos

import "time"

type AuthorizeDTO struct {
	UserID        int64
	UserName      string
	UserPhotoURL  *string
	TokenSecret   string
	TokenLifetime time.Duration
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

func (dto AuthorizeDTO) GetTokenSecret() string {
	return dto.TokenSecret
}

func (dto AuthorizeDTO) GetTokenLifetime() time.Duration {
	return dto.TokenLifetime
}
