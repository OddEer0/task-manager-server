package appValidator

import "github.com/go-playground/validator/v10"

func New() *validator.Validate {
	validate := validator.New()

	_ = validate.RegisterValidation("isPriority", isPriority)
	_ = validate.RegisterValidation("uuidv4", uuidv4)
	_ = validate.RegisterValidation("email", email)
	_ = validate.RegisterValidation("userRole", userRole)
	_ = validate.RegisterValidation("dateIsLessNow", dateIsLessNow)
	_ = validate.RegisterValidation("isLink", isLink)
	_ = validate.RegisterValidation("isJwtToken", isJwtToken)

	return validate
}
