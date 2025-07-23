package place

import (
	"fmt"
	"math"

	"github.com/sunsetsavorer/eat-mate-api/internal/dtos"
	"github.com/sunsetsavorer/eat-mate-api/internal/exceptions"
	"github.com/sunsetsavorer/eat-mate-api/internal/usecases"
)

type GetPlaceBranchesUseCase struct {
	log                     usecases.LoggerInterface
	placeBranchesRepository PlaceBranchRepositoryInterface
}

func NewGetPlaceBranchesUseCase(
	log usecases.LoggerInterface,
	placeBranchRepository PlaceBranchRepositoryInterface,
) *GetPlaceBranchesUseCase {

	return &GetPlaceBranchesUseCase{
		log,
		placeBranchRepository,
	}
}

func (uc GetPlaceBranchesUseCase) Exec(dto dtos.GetPlaceBranchesDTO) (GetPlaceBranchesResponse, error) {

	limit := dto.GetLimit()
	if limit <= 0 {
		limit = 4
	}

	page := dto.GetPage()
	if page <= 1 {
		page = 1
	}

	filter := PlaceBranchFilter{
		Page:  page,
		Limit: limit,
		Query: dto.GetQuery(),
	}

	placeBranches, total, err := uc.placeBranchesRepository.GetList(filter)
	if err != nil {
		uc.log.Errorf("failed to get place branches: %v", err)
		return GetPlaceBranchesResponse{}, exceptions.NewBadRequestError(fmt.Errorf("failed to get place branches"))
	}

	res := GetPlaceBranchesResponse{}

	for _, e := range placeBranches {
		responseItem := GetPlaceBranchesResponseItem{
			ID:           e.GetID(),
			IconPath:     e.Place.GetIconPath(),
			Name:         e.Place.GetName(),
			Address:      e.GetAddress(),
			ContactPhone: e.GetContactPhone(),
		}

		res.PlaceBranches = append(res.PlaceBranches, responseItem)
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
