package userService

import (
	"context"
	"encoding/base64"
	"github.com/OddEer0/task-manager-server/internal/common/constants"
	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
	"github.com/OddEer0/task-manager-server/internal/domain/models"
	"github.com/OddEer0/task-manager-server/internal/domain/valuesobject"
	"github.com/OddEer0/task-manager-server/internal/presentation/dto"
	"github.com/OddEer0/task-manager-server/pkg/app_errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (u *userService) Create(ctx context.Context, data dto.RegistrationInputDto) (*aggregate.UserAggregate, error) {
	candidate, err := u.UserRepository.HasUserByNick(ctx, data.Nick)
	if err != nil {
		return nil, appErrors.InternalServerError("")
	}
	if candidate {
		return nil, appErrors.Conflict(constants.UserNickExist)
	}
	candidate, err = u.UserRepository.HasUserByEmail(ctx, data.Email)
	if err != nil {
		return nil, appErrors.InternalServerError("")
	}
	if candidate {
		return nil, appErrors.Conflict(constants.UserEmailExist)
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, appErrors.InternalServerError("")
	}

	activationLink := uuid.New()

	createUserAggregate, err := aggregate.NewUserAggregate(models.User{
		Nick:           data.Nick,
		Password:       base64.StdEncoding.EncodeToString(hashPassword),
		Email:          valuesobject.Email{Value: data.Email},
		FirstName:      data.FirstName,
		LastName:       data.LastName,
		Role:           constants.User,
		ActivationLink: activationLink.String(),
	})
	if err != nil {
		return nil, appErrors.UnprocessableEntity("")
	}

	userAggregate, err := u.UserRepository.Create(ctx, createUserAggregate)
	if err != nil {
		return nil, appErrors.InternalServerError("")
	}
	return userAggregate, nil
}
