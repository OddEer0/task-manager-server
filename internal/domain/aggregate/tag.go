package aggregate

import "task-manager-server/internal/domain/models"

type TagAggregate struct {
	Tag     models.Tag
	Tasks   []*models.Task
	Project *models.Project
}
