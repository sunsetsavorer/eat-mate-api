package user

import (
	"github.com/sunsetsavorer/eat-mate-api/internal/dtos"
	"github.com/sunsetsavorer/eat-mate-api/internal/entities"
	"github.com/sunsetsavorer/eat-mate-api/internal/usecases"
)

type UpdateUserUseCase struct {
	logger         usecases.LoggerInterface
	userRepository UserRepositoryInterface
}

func NewUpdateUserUseCase(
	logger usecases.LoggerInterface,
	userRepository UserRepositoryInterface,
) *UpdateUserUseCase {

	return &UpdateUserUseCase{
		logger:         logger,
		userRepository: userRepository,
	}
}

func (uc UpdateUserUseCase) Exec(dto dtos.UpdateUserDTO) error {

	entity := entities.UserEntity{
		ID:       dto.GetUserID(),
		Name:     dto.GetName(),
		PhotoURL: dto.GetPhotoURL(),
	}

	err := uc.userRepository.Update(entity)
	if err != nil {
		uc.logger.Errorf("error while updating user: %v", err)
		return err
	}

	return nil
}
