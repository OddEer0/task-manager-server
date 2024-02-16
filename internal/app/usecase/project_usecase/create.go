package projectUsecase

import (
	"context"
	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
)

func (uc *projectUseCase) Create(ctx context.Context, data *aggregate.ProjectAggregate) (*aggregate.ProjectAggregate, error) {
	return nil, nil
}
