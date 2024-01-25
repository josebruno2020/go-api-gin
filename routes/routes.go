package routes

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/josebruno2020/go-api-gin/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}

func handlerNotFound(ctx *gin.Context) {
	ctx.HTML(http.StatusNotFound, "404.html", nil)
}

func setupSwagger(r *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func setupReleaseMode() {
	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func HandleRequest() *gin.Engine {
	setupReleaseMode()
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/assets", "./assets")
	r.NoRoute(handlerNotFound)

	r.GET("/", healthCheck)

	apiGroup := r.Group("/api/v1")
	studentsRoutes(apiGroup)
	setupSwagger(r)

	r.Use(cors.Default())
	return r
}
