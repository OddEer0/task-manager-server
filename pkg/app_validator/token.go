package appValidator

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
)

func isJwtToken(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	_, _, err := new(jwt.Parser).ParseUnverified(value, jwt.MapClaims{})
	return err == nil
}
