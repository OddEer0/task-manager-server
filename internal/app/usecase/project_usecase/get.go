package projectUsecase

import (
	"context"
	"database/sql"
	appErrors "github.com/OddEer0/task-manager-server/internal/common/lib/app_errors"
	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
)

func (uc *projectUseCase) GetById(ctx context.Context, id string) (*aggregate.ProjectAggregate, error) {
	projectAggregate, err := uc.ProjectRepository.GetById(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, appErrors.NotFound("", "target: project use case, method: get by id, error: ", err.Error())
		}
		return nil, appErrors.InternalServerError("", "target: project use case, method: get by id, error: ", err.Error())
	}
	return projectAggregate, nil
}
