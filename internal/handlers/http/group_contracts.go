package http

import "github.com/google/uuid"

type (
	CreateGroupRequest struct {
		Name          string        `json:"name" validate:"required"`
		SelectionMode string        `json:"selection_mode" validate:"required"`
		IsPublic      bool          `json:"is_public" validate:"required"`
		BranchID      uuid.NullUUID `json:"branch_id"`
		BranchOptions []uuid.UUID   `json:"branch_options"`
	}
)
