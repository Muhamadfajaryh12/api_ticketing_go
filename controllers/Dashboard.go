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

	 for _, r := range categoryStatus {
        grouped[r.Category] = append(grouped[r.Category], map[string]interface{}{
            "status":        r.Status,
            "total_ticket":  r.TotalTicket,
        })
    }

	var categoryStatusResponse []map[string]interface{}
    for category, status := range grouped {
        response = append(response, map[string]interface{}{
            "category": category,
            "status": status,
        })
    }
	c.JSON(http.StatusOK,gin.H{"status":"success","data":gin.H{
		"summary_ticket":data,
		"summary_category_priority":response,
		"summary_category_status":categoryStatusResponse,
	}})
}