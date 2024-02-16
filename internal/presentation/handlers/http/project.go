package http

import (
	tokenService "github.com/OddEer0/task-manager-server/internal/app/service/token_service"
	projectUseCase "github.com/OddEer0/task-manager-server/internal/app/usecase/project_usecase"
	appErrors "github.com/OddEer0/task-manager-server/internal/common/lib/app_errors"
	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
	"github.com/OddEer0/task-manager-server/internal/domain/models"
	"github.com/OddEer0/task-manager-server/internal/presentation/dto"
	httpUtils "github.com/OddEer0/task-manager-server/pkg/http_utils"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
)

type (
	ProjectHandler interface {
		GetProjectsById(res http.ResponseWriter, req *http.Request) error
		CreateProject(res http.ResponseWriter, req *http.Request) error
		UpdateProject(res http.ResponseWriter, req *http.Request) error
		DeleteProject(res http.ResponseWriter, req *http.Request) error
	}

	projectHandler struct {
		projectUseCase.ProjectUseCase
	}
)

func NewProjectHandler(projectUseCase projectUseCase.ProjectUseCase) ProjectHandler {
	return &projectHandler{
		ProjectUseCase: projectUseCase,
	}
}

func (p projectHandler) GetProjectsById(res http.ResponseWriter, req *http.Request) error {
	id := chi.URLParam(req, "id")
	if id == "" {
		return appErrors.BadRequest("id is required")
	}
	projectAggregate, err := p.ProjectUseCase.GetById(req.Context(), id)
	if err != nil {
		return err
	}
	httpUtils.SendJson(res, http.StatusOK, projectAggregate.Project)
	return nil
}

func (p projectHandler) CreateProject(res http.ResponseWriter, req *http.Request) error {
	var body dto.CreateProjectDto
	err := httpUtils.BodyJson(req, &body)
	if err != nil {
		return appErrors.BadRequest("", err.Error())
	}
	defer func() {
		_ = req.Body.Close()
	}()

	val := req.Context().Value("user")
	if val == nil {
		return appErrors.Unauthorized("")
	}
	user := val.(*tokenService.JwtUserData)

	projectAggregate, err := aggregate.NewProjectAggregate(models.Project{
		Id:     uuid.New().String(),
		Name:   body.Name,
		Color:  body.Color,
		Bg:     body.Bg,
		Order:  1,
		UserId: user.Id,
	})

	if err != nil {
		return appErrors.UnprocessableEntity("")
	}

	createdProjectAggregate, err := p.ProjectUseCase.Create(req.Context(), projectAggregate)
	if err != nil {
		return err
	}
	httpUtils.SendJson(res, http.StatusCreated, createdProjectAggregate.Project)
	return nil
}

func (p projectHandler) UpdateProject(res http.ResponseWriter, req *http.Request) error {
	var body models.Project
	err := httpUtils.BodyJson(req, &body)
	if err != nil {
		return appErrors.BadRequest("", err.Error())
	}
	defer func() {
		_ = req.Body.Close()
	}()

	return nil
}

func (p projectHandler) DeleteProject(res http.ResponseWriter, req *http.Request) error {
	//TODO implement me
	panic("implement me")
}
