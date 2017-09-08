package errors

import (
	"net/http"

	"gopkg.in/go-playground/validator.v8"
)

//APIError not found error
type APIError struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details,omitempty"`
}

func (e *APIError) Error() string {
	m := e.Message
	if m == "" {
		m = http.StatusText(e.Code)
	}
	return m
}

//New error
func New(code int, message string) *APIError {
	return &APIError{code, message, make(map[string]interface{})}
}

//NotFound error
func NotFound() *APIError {
	return New(http.StatusNotFound, "")
}

//Unauthorized error
func Unauthorized() *APIError {
	return New(http.StatusUnauthorized, "")
}

//BadRequest error
func BadRequest() *APIError {
	return New(http.StatusBadRequest, "")
}

//InternalServerError error
func InternalServerError() *APIError {
	return New(http.StatusInternalServerError, "")
}

//LoginFailed error
func LoginFailed(msg string) *APIError {
	return New(http.StatusBadRequest, msg)
}

//RegistrationFailed error
func RegistrationFailed(msg string) *APIError {
	return New(http.StatusBadRequest, msg)
}

//Validation error
func Validation(err error) *APIError {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errMsg := BadRequest()
		errMsg.Message = "Validation error"
		for key, verr := range validationErrors {
			errMsg.Details[key] = map[string]interface{}{verr.Tag: verr.Param}
		}
		return errMsg
	}
	return New(http.StatusBadRequest, err.Error())
}
