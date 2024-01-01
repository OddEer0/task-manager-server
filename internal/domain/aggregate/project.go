package aggregate

import "task-manager-server/internal/domain/models"

type ProjectAggregate struct {
	Project models.Project
	Tags    []*models.Tag
	Columns []*models.Column
	Owner   *models.User
}
