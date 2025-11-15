package model

import "time"

type User struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	Email     string `gorm:"unique" json:"email"`
	Password  string `json:"-"`
	Role      string `json:"role"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type UserForm struct{
	Name string `form:"name"`
	Email string `form:"email"`
	Password string `form:"password"`
	Role string `form:"role"`
}

type LoginForm struct{
	Email string `form:"email"`
	Password string `form:"password"`
}

type UserResponse struct{
	ID uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Role string `json:"role"`
}

type TeknisiResponse struct{
	ID uint `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
	Status string `json:"status"`
}