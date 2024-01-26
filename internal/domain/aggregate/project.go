package aggregate

import (
	appValidator "github.com/OddEer0/task-manager-server/internal/common/lib/app_validator"
	"github.com/OddEer0/task-manager-server/internal/domain/models"
)

type ProjectAggregate struct {
	Project models.Project
	Tags    []*models.Tag
	Columns []*models.Column
	Owner   *models.User
}

func (p *ProjectAggregate) Validation() error {
	validator := appValidator.New()
	err := validator.Struct(p.Project)
	if err != nil {
		return err
	}
	return nil
}

func NewProjectAggregate(project models.Project) (*ProjectAggregate, error) {
	result := ProjectAggregate{Project: project}
	err := result.Validation()
	if err != nil {
		return nil, err
	}
	return &result, nil
}
