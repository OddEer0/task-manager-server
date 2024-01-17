package authUsecase

import (
	"context"

	appDto "github.com/OddEer0/task-manager-server/internal/app/app_dto"
	"github.com/OddEer0/task-manager-server/internal/presentation/mapper"
	appErrors "github.com/OddEer0/task-manager-server/pkg/app_errors"
)

func (a *authUseCase) Registration(ctx context.Context, data appDto.RegistrationUseCaseDto) (*AuthResult, error) {
	userAggregate, err := a.UserService.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	tokens, err := a.TokenService.Generate(appDto.GenerateTokenServiceDto{Id: userAggregate.User.Id, Role: userAggregate.User.Role})
	if err != nil {
		return nil, appErrors.InternalServerError("")
	}
	err = userAggregate.SetToken(tokens.RefreshToken)
	if err != nil {
		return nil, appErrors.UnprocessableEntity("")
	}

	dbUserAggregate, err := a.UserRepository.Create(ctx, userAggregate)
	if err != nil {
		return nil, appErrors.InternalServerError("")
	}
	userMapper := mapper.NewUserAggregateMapper()
	responseUser := userMapper.ToResponseUserDto(dbUserAggregate)

	return &AuthResult{User: &responseUser, Tokens: *tokens}, nil
}
