package usecases

import (
	"github.com/sunsetsavorer/eat-mate-api/internal/domains/user/dtos"
	"github.com/sunsetsavorer/eat-mate-api/internal/domains/user/entities"
	"github.com/sunsetsavorer/eat-mate-api/internal/domains/user/interfaces"
	"github.com/sunsetsavorer/eat-mate-api/pkg/jwt"
)

type AuthorizeUseCase struct {
	UserRepository interfaces.UserRepositoryInterface
}

func NewAuthorizeUseCase(
	userRepository interfaces.UserRepositoryInterface,
) *AuthorizeUseCase {

	return &AuthorizeUseCase{
		userRepository,
	}
}

func (uc AuthorizeUseCase) Exec(dto dtos.AuthorizeDTO) (string, error) {

	user, err := uc.UserRepository.GetByID(dto.GetUserID())
	if err != nil {
		user = entities.UserEntity{
			ID:       dto.GetUserID(),
			Name:     dto.GetUserName(),
			PhotoURL: dto.GetUserPhotoURL(),
		}

		err := uc.UserRepository.Create(user)
		if err != nil {
			return "", err
		}
	}

	token, err := jwt.GenerateToken(
		user.GetID(),
		dto.GetTokenSecret(),
		dto.GetTokenLifetime(),
	)
	if err != nil {
		return "", err
	}

	return token.Value, nil
}
