package http

type AuthorizeRequest struct {
	UserID       int64   `form:"telegram_id" validate:"required"`
	UserName     string  `form:"name" validate:"required"`
	UserPhotoURL *string `form:"photo_url"`
}
