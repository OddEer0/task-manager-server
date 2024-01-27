package tests

import (
	"context"

	userService "github.com/OddEer0/task-manager-server/internal/app/service/user_service"
	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
	"github.com/OddEer0/task-manager-server/internal/infrastructure/storage/mock_repository"
	"github.com/OddEer0/task-manager-server/internal/presentation/mock"
)

type MockServices struct {
	UserService userService.Service
}

func NewServices() *MockServices {
	userRepo := mock_repository.NewUserRepository()
	memUser := mock.NewMockUser()
	for _, user := range memUser.Users {
		userRepo.Create(context.Background(), &aggregate.UserAggregate{User: *user})
	}

	return &MockServices{
		UserService: userService.New(userRepo),
	}
}
