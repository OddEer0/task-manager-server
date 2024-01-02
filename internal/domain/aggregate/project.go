package aggregate

import "github.com/OddEer0/task-manager-server/internal/domain/models"

type ProjectAggregate struct {
	Project models.Project
	Tags    []*models.Tag
	Columns []*models.Column
	Owner   *models.User
}
