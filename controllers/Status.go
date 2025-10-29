package controllers

import (
	"api_ticketing_web/config"
	"api_ticketing_web/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertStatus(c *gin.Context) {
	var statusForm model.StatusForm
	var status model.Status

	if err := c.ShouldBind(&statusForm); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	status.Status = statusForm.Status
	if err := config.DB.Create(&status).Error; err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated,gin.H{
		"status":"success",
		"message":"Berhasil membuat Status",
	})
}

func GetStatus(c *gin.Context){
	var status []model.Status

	if err := config.DB.Find(&status).Error; err != nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusCreated,gin.H{
		"status":"success",
		"data":status,
	})
}