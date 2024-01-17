package repository

import (
	"context"
	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
	"github.com/OddEer0/task-manager-server/pkg/shared"
)

type UserRepository interface {
	shared.CRUDRepository[*aggregate.UserAggregate, *aggregate.UserAggregate]
	HasUserByNick(ctx context.Context, nick string) (bool, error)
	HasUserByEmail(ctx context.Context, email string) (bool, error)
}
