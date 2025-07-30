package http

type (
	GetBranchesRequest struct {
		Page  int     `form:"page" validate:"required"`
		Limit int     `form:"limit" validate:"required"`
		Query *string `form:"query"`
	}
)
