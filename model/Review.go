package model

type Review struct {
	ID       uint `gorm:"primaryKey" json:"id"`
	Rating   int  `json:"rating"`
	TicketID uint `json:"ticket_id"`
}

type ReviewForm struct {
	Rating   int  `form:"rating"`
	TicketID uint `form:"ticket_id"`
}

type ReviewResponse struct {
	ID       uint   `json:"id"`
	Rating   int    `json:"rating"`
	TicketID uint   `json:"ticket_id"`
	Name     string `json:"name"`
}