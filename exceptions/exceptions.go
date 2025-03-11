package exceptions

import (
	"fmt"
	"time"
)

type Error struct {
	Name           string      `json:"name"`
	Message        interface{} `json:"message"`
	CustomCode     string      `json:"code"`
	HttpStatusCode int         `json:"httpStatusCode"`
	Date           time.Time   `json:"date"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%v", e.Message)
}