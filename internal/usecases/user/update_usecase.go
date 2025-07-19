package user

import (
	"github.com/sunsetsavorer/eat-mate-api/internal/dtos"
	"github.com/sunsetsavorer/eat-mate-api/internal/entities"
	"github.com/sunsetsavorer/eat-mate-api/internal/usecases"
)

type UpdateUserUseCase struct {
	log            usecases.LoggerInterface
	userRepository UserRepositoryInterface
}

func NewUpdateUserUseCase(
	log usecases.LoggerInterface,
	userRepository UserRepositoryInterface,
) *UpdateUserUseCase {

	return &UpdateUserUseCase{
		log:            log,
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
		uc.log.Errorf("error while updating user: %v", err)
		return err
	}

	return nil
}
