package controller

import (
	"gocourseCRUD/model"
	usecase_course "gocourseCRUD/use-cases/course"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteCourse(c *gin.Context) {
	id := c.Param("id")

	err := usecase_course.DeleteCourse(id)
	if err != nil {
		c.JSON(err.(model.FormateError).Code, err)
		return
	}


	c.Status(http.StatusNoContent)
}
