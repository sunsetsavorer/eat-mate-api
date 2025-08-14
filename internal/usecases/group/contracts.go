package group

import (
	"github.com/google/uuid"
	"github.com/sunsetsavorer/eat-mate-api/internal/entities"
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
		Page  int
		Limit int
	}

	GetGroupsResponse struct {
		Groups     []GetGroupsResponseItem     `json:"groups"`
		Pagination usecases.PaginationResponse `json:"pagination"`
	}

	GetGroupsResponseItem struct {
		ID            uuid.UUID             `json:"id"`
		Name          string                `json:"name"`
		SelectionMode string                `json:"selection_mode"`
		BrandName     *string               `json:"brand_name,omitempty"`
		BrandIconPath *string               `json:"brand_icon_path,omitempty"`
		Address       *string               `json:"address,omitempty"`
		Members       []entities.UserEntity `json:"members"`
	}

	GetGroupResponse struct {
		ID            uuid.UUID           `json:"id"`
		Name          string              `json:"name"`
		SelectionMode string              `json:"selection_mode"`
		BrandName     string              `json:"brand_name,omitempty"`
		BrandIconPath *string             `json:"brand_icon_path,omitempty"`
		Address       *string             `json:"address,omitempty"`
		ContactPhone  *string             `json:"contact_phone,omitempty"`
		Members       []GroupMember       `json:"members"`
		BranchOptions []GroupBranchOption `json:"branch_options,omitempty"`
	}

	GroupMember struct {
		ID       int64   `json:"id"`
		Name     string  `json:"name"`
		PhotoURL *string `json:"photo_url"`
		Role     string  `json:"role"`
	}

	GroupBranchOption struct {
		ID            uuid.UUID `json:"id"`
		BrandName     string    `json:"brand_name"`
		BrandIconPath *string   `json:"brand_icon_path"`
		Members       []int64   `json:"members"`
	}
)
