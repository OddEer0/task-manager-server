package repository

import (
	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
	"github.com/OddEer0/task-manager-server/internal/pkg/shared"
	"github.com/OddEer0/task-manager-server/internal/presentation/dto"
)

type UserRepository interface {
	shared.CRUDRepository[*aggregate.UserAggregate, dto.CreateUserDto]
}
