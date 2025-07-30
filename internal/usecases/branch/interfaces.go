package branch

import "github.com/sunsetsavorer/eat-mate-api/internal/entities"

type (
	BranchRepositoryInterface interface {
		GetList(filter BranchFilter) ([]entities.BranchEntity, int64, error)
	}
)
