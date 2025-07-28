package http

type (
	UpdateUserRequest struct {
		Name     string  `json:"name" validate:"required"`
		PhotoURL *string `json:"photo_url"`
	}
)
