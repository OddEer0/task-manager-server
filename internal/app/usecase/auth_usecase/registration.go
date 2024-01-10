package authUsecase

import (
	"context"
	"github.com/OddEer0/task-manager-server/internal/presentation/dto"
	"github.com/OddEer0/task-manager-server/internal/presentation/mapper"
	"github.com/OddEer0/task-manager-server/pkg/app_errors"
)

func (a *authUseCase) Registration(ctx context.Context, data dto.RegistrationInputDto) (*AuthResult, error) {
	userAggregate, err := a.UserService.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	tokens := a.TokenService.Generate(dto.GenerateTokenDto{Id: userAggregate.User.Id, Roles: userAggregate.User.Role})
	_, err = a.TokenService.Save(ctx, dto.SaveTokenDto{Id: userAggregate.User.Id, RefreshToken: tokens.RefreshToken})
	if err != nil {
		return nil, appErrors.InternalServerError("")
	}
	userMapper := mapper.NewUserAggregateMapper()
	responseUser := userMapper.ToResponseUserDto(userAggregate)

	return &AuthResult{User: &responseUser, Tokens: tokens}, nil
}
