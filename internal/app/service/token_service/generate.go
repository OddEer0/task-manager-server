package tokenService

import (
	"github.com/OddEer0/task-manager-server/config"
	"github.com/OddEer0/task-manager-server/internal/presentation/dto"
	appErrors "github.com/OddEer0/task-manager-server/pkg/app_errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func (t *tokenService) Generate(data dto.GenerateTokenDto) (*JwtTokens, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, appErrors.InternalServerError("")
	}
	accessDuration, _ := time.ParseDuration(cfg.AccessTokenTime)
	refreshDuration, _ := time.ParseDuration(cfg.RefreshTokenTime)
	accessClaims := CustomClaims{
		JwtUserData:    data,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(accessDuration).Unix()},
	}
	refreshClaims := CustomClaims{
		JwtUserData:    data,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(refreshDuration).Unix()},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	return &JwtTokens{
		AccessToken:  accessToken.Signature,
		RefreshToken: refreshToken.Signature,
	}, nil
}
