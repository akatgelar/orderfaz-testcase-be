//go:debug x509negativeserial=1
package main

import (
	"akatgelar/orderfaz-testcase-be/database"
	"akatgelar/orderfaz-testcase-be/models"
	routes "akatgelar/orderfaz-testcase-be/routes"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)
 
func main() {
	r := gin.Default()

    // DB init 
    database.ConnnectDatabasePostgres() 

	// CORS Configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})) 
 
    // hello
    fmt.Println("Hello World")

    // router
	// @securityDefinitions.apikey Bearer
	// @in header
	// @name Authorization
	// @description Type "Bearer" followed by a space and JWT token.
	r.GET("/", func(c *gin.Context) {  
		c.JSON(
            200, 
            models.BaseResponse{
                Status: true,
                Message: "Welcome to API",
                Data: gin.H{ 
                    "swagger": "/docs/index.html",
                }, 
            }, 
        )
	}) 
    root := r.Group("/")
    {     
        routes.Swager(root.Group("/"))  
        routes.Auth(root.Group("/")) 
        routes.Logistic(root.Group("/"))  
    } 

	r.Run(":8080")
}