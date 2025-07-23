package place

import "github.com/sunsetsavorer/eat-mate-api/internal/entities"

type (
	PlaceBranchRepositoryInterface interface {
		GetList(filter PlaceBranchFilter) ([]entities.PlaceBranchEntity, int64, error)
	}
)
