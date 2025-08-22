package group

import (
	"fmt"

	"github.com/sunsetsavorer/eat-mate-api/internal/dtos"
	"github.com/sunsetsavorer/eat-mate-api/internal/entities"
	"github.com/sunsetsavorer/eat-mate-api/internal/exceptions"
	"github.com/sunsetsavorer/eat-mate-api/internal/usecases"
)

type RevokeVoteUseCase struct {
	logger          usecases.LoggerInterface
	groupRepository GroupRepositoryInterface
}

func NewRevokeVoteUseCase(
	logger usecases.LoggerInterface,
	groupRepositry GroupRepositoryInterface,
) *RevokeVoteUseCase {

	return &RevokeVoteUseCase{
		logger:          logger,
		groupRepository: groupRepositry,
	}
}

func (uc RevokeVoteUseCase) Exec(dto dtos.RevokeVoteDTO) error {

	group, err := uc.groupRepository.GetByID(dto.GetGroupID())
	if err != nil {
		uc.logger.Errorf("failed to get group: %v", err)
		return exceptions.NewNotFoundError(err)
	}

	if group.GetSelectionMode() != VOTING_SELECTION_MODE {
		return exceptions.NewBadRequestError(fmt.Errorf("you can't vote in this group"))
	}

	_, err = uc.groupRepository.GetMemberByID(dto.GetUserID(), dto.GetGroupID())
	if err != nil {
		return exceptions.NewBadRequestError(fmt.Errorf("you aren't in this group"))
	}

	voteEntity := entities.VoteEntity{
		GroupID: dto.GetGroupID(),
		UserID:  dto.GetUserID(),
	}

	err = uc.groupRepository.DeleteVote(voteEntity)
	if err != nil {
		uc.logger.Errorf("failed to delete vote: %v", err)
		return exceptions.NewBadRequestError(fmt.Errorf("failed to delete vote"))
	}

	return nil
}
