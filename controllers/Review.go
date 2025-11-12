package controllers

import (
	"api_ticketing_web/config"
	"api_ticketing_web/model"
	"api_ticketing_web/services"
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

	if err := config.DB.Model(&model.Ticket{}).
		Where("id = ?", review.TicketID).
		Update("status_id", 5).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	if err:= InsertTicketLog(review.TicketID, 5); err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"status":"error","message":err})
		return
	}


	ticketDetail, err := services.GetDetailTicket(uint(idUint));
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"status":"error","message":err})
		return
	}
	c.JSON(http.StatusCreated,gin.H{"status":"success","message":"Berhasil memberikan rating","data":ticketDetail})
}

func GetReview(c *gin.Context){
	var review []model.ReviewResponse

	query := `
	SELECT
	reviews.id, tickets.id as ticket_id, reviews.rating, users.name
	FROM reviews
	JOIN tickets ON reviews.ticket_id = tickets.id
	JOIN users ON tickets.user_id = users.id
	ORDER BY reviews.id DESC
	`

	if err:= config.DB.Raw(query).Scan(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status":"error","message":err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status":"success","data":review})
}