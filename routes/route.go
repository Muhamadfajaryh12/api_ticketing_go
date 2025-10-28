package routes

import (
	"api_ticketing_web/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
	r:=gin.Default()
	api := r.Group("/api")

	CategoryRoute := api.Group("/category")
	CategoryRoute.POST("/",controllers.InsertCategory)
	CategoryRoute.GET("/",controllers.GetCategory)

	PriorityRoute := api.Group("/priority")
	PriorityRoute.POST("/",controllers.InsertPriority)
	PriorityRoute.GET("/",controllers.GetPriority)
	return r
}