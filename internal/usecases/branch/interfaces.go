package branch

import "github.com/sunsetsavorer/eat-mate-api/internal/entities"

type (
	BranchRepositoryInterface interface {
		GetList(filter BranchesFilter) ([]entities.BranchEntity, int64, error)
	}
)
