package model

type DashboardResponse struct {
	TotalTicket     int `json:"total_ticket"`
	TotalOpen       int `json:"total_open"`
	TotalInProgress int `json:"total_in_progress"`
	TotalClose      int `json:"total_close"`
}
type TicketMonthResponse struct {
	Date        string `json:"date"`
	TotalTicket int    `json:"total_ticket"`
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

type DashboardTotalUsers struct {
	TotalTechnician int `json:"total_technician"`
	TotalGeneral    int `json:"total_general"`
}
type PerfomanceResponse struct {
	AvgInProgress  int              `json:"avg_in_progress"`
	AvgResolved    int              `json:"avg_resolved"`
	TeknisiTicket  []TeknisiTicket  `json:"teknisi_ticket"`
	TeknisiReview  []TeknisiReview  `json:"teknisi_review"`
	TeknisiAvgTime []TeknisiAvgTime `json:"teknisi_avg_time"`
}

type TeknisiTicket struct {
	Name        string `json:"name"`
	TotalTicket int    `json:"total_ticket"`
}

type TeknisiReview struct {
	Name      string `json:"name"`
	AvgReview int    `json:"avg_review"`
}

type TeknisiAvgTime struct {
	Name        string `json:"name"`
	AvgResponse string `json:"avg_response"`
	AvgResolved string `json:"avg_resolved"`
}