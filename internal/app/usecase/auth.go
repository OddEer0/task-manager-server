package usecase

import (
	"context"
	"task-manager-server/internal/domain/models"
	"task-manager-server/internal/presentation/dto"
)

type AuthUseCase interface {
	registration(ctx context.Context, data dto.CreateUserDto) (*models.User, error)
	login(ctx context.Context, data dto.LoginInputDto)
}
