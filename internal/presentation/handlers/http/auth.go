package http

import (
	"net/http"
	"time"

	"github.com/OddEer0/task-manager-server/config"
	appDto "github.com/OddEer0/task-manager-server/internal/app/app_dto"
	authUsecase "github.com/OddEer0/task-manager-server/internal/app/usecase/auth_usecase"
	"github.com/OddEer0/task-manager-server/internal/common/lib/app_errors"
	httpUtils "github.com/OddEer0/task-manager-server/pkg/http_utils"
)

type AuthHandler interface {
	Registration(res http.ResponseWriter, req *http.Request)
	Login(res http.ResponseWriter, req *http.Request)
}

type authHandler struct {
	authUsecase.AuthUseCase
}

func NewAuthHandler(authUseCase authUsecase.AuthUseCase) AuthHandler {
	return authHandler{
		AuthUseCase: authUseCase,
	}
}

func (a authHandler) setToken(res http.ResponseWriter, refreshToken string, accessToken string) error {
	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}
	refreshTokenTime, err := time.ParseDuration(cfg.RefreshTokenTime)
	if err != nil {
		return err
	}
	accessTokenTime, err := time.ParseDuration(cfg.AccessTokenTime)
	if err != nil {
		return err
	}

	refreshCookie := http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		Path:     "/",
		MaxAge:   int(refreshTokenTime.Minutes()),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	accessCookie := http.Cookie{
		Name:     "accessToken",
		Value:    accessToken,
		Path:     "/",
		MaxAge:   int(accessTokenTime.Minutes()),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(res, &accessCookie)
	http.SetCookie(res, &refreshCookie)
	return nil
}

func (a authHandler) Registration(res http.ResponseWriter, req *http.Request) {
	var body appDto.RegistrationUseCaseDto
	err := httpUtils.BodyJson(req, &body)
	if err != nil {
		appErrors.ErrorHandler(res, appErrors.BadRequest(""))
		return
	}
	defer func() {
		_ = req.Body.Close()
	}()

	registerResult, err := a.AuthUseCase.Registration(req.Context(), body)
	if err != nil {
		appErrors.ErrorHandler(res, err)
		return
	}

	err = a.setToken(res, registerResult.Tokens.RefreshToken, registerResult.Tokens.AccessToken)
	if err != nil {
		appErrors.ErrorHandler(res, appErrors.InternalServerError(err.Error()))
		return
	}
	httpUtils.SendJson(res, http.StatusOK, registerResult.User)
}

func (a authHandler) Login(res http.ResponseWriter, req *http.Request) {
	var body appDto.LoginUseCaseDto
	err := httpUtils.BodyJson(req, &body)
	if err != nil {
		appErrors.ErrorHandler(res, appErrors.BadRequest(""))
		return
	}
	defer func() {
		_ = req.Body.Close()
	}()

	loginResult, err := a.AuthUseCase.Login(req.Context(), body)
	if err != nil {
		appErrors.ErrorHandler(res, err)
		return
	}

	err = a.setToken(res, loginResult.Tokens.RefreshToken, loginResult.Tokens.AccessToken)
	if err != nil {
		appErrors.ErrorHandler(res, appErrors.InternalServerError("set token error"))
		return
	}

	httpUtils.SendJson(res, http.StatusOK, loginResult.User)
}
