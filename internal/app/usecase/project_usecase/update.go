package projectUsecase

import (
	"context"
	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
)

func (uc *projectUseCase) Update(ctx context.Context, id string, data *aggregate.ProjectAggregate) (*aggregate.ProjectAggregate, error) {
	return nil, nil
}
