package models

import "net/http"

type HttpRequestError struct {
	Message string
	Status  int
}

func InternalServerError(message string) *HttpRequestError {
	return &HttpRequestError{
		Message: message,
		Status:  http.StatusInternalServerError,
	}
}
