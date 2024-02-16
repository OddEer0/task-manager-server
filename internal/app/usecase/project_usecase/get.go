package projectUsecase

import (
	"context"
	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
)

func (uc *projectUseCase) GetById(ctx context.Context, id string) (*aggregate.ProjectAggregate, error) {
	return nil, nil
}
