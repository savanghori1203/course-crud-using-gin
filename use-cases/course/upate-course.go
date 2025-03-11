package usecase_course

import (
	"context"
	db "gocourseCRUD/data-access"
	"gocourseCRUD/exceptions"
	"gocourseCRUD/model"

	"github.com/google/uuid"
)

func UpdateCourse(id string, courseDetail model.Course) error {
	_, err := uuid.Parse(id)

	if err != nil {
		return exceptions.ValidationError("EX-0001", err.Error())
	}

	if courseDetail.IsEmpty() {
		return exceptions.ValidationError("EX-0001", "Please enter course detail")
	} else {
		courseData, err := GetCourse(id)
		if err != nil {
			return err
		}
		if len(courseData) == 0 {
			return exceptions.ObjectNotFoundError("EX-0003", "No Course for given id provided")
		}

		if courseDetail.Name != "" {
			IsCourseExist, err := IsCourseExist(courseDetail.Name)
			if err != nil {
				return err
			}
			if IsCourseExist {
				return exceptions.AlreadyExistError("EX-002", "Course already exist")
			}
		}

		return db.UpdateCourse(context.TODO(), id, courseDetail)

	}
}
