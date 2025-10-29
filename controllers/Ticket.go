package controllers

import (
	"api_ticketing_web/config"
	"api_ticketing_web/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertTicket(c *gin.Context) {
	userID, exist := c.Get("user_ID")

	if !exist {
		c.JSON(http.StatusUnauthorized,gin.H{"status":"error","message":"Unauthorized"})
		return
	}

	var input model.TicketForm
	if err := c.ShouldBind(&input); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"status":"error","message":err.Error()})
		return
	}

	var ticket model.Ticket

	ticket = model.Ticket{
		Title: input.Title,
		Description: input.Description,
		CategoryID: input.CategoryID,
		StatusID: input.StatusID,
		PriorityID: input.PriorityID,
		UserID: userID.(uint),
	}

	if err := config.DB.Create(&ticket).Error; err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"status":"error","message":err.Error()})
		return
	}

	if err := InsertTicketLog(ticket.ID, ticket.StatusID); err != nil {
       c.JSON(http.StatusInternalServerError, gin.H{"status":"error","message":err.Error()})
        return
	}

	c.JSON(http.StatusCreated,gin.H{"status":"success","message":"Berhasil membuat Ticket"})
}

func UpdateTicket(c *gin.Context){
	id := c.Params("id")

}