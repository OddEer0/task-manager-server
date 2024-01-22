package validator

import (
	"github.com/OddEer0/task-manager-server/internal/common/app_validator"
)

func ValidateDto(dto interface{}) error {
	validator := appValidator.appValidator.New()
	err := validator.Struct(dto)
	if err != nil {
		return err
	}
	return nil
}
