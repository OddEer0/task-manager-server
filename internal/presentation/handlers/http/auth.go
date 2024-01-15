package http

import (
	"github.com/OddEer0/task-manager-server/internal/presentation/dto"
	"github.com/OddEer0/task-manager-server/pkg/app_errors"
	httpUtils "github.com/OddEer0/task-manager-server/pkg/http_utils"
	"net/http"
)

type AuthHandler interface {
	Registration(res http.ResponseWriter, req *http.Request)
}

type authHandler struct {
}

func NewAuthHandler() AuthHandler {
	return authHandler{}
}

func (a authHandler) Registration(res http.ResponseWriter, req *http.Request) {
	var body dto.RegistrationInputDto
	err := httpUtils.BodyJson(req, &body)

	if err != nil {
		appErrors.ErrorHandler(res, appErrors.BadRequest(""))
		return
	}
	defer func() {
		_ = req.Body.Close()
	}()

	httpUtils.SendJson(res, http.StatusOK, nil)
}
