package userService

import (
	"context"
	"fmt"
	"github.com/OddEer0/task-manager-server/internal/common/constants"
	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
	"github.com/OddEer0/task-manager-server/internal/domain/models"
	"github.com/OddEer0/task-manager-server/internal/domain/valuesobject"
	"github.com/OddEer0/task-manager-server/internal/presentation/dto"
	"github.com/OddEer0/task-manager-server/pkg/app_errors"
	"github.com/google/uuid"
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

	if err != nil {
		return nil, appErrors.UnprocessableEntity("")
	}
	hashPassword, err := valuesobject.NewPassword(data.Password)
	if err != nil {
		return nil, appErrors.UnprocessableEntity("")
	}
	activationLink := uuid.New()
	activationURL := fmt.Sprintf("%s/%s", constants.ActivationLinkURL, activationLink.String())
	id := uuid.New()

	userAggregate, err := aggregate.NewUserAggregate(models.User{
		Id:             id.String(),
		Nick:           data.Nick,
		Password:       hashPassword,
		Email:          "odd@",
		FirstName:      data.FirstName,
		LastName:       data.LastName,
		Role:           constants.User,
		ActivationLink: activationURL,
	})
	if err != nil {
		return nil, appErrors.UnprocessableEntity("")
	}

	return userAggregate, nil
}
