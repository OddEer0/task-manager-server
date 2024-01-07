package app_errors

import (
	"encoding/json"
	"net/http"
)

const (
	DefaultBadRequestMessage          = "Bad request"
	DefaultForbiddenMessage           = "Forbidden"
	DefaultInternalServerErrorMessage = "Server error"
	DefaultInternalServerErrorJson    = "{\"code\": 500, \"message\": \"" + DefaultInternalServerErrorMessage + "\"}"
)

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func HttpError(res http.ResponseWriter, message string, code int) {
	res.WriteHeader(code)
	errorData := ResponseError{Code: code, Message: DefaultBadRequestMessage}
	if err := json.NewEncoder(res).Encode(errorData); err != nil {
		http.Error(res, DefaultInternalServerErrorJson, http.StatusInternalServerError)
	}
}

func BadRequest(res http.ResponseWriter, message string) {
	if message == "" {
		message = DefaultBadRequestMessage
	}
	HttpError(res, message, http.StatusBadRequest)
}

func InternalServerError(res http.ResponseWriter, message string) {
	if message == "" {
		message = DefaultInternalServerErrorMessage
	}
	HttpError(res, message, http.StatusInternalServerError)
}

func Forbidden(res http.ResponseWriter, message string) {
	if message == "" {
		message = DefaultForbiddenMessage
	}
	HttpError(res, message, http.StatusForbidden)
}
