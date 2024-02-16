package projectUsecase

import (
	"context"
	"database/sql"
	appErrors "github.com/OddEer0/task-manager-server/internal/common/lib/app_errors"
	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
)

func (uc *projectUseCase) Update(ctx context.Context, id string, data *aggregate.ProjectAggregate) (*aggregate.ProjectAggregate, error) {
	projectAggregate, err := uc.ProjectRepository.Update(ctx, id, data)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, appErrors.NotFound("", "target: project use case, method: update, error: ", err.Error())
		}
		return nil, appErrors.InternalServerError("", "target: project use case, method: update, error: %s", err.Error())
	}
	return projectAggregate, nil
}
