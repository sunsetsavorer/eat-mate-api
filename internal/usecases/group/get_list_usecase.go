package group

import (
	"fmt"
	"math"

	"github.com/sunsetsavorer/eat-mate-api/internal/dtos"
	"github.com/sunsetsavorer/eat-mate-api/internal/entities"
	"github.com/sunsetsavorer/eat-mate-api/internal/exceptions"
	"github.com/sunsetsavorer/eat-mate-api/internal/usecases"
)

type GetGroupsUseCase struct {
	logger          usecases.LoggerInterface
	groupRepository GroupRepositoryInterface
}

func NewGetGroupsUseCase(
	logger usecases.LoggerInterface,
	groupRepository GroupRepositoryInterface,
) *GetGroupsUseCase {

	return &GetGroupsUseCase{
		logger:          logger,
		groupRepository: groupRepository,
	}
}

func (uc GetGroupsUseCase) Exec(dto dtos.GetGroupsDTO) (GetGroupsResponse, error) {

	page := dto.GetPage()

	if page <= 0 {
		page = 1
	}

	limit := dto.GetLimit()

	if limit <= 0 {
		limit = 4
	}

	filter := GroupsFilter{
		PaginationFilter: usecases.PaginationFilter{
			Page:  page,
			Limit: limit,
		},
	}

	groups, total, err := uc.groupRepository.GetList(filter)
	if err != nil {
		uc.logger.Errorf("failed to get groups: %v", err)
		return GetGroupsResponse{}, exceptions.NewBadRequestError(fmt.Errorf("failed to get groups"))
	}

	res := GetGroupsResponse{}

	for _, e := range groups {
		responseItem := GetGroupsResponseItem{
			ID:            e.GetID(),
			Name:          e.GetName(),
			SelectionMode: e.GetSelectionMode(),
		}

		groupMembers := e.GetMembers()
		responseItem.Members = make([]entities.UserEntity, len(groupMembers))

		for i, e := range groupMembers {
			responseItem.Members[i] = e.GetUser()
		}

		if e.GetSelectionMode() == DEFINED_SELECTION_MODE {
			brand := e.GetBranch().GetBrand()
			brandName := brand.GetName()

			responseItem.BrandName = &brandName
			responseItem.BrandIconPath = brand.GetIconPath()
			responseItem.Address = e.GetBranch().GetAddress()
		}

		res.Groups = append(res.Groups, responseItem)
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
