package main

import (
	"gocourseCRUD/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/course", controller.AddCourse)

	r.GET("/course/:id", controller.GetCourse)

	r.GET("/course", controller.GetAllCouseList)

	r.DELETE("/course/:id", controller.DeleteCourse)
	
	r.PATCH("/course/:id", controller.UpdateCourse)

	return r
}
