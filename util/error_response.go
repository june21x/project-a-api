package util

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// ErrorRes is a struct representing an error response
type ErrorRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewErrorResponse creates a new ErrorResponse instance from an error
func NewErrorResponse(code int, err error) *ErrorRes {
	// Default values
	var _code int
	var message string

	// Check if the error is a known error type
	switch typedErr := err.(type) {
	case validator.ValidationErrors:
		_code = http.StatusBadRequest

		fieldErrs := typedErr
		for _, fieldErr := range fieldErrs {
			message = fmt.Sprint(message, fieldErr.Translate(errorTranslator))
		}
	case *validator.ValidationErrors:
		_code = http.StatusBadRequest

		fieldErrs := *typedErr
		for _, fieldErr := range fieldErrs {
			message = fmt.Sprint(message, fieldErr.Translate(errorTranslator))
		}
	// case *MyCustomErrorType:
	// 	code = http.StatusBadRequest
	// 	message = err.Error()
	// case *AnotherCustomErrorType:
	// 	code = http.StatusNotFound
	// 	message = err.Error()
	default:
		{
			_code = code
			message = err.Error()
		}
	}

	// Create and return the error response
	return &ErrorRes{
		Code:    _code,
		Message: message,
	}
}
