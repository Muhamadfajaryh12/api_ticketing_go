package controllers

import (
	"api_ticketing_web/config"
	"api_ticketing_web/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertPriority(c *gin.Context){
	var priorityForm  model.PriorityForm
	var priority model.Priority

	if err:= c.ShouldBind(&priorityForm); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	priority.Priority = priorityForm.Priority
	if err := config.DB.Create(&priority).Error; err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusCreated,gin.H{
		"status":"success",
		"message":"Berhasil membuat Priority",
	})
}

func GetPriority(c *gin.Context){
	var priority []model.Priority

	if err := config.DB.Find(&priority).Error; err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	
	c.JSON(http.StatusOK,gin.H{
		"success":true,
		"data":priority,
	})

}