package routes

import (
	"akatgelar/orderfaz-testcase-be/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)
func Swager(g *gin.RouterGroup) { 
	docs.SwaggerInfo.Title = "Orderfaz Testcase API"
	docs.SwaggerInfo.Description = "This is a sample server."
	docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Host = "orderfaz-testcase-be.akatgelar.app"
	docs.SwaggerInfo.BasePath = "/"
	// docs.SwaggerInfo.Schemes = []string{"http"} 
	docs.SwaggerInfo.Schemes = []string{"https"} 
	g.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}