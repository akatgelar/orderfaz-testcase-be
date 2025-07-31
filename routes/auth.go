package routes

import (
	"akatgelar/orderfaz-testcase-be/controllers"
	"akatgelar/orderfaz-testcase-be/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)
func Auth(g *gin.RouterGroup) {  

	g.POST("/auth/register", func(c *gin.Context) {    
		// response 
		response := controllers.AuthRegister(c) 
		
		if response.Status {
			c.JSON(http.StatusOK, response)
		} else { 
			c.JSON(http.StatusInternalServerError, response)
		}
	})

	g.POST("/auth/login", func(c *gin.Context) {   
		// response
		response := controllers.AuthLogin(c) 
		if response.Status {
			c.JSON(http.StatusOK, response)
		} else { 
			c.JSON(http.StatusInternalServerError, response)
		}
	})
 

	g.GET("/auth/validate", helpers.JwtMiddleware(), 
	func(c *gin.Context) {   
		// response
		response := controllers.AuthValidate(c) 
		if response.Status {
			c.JSON(http.StatusOK, response)
		} else { 
			c.JSON(http.StatusInternalServerError, response)
		}
	})
}