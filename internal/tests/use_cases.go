package tests

import (
	"context"

	tokenService "github.com/OddEer0/task-manager-server/internal/app/service/token_service"
	userService "github.com/OddEer0/task-manager-server/internal/app/service/user_service"
	authUsecase "github.com/OddEer0/task-manager-server/internal/app/usecase/auth_usecase"
	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
	"github.com/OddEer0/task-manager-server/internal/infrastructure/storage/mock_repository"
	"github.com/OddEer0/task-manager-server/internal/presentation/mock"
)

type MockUseCases struct {
	AuthUseCase authUsecase.AuthUseCase
}

func NewUseCases() *MockUseCases {
	userRepo := mock_repository.NewUserRepository()
	tokenRepo := mock_repository.NewTokenRepository()
	memUser := mock.NewMockUser()
	for _, user := range memUser.Users {
		userRepo.Create(context.Background(), &aggregate.UserAggregate{User: *user})
	}

	return &MockUseCases{
		AuthUseCase: authUsecase.New(userService.New(userRepo), tokenService.New(tokenRepo), userRepo),
	}
}
