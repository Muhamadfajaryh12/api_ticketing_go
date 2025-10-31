package controllers

import (
	"api_ticketing_web/config"
	"api_ticketing_web/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertTicketLog( ticketID uint, statusID uint) error {

	ticketLog := model.TicketLog{
		TicketID: ticketID,
		StatusID: statusID,
	}
	
	if err := config.DB.Create(&ticketLog).Error; err != nil{
		return err
	}
	
	return nil
}

func GetTicketLog(c *gin.Context){
	var ticket []model.Ticket

	if err := config.DB.Find(&ticket).Error;err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"status":"error","message":err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status":"success","data":ticket})
}