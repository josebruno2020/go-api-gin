package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/josebruno2020/go-api-gin/controllers"
)

func studentsRoutes(r *gin.Engine) {
	s := r.Group("/api/v1/students")

	studentController := controllers.StudentController{}

	s.GET("", studentController.FindAll)
	s.GET("/cpf/:cpf", studentController.FindByCPF)
	s.POST("", studentController.Create)
	s.GET("/:id", studentController.Find)
	s.PATCH("/:id", studentController.Update)
	s.DELETE("/:id", studentController.Delete)

	r.GET("/index", studentController.IndexPage)
}
