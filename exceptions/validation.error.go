package exceptions

import (
	"net/http"
	"time"
)

func ValidationError(customCode string, message interface{}) Error {
	if customCode == "" {
		customCode = "EX-0001"
	}
	if message == nil {
		message = "Something went wrong"
	}

	return Error{
		Name:           "ValidationError",
		Message:        message,
		CustomCode:     customCode,
		HttpStatusCode: http.StatusBadRequest,
		Date:           time.Now(),
	}
}
