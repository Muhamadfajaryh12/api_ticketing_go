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
	var ticketLog []model.TicketLogResponseList

	query := `
	SELECT 
	ticket_logs.id,
	ticket_logs.ticket_id,
	ticket_logs.status_at,
	statuses.status,
	users.name
	FROM ticket_logs
	INNER JOIN tickets ON ticket_logs.ticket_id = tickets.id
	INNER JOIN users ON tickets.assigned_id = users.id
	INNER JOIN statuses ON ticket_logs.status_id = statuses.id
	ORDER BY ticket_logs.id DESC
	`
	if err := config.DB.Raw(query).Scan(&ticketLog).Error;err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"status":"error","message":err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status":"success","data":ticketLog})
}