package place

import (
	"github.com/google/uuid"
	"github.com/sunsetsavorer/eat-mate-api/internal/usecases"
)

type (
	PlaceBranchFilter struct {
		Page  int
		Limit int
		Query *string
	}

	GetPlaceBranchesResponse struct {
		PlaceBranches []GetPlaceBranchesResponseItem `json:"place_branches"`
		Pagination    usecases.PaginationResponse    `json:"pagination"`
	}

	GetPlaceBranchesResponseItem struct {
		ID           uuid.UUID `json:"id"`
		IconPath     *string   `json:"icon_path"`
		Name         string    `json:"name"`
		Address      *string   `json:"address"`
		ContactPhone *string   `json:"contact_phone"`
	}
)
