package controllers

import (
	"api_ticketing_web/config"
	"api_ticketing_web/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertStatus(c *gin.Context) {
	var status model.StatusForm

	if err := c.ShouldBind(&status); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	if err := config.DB.Create(&status); err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err})
		return
	}
	
	c.JSON(http.StatusCreated,gin.H{
		"status":true,
	})
}

func GetStatus(c *gin.Context){
	var status []model.Status

	if err := config.DB.Find(&status); err != nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":err})
		return
	}

	c.JSON(http.StatusCreated,gin.H{
		"status":true,
		"data":status,
	})
}