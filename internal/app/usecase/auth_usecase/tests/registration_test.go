package tests

import (
	"context"
	"fmt"
	tokenService "github.com/OddEer0/task-manager-server/internal/app/service/token_service"
	userService "github.com/OddEer0/task-manager-server/internal/app/service/user_service"
	"github.com/OddEer0/task-manager-server/internal/app/usecase/auth_usecase"
	"github.com/OddEer0/task-manager-server/internal/app/usecase/auth_usecase/tests/mocks"
	"github.com/OddEer0/task-manager-server/internal/infrastructure/storage/mock_repository"
	"github.com/OddEer0/task-manager-server/internal/presentation/dto"
	"testing"
)

func TestAuthUseCaseRegistration(t *testing.T) {
	mockData := mocks.NewAuthUseCaseDataMock().Registration

	testCases := []struct {
		name           string
		inputData      dto.RegistrationInputDto
		expectedResult authUsecase.AuthResult
		expectedError  error
		isError        bool
	}{
		{
			name:           "Should register user",
			inputData:      mockData.CorrectRegInput1,
			expectedResult: mockData.CorrectRegInput1Result1,
			expectedError:  mockData.CorrectRegInput1Result2,
			isError:        false,
		},
	}
	userRepo := mock_repository.NewUserRepository()
	tokenRepo := mock_repository.NewTokenRepository()
	userServ := userService.NewUserService(userRepo)
	tokenServ := tokenService.NewTokenService(tokenRepo)
	authUseCase := authUsecase.NewAuthUseCase(userServ, tokenServ, userRepo)

	//isEqualUser := func(a *dto.ResponseUserDto, b *dto.ResponseUserDto) bool {
	//	aType := reflect.TypeOf(*a)
	//	aValue := reflect.ValueOf(*a)
	//	bValue := reflect.ValueOf(*a)
	//	for i := 0; i < aType.NumField(); i++ {
	//		aField := aType.Field(i)
	//		aVal := aValue.Field(i).Interface()
	//		bVal := bValue.Field(i).Interface()
	//		if aField.Name == "Id" {
	//			continue
	//		}
	//		if aVal != bVal {
	//			return false
	//		}
	//	}
	//	return true
	//}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := authUseCase.Registration(context.Background(), testCase.inputData)
			fmt.Printf("%v %v", result, err)

			//assert.Equal(t, err, testCase.expectedError)
			//assert.NotEmpty(t, result.Tokens.RefreshToken)
			//assert.NotEmpty(t, result.Tokens.AccessToken)
			//assert.Nil(t, uuid.Validate(result.User.Id))
			//assert.True(t, isEqualUser(testCase.expectedResult.User, result.User))
		})
	}
}
