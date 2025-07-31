package group

import (
	"github.com/google/uuid"
	"github.com/sunsetsavorer/eat-mate-api/internal/usecases"
)

const (
	OWNER_ROLE                       = "owner"
	PARTICIPANT_ROLE                 = "participant"
	DEFINED_SELECTION_MODE           = "defined"
	RANDOM_SELECTION_MODE            = "random"
	VOTING_SELECTION_MODE            = "voting"
	BRANCH_OPTIONS_LIST_MINIMAL_SIZE = 2
)

type (
	CreateGroupResponse struct {
		GroupID uuid.UUID `json:"group_id"`
	}

	GroupsFilter struct {
		usecases.PaginationFilter
	}
)
