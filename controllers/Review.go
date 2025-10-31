package controllers

import (
	"api_ticketing_web/config"
	"api_ticketing_web/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InsertReview(c *gin.Context) {
	id := c.Param("id")
	idUint,err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid ticket ID"})
		return
	}
	
	var input model.ReviewForm

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"status":"error", "message":err.Error()})
		return
	}

	review:= model.Review{
		TicketID: uint(idUint),
		Rating: input.Rating,
	}

	if err := config.DB.Create(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status":"error","message":err.Error()})
		return
	}

	c.JSON(http.StatusCreated,gin.H{"status":"success","message":"Berhasil memberikan rating"})

}