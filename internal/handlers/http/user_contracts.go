package http

type (
	UpdateUserRequest struct {
		Name     string  `form:"name" validate:"required"`
		PhotoURL *string `form:"photo_url"`
	}
)
