package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josebruno2020/go-api-gin/controllers"
)

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}

func StudentsRoutes(r *gin.Engine) {
	s := r.Group("/students")

	s.GET("", controllers.FindAll)
	s.POST("", controllers.Create)
	s.GET("/:id", controllers.Find)
}

func HandleRequest() {
	r := gin.Default()

	r.GET("/", healthCheck)

	StudentsRoutes(r)

	r.Run(":3000")
}
