package aggregate

import (
	appValidator "github.com/OddEer0/task-manager-server/internal/common/lib/app_validator"
	"github.com/OddEer0/task-manager-server/internal/domain/models"
)

type ColumnAggregate struct {
	Column  models.Column
	Tasks   []*models.Task
	Project *models.Project
}

func (c *ColumnAggregate) Validation() error {
	validator := appValidator.New()
	err := validator.Struct(c.Column)
	if err != nil {
		return err
	}
	return nil
}

func NewColumnAggregate(column models.Column) (*ColumnAggregate, error) {
	result := ColumnAggregate{Column: column}
	err := result.Validation()
	if err != nil {
		return nil, err
	}
	return &result, nil
}
