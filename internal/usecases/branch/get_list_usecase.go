package branch

import (
	"fmt"
	"math"

	"github.com/sunsetsavorer/eat-mate-api/internal/dtos"
	"github.com/sunsetsavorer/eat-mate-api/internal/exceptions"
	"github.com/sunsetsavorer/eat-mate-api/internal/usecases"
)

type GetBranchesUseCase struct {
	logger           usecases.LoggerInterface
	branchRepository BranchRepositoryInterface
}

func NewGetBranchesUseCase(
	logger usecases.LoggerInterface,
	branchRepository BranchRepositoryInterface,
) *GetBranchesUseCase {

	return &GetBranchesUseCase{
		logger:           logger,
		branchRepository: branchRepository,
	}
}

func (uc GetBranchesUseCase) Exec(dto dtos.GetBranchesDTO) (GetBranchesResponse, error) {

	limit := dto.GetLimit()
	if limit <= 0 {
		limit = 4
	}

	page := dto.GetPage()
	if page <= 1 {
		page = 1
	}

	filter := BranchesFilter{
		Page:  page,
		Limit: limit,
		Query: dto.GetQuery(),
	}

	branches, total, err := uc.branchRepository.GetList(filter)
	if err != nil {
		uc.logger.Errorf("failed to get branches: %v", err)
		return GetBranchesResponse{}, exceptions.NewBadRequestError(fmt.Errorf("failed to get branches"))
	}

	res := GetBranchesResponse{}

	for _, e := range branches {
		responseItem := GetBranchesResponseItem{
			ID:           e.GetID(),
			Name:         e.Brand.GetName(),
			IconPath:     e.Brand.GetIconPath(),
			Address:      e.GetAddress(),
			ContactPhone: e.GetContactPhone(),
		}

		res.Branches = append(res.Branches, responseItem)
	}

	lastPage := int(math.Ceil(float64(total) / float64(dto.GetLimit())))

	nextPage := dto.GetPage() + 1
	if nextPage > lastPage {
		nextPage = 0
	}

	res.Pagination = usecases.PaginationResponse{
		Limit:       limit,
		CurrentPage: page,
		NextPage:    nextPage,
		LastPage:    lastPage,
	}

	return res, nil
}
