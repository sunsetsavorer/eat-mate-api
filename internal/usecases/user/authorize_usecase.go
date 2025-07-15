package user

import (
	"fmt"

	"github.com/sunsetsavorer/eat-mate-api/internal/dtos"
	"github.com/sunsetsavorer/eat-mate-api/internal/entities"
	"github.com/sunsetsavorer/eat-mate-api/internal/exceptions"
	"github.com/sunsetsavorer/eat-mate-api/internal/usecases"
	"github.com/sunsetsavorer/eat-mate-api/pkg/jwt"
)

type AuthorizeUseCase struct {
	log            usecases.LoggerInterface
	userRepository UserRepositoryInterface
}

func NewAuthorizeUseCase(
	log usecases.LoggerInterface,
	userRepository UserRepositoryInterface,
) *AuthorizeUseCase {

	return &AuthorizeUseCase{
		log,
		userRepository,
	}
}

func (uc AuthorizeUseCase) Exec(dto dtos.AuthorizeDTO) (TokenResponse, error) {

	user, err := uc.userRepository.GetByID(dto.GetUserID())
	if err != nil {
		user = entities.UserEntity{
			ID:       dto.GetUserID(),
			Name:     dto.GetUserName(),
			PhotoURL: dto.GetUserPhotoURL(),
		}

		err := uc.userRepository.Create(user)
		if err != nil {
			uc.log.Errorf("failed to create user: %v", err)
			return TokenResponse{}, exceptions.NewBadRequestError(fmt.Errorf("failed to create user"))
		}
	}

	token, err := jwt.GenerateToken(
		user.GetID(),
		dto.GetTokenSecret(),
		dto.GetTokenLifetime(),
	)
	if err != nil {
		uc.log.Errorf("failed to generate authorization token: %v", err)
		return TokenResponse{}, exceptions.NewBadRequestError(fmt.Errorf("failed to generate authorization token"))
	}

	return TokenResponse{token.Value}, nil
}
