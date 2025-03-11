package exceptions

import (
	"net/http"
	"time"
)

func ObjectNotFoundError(customCode string, message interface{}) Error {
	if customCode == "" {
		customCode = "EX-0003"
	}
	if message == nil {
		message = "Something went wrong"
	}

	return Error{
		Name:           "ObjectNotFoundError",
		Message:        message,
		CustomCode:     customCode,
		HttpStatusCode: http.StatusNotFound,
		Date:           time.Now(),
	}
}
