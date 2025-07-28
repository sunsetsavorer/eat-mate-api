package http

type AuthorizeRequest struct {
	TelegramID int64   `json:"telegram_id" validate:"required"`
	Name       string  `json:"name" validate:"required"`
	PhotoURL   *string `json:"photo_url"`
}
