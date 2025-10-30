package model

import "time"

type Ticket struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
	AssignedID  *uint   `json:"assigned_id"`
	CategoryID  uint   `json:"category_id"`
	PriorityID  uint   `json:"priority_id"`
	StatusID    uint   `json:"status_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TicketForm struct {
	Title       string `form:"title"`
	Description string `form:"description"`
	CategoryID  uint   `form:"category_id"`
	PriorityID  uint   `form:"priority_id"`
	StatusID    uint   `form:"status_id"`
	AssignedID *uint `form:"assigned_id"`
}

type TicketResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	User        string    `json:"user"`       
	Assigned    *string   `json:"assigned"`   
	Category    string    `json:"category"`
	Priority    string    `json:"priority"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}