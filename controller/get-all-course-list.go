package controller

import (
	"gocourseCRUD/model"
	usecase_course "gocourseCRUD/use-cases/course"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCouseList(c *gin.Context) {
	courses, err := usecase_course.GetAllCouseList()
	if err != nil {
		c.JSON(err.(model.FormateError).Code, err)
	}

	if len(courses) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "No Course Added",
			"name":    "ValidationError",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_count": len(courses),
		"Data":        courses,
	})
}
