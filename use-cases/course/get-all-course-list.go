package usecase_course

import (
	"context"
    "gocourseCRUD/model"
    db "gocourseCRUD/data-access"
)

func GetAllCouseList() ([]model.Course, error) {
   return db.GetAllCouseList(context.TODO())
}
