package authUsecase

import (
	"context"

	appDto "github.com/OddEer0/task-manager-server/internal/app/app_dto"
	tokenService "github.com/OddEer0/task-manager-server/internal/app/service/token_service"
	"github.com/OddEer0/task-manager-server/internal/common/lib/app_errors"
	"github.com/OddEer0/task-manager-server/internal/presentation/mapper"
)

func (a *authUseCase) Registration(ctx context.Context, data appDto.RegistrationUseCaseDto) (*AuthResult, error) {
	userAggregate, err := a.UserService.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	tokens, err := a.TokenService.Generate(tokenService.JwtUserData{Id: userAggregate.User.Id, Role: userAggregate.User.Role})
	if err != nil {
		return nil, err
	}
	err = userAggregate.SetToken(tokens.RefreshToken)
	if err != nil {
		return nil, appErrors.UnprocessableEntity("", "target: AuthUseCase, method: Registration. ", "Aggregate SetToken method error: ", err.Error())
	}

	dbUserAggregate, err := a.UserRepository.Create(ctx, userAggregate)
	if err != nil {
		return nil, appErrors.InternalServerError("", "target: AuthUseCase, method: Registration. ", "UserRepository create user error: ", err.Error())
	}
	userMapper := mapper.NewUserAggregateMapper()
	responseUser := userMapper.ToResponseUserDto(dbUserAggregate)

	return &AuthResult{User: &responseUser, Tokens: *tokens}, nil
}
