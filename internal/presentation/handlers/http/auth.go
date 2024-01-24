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
	Registration(res http.ResponseWriter, req *http.Request) error
	Login(res http.ResponseWriter, req *http.Request) error
	Logout(res http.ResponseWriter, req *http.Request) error
	Refresh(res http.ResponseWriter, req *http.Request) error
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

func (a authHandler) removeToken(res http.ResponseWriter) {
	refreshCookie := http.Cookie{
		Name:     "refreshToken",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}

	accessCookie := http.Cookie{
		Name:     "accessToken",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}

	http.SetCookie(res, &accessCookie)
	http.SetCookie(res, &refreshCookie)
}

func (a authHandler) Registration(res http.ResponseWriter, req *http.Request) error {
	var body appDto.RegistrationUseCaseDto
	err := httpUtils.BodyJson(req, &body)
	if err != nil {
		return appErrors.BadRequest("")
	}
	defer func() {
		_ = req.Body.Close()
	}()

	registerResult, err := a.AuthUseCase.Registration(req.Context(), body)
	if err != nil {
		return err
	}

	err = a.setToken(res, registerResult.Tokens.RefreshToken, registerResult.Tokens.AccessToken)
	if err != nil {
		return appErrors.InternalServerError(err.Error())
	}
	httpUtils.SendJson(res, http.StatusOK, registerResult.User)
	return nil
}

func (a authHandler) Login(res http.ResponseWriter, req *http.Request) error {
	var body appDto.LoginUseCaseDto
	err := httpUtils.BodyJson(req, &body)
	if err != nil {
		return appErrors.BadRequest("")
	}
	defer func() {
		_ = req.Body.Close()
	}()

	loginResult, err := a.AuthUseCase.Login(req.Context(), body)
	if err != nil {
		return err
	}

	err = a.setToken(res, loginResult.Tokens.RefreshToken, loginResult.Tokens.AccessToken)
	if err != nil {
		return appErrors.InternalServerError("set token error")
	}

	httpUtils.SendJson(res, http.StatusOK, loginResult.User)
	return nil
}

func (a authHandler) Logout(res http.ResponseWriter, req *http.Request) error {
	token, err := req.Cookie("refreshToken")
	if err != nil {
		return appErrors.BadRequest("")
	}
	err = a.AuthUseCase.Logout(req.Context(), token.Value)
	if err != nil {
		return err
	}

	a.removeToken(res)
	return nil
}

func (a authHandler) Refresh(res http.ResponseWriter, req *http.Request) error {
	token, err := req.Cookie("refreshToken")
	if err != nil {
		return appErrors.BadRequest("")
	}
	result, err := a.AuthUseCase.Refresh(req.Context(), token.Value)
	if err != nil {
		return err
	}
	err = a.setToken(res, result.Tokens.RefreshToken, result.Tokens.AccessToken)
	if err != nil {
		return appErrors.InternalServerError("set token error")
	}
	httpUtils.SendJson(res, http.StatusOK, result.User)
	return nil
}
