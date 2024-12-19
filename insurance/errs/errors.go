package errs

import "net/http"

type AppError struct {
	Code    int    `json:"code,omitempty"` // Corrected JSON tag syntax
	Message string `json:"message"`        // Added JSON tag for the message
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewUnxpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}
