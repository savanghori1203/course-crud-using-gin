package controller

import (
	"gocourseCRUD/exceptions"
	"gocourseCRUD/model"
	usecase_course "gocourseCRUD/use-cases/course"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddCourse(c *gin.Context) {
	var course model.Course

	err := c.BindJSON(&course)

	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.ValidationError("EX-0001", err.Error()))
	}

	id, err := usecase_course.AddCourse(course.Name, course.Platform, course.Price)

	if err != nil {
		c.JSON(err.(exceptions.Error).HttpStatusCode, err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}
