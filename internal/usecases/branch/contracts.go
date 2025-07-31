package branch

import (
	"github.com/google/uuid"
	"github.com/sunsetsavorer/eat-mate-api/internal/usecases"
)

type (
	BranchesFilter struct {
		usecases.PaginationFilter
		Query *string
	}

	GetBranchesResponse struct {
		Branches   []GetBranchesResponseItem   `json:"branches"`
		Pagination usecases.PaginationResponse `json:"pagination"`
	}

	GetBranchesResponseItem struct {
		ID           uuid.UUID `json:"id"`
		Name         string    `json:"name"`
		IconPath     *string   `json:"icon_path"`
		Address      *string   `json:"address"`
		ContactPhone *string   `json:"contact_phone"`
	}
)
