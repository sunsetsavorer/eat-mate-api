package http

type AuthorizeRequest struct {
	TelegramID int64   `form:"telegram_id" validate:"required"`
	Name       string  `form:"name" validate:"required"`
	PhotoURL   *string `form:"photo_url"`
}
