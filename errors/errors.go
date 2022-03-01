package errors

import (
	"net/http"
)

type AppError struct {
	Message string `json:"message"`
	Code    int    `json:"-"`
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewInternalServerError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NewBadRequestError(message string) *AppError {
	return &AppError{
		Message: message,
		Code: http.StatusBadRequest,
	}
}
