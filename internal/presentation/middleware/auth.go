package middleware

import (
	"context"
	"net/http"

	"github.com/OddEer0/task-manager-server/config"
	tokenService "github.com/OddEer0/task-manager-server/internal/app/service/token_service"
	appErrors "github.com/OddEer0/task-manager-server/internal/common/lib/app_errors"
	httpUtils "github.com/OddEer0/task-manager-server/pkg/http_utils"
	"github.com/dgrijalva/jwt-go"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		accessToken, err := req.Cookie("accessToken")
		if err != nil {
			unauthorized(res)
			return
		}
		cfg, err := config.NewConfig()
		if err != nil {
			httpUtils.SendJson(res, http.StatusInternalServerError, appErrors.ResponseError{Code: http.StatusInternalServerError, Message: appErrors.DefaultInternalServerErrorMessage})
			return
		}
		token, err := jwt.ParseWithClaims(accessToken.Value, &tokenService.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.ApiKey), nil
		})
		if err != nil {
			unauthorized(res)
			return
		}
		ctx := context.WithValue(req.Context(), "user", token.Claims.(*tokenService.CustomClaims).JwtUserData)
		req = req.WithContext(ctx)
		next.ServeHTTP(res, req)
	})
}

func unauthorized(res http.ResponseWriter) {
	httpUtils.SendJson(res, http.StatusUnauthorized, appErrors.ResponseError{Code: http.StatusUnauthorized, Message: appErrors.DefaultUnauthorizedMessage})
}
