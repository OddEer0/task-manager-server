package mocks

import (
	"github.com/OddEer0/task-manager-server/internal/app/service/token_service"
	"github.com/OddEer0/task-manager-server/internal/app/usecase/auth_usecase"
	"github.com/OddEer0/task-manager-server/internal/common/constants"
	"github.com/OddEer0/task-manager-server/internal/presentation/dto"
	"github.com/google/uuid"
)

type AuthUseCaseRegistrationDataMock struct {
	CorrectRegInput1        dto.RegistrationInputDto
	CorrectRegInput1Result1 authUsecase.AuthResult
	CorrectRegInput1Result2 error
}

type AuthUseCaseDataMock struct {
	Registration *AuthUseCaseRegistrationDataMock
}

func NewAuthUseCaseDataMock() *AuthUseCaseDataMock {
	return &AuthUseCaseDataMock{
		Registration: &AuthUseCaseRegistrationDataMock{
			CorrectRegInput1: dto.RegistrationInputDto{
				Nick:      "NewEer0",
				Email:     "eer0@gmail.com",
				Password:  "Supperpupper123",
				FirstName: "Marlen",
				LastName:  "Karimov",
			},
			CorrectRegInput1Result1: authUsecase.AuthResult{
				User: &dto.ResponseUserDto{
					Id:        uuid.New().String(),
					Nick:      "NewEer0",
					Email:     "Lolkek@gmail.com",
					FirstName: "Marlen",
					LastName:  "Karimov",
					Role:      constants.User,
				},
				Tokens: tokenService.JwtTokens{AccessToken: "dsads", RefreshToken: "dsadsad"},
			},
			CorrectRegInput1Result2: nil,
		},
	}
}
