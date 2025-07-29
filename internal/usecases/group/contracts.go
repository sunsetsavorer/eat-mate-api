package group

import "github.com/google/uuid"

type (
	CreateGroupResponse struct {
		GroupID uuid.UUID `json:"group_id"`
	}
)
