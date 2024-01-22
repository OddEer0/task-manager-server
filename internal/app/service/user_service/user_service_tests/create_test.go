package user_service_tests

import (
	"context"
	"errors"
	"net/http"
	"testing"

	appDto "github.com/OddEer0/task-manager-server/internal/app/app_dto"
	"github.com/OddEer0/task-manager-server/internal/app/service/user_service"
	"github.com/OddEer0/task-manager-server/internal/common/constants"
	"github.com/OddEer0/task-manager-server/internal/common/constants/app_errors"
	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
	"github.com/OddEer0/task-manager-server/internal/infrastructure/storage/mock_repository"
	"github.com/OddEer0/task-manager-server/internal/presentation/mock"
	"github.com/stretchr/testify/assert"
)

func TestUserServiceCreate(t *testing.T) {
	memMockUser := mock.NewMockUser()
	testCases := []struct {
		name    string
		newUser appDto.RegistrationUseCaseDto
		isError bool
		errCode int
	}{
		{
			name: "Should create user aggregate",
			newUser: appDto.RegistrationUseCaseDto{
				Nick:      "new_user",
				Password:  "CorrectPassword123",
				Email:     "newuser@gmail.com",
				FirstName: "FakeMarlen",
				LastName:  "FakeLName",
			},
			isError: false,
		},
		{
			name: "Should conflict error by nick user",
			newUser: appDto.RegistrationUseCaseDto{
				Nick:      memMockUser.FullUser.Nick,
				Password:  "CorrectPassword123",
				Email:     "correct@gmail.com",
				FirstName: "FakeMarlen",
				LastName:  "FakeLName",
			},
			isError: true,
			errCode: http.StatusConflict,
		},
		{
			name: "Should conflict error by email user",
			newUser: appDto.RegistrationUseCaseDto{
				Nick:      "new_user_2",
				Password:  "CorrectPassword123",
				Email:     memMockUser.FullUser.Email,
				FirstName: "FakeMarlen",
				LastName:  "FakeLName",
			},
			isError: true,
			errCode: http.StatusConflict,
		},
	}

	userRepo := mock_repository.NewUserRepository()
	userServ := userService.New(userRepo)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := userServ.Create(context.Background(), tc.newUser)
			if !tc.isError {
				if assert.Equal(t, nil, err) {
					assert.NotEmpty(t, result.User.Id)
					assert.NotEmpty(t, result.User.ActivationLink)
					assert.Equal(t, constants.User, result.User.Role)
					assert.Equal(t, tc.newUser.Nick, result.User.Nick)
					assert.Equal(t, tc.newUser.FirstName, result.User.FirstName)
					assert.Equal(t, tc.newUser.LastName, result.User.LastName)
					assert.Equal(t, tc.newUser.Email, result.User.Email)
				}
			} else {
				assert.Equal(t, (*aggregate.UserAggregate)(nil), result)
				var appErr *appErrors.appErrors
				if errors.As(err, &appErr) {
					assert.Equal(t, tc.errCode, appErr.Code)
				}
			}
		})
	}
}
