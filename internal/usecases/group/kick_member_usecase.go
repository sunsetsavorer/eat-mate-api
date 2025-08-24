package group

import (
	"fmt"

	"github.com/sunsetsavorer/eat-mate-api/internal/dtos"
	"github.com/sunsetsavorer/eat-mate-api/internal/exceptions"
	"github.com/sunsetsavorer/eat-mate-api/internal/usecases"
)

type KickMemberUseCase struct {
	logger          usecases.LoggerInterface
	groupRepository GroupRepositoryInterface
}

func NewKickMemberUseCase(
	logger usecases.LoggerInterface,
	groupRepository GroupRepositoryInterface,
) *KickMemberUseCase {

	return &KickMemberUseCase{
		logger:          logger,
		groupRepository: groupRepository,
	}
}

func (uc KickMemberUseCase) Exec(dto dtos.KickMemberDTO) error {

	_, err := uc.groupRepository.GetByID(dto.GetGroupID())
	if err != nil {
		uc.logger.Errorf("failed to get group: %v", err)
		return exceptions.NewNotFoundError(fmt.Errorf("group with specified id wasn't found"))
	}

	user, err := uc.groupRepository.GetMemberByID(dto.GetUserID(), dto.GetGroupID())
	if err != nil {
		uc.logger.Errorf("failed to get member: %v", err)
		return exceptions.NewBadRequestError(fmt.Errorf("failed to get member"))
	}

	if user.GetRole() != OWNER_ROLE {
		return exceptions.NewForbiddenError(fmt.Errorf("you aren't owner of this group"))
	}

	if dto.GetMemberID() == dto.GetUserID() {
		return exceptions.NewBadRequestError(fmt.Errorf("you can't kick yourself"))
	}

	member, err := uc.groupRepository.GetMemberByID(dto.GetMemberID(), dto.GetGroupID())
	if err != nil {
		uc.logger.Errorf("failed to get member: %v", err)
		return exceptions.NewBadRequestError(fmt.Errorf("failed to get member"))
	}

	err = uc.groupRepository.RemoveMember(member)
	if err != nil {
		uc.logger.Errorf("failed to kick member: %v", err)
		return exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))
	}

	return nil
}
