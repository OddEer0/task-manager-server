package userService

import (
	"context"
	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
	"github.com/OddEer0/task-manager-server/internal/domain/repository"
	"github.com/OddEer0/task-manager-server/internal/presentation/dto"
)

type (
	Service interface {
		Create(ctx context.Context, data dto.RegistrationInputDto) (*aggregate.UserAggregate, error)
	}

	userService struct {
		repository.UserRepository
	}
)

func NewUserService(userRepo repository.UserRepository) Service {
	return &userService{UserRepository: userRepo}
}
