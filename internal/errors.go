package internal

import (
	"log"
	"net/http"
)

type Error struct {
	Message    string
	StatusCode int
}

func NewError(message string, statusCode int) *Error {
	return &Error{
		Message:    message,
		StatusCode: statusCode,
	}
}

func (e Error) Error() string {
	return e.Message
}

func CoverError(service string, error *Error) *Error {
	if error.StatusCode == http.StatusInternalServerError {
		log.Println(service, error)
		return NewError("service temporarily unavailable", error.StatusCode)
	}
	return error
}
