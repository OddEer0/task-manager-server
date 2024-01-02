package usecase

import (
	"context"
	"github.com/OddEer0/task-manager-server/internal/domain/models"
	"github.com/OddEer0/task-manager-server/internal/presentation/dto"
)

type AuthUseCase interface {
	registration(ctx context.Context, data dto.CreateUserDto) (*models.User, error)
	login(ctx context.Context, data dto.LoginInputDto)
}
