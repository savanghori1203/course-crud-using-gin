package exceptions

import (
	"net/http"
	"time"
)

func AlreadyExistError(customCode string, message interface{}) Error {
	if customCode == "" {
		customCode = "EX-002"
	}
	if message == nil {
		message = "Something went wrong"
	}

	return Error{
		Name:           "AlreadyExistError",
		Message:        message,
		CustomCode:     customCode,
		HttpStatusCode: http.StatusBadRequest,
		Date:           time.Now(),
	}
}
