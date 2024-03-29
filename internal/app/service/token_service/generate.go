package tokenService

import (
	"time"

	"github.com/OddEer0/task-manager-server/config"
	appDto "github.com/OddEer0/task-manager-server/internal/app/app_dto"
	appErrors "github.com/OddEer0/task-manager-server/pkg/app_errors"
	"github.com/dgrijalva/jwt-go"
)

func (t *tokenService) Generate(data appDto.GenerateTokenServiceDto) (*JwtTokens, error) {
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
	accessTokenString, err := accessToken.SignedString([]byte(cfg.ApiKey))
	if err != nil {
		return nil, appErrors.InternalServerError("")
	}
	refreshTokenString, err := refreshToken.SignedString([]byte(cfg.ApiKey))
	if err != nil {
		return nil, appErrors.InternalServerError("")
	}
	return &JwtTokens{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
