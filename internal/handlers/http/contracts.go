package http

type PaginationRequest struct {
	Page  int `form:"page" validate:"required"`
	Limit int `form:"limit" validate:"required"`
}
