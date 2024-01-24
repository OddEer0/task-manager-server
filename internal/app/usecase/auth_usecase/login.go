package authUsecase

import (
	"context"

	appDto "github.com/OddEer0/task-manager-server/internal/app/app_dto"
	tokenService "github.com/OddEer0/task-manager-server/internal/app/service/token_service"
	"github.com/OddEer0/task-manager-server/internal/common/constants"
	"github.com/OddEer0/task-manager-server/internal/common/lib/app_errors"
	"github.com/OddEer0/task-manager-server/internal/presentation/mapper"
	"golang.org/x/crypto/bcrypt"
)

func (a *authUseCase) Login(ctx context.Context, data appDto.LoginUseCaseDto) (*AuthResult, error) {
	candidate, err := a.UserRepository.HasUserByNick(ctx, data.Nick)
	if err != nil {
		return nil, appErrors.InternalServerError("")
	}
	if !candidate {
		return nil, appErrors.Forbidden(constants.NickOrPasswordIncorrect)
	}
	userAggregate, err := a.UserRepository.GetByNick(ctx, data.Nick)
	if err != nil {
		return nil, appErrors.InternalServerError("")
	}

	isEqualPassword := bcrypt.CompareHashAndPassword([]byte(userAggregate.User.Password.Value), []byte(data.Password))
	if isEqualPassword != nil {
		return nil, appErrors.Forbidden(constants.NickOrPasswordIncorrect)
	}

	tokens, err := a.TokenService.Generate(tokenService.JwtUserData{Id: userAggregate.User.Id, Role: userAggregate.User.Role})
	if err != nil {
		return nil, err
	}

	_, err = a.TokenService.Save(ctx, appDto.SaveTokenServiceDto{Id: userAggregate.User.Id, RefreshToken: tokens.RefreshToken})
	if err != nil {
		return nil, err
	}

	aggregateMapper := mapper.NewUserAggregateMapper()
	responseUser := aggregateMapper.ToResponseUserDto(userAggregate)

	return &AuthResult{User: &responseUser, Tokens: *tokens}, nil
}
