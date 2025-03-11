package controller

import (
	"gocourseCRUD/model"
	usecase_course "gocourseCRUD/use-cases/course"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateCourse(c *gin.Context) {
	var courseDetail model.Course

	id := c.Param("id")

	err := c.BindJSON(&courseDetail)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"name":    "ValidationError",
		})
		return
	}

	err = usecase_course.UpdateCourse(id, courseDetail)

	if err != nil {
		c.JSON(err.(model.FormateError).Code, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Course updated successfully",
	})
}
