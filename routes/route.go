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

	StatusRoute := api.Group("/status")
	StatusRoute.POST("/",controllers.InsertStatus)
	StatusRoute.GET("/",controllers.GetStatus)
	
	UserRoute :=api.Group("/user")
	UserRoute.POST("/register",controllers.Register)
	UserRoute.POST("/login",controllers.Login)

	return r
}