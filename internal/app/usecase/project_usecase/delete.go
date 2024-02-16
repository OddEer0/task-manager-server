package projectUsecase

import (
	"context"
	"database/sql"
	appErrors "github.com/OddEer0/task-manager-server/internal/common/lib/app_errors"
)

func (uc *projectUseCase) Delete(ctx context.Context, id string) error {
	err := uc.ProjectRepository.Delete(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return appErrors.NotFound("", "target: project use case, method: delete, error: %s", err.Error())
		}
		return appErrors.InternalServerError("", "target: project use case, method: delete, error: ", err.Error())
	}
	return nil
}
