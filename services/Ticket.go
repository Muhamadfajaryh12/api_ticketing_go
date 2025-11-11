package services

import (
	"api_ticketing_web/config"
	"api_ticketing_web/model"
)

func GetDetailTicket(id uint) (model.TicketDetailResponse, error) {
	var ticket model.TicketResponse
	queryTicket := `
	SELECT 
		tickets.id,
		tickets.title,
		tickets.description,
		user.name AS user,
		assigned.name AS assigned,
		categories.category,
		priorities.priority,
		statuses.status,
		tickets.category_id, 
		tickets.priority_id,
		tickets.assigned_id,
		tickets.status_id
	FROM tickets
	LEFT JOIN users AS user ON tickets.user_id = user.id
	LEFT JOIN users AS assigned ON tickets.assigned_id = assigned.id
	LEFT JOIN categories ON tickets.category_id = categories.id
	LEFT JOIN priorities ON tickets.priority_id = priorities.id
	LEFT JOIN statuses ON tickets.status_id = statuses.id
	WHERE tickets.id = ?;
	`

	if err := config.DB.Raw(queryTicket, id).Scan(&ticket).Error;err != nil{
		return model.TicketDetailResponse{} , err
	}

	var ticketLog []model.TicketLogResponse
	queryLog := `
	SELECT 
	ticket_logs.id,
	statuses.status, ticket_logs.status_at
	FROM ticket_logs 
	JOIN statuses ON ticket_logs.status_id = statuses.id
	WHERE ticket_logs.ticket_id = ?
	ORDER BY ticket_logs.id ASC
	`
	if err := config.DB.Raw(queryLog, id).Scan(&ticketLog).Error;err != nil{
		return model.TicketDetailResponse{} , err
	}

	response := model.TicketDetailResponse{
		ID:          ticket.ID,
		Title:       ticket.Title,
		Description: ticket.Description,
		User:        ticket.User,
		Assigned:    ticket.Assigned,
		Category:    ticket.Category,
		CategoryID:  uint(ticket.CategoryID),
		Priority:    ticket.Priority,
		PriorityID:  uint(ticket.PriorityID),
		AssignedID:  uint(ticket.AssignedID),
		Status:      ticket.Status,
		StatusID: 	 ticket.StatusID,
		CreatedAt:   ticket.CreatedAt,
		UpdatedAt:   ticket.UpdatedAt,
		TicketLog:   ticketLog,
	}
	return response, nil
}