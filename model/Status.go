package model

import "time"

type Status struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Status    string `json:"status"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type StatusForm struct{
	Status string `form:"status"`
}