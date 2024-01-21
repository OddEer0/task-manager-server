package appErrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	DefaultBadRequestMessage          = "Bad request"
	DefaultNotFoundErrorMessage       = "Not found"
	DefaultForbiddenMessage           = "Forbidden"
	DefaultInternalServerErrorMessage = "Server error"
	DefaultConflictMessage            = "Conflict"
	DefaultUnauthorizedMessage        = "Unauthorized"
	DefaultUnprocessableEntity        = "UnprocessableEntity"
	DefaultInternalServerErrorJson    = "{\"code\": 500, \"message\": \"" + DefaultInternalServerErrorMessage + "\"}"
)

type (
	AppError struct {
		Code       int    `json:"code"`
		Message    string `json:"message,omitempty"`
		DevMessage string `json:"-"`
	}
)

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

func HttpAppError(message string, code int, devMessages ...string) error {
	return &AppError{
		Code:       code,
		Message:    message,
		DevMessage: strings.Join(devMessages, ""),
	}
}

func BadRequest(message string, devMessages ...string) error {
	if message == "" {
		message = DefaultBadRequestMessage
	}
	return HttpAppError(message, http.StatusBadRequest, devMessages...)
}

func InternalServerError(message string, devMessages ...string) error {
	if message == "" {
		message = DefaultInternalServerErrorMessage
	}
	return HttpAppError(message, http.StatusInternalServerError, devMessages...)
}

func Forbidden(message string, devMessages ...string) error {
	if message == "" {
		message = DefaultForbiddenMessage
	}
	return HttpAppError(message, http.StatusForbidden, devMessages...)
}

func Conflict(message string, devMessages ...string) error {
	if message == "" {
		message = DefaultConflictMessage
	}
	return HttpAppError(message, http.StatusConflict, devMessages...)
}

func Unauthorized(message string, devMessages ...string) error {
	if message == "" {
		message = DefaultUnauthorizedMessage
	}
	return HttpAppError(message, http.StatusUnauthorized, devMessages...)
}

func UnprocessableEntity(message string, devMessages ...string) error {
	if message == "" {
		message = DefaultUnprocessableEntity
	}
	return HttpAppError(message, http.StatusUnprocessableEntity, devMessages...)
}

func NotFound(message string, devMessages ...string) error {
	if message == "" {
		message = DefaultNotFoundErrorMessage
	}
	return HttpAppError(message, http.StatusNotFound, devMessages...)
}
