package controllers

import (
	"akatgelar/orderfaz-testcase-be/database"
	"akatgelar/orderfaz-testcase-be/models"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

// AuthRegister godoc
// @Summary      Post register payload
// @Description  Post register payload
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        auth body  models.AuthRegister true "register"
// @Success      200  {object}  models.BaseResponse
// @Failure      400  {object}  models.BaseResponse
// @Failure      404  {object}  models.BaseResponse
// @Failure      500  {object}  models.BaseResponse
// @Router       /auth/register [post]
func AuthRegister(c *gin.Context) models.BaseResponse {
	db := database.DB_POSTGRES
	  
	var tempAuth models.Auth

	// prepare
	if err := c.BindJSON(&tempAuth); err != nil { 
		response := models.BaseResponse{Status: false, Message: "Internal server error, " + "Invalid request payload"} 
		return response
	} 

	// check if key exist
	if tempAuth.Msisdn == "" || tempAuth.Username == "" || tempAuth.Password == "" || tempAuth.Name == "" { 
		response := models.BaseResponse{Status: false, Message: "Required msisdn, username, password, name"} 
		return response
	}

	// check if data exist 
	var checkAuth []models.Auth
	query := db.Model(&models.Auth{})
	query.Where("msisdn = ? or username = ? ", tempAuth.Msisdn, tempAuth.Username) 
	if err := query.Find(&checkAuth).Scan(&checkAuth).Error; err != nil { 
		response := models.BaseResponse{Status: false, Message: "Internal server error: " + err.Error()} 
		return response
	}
	if len(checkAuth) > 1 {
		response := models.BaseResponse{Status: false, Message: "Data msisdn or username already exist"} 
		return response
	} 

	// check prefix msidn
	if !strings.HasPrefix(tempAuth.Msisdn, "62") {
		response := models.BaseResponse{Status: false, Message: "Msisdn should start with 62"} 
		return response
	}

	// fill data
	tempAuth.Id = uuid.New().String()
	tempAuth.CreatedAt = time.Now()
	bytes, _ := bcrypt.GenerateFromPassword([]byte(tempAuth.Password), bcrypt.DefaultCost) 
	tempAuth.Password = string(bytes)
  
	// query
	if err := db.Create(&tempAuth).Error; err != nil {
		response := models.BaseResponse{Status: false, Message: "Internal server error: " + err.Error()} 
		return response
	}
	
	// remove password
	tempAuth.Password = ""
 
	// response
	response := models.BaseResponse{Status: true, Message: "Create data success", Data: tempAuth}
	return response
}
 
// AuthLogin godoc
// @Summary      Post login payload
// @Description  Post login payload
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        auth body  models.AuthLogin  true "register"
// @Success      200  {object}  models.BaseResponse
// @Failure      400  {object}  models.BaseResponse
// @Failure      404  {object}  models.BaseResponse
// @Failure      500  {object}  models.BaseResponse
// @Router       /auth/login [post]
func AuthLogin(c *gin.Context) models.BaseResponse  {
	db := database.DB_POSTGRES

	err := godotenv.Load()
	if err != nil { 
	  fmt.Println(err)
	}  
	secretKey := []byte(os.Getenv("SECRET_KEY")) 
	  
	result := make(map[string]any) 
	var tempAuth models.Auth

	// prepare
	if err := c.BindJSON(&tempAuth); err != nil { 
		response := models.BaseResponse{Status: false, Message: "Internal server error, " + "Invalid request payload"} 
		return response
	}  

	// check if key exist
	if tempAuth.Msisdn == "" || tempAuth.Password == "" { 
		response := models.BaseResponse{Status: false, Message: "Required msisdn, password"} 
		return response
	}

	// check data 
	var checkAuth models.Auth
	query := db.Model(&models.Auth{}) 
	query.Where("msisdn = ?", tempAuth.Msisdn) 
	if err := query.Find(&checkAuth).Limit(1).Scan(&checkAuth).Error; err != nil { 
		response := models.BaseResponse{Status: false, Message: "Internal server error: " + err.Error()} 
		return response
	}
	
	// check data
	if checkAuth.Id == "" { 
		response := models.BaseResponse{Status: false, Message: "Login failed, msisdn not found"} 
		return response
	} 
   
	// check password  
	err = bcrypt.CompareHashAndPassword([]byte(checkAuth.Password), []byte(tempAuth.Password))
	if err != nil { 
		response := models.BaseResponse{Status: false, Message: "Login failed, password not match"} 
		return response
	}

	// create jwt 
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
			"id": checkAuth.Id, 
			"username": checkAuth.Username, 
			"msisdn": checkAuth.Msisdn, 
			"name": checkAuth.Name, 
			"exp": time.Now().Add(time.Hour * 24).Unix(), 
        })

	checkAuth.Password = "" 
    tokenString, _ := token.SignedString(secretKey)
	result["payload"] = checkAuth
	result["token"] = tokenString
	
	// response
	response := models.BaseResponse{Status: true, Message: "Login", Data: result}
	return response
}
 
 
// AuthValidate godoc
// @Summary      Get payload data auth
// @Description  Get payload data auth
// @Tags         Auth
// @Accept       json
// @Produce      json 
// @Success      200  {object}  models.BaseResponse
// @Failure      400  {object}  models.BaseResponse
// @Failure      404  {object}  models.BaseResponse
// @Failure      500  {object}  models.BaseResponse
// @Router       /auth/validate [get]
// @Security Bearer
func AuthValidate(c *gin.Context) models.BaseResponse  {
	err := godotenv.Load()
	if err != nil { 
		fmt.Println(err)
	} 
	secretKey := []byte(os.Getenv("SECRET_KEY")) 
	
	var message string
	header := c.GetHeader("Authorization")
	header_split := strings.Split(header, " ") 
  
	tokenString := header_split[1] 

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil 
	}) 
	
	if err != nil {
		message = "Internal error, " + err.Error()
		response := models.BaseResponse{Status: false, Message: message} 
		return response
	}
	
	if !token.Valid {
		message = "Internal error, Token not valid."
		response := models.BaseResponse{Status: false, Message: message} 
		return response
	} 
 
	// response
	response := models.BaseResponse{Status: true, Message: "Token valid", Data: token.Claims}
	return response
}
 