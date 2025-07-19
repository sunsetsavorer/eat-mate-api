package user

import (
	"fmt"

	"github.com/sunsetsavorer/eat-mate-api/internal/dtos"
	"github.com/sunsetsavorer/eat-mate-api/internal/entities"
	"github.com/sunsetsavorer/eat-mate-api/internal/exceptions"
	"github.com/sunsetsavorer/eat-mate-api/internal/usecases"
)

type AuthorizeUseCase struct {
	log            usecases.LoggerInterface
	userRepository UserRepositoryInterface
	jwtService     usecases.JWTServiceInterface
}

func NewAuthorizeUseCase(
	log usecases.LoggerInterface,
	userRepository UserRepositoryInterface,
	jwtService usecases.JWTServiceInterface,
) *AuthorizeUseCase {

	return &AuthorizeUseCase{
		log,
		userRepository,
		jwtService,
	}
}

func (uc AuthorizeUseCase) Exec(dto dtos.AuthorizeDTO) (TokenResponse, error) {

	user, err := uc.userRepository.GetByID(dto.GetTelegramID())
	if err != nil {
		user = entities.UserEntity{
			ID:       dto.GetTelegramID(),
			Name:     dto.GetName(),
			PhotoURL: dto.GetPhotoURL(),
		}

		err := uc.userRepository.Create(user)
		if err != nil {
			uc.log.Errorf("failed to create user: %v", err)
			return TokenResponse{}, exceptions.NewBadRequestError(fmt.Errorf("failed to create user"))
		}
	}

	token, err := uc.jwtService.GenerateTokenByUserID(user.GetID())
	if err != nil {
		uc.log.Errorf("failed to generate authorization token: %v", err)
		return TokenResponse{}, exceptions.NewBadRequestError(fmt.Errorf("failed to generate authorization token"))
	}

	return TokenResponse{token.Value}, nil
}
