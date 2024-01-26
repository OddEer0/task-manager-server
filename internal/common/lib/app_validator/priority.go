package appValidator

import (
	"github.com/OddEer0/task-manager-server/internal/common/constants"
	"github.com/go-playground/validator/v10"
)

func isPriority(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return value == constants.PriorityExtra || value == constants.PriorityHigh || value == constants.PriorityNormal || value == constants.PriorityLow
}
