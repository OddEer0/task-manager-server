package projectUsecase

import (
	"context"
	appErrors "github.com/OddEer0/task-manager-server/internal/common/lib/app_errors"
	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
)

func (uc *projectUseCase) Create(ctx context.Context, data *aggregate.ProjectAggregate) (*aggregate.ProjectAggregate, error) {
	projectAggregate, err := uc.ProjectRepository.Create(ctx, data)
	if err != nil {
		return nil, appErrors.InternalServerError("", "target: project use case, method: create, error: %s", err.Error())
	}
	return projectAggregate, nil
}
