package group

import (
	"fmt"

	"github.com/sunsetsavorer/eat-mate-api/internal/dtos"
	"github.com/sunsetsavorer/eat-mate-api/internal/entities"
	"github.com/sunsetsavorer/eat-mate-api/internal/exceptions"
	"github.com/sunsetsavorer/eat-mate-api/internal/usecases"
)

type JoinGroupUseCase struct {
	logger          usecases.LoggerInterface
	groupRepository GroupRepositoryInterface
	userRepository  UserRepositoryInterface
}

func NewJoinGroupUseCase(
	logger usecases.LoggerInterface,
	groupRepository GroupRepositoryInterface,
	userRepository UserRepositoryInterface,
) *JoinGroupUseCase {

	return &JoinGroupUseCase{
		logger:          logger,
		groupRepository: groupRepository,
		userRepository:  userRepository,
	}
}

func (uc JoinGroupUseCase) Exec(dto dtos.JoinGroupDTO) error {

	isInAnyGroup, err := uc.userRepository.IsInAnyGroup(dto.GetUserID())
	if err != nil {
		return exceptions.NewBadRequestError(fmt.Errorf("something went wrong. try later"))
	}

	if isInAnyGroup {
		return exceptions.NewBadRequestError(fmt.Errorf("user already in group"))
	}

	_, err = uc.groupRepository.GetByID(dto.GetGroupID())
	if err != nil {
		return err
	}

	memberEntity := entities.GroupMemberEntity{
		UserID:  dto.GetUserID(),
		GroupID: dto.GetGroupID(),
		Role:    "participant",
	}

	err = uc.groupRepository.AddMember(memberEntity)
	if err != nil {
		return exceptions.NewBadRequestError(fmt.Errorf("failed to join group"))
	}

	return nil
}
