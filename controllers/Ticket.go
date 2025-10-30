package controllers

import (
	"api_ticketing_web/config"
	"api_ticketing_web/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InsertTicket(c *gin.Context) {
	userID, exist := c.Get("user_id")

	if !exist {
		c.JSON(http.StatusUnauthorized,gin.H{"status":"error","message":"Unauthorized"})
		return
	}
	userIDFloat := userID.(float64)

	var input model.TicketForm
	if err := c.ShouldBind(&input); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"status":"error","message":err.Error()})
		return
	}

	ticket := model.Ticket{
		Title: input.Title,
		Description: input.Description,
		CategoryID: input.CategoryID,
		StatusID: input.StatusID,
		PriorityID: input.PriorityID,
		UserID: uint(userIDFloat),
	}

	if err := config.DB.Create(&ticket).Error; err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"status":"error","message":err.Error()})
		return
	}

	if err := InsertTicketLog(ticket.ID, ticket.StatusID); err != nil {
       c.JSON(http.StatusInternalServerError, gin.H{"status":"error","message":err})
        return
	}

	c.JSON(http.StatusCreated,gin.H{"status":"success","message":"Berhasil membuat Ticket"})
}

func GetTicket( c* gin.Context){
	var ticket []model.TicketResponse

	if err := config.DB.Raw(
		`SELECT 
		tickets.id,
		tickets.title,
		tickets.description,
		user.name as user, 
		assigned.name as assigned,
		categories.category,
		priorities.priority,
		statuses.status
		FROM tickets
		LEFT JOIN users as user ON tickets.user_id = user.id
		LEFT JOIN users as assigned ON tickets.assigned_id = assigned.id
		LEFT JOIN categories ON  tickets.category_id =  categories.id
		LEFT JOIN priorities ON tickets.priority_id =  priorities.id
		LEFT JOIN statuses ON  tickets.status_id= statuses.id
		`).Scan(&ticket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status":"error","message":err.Error()})
		return
	}

	c.JSON(http.StatusOK,gin.H{"status":"success","data":ticket})
}

func UpdateTicket(c *gin.Context){
	id := c.Param("id")
	var input model.TicketForm

	userIDRaw, isExist := c.Get("user_id")
	
	if !isExist {
		c.JSON(http.StatusUnauthorized,gin.H{"status":"error","message":"Unauthorized"})
		return
	}

	userID := uint(userIDRaw.(float64))

	role, isExist := c.Get("role")

	if !isExist {
		c.JSON(http.StatusUnauthorized,gin.H{"status":"error","message":"Unauthorized"})
		return
	}

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid ticket ID"})
		return
	}

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"status":"error","message":err})
		return
	}

	if role == "teknisi"{
		input.AssignedID = &userID
	}

	var ticket model.Ticket
	if err := config.DB.Model(&ticket).Where("id =  ?",idUint).Updates(input).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"status":"error","message":err.Error()})
		return
	}

	if err:= InsertTicketLog(uint(idUint), input.StatusID); err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"status":"error","message":err})
		return
	}

	c.JSON(http.StatusOK,gin.H{"status":"success","message":"Berhasil mengupdate ticket"})

}