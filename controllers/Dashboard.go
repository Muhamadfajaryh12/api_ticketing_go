package controllers

import (
	"api_ticketing_web/config"
	"api_ticketing_web/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDashhoard(c *gin.Context) {
	var data  model.DashboardResponse

	query := `
	SELECT 
	COUNT(id) as total_ticket,
	SUM(CASE WHEN status_id = 1 THEN 1 ELSE 0 END) AS total_open,
	SUM(CASE WHEN status_id = 2 THEN 1 ELSE 0 END) AS total_in_progress,
	SUM(CASE WHEN status_id = 5 THEN 1 ELSE 0 END) AS total_close
	FROM tickets
	`

	if err := config.DB.Raw(query).Scan(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status":"error","message":err.Error() })
		return
	}

	var CategoryPriority []model.DashboardCategoryPriorityResponse
	queryGetCategoryPriority := `
		SELECT 
			categories.category,
			priorities.priority,
			COUNT(tickets.id) AS total_ticket
		FROM categories 
		CROSS JOIN priorities 
		LEFT JOIN tickets  
			ON tickets.category_id = categories.id AND tickets.priority_id = priorities.id
		GROUP BY categories.category, priorities.priority
		ORDER BY categories.category, priorities.priority;
	`

	if err := config.DB.Raw(queryGetCategoryPriority).Scan(&CategoryPriority).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status":"error","message":err.Error() })
		return
	}

	grouped := make(map[string][]map[string]interface{})
    for _, r := range CategoryPriority {
        grouped[r.Category] = append(grouped[r.Category], map[string]interface{}{
            "priority":        r.Priority,
            "total_ticket":  r.TotalTicket,
        })
    }

	var response []map[string]interface{}
    for category, priority := range grouped {
        response = append(response, map[string]interface{}{
            "category": category,
            "priority": priority,
        })
    }


	var categoryStatus []model.DashboardCategoryStatusResponse
	queryGetCategoryStatus := `
		SELECT 
			categories.category,
			statuses.status,
			COUNT(tickets.id) AS total_ticket
		FROM categories 
		CROSS JOIN statuses 
		LEFT JOIN tickets  
			ON tickets.category_id = categories.id AND tickets.status_id = statuses.id
		GROUP BY categories.category, statuses.status
		ORDER BY categories.category, statuses.status;
	`
	if err := config.DB.Raw(queryGetCategoryStatus).Scan(&categoryStatus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status":"error","message":err.Error() })
		return
	}
	
	groupedStatus := make(map[string][]map[string]interface{})
	 for _, r := range categoryStatus {
        groupedStatus[r.Category] = append(groupedStatus[r.Category], map[string]interface{}{
            "status":        r.Status,
            "total_ticket":  r.TotalTicket,
        })
    }


	var categoryStatusResponse []map[string]interface{}
    for category, status := range groupedStatus {
        categoryStatusResponse = append(categoryStatusResponse, map[string]interface{}{
            "category": category,
            "status": status,
        })
    }

	queryGetTicketMonth := `
		WITH RECURSIVE date_series AS (
		SELECT DATE_FORMAT(CURDATE(), '%Y-%m-01') AS date
		UNION ALL
		SELECT DATE_ADD(date, INTERVAL 1 DAY)
		FROM date_series
		WHERE DATE_ADD(date, INTERVAL 1 DAY) <= LAST_DAY(CURDATE())
		)
		SELECT 
		ds.date,
		COUNT(t.id) AS total_ticket
		FROM date_series ds
		LEFT JOIN tickets t 
		ON DATE(t.created_at) = ds.date
		GROUP BY ds.date
		ORDER BY ds.date;
	`

	var ticketMonth []model.TicketMonthResponse

	if err := config.DB.Raw(queryGetTicketMonth).Scan(&ticketMonth).Error; err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"status":"error","message":err.Error()})
		return
	}

	queryTotalUsers := `
	SELECT 
	COUNT(CASE WHEN role = "Technician" THEN 1 END) as total_technician,
	COUNT(CASE WHEN role = "General" THEN 1 END) as total_general
	FROM users
	WHERE is_delete = 0
	`
	var totalUsers model.DashboardTotalUsers
	if err := config.DB.Raw(queryTotalUsers).Scan(&totalUsers).Error; err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"status":"error","message":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"status":"success","data":gin.H{
		"summary_ticket":data,
		"summary_category_priority":response,
		"summary_category_status":categoryStatusResponse,
		"summary_ticket_month":ticketMonth,
		"summary_users":totalUsers,
	}})
}

func GetPerformance(c *gin.Context){
	var response  model.PerfomanceResponse

	queryGetAverageProgress := `
	SELECT 
	ROUND(AVG(TIMESTAMPDIFF(MINUTE,t1.status_at, t2.status_at))) as avg_in_progress
	FROM ticket_logs t1
	JOIN ticket_logs t2 ON t1.ticket_id = t2.ticket_id
	WHERE t1.status_id = 1 AND t2.status_id = 2;
	`
	var avgInProgress int
	if err := config.DB.Raw(queryGetAverageProgress).Scan(&avgInProgress).Error;err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"status":"error", "message":err.Error()})
		return
	}
	response.AvgInProgress = avgInProgress

	queryGetAverageResolved := `
	SELECT 
	ROUND(AVG(TIMESTAMPDIFF(MINUTE,t1.status_at, t2.status_at))) as avg_resolved
	FROM ticket_logs t1
	JOIN ticket_logs t2 ON t1.ticket_id = t2.ticket_id
	WHERE t1.status_id = 2 AND t2.status_id = 3;
	`	
	var avgResolved int
	if err := config.DB.Raw(queryGetAverageResolved).Scan(&avgResolved).Error;err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"status":"error", "message":err.Error()})
		return
	}
	response.AvgResolved = avgResolved	

	queryGetTeknisiTicket := `
	SELECT 
	users.name,
	COUNT(tickets.id) AS total_ticket
	FROM users
	LEFT JOIN tickets ON users.id = tickets.assigned_id
	WHERE users.role = "Technician"
	GROUP BY users.id, users.name;`

	var teknisiTicket []model.TeknisiTicket
	if err := config.DB.Raw(queryGetTeknisiTicket).Scan(&teknisiTicket).Error;err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"status":"error", "message":err.Error()})
		return
	}
	response.TeknisiTicket = teknisiTicket

	queryGetTeknisiReview := `
	SELECT 
	users.name,
	ROUND(AVG(reviews.rating)) AS avg_review
	FROM users
	LEFT JOIN tickets ON users.id = tickets.assigned_id
	LEFT JOIN reviews ON tickets.id = reviews.ticket_id
	WHERE users.role = "Technician"
	GROUP BY users.id, users.name;`

	var TeknisiReview []model.TeknisiReview
	if err := config.DB.Raw(queryGetTeknisiReview).Scan(&TeknisiReview).Error;err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"status":"error", "message":err.Error()})
		return
	}

	response.TeknisiReview = TeknisiReview

	queryGetTeknisiAvgTime := `
	SELECT
	users.name,
	ROUND(AVG(TIMESTAMPDIFF(MINUTE,t_open.status_at, t_progress.status_at))) as avg_response,
	ROUND(AVG(TIMESTAMPDIFF(MINUTE,t_progress.status_at, t_resolved.status_at))) as avg_resolved
	from users
	LEFT JOIN tickets ON users.id = tickets.assigned_id
	LEFT JOIN ticket_logs t_open ON tickets.id = t_open.ticket_id AND t_open.status_id = 1
	LEFT JOIN ticket_logs t_progress ON tickets.id = t_progress.ticket_id AND t_progress.status_id = 2
	LEFT JOIN ticket_logs t_resolved ON tickets.id = t_resolved.ticket_id AND t_resolved.status_id = 3
 	WHERE users.role = "Technician"
	GROUP BY users.id, users.name	
	`

	var teknisiAvgTime []model.TeknisiAvgTime

	if err:= config.DB.Raw(queryGetTeknisiAvgTime).Scan(&teknisiAvgTime).Error;err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status":"error", "message":err.Error()})
		return
	}

	response.TeknisiAvgTime = teknisiAvgTime

	c.JSON(http.StatusOK, gin.H{"status":"success","data":response})
}