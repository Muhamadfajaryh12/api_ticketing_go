package model

type DashboardResponse struct {
	TotalTicket     int `json:"total_ticket"`
	TotalOpen       int `json:"total_open"`
	TotalInProgress int `json:"total_in_progress"`
	TotalClose      int `json:"total_close"`
}

type DashboardCategoryPriorityResponse struct {
	Category    string `json:"category"`
	Priority    string `json:"priority"`
	TotalTicket int    `json:"total_ticket"`
}

type DashboardCategoryStatusResponse struct {
	Category    string `json:"category"`
	Status      string `json:"status"`
	TotalTicket int    `json:"total_ticket"`
}