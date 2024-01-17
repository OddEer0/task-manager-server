package http

import (
	"net/http"

	appDto "github.com/OddEer0/task-manager-server/internal/app/app_dto"
	authUsecase "github.com/OddEer0/task-manager-server/internal/app/usecase/auth_usecase"
	"github.com/OddEer0/task-manager-server/pkg/app_errors"
	httpUtils "github.com/OddEer0/task-manager-server/pkg/http_utils"
)

type AuthHandler interface {
	Registration(res http.ResponseWriter, req *http.Request)
}

type authHandler struct {
	authUsecase.AuthUseCase
}

func NewAuthHandler(authUseCase authUsecase.AuthUseCase) AuthHandler {
	return authHandler{
		AuthUseCase: authUseCase,
	}
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

	httpUtils.SendJson(res, http.StatusOK, registerResult.User)
}
