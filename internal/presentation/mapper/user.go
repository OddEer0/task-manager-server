package mapper

import (
	"time"

	appDto "github.com/OddEer0/task-manager-server/internal/app/app_dto"
	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
)

type (
	UserAggregateMapper interface {
		ToResponseUserDto(aggregate *aggregate.UserAggregate) appDto.ResponseUserDto
	}

	userAggregateMapper struct{}
)

func NewUserAggregateMapper() UserAggregateMapper {
	return userAggregateMapper{}
}

func (u userAggregateMapper) ToResponseUserDto(aggregate *aggregate.UserAggregate) appDto.ResponseUserDto {
	var brthDay *time.Time = nil
	if aggregate.User.Birthday != nil {
		brthDay = &aggregate.User.Birthday.Value
	}

	return appDto.ResponseUserDto{
		Id:         aggregate.User.Id,
		Nick:       aggregate.User.Nick,
		Email:      aggregate.User.Email,
		FirstName:  aggregate.User.FirstName,
		LastName:   aggregate.User.LastName,
		SubTitle:   aggregate.User.SubTitle,
		Avatar:     aggregate.User.Avatar,
		Birthday:   brthDay,
		Role:       aggregate.User.Role,
		IsActivate: aggregate.User.IsActivate,
		IsBanned:   aggregate.User.IsBanned,
		BanReason:  aggregate.User.BanReason,
		CreatedAt:  &aggregate.User.CreatedAt,
		UpdatedAt:  &aggregate.User.UpdatedAt,
	}
}
