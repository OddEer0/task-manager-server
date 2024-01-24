package authUsecase

import (
	"context"

	appDto "github.com/OddEer0/task-manager-server/internal/app/app_dto"
	"github.com/OddEer0/task-manager-server/internal/common/constants"
	appErrors "github.com/OddEer0/task-manager-server/internal/common/lib/app_errors"
	"github.com/OddEer0/task-manager-server/internal/presentation/mapper"
)

func (a *authUseCase) Refresh(ctx context.Context, refreshToken string) (*AuthResult, error) {
	if refreshToken == "" {
		return nil, appErrors.Unauthorized(constants.Unauthorized)
	}

	jwtUserData, err := a.TokenService.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}
	has, err := a.TokenService.HasByValue(ctx, refreshToken)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, appErrors.Unauthorized(constants.Unauthorized)
	}
	userAggregate, err := a.UserRepository.GetById(ctx, jwtUserData.Id)
	if err != nil {
		return nil, appErrors.InternalServerError("", "target: AuthUseCase, method: Refresh", "get user by id error", err.Error())
	}

	tokens, err := a.TokenService.Generate(*jwtUserData)
	if err != nil {
		return nil, err
	}
	_, err = a.TokenService.Save(ctx, appDto.SaveTokenServiceDto{Id: userAggregate.User.Id, RefreshToken: tokens.RefreshToken})
	if err != nil {
		return nil, err
	}
	userMapper := mapper.NewUserAggregateMapper()
	responseUser := userMapper.ToResponseUserDto(userAggregate)

	return &AuthResult{User: &responseUser, Tokens: *tokens}, nil
}
