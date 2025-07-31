package http

type (
	GetBranchesRequest struct {
		PaginationRequest
		Query *string `form:"query"`
	}
)
