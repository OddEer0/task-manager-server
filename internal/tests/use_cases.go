package tests

import (
	tokenService "github.com/OddEer0/task-manager-server/internal/app/service/token_service"
	userService "github.com/OddEer0/task-manager-server/internal/app/service/user_service"
	authUsecase "github.com/OddEer0/task-manager-server/internal/app/usecase/auth_usecase"
	"github.com/OddEer0/task-manager-server/internal/infrastructure/storage/mock_repository"
)

type MockUseCases struct {
	AuthUseCase authUsecase.AuthUseCase
}

func NewUseCases() *MockUseCases {
	userRepo := mock_repository.NewUserRepository()
	tokenRepo := mock_repository.NewTokenRepository()

	return &MockUseCases{
		AuthUseCase: authUsecase.New(userService.New(userRepo), tokenService.New(tokenRepo), userRepo),
	}
}
