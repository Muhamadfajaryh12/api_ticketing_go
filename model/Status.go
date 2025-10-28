package model

import "time"

type Status struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Status  string `json:"priority"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type StatusForm struct{
	Status string `form:"status"`
}