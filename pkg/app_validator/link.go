package appValidator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func isLink(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	urlPattern := regexp.MustCompile(`(https?://[^\s]+)`)
	matches := urlPattern.FindStringSubmatch(value)
	if len(matches) > 0 {
		return true
	}
	return false
}
