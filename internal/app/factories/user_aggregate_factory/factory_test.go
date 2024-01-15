package userAggregateFactory

import (
	"testing"
	"time"

	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
	"github.com/OddEer0/task-manager-server/internal/domain/models"
	"github.com/OddEer0/task-manager-server/internal/presentation/mock"
	"github.com/stretchr/testify/assert"
)

func convertUserModelToCreateUserAggregateData(user models.User) CreateUserAggregateData {
	var (
		birthday time.Time
		pass     string
	)

	if len(user.Password.Value) > 35 {
		pass = user.Password.Value[0:33]
	} else {
		pass = user.Password.Value
	}

	if user.Birthday != nil {
		birthday = user.Birthday.Value
	}

	return CreateUserAggregateData{
		user.Id,
		user.Nick,
		pass,
		user.Email,
		user.FirstName,
		user.LastName,
		user.SubTitle,
		user.Avatar,
		birthday,
		user.Role,
		user.ActivationLink,
		user.IsActivate,
		user.IsBanned,
		user.BanReason,
		user.CreatedAt,
		user.UpdatedAt,
	}
}

func TestCreateUserAggregate(t *testing.T) {
	memMockUser := mock.NewMockUser()
	testCases := []struct {
		name      string
		userModel models.User
		isError   bool
	}{
		{
			name:      "Should creat new user aggregate",
			userModel: memMockUser.User,
			isError:   false,
		},
		{
			name:      "Should required errors",
			userModel: memMockUser.RequiredErrUser,
			isError:   true,
		},
		{
			name:      "Should min value errors",
			userModel: memMockUser.MinErrUser,
			isError:   true,
		},
		{
			name:      "Should max value errors",
			userModel: memMockUser.MaxErrUser,
			isError:   true,
		},
		{
			name:      "Should invalid id error",
			userModel: memMockUser.IdErrUser,
			isError:   true,
		},
		{
			name:      "Should invalid email error",
			userModel: memMockUser.EmailErrUser,
			isError:   true,
		},
		{
			name:      "Should invalid link avatar activation-link error",
			userModel: memMockUser.LinkErrUser,
			isError:   true,
		},
		{
			name:      "Should invalid role error",
			userModel: memMockUser.RoleErrUser,
			isError:   true,
		},
		{
			name:      "Should invalid createdAt updatedAt time error",
			userModel: memMockUser.DateCreatedAndUpdatedAtErrUser,
			isError:   true,
		},
	}

	factory := UserAggregateFactory{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data := convertUserModelToCreateUserAggregateData(tc.userModel)
			if tc.isError {
				userAggregate, err := factory.CreateUserAggregate(data)
				assert.Equal(t, (*aggregate.UserAggregate)(nil), userAggregate)
				assert.Error(t, err)
			} else {
				userAggregate, err := factory.CreateUserAggregate(data)
				assert.Nil(t, err)
				newUserAggregate, err := aggregate.NewUserAggregate(tc.userModel)
				assert.Nil(t, err)
				newUserAggregate.User.Password.Value = userAggregate.User.Password.Value
				assert.Equal(t, newUserAggregate, userAggregate)
			}
		})
	}
}
