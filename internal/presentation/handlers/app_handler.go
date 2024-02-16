package handlers

import (
	tokenService "github.com/OddEer0/task-manager-server/internal/app/service/token_service"
	userService "github.com/OddEer0/task-manager-server/internal/app/service/user_service"
	authUsecase "github.com/OddEer0/task-manager-server/internal/app/usecase/auth_usecase"
	projectUsecase "github.com/OddEer0/task-manager-server/internal/app/usecase/project_usecase"
	"github.com/OddEer0/task-manager-server/internal/infrastructure/storage/mock_repository"
	"github.com/OddEer0/task-manager-server/internal/presentation/handlers/http"
)

type AppHandler struct {
	http.AuthHandler
	http.ProjectHandler
}

func NewAppHandler() AppHandler {
	tokenRepo := mock_repository.NewTokenRepository()
	userRepo := mock_repository.NewUserRepository()
	projectRepo := mock_repository.NewProjectRepository()

	appHandler := AppHandler{
		AuthHandler:    http.NewAuthHandler(authUsecase.New(userService.New(userRepo), tokenService.New(tokenRepo), userRepo)),
		ProjectHandler: http.NewProjectHandler(projectUsecase.New(projectRepo)),
	}

	return appHandler
}
