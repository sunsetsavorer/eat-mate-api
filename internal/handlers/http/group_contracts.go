package http

import "github.com/google/uuid"

type (
	CreateGroupRequest struct {
		Name               string        `json:"name" validate:"required"`
		SelectionMode      string        `json:"selection_mode" validate:"required"`
		IsPublic           bool          `json:"is_public" validate:"required"`
		PlaceBranchID      uuid.NullUUID `json:"place_branch_id"`
		PlaceBranchOptions []uuid.UUID   `json:"place_branch_options"`
	}
)
