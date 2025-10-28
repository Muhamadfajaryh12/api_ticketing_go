package model

import "time"

type Category struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CategoryForm struct {
	Category string `form:"category"`
}