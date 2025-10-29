package controllers

import (
	"api_ticketing_web/config"
	"api_ticketing_web/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertCategory(c *gin.Context){
	var categoryForm model.CategoryForm
	var category model.Category

	if err:= c.ShouldBind(&categoryForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	category.Category = categoryForm.Category
	if err := config.DB.Create(&category).Error;err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
	}

	c.JSON(http.StatusCreated,gin.H{"status":"success","message":"Berhasil membuat category"})
}

func GetCategory(c *gin.Context){
	var category []model.Category

	if err:= config.DB.Find(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"status":"success", "data":category})
}

func DeleteCategory(c *gin.Context){

}

