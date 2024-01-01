package aggregate

import "task-manager-server/internal/domain/models"

type TaskAggregate struct {
	Task   models.Task
	Tags   []*models.Tag
	Column *models.Column
}
