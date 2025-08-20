package group

import (
	"fmt"

	"github.com/sunsetsavorer/eat-mate-api/internal/dtos"
	"github.com/sunsetsavorer/eat-mate-api/internal/exceptions"
	"github.com/sunsetsavorer/eat-mate-api/internal/usecases"
)

type LeaveGroupUseCase struct {
	logger          usecases.LoggerInterface
	groupRepository GroupRepositoryInterface
	userRepository  UserRepositoryInterface
}

func NewLeaveGroupUseCase(
	logger usecases.LoggerInterface,
	groupRepository GroupRepositoryInterface,
	userRepository UserRepositoryInterface,
) *LeaveGroupUseCase {

	return &LeaveGroupUseCase{
		logger:          logger,
		groupRepository: groupRepository,
		userRepository:  userRepository,
	}
}

func (uc LeaveGroupUseCase) Exec(dto dtos.LeaveGroupDTO) error {

	_, err := uc.groupRepository.GetByID(dto.GetGroupID())
	if err != nil {
		return err
	}

	member, err := uc.groupRepository.GetMemberByID(dto.GetUserID(), dto.GetGroupID())
	if err != nil {
		uc.logger.Errorf("get error while fetching group member: %v", err)
		return exceptions.NewBadRequestError(fmt.Errorf("you can't leave group you're not in"))
	}

	isGroupOwner := member.GetRole() == OWNER_ROLE

	if isGroupOwner {
		err = uc.groupRepository.DeactivateByID(dto.GetGroupID())
	} else {
		err = uc.groupRepository.RemoveMember(member)
	}

	if err != nil {
		uc.logger.Errorf("get error while removing member: %v", err)
		return exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))
	}

	return nil
}
