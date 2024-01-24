package tokenService

import (
	"github.com/OddEer0/task-manager-server/config"
	appErrors "github.com/OddEer0/task-manager-server/internal/common/lib/app_errors"
	"github.com/dgrijalva/jwt-go"
)

func (t *tokenService) ValidateRefreshToken(refreshToken string) (*JwtUserData, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, appErrors.InternalServerError("", "target: TokenService, method: ValidateRefreshToken .", "error: ", err.Error())
	}
	token, err := jwt.ParseWithClaims(refreshToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.ApiKey), nil
	})
	if err != nil {
		return nil, appErrors.InternalServerError("", "target: TokenService, method: ValidateRefreshToken .", "jwt parse error: ", err.Error())
	}
	claim := token.Claims.(*CustomClaims)
	return &claim.JwtUserData, nil
}
