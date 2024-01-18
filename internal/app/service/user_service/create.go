package userService

import (
	"context"
	"fmt"
	"time"

	"github.com/OddEer0/task-manager-server/internal/app/app_dto"
	"github.com/OddEer0/task-manager-server/internal/app/factories/user_aggregate_factory"
	"github.com/OddEer0/task-manager-server/internal/common/constants"
	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
	"github.com/OddEer0/task-manager-server/pkg/app_errors"
	"github.com/google/uuid"
)

func (u *userService) Create(ctx context.Context, data appDto.RegistrationUseCaseDto) (*aggregate.UserAggregate, error) {
	candidate, err := u.UserRepository.HasUserByNick(ctx, data.Nick)
	if err != nil {
		return nil, appErrors.InternalServerError("", "target: UserService, method: Create. ", "user repository error: ", err.Error())
	}
	if candidate {
		return nil, appErrors.Conflict(constants.UserNickExist, "target: UserService, method: Create. ", "Nick conflict")
	}
	candidate, err = u.UserRepository.HasUserByEmail(ctx, data.Email)
	if err != nil {
		return nil, appErrors.InternalServerError("", "target: UserService, method: Create. ", "user repository error: ", err.Error())
	}
	if candidate {
		return nil, appErrors.Conflict(constants.UserEmailExist, "target: UserService, method: Create. ", "Email conflict")
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
		return nil, appErrors.UnprocessableEntity("", "target: UserService, method: Create. ", "Factory create user aggregate error: ", err.Error())
	}

	return userAggregate, nil
}
