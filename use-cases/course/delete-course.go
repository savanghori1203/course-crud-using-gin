package usecase_course

import (
	"context"
	db "gocourseCRUD/data-access"
	"gocourseCRUD/exceptions"

	"github.com/google/uuid"
)

func DeleteCourse(id string) error {
	_, err := uuid.Parse(id)

	if err != nil {
		return  exceptions.ValidationError("EX-0001", err.Error())
	}

	couseDetail, err := GetCourse(id)
	if err != nil {
		return err
	}

	if len(couseDetail) == 0 {
		return exceptions.ObjectNotFoundError("EX-0003", "No Course for given id provided")
	}

	return db.DeleteCourse(context.TODO(), id)
}
