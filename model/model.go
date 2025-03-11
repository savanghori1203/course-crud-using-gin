package model

import (
	"fmt"
	"github.com/google/uuid"
)

type Course struct {
	Id       uuid.UUID `json:"id,omitempty"`
	Name     string    `json:"name,omitempty" validate:"required,min=2"`
	Platform string    `json:"platform,omitempty" validate:"required,min=2"`
	Price    string    `json:"price,omitempty" validate:"required,number"`
}

type FormateError struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Name    string      `json:"name"`
}

func (e FormateError) Error() string {
	return fmt.Sprintf("%v", e.Message)
}

func (c Course) IsEmpty() bool {
	return c.Name == "" && c.Platform == "" && c.Price == ""
}
