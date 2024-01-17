package token_service_tests

import (
	"fmt"
	"testing"

	"github.com/OddEer0/task-manager-server/config"
	appDto "github.com/OddEer0/task-manager-server/internal/app/app_dto"
	tokenService "github.com/OddEer0/task-manager-server/internal/app/service/token_service"
	"github.com/OddEer0/task-manager-server/internal/infrastructure/storage/mock_repository"
	"github.com/OddEer0/task-manager-server/internal/presentation/mock"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestTokenServiceGenerate(t *testing.T) {
	memMockUser := mock.NewMockUser()
	cfg, err := config.NewConfig()
	if err != nil {
		t.Fatalf("config load error: %v", err)
	}
	testCases := []struct {
		name string
		data appDto.GenerateTokenServiceDto
	}{
		{
			name: "Should generate token",
			data: appDto.GenerateTokenServiceDto{Id: memMockUser.User.Id, Role: memMockUser.User.Role},
		},
	}

	tokenServ := tokenService.New(mock_repository.NewTokenRepository())

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tokens, _ := tokenServ.Generate(tc.data)
			assert.IsType(t, "", tokens.AccessToken)
			assert.IsType(t, "", tokens.RefreshToken)
			accessToken, err := jwt.Parse(tokens.AccessToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(cfg.ApiKey), nil
			})
			assert.Equal(t, nil, err)
			assert.True(t, accessToken.Valid)
			refreshToken, err := jwt.Parse(tokens.RefreshToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(cfg.ApiKey), nil
			})
			assert.Equal(t, nil, err)
			assert.True(t, refreshToken.Valid)
		})
	}
}
