package group

import (
	"github.com/google/uuid"
	"github.com/sunsetsavorer/eat-mate-api/internal/entities"
	"github.com/sunsetsavorer/eat-mate-api/internal/usecases"
)

type GetGroupUseCase struct {
	logger          usecases.LoggerInterface
	groupRepository GroupRepositoryInterface
}

func NewGetGroupUseCase(
	logger usecases.LoggerInterface,
	groupRepository GroupRepositoryInterface,
) *GetGroupUseCase {

	return &GetGroupUseCase{
		logger:          logger,
		groupRepository: groupRepository,
	}
}

func (uc GetGroupUseCase) Exec(ID uuid.UUID) (GetGroupResponse, error) {

	groupEntity, err := uc.groupRepository.GetByID(ID)
	if err != nil {
		uc.logger.Errorf("failed to get group: %v", err)
		return GetGroupResponse{}, err
	}

	var members []GroupMember

	for _, member := range groupEntity.GetMembers() {

		user := member.GetUser()

		members = append(
			members,
			GroupMember{
				ID:       user.GetID(),
				Name:     user.GetName(),
				PhotoURL: user.GetPhotoURL(),
				Role:     member.GetRole(),
			},
		)
	}

	selectionMode := groupEntity.GetSelectionMode()

	result := GetGroupResponse{
		ID:            groupEntity.GetID(),
		Name:          groupEntity.GetName(),
		SelectionMode: selectionMode,
		Members:       members,
	}

	switch selectionMode {
	case DEFINED_SELECTION_MODE:
		branch := groupEntity.GetBranch()
		brand := branch.GetBrand()

		result.BrandName = brand.GetName()
		result.BrandIconPath = brand.GetIconPath()
		result.Address = branch.GetAddress()
		result.ContactPhone = branch.GetContactPhone()

	case VOTING_SELECTION_MODE:
		options := uc.getBranchOptions(groupEntity)
		votes := groupEntity.GetVotes()

		for i := range options {
			for _, vote := range votes {

				if options[i].ID.String() == vote.GetBranchID().String() {
					options[i].Members = append(options[i].Members, vote.GetUserID())
				}
			}
		}

		result.BranchOptions = options

	case RANDOM_SELECTION_MODE:
		result.BranchOptions = uc.getBranchOptions(groupEntity)
	}

	return result, nil
}

func (uc GetGroupUseCase) getBranchOptions(groupEntity entities.GroupEntity) []GroupBranchOption {

	options := groupEntity.GetBranchOptions()

	branchOptions := make([]GroupBranchOption, 0, len(options))

	for _, branch := range options {

		brand := branch.GetBrand()

		branchOption := GroupBranchOption{
			ID:            branch.GetID(),
			BrandName:     brand.GetName(),
			BrandIconPath: brand.GetIconPath(),
		}

		branchOptions = append(branchOptions, branchOption)
	}

	return branchOptions
}
