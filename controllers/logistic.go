package controllers

import (
	"akatgelar/orderfaz-testcase-be/database"
	"akatgelar/orderfaz-testcase-be/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// LogisticRead godoc
// @Summary      Get logistics data
// @Description  Get logistics data
// @Tags         Logistic
// @Accept       json
// @Produce      json
// @Param        logistic_name    query     boolean  false  "logistic_name=jne"
// @Param        origin_name      query     boolean  false  "origin_name=jakarta"
// @Param        destination_name      query     boolean  false  "destination_name=bandung"
// @Success      200  {object}  models.BaseResponse
// @Failure      400  {object}  models.BaseResponse
// @Failure      404  {object}  models.BaseResponse
// @Failure      500  {object}  models.BaseResponse
// @Router       /logistic [get]
// @Security Token
func LogisticRead(c *gin.Context, logistic_name string, origin_name string, destination_name string) models.BaseResponse {
	db := database.DB_POSTGRES
	   
	// default value  
	var result []models.Logistics

	// select
	query := db.Model(&models.Logistics{}) 
	if logistic_name != "ALL" {
		query.Where("logistic_name ilike ?", logistic_name)
	}
	if origin_name != "ALL" {
		query.Where("origin_name ilike ?", origin_name)
	}
	if destination_name != "ALL" {
		query.Where("destination_name ilike ?", destination_name)
	}

	if err := query.Find(&result).Scan(&result).Error; err != nil { 
		response := models.BaseResponse{Status: false, Message: "Internal server error: " + err.Error()} 
		return response
	}
 
	// response
	response := models.BaseResponse{Status: true, Message: "Get data success", Data: result}
	return response
}
 
 
// LogisticCreate godoc
// @Summary      Create logistics data
// @Description  Create logistics data
// @Tags         Logistic
// @Accept       json
// @Produce      json
// @Param        auth body  models.LogisticsCreate  true "logistic"
// @Success      200  {object}  models.BaseResponse
// @Failure      400  {object}  models.BaseResponse
// @Failure      404  {object}  models.BaseResponse
// @Failure      500  {object}  models.BaseResponse
// @Router       /logistic [post] 
// @Security Token
func LogisticCreate(c *gin.Context) models.BaseResponse {
	db := database.DB_POSTGRES
	  
	var tempLogistic models.Logistics
	tempLogistic.Id = uuid.New().String()
	tempLogistic.CreatedAt = time.Now()

	// prepare
	if err := c.BindJSON(&tempLogistic); err != nil { 
		response := models.BaseResponse{Status: false, Message: "Internal server error, " + "Invalid request payload"} 
		return response
	}

	// query
	if err := db.Create(&tempLogistic).Error; err != nil {
		response := models.BaseResponse{Status: false, Message: "Internal server error: " + err.Error()} 
		return response
	}
 
	// response
	response := models.BaseResponse{Status: true, Message: "Create data success", Data: tempLogistic}
	return response
}
 
 