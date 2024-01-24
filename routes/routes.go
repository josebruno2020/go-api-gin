package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josebruno2020/go-api-gin/controllers"
	"github.com/josebruno2020/go-api-gin/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}

func HandlerNotFound(ctx *gin.Context) {
	ctx.HTML(http.StatusNotFound, "404.html", nil)
}

func studentsRoutes(r *gin.Engine) {
	s := r.Group("/students")

	s.GET("", controllers.FindAll)
	s.GET("/cpf/:cpf", controllers.FindByCPF)
	s.POST("", controllers.Create)
	s.GET("/:id", controllers.Find)
	s.PATCH("/:id", controllers.Update)
	s.DELETE("/:id", controllers.Delete)

	s.GET("/index", controllers.IndexPage)
}

func setupSwagger(r *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func HandleRequest() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/assets", "./assets")
	r.NoRoute(HandlerNotFound)

	r.GET("/", healthCheck)

	studentsRoutes(r)
	setupSwagger(r)

	return r
}
