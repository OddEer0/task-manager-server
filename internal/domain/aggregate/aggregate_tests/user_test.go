package aggregate_tests

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
	"github.com/OddEer0/task-manager-server/internal/domain/models"
	"github.com/OddEer0/task-manager-server/internal/presentation/dto"
	"github.com/OddEer0/task-manager-server/internal/presentation/mock"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type CustomClaims struct {
	JwtUserData dto.GenerateTokenDto
	jwt.StandardClaims
}

func TestNewUserAggregate(t *testing.T) {
	memMockUser := mock.NewMockUser()

	testCases := []struct {
		name          string
		userModel     models.User
		expectedValue *aggregate.UserAggregate
		isError       bool
		errFields     []string
	}{
		{
			name:          "Should correct create aggregate",
			userModel:     memMockUser.User,
			expectedValue: &aggregate.UserAggregate{User: memMockUser.User},
			isError:       false,
		},
		{
			name:          "Should required errors",
			userModel:     memMockUser.RequiredErrUser,
			expectedValue: nil,
			isError:       true,
			errFields:     []string{"Id", "Nick", "Value", "Email", "FirstName", "LastName", "Role", "ActivationLink", "CreatedAt", "UpdatedAt"},
		},
		{
			name:          "Should min value errors",
			userModel:     memMockUser.MinErrUser,
			expectedValue: nil,
			isError:       true,
			errFields:     []string{"Nick", "FirstName", "LastName", "SubTitle", "BanReason"},
		},
		{
			name:          "Should max value errors",
			userModel:     memMockUser.MaxErrUser,
			expectedValue: nil,
			isError:       true,
			errFields:     []string{"Nick", "FirstName", "LastName", "SubTitle", "BanReason"},
		},
		{
			name:          "Should invalid id error",
			userModel:     memMockUser.IdErrUser,
			expectedValue: nil,
			isError:       true,
			errFields:     []string{"Id"},
		},
		{
			name:          "Should invalid email error",
			userModel:     memMockUser.EmailErrUser,
			expectedValue: nil,
			isError:       true,
			errFields:     []string{"Email"},
		},
		{
			name:          "Should invalid link avatar activation-link error",
			userModel:     memMockUser.LinkErrUser,
			expectedValue: nil,
			isError:       true,
			errFields:     []string{"Avatar", "ActivationLink"},
		},
		{
			name:          "Should invalid role error",
			userModel:     memMockUser.RoleErrUser,
			expectedValue: nil,
			isError:       true,
			errFields:     []string{"Role"},
		},
		{
			name:          "Should invalid createdAt updatedAt time error",
			userModel:     memMockUser.DateCreatedAndUpdatedAtErrUser,
			expectedValue: nil,
			isError:       true,
			errFields:     []string{"CreatedAt", "UpdatedAt"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := aggregate.NewUserAggregate(tc.userModel)

			if tc.isError {
				assert.Error(t, err)
				var validationError validator.ValidationErrors
				ok := errors.As(err, &validationError)
				assert.True(t, ok)
				for i, e := range validationError {
					assert.Equal(t, tc.errFields[i], e.Field())
				}
			} else {
				assert.Equal(t, nil, err)
				assert.Equal(t, tc.expectedValue, result)
			}
		})
	}
}

func TestUserAggregateSetToken(t *testing.T) {
	memMockUser := mock.NewMockUser()
	accessDuration, _ := time.ParseDuration("15m")
	accessClaims := CustomClaims{
		JwtUserData:    dto.GenerateTokenDto{Id: memMockUser.User.Id, Roles: memMockUser.User.Role},
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(accessDuration).Unix()},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	tokenStr, err := accessToken.SignedString([]byte("super-secret"))
	if err != nil {
		return
	}
	invalidToken := "invalid-error"

	testCases := []struct {
		name    string
		token   string
		isError bool
	}{
		{
			name:    "Should set token",
			token:   tokenStr,
			isError: false,
		},
		{
			name:    "Should error required token",
			token:   "",
			isError: true,
		},
		{
			name:    "Should error invalid token",
			token:   invalidToken,
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userAggregate, _ := aggregate.NewUserAggregate(memMockUser.User)
			err := userAggregate.SetToken(tc.token)

			fmt.Println("token:", tc.token)

			if tc.isError {
				assert.Error(t, err)
				var validationError validator.ValidationErrors
				ok := errors.As(err, &validationError)
				assert.True(t, ok)
				for _, e := range validationError {
					assert.Equal(t, "Value", e.Field())
				}
			} else {
				assert.Equal(t, nil, err)
			}
		})
	}
}
