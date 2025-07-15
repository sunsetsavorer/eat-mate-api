package http

type AuthorizeRequest struct {
	UserID       int64   `form:"telegram_id"`
	UserName     string  `form:"name"`
	UserPhotoURL *string `form:"photo_url"`
}
