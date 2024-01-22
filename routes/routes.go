package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josebruno2020/go-api-gin/controllers"
)

func HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}

func StudentsRoutes(r *gin.Engine) {
	s := r.Group("/students")

	s.GET("", controllers.FindAll)
	s.GET("/cpf/:cpf", controllers.FindByCPF)
	s.POST("", controllers.Create)
	s.GET("/:id", controllers.Find)
	s.PATCH("/:id", controllers.Update)
	s.DELETE("/:id", controllers.Delete)
}

func HandleRequest() *gin.Engine {
	r := gin.Default()

	r.GET("/", HealthCheck)

	StudentsRoutes(r)

	return r
}
