package validator

import appValidator "github.com/OddEer0/task-manager-server/pkg/app_validator"

func ValidateDto(dto interface{}) error {
	validator := appValidator.New()
	err := validator.Struct(dto)
	if err != nil {
		return err
	}
	return nil
}
