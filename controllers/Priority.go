package controllers

import (
	"api_ticketing_web/config"
	"api_ticketing_web/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertPriority(c *gin.Context){
	var priority  model.PriorityForm

	if err:= c.ShouldBind(&priority); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	if err := config.DB.Create(&priority); err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err})
		return
	}

	c.JSON(http.StatusCreated,gin.H{
		"success":true,
	})
}

func GetPriority(c *gin.Context){
	var priority []model.Priority

	if err := config.DB.Find(&priority); err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err})
		return
	}
	
	c.JSON(http.StatusOK,gin.H{
		"success":true,
		"data":priority,
	})

}