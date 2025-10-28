package controllers

import (
	"api_ticketing_web/config"
	"api_ticketing_web/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertCategory(c *gin.Context){
	var category model.CategoryForm

	if err:= c.ShouldBind(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	if err := config.DB.Create(&category);err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err})
	}

	c.JSON(http.StatusCreated,gin.H{"success":true})
}

func GetCategory(c *gin.Context){
	var category []model.Category

	if err:= config.DB.Find(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err})
	}

	c.JSON(http.StatusOK, gin.H{"success":true , "data":category})
}

func DeleteCategory(c *gin.Context){

}

