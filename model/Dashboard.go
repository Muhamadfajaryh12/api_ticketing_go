package model

type DashboardResponse struct {
	TotalTicket     int `json:"total_ticket"`
	TotalOpen       int `json:"total_open"`
	TotalInProgress int `json:"total_in_progress"`
	TotalClose      int `json:"total_close"`
}