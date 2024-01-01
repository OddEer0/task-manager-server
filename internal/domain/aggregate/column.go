package aggregate

import "task-manager-server/internal/domain/models"

type ColumnAggregate struct {
	Column  models.Column
	Tasks   []*models.Task
	Project *models.Project
}
