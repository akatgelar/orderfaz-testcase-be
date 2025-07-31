package routes

import (
	"akatgelar/orderfaz-testcase-be/controllers"
	"akatgelar/orderfaz-testcase-be/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)
func Logistic(g *gin.RouterGroup) {  

	g.GET("/logistic", helpers.JwtMiddleware(), 
	func(c *gin.Context) {  
		// parameter
		logistic_name := c.DefaultQuery("logistic_name", "ALL")
		origin_name := c.DefaultQuery("origin_name", "ALL")
		destination_name := c.DefaultQuery("destination_name", "ALL") 

		// response 
		response := controllers.LogisticRead(c, logistic_name, origin_name, destination_name) 
		
		if response.Status {
			c.JSON(http.StatusOK, response)
		} else { 
			c.JSON(http.StatusInternalServerError, response)
		}
	})

	g.POST("/logistic", helpers.JwtMiddleware(), 
	func(c *gin.Context) {   
		// response
		response := controllers.LogisticCreate(c) 
		if response.Status {
			c.JSON(http.StatusOK, response)
		} else { 
			c.JSON(http.StatusInternalServerError, response)
		}
	})
 
}