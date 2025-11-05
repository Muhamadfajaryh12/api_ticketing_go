package model

import "time"

type TicketLog struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	TicketID uint      `json:"ticket_id"`
	StatusID uint `json:"status_id"`
	StatusAt time.Time `gorm:"autoCreateTime" json:"status_at"`
}

type TicketLogForm struct{
	TicketID uint `form:"ticket_id"`
	StatusID uint `form:"status_id"`
}

type TicketLogResponse struct{
	ID       uint      `json:"id"`
	Status string `json:"status"`
	StatusAt time.Time `json:"status_at"`
}

type TicketLogResponseList struct{
		ID       uint      `json:"id"`
		TicketID uint `json:"ticket_id"`
	Status string `json:"status"`
	StatusAt time.Time `json:"status_at"`
	Name string `json:"name"`
}
