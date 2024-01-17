package auth_usecase_tests

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/OddEer0/task-manager-server/config"
	appDto "github.com/OddEer0/task-manager-server/internal/app/app_dto"
	authUsecase "github.com/OddEer0/task-manager-server/internal/app/usecase/auth_usecase"
	"github.com/OddEer0/task-manager-server/internal/app/usecase/auth_usecase/auth_usecase_tests/testdata"
	"github.com/OddEer0/task-manager-server/internal/tests"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

// TestAuthUseCaseRegistration - Можно добавить ещё тестов
func TestAuthUseCaseRegistration(t *testing.T) {
	mockData := testdata.NewAuthUseCaseDataMock().Registration

	testCases := []struct {
		name           string
		inputData      appDto.RegistrationUseCaseDto
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
	cfg, err := config.NewConfig()
	if err != nil {
		t.Fatalf("cfg not load: %v", err)
	}

	isEqualUser := func(a *appDto.ResponseUserDto, b *appDto.ResponseUserDto) bool {
		aType := reflect.TypeOf(*a)
		aValue := reflect.ValueOf(*a)
		bValue := reflect.ValueOf(*a)
		for i := 0; i < aType.NumField(); i++ {
			aField := aType.Field(i)
			aVal := aValue.Field(i).Interface()
			bVal := bValue.Field(i).Interface()
			if aField.Name == "Id" {
				continue
			}
			if aVal != bVal {
				return false
			}
		}
		return true
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := authUseCase.Registration(context.Background(), testCase.inputData)
			refreshToken, err := jwt.Parse(result.Tokens.RefreshToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(cfg.ApiKey), nil
			})
			assert.Equal(t, nil, err)
			assert.True(t, refreshToken.Valid)
			assert.Equal(t, testCase.expectedError, err)
			assert.NotEmpty(t, result.Tokens.RefreshToken)
			assert.NotEmpty(t, result.Tokens.AccessToken)
			assert.True(t, isEqualUser(testCase.expectedResult.User, result.User))
		})
	}
}
