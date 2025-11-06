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

	c.JSON(http.StatusOK,gin.H{"status":"success","data":data})
}