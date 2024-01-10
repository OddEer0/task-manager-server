package mapper

import (
	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
	"github.com/OddEer0/task-manager-server/internal/presentation/dto"
)

type (
	UserAggregateMapper interface {
		ToResponseUserDto(aggregate *aggregate.UserAggregate) dto.ResponseUserDto
	}

	userAggregateMapper struct{}
)

func NewUserAggregateMapper() UserAggregateMapper {
	return userAggregateMapper{}
}

func (u userAggregateMapper) ToResponseUserDto(aggregate *aggregate.UserAggregate) dto.ResponseUserDto {
	return dto.ResponseUserDto{}
}
