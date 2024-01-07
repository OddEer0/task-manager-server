package http

import (
	"github.com/OddEer0/task-manager-server/internal/presentation/dto"
	"github.com/OddEer0/task-manager-server/internal/presentation/mock"
	"github.com/OddEer0/task-manager-server/pkg/app_errors"
	"github.com/OddEer0/task-manager-server/pkg/http_utils"
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
	err := http_utils.BodyJson(req, &body)
	if err != nil {
		app_errors.BadRequest(res, "")
		return
	}
	defer func() {
		_ = req.Body.Close()
	}()
	mockUser := mock.NewMockUser()
	http_utils.SendJson(res, http.StatusOK, mockUser.AdminUser)
}
