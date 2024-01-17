package authUsecase

import (
	"context"

	appDto "github.com/OddEer0/task-manager-server/internal/app/app_dto"
)

func (a *authUseCase) Login(ctx context.Context, data appDto.LoginUseCaseDto) (*AuthResult, error) {
	return nil, nil
}
