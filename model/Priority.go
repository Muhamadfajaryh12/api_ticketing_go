package model

import "time"

type Priority struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Priority  string `json:"priority"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PriorityForm struct {
	Priority string `form:"priority"`
}