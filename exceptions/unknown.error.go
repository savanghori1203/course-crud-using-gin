package exceptions

import (
	"net/http"
	"time"
)

func UnknownError(customCode string, message interface{}) Error {
	if customCode == "" {
		customCode = "EX-0004"
	}
	if message == nil {
		message = "Something went wrong"
	}

	return Error{
		Name:           "UnknownError",
		Message:        "Something Went Wrong",
		CustomCode:     customCode,
		HttpStatusCode: http.StatusBadRequest,
		Date:           time.Now(),
	}
}
