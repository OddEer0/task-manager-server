package appErrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	DefaultBadRequestMessage          = "Bad request"
	DefaultForbiddenMessage           = "Forbidden"
	DefaultInternalServerErrorMessage = "Server error"
	DefaultConflictMessage            = "Conflict"
	DefaultUnauthorizedMessage        = "Unauthorized"
	DefaultUnprocessableEntity        = "UnprocessableEntity"
	DefaultInternalServerErrorJson    = "{\"code\": 500, \"message\": \"" + DefaultInternalServerErrorMessage + "\"}"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func ErrorHandler(res http.ResponseWriter, err error) {
	res.Header().Set("Content-Type", "application/json")

	var statusCode int
	var errorMessage string

	var e *AppError
	if errors.As(err, &e) {
		statusCode = e.Code
		errorMessage = e.Message
	}

	res.WriteHeader(statusCode)

	jsonError := &AppError{
		Code:    statusCode,
		Message: errorMessage,
	}

	encoder := json.NewEncoder(res)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(jsonError)
	if err != nil {
		http.Error(res, DefaultInternalServerErrorJson, http.StatusInternalServerError)
	}
}

func HttpAppError(message string, code int) error {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func BadRequest(message string) error {
	if message == "" {
		message = DefaultBadRequestMessage
	}
	return HttpAppError(message, http.StatusBadRequest)
}

func InternalServerError(message string) error {
	if message == "" {
		message = DefaultInternalServerErrorMessage
	}
	return HttpAppError(message, http.StatusInternalServerError)
}

func Forbidden(message string) error {
	if message == "" {
		message = DefaultForbiddenMessage
	}
	return HttpAppError(message, http.StatusForbidden)
}

func Conflict(message string) error {
	if message == "" {
		message = DefaultConflictMessage
	}
	return HttpAppError(message, http.StatusConflict)
}

func Unauthorized(message string) error {
	if message == "" {
		message = DefaultUnauthorizedMessage
	}
	return HttpAppError(message, http.StatusUnauthorized)
}

func UnprocessableEntity(message string) error {
	if message == "" {
		message = DefaultUnauthorizedMessage
	}
	return HttpAppError(message, http.StatusUnprocessableEntity)
}
