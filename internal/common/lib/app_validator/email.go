package appValidator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func email(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(regex, value)
	return match && err == nil
}
