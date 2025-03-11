package usecase_course

import (
	"context"
	db "gocourseCRUD/data-access"
)

func IsCourseExist(name string) (bool, error) {
	cousre, err := db.GetCourseByName(context.TODO(), name)

	if err != nil {
		return false, err
	}
	return len(cousre) != 0, nil
}
