package controller

import (
	"gocourseCRUD/model"
	usecase_course "gocourseCRUD/use-cases/course"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCourse(c *gin.Context) {
	id := c.Param("id")

	course, err := usecase_course.GetCourse(id)
	if err != nil {
		c.JSON(err.(model.FormateError).Code, err)
		return
	}

	if len(course) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "No Course for given id provided",
			"name":    "ValidationError",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Data": course,
	})

}
