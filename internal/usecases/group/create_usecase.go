package group

import (
	"fmt"

	"github.com/sunsetsavorer/eat-mate-api/internal/dtos"
	"github.com/sunsetsavorer/eat-mate-api/internal/entities"
	"github.com/sunsetsavorer/eat-mate-api/internal/exceptions"
	"github.com/sunsetsavorer/eat-mate-api/internal/usecases"
)

type CreateGroupUseCase struct {
	logger          usecases.LoggerInterface
	groupRepository GroupRepositoryInterface
	userRepository  UserRepositoryInterface
}

func NewCreateGroupUseCase(
	logger usecases.LoggerInterface,
	groupRepository GroupRepositoryInterface,
	userRepository UserRepositoryInterface,
) *CreateGroupUseCase {

	return &CreateGroupUseCase{
		logger:          logger,
		groupRepository: groupRepository,
		userRepository:  userRepository,
	}
}

func (uc CreateGroupUseCase) Exec(dto dtos.CreateGroupDTO) (CreateGroupResponse, error) {

	isAlreadyInGroup, err := uc.userRepository.IsInAnyGroup(dto.GetOwnerID())
	if err != nil {
		uc.logger.Errorf("failed to check user groups: %v", err)
		return CreateGroupResponse{}, exceptions.NewBadRequestError(fmt.Errorf("failed to check user groups"))
	}

	if isAlreadyInGroup {
		uc.logger.Errorf("user already in group: %v", dto.GetOwnerID())
		return CreateGroupResponse{}, exceptions.NewBadRequestError(fmt.Errorf("user already in group"))
	}

	entity := entities.GroupEntity{}

	entity.SetName(dto.GetName())
	entity.SetSelectionMode(dto.GetSelectionMode())
	entity.SetIsPublic(dto.GetIsPublic())
	entity.SetIsActive(true)

	groupMembers := []entities.GroupMemberEntity{
		{
			UserID: dto.GetOwnerID(),
			Role:   OWNER_ROLE,
		},
	}

	entity.SetMembers(groupMembers)

	if dto.GetSelectionMode() == DEFINED_SELECTION_MODE {

		if !dto.GetBranchID().Valid {
			return CreateGroupResponse{}, exceptions.NewBadRequestError(fmt.Errorf("not valid branch uuid"))
		}

		entity.SetBranchID(dto.GetBranchID())
	} else {

		dtoBranchOptions := dto.GetBranchOptions()

		if len(dtoBranchOptions) < BRANCH_OPTIONS_LIST_MINIMAL_SIZE {
			return CreateGroupResponse{}, exceptions.NewBadRequestError(fmt.Errorf("too few options"))
		}

		branchOptions := make([]entities.BranchEntity, len(dto.GetBranchOptions()))

		for i := range branchOptions {
			branchOptions[i].SetID(dtoBranchOptions[i])
		}

		entity.SetBranchOptions(branchOptions)
	}

	group, err := uc.groupRepository.Create(entity)
	if err != nil {
		uc.logger.Errorf("failed to create group: %v", err)
		return CreateGroupResponse{}, exceptions.NewBadRequestError(fmt.Errorf("failed to create group"))
	}

	return CreateGroupResponse{GroupID: group.ID}, nil
}
