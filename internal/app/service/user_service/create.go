package userService

import (
	"context"
	"fmt"
	"time"

	appDto "github.com/OddEer0/task-manager-server/internal/app/app_dto"
	userAggregateFactory "github.com/OddEer0/task-manager-server/internal/app/factories/user_aggregate_factory"
	"github.com/OddEer0/task-manager-server/internal/common/constants"
	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
	"github.com/OddEer0/task-manager-server/pkg/app_errors"
	"github.com/google/uuid"
)

func (u *userService) Create(ctx context.Context, data appDto.RegistrationUseCaseDto) (*aggregate.UserAggregate, error) {
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

	factory := userAggregateFactory.UserAggregateFactory{}
	userAggregate, err := factory.CreateUserAggregate(userAggregateFactory.CreateUserAggregateData{
		Id:             uuid.New().String(),
		Nick:           data.Nick,
		Password:       data.Password,
		Email:          data.Email,
		FirstName:      data.FirstName,
		LastName:       data.LastName,
		Role:           constants.User,
		ActivationLink: fmt.Sprintf("%s/%s", constants.ActivationLinkURL, uuid.New().String()),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	})
	if err != nil {
		return nil, appErrors.UnprocessableEntity("")
	}

	return userAggregate, nil
}
