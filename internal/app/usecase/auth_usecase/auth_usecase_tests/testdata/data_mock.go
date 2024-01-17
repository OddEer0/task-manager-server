package testdata

import (
	appDto "github.com/OddEer0/task-manager-server/internal/app/app_dto"
	"github.com/OddEer0/task-manager-server/internal/app/service/token_service"
	"github.com/OddEer0/task-manager-server/internal/app/usecase/auth_usecase"
	"github.com/OddEer0/task-manager-server/internal/common/constants"
	"github.com/google/uuid"
)

type AuthUseCaseRegistrationDataMock struct {
	CorrectRegInput1        appDto.RegistrationUseCaseDto
	CorrectRegInput1Result1 authUsecase.AuthResult
	CorrectRegInput1Result2 error
}

type AuthUseCaseDataMock struct {
	Registration *AuthUseCaseRegistrationDataMock
}

func NewAuthUseCaseDataMock() *AuthUseCaseDataMock {
	return &AuthUseCaseDataMock{
		Registration: &AuthUseCaseRegistrationDataMock{
			CorrectRegInput1: appDto.RegistrationUseCaseDto{
				Nick:      "NewEer0",
				Email:     "eer0@gmail.com",
				Password:  "Supperpupper123",
				FirstName: "Marlen",
				LastName:  "Karimov",
			},
			CorrectRegInput1Result1: authUsecase.AuthResult{
				User: &appDto.ResponseUserDto{
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
