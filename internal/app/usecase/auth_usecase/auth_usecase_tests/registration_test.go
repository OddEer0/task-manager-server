package auth_usecase_tests

import (
	"context"
	"testing"

	"github.com/OddEer0/task-manager-server/internal/app/usecase/auth_usecase"
	"github.com/OddEer0/task-manager-server/internal/app/usecase/auth_usecase/auth_usecase_tests/mocks"
	"github.com/OddEer0/task-manager-server/internal/presentation/dto"
	"github.com/OddEer0/task-manager-server/internal/tests"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
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

	authUseCase := tests.NewUseCases().AuthUseCase

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
			assert.Equal(t, err, testCase.expectedError)
			assert.NotEmpty(t, result.Tokens.RefreshToken)
			assert.NotEmpty(t, result.Tokens.AccessToken)
			assert.Nil(t, uuid.Validate(result.User.Id))
			//assert.True(t, isEqualUser(testCase.expectedResult.User, result.User))
		})
	}
}
