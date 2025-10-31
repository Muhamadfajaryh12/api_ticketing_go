package routes

import (
	"api_ticketing_web/controllers"
	"api_ticketing_web/middleware"

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

	
	TicketRoute := api.Group("/ticket",middleware.Authorization())
	TicketRoute.GET("/",controllers.GetTicket)
	TicketRoute.POST("/",controllers.InsertTicket)
	TicketRoute.PATCH("/:id",controllers.UpdateTicket)

	TicketLogRoute := api.Group("/ticket-log",middleware.Authorization())
	TicketLogRoute.GET("/",controllers.GetTicketLog)

	ReviewRoute := api.Group("/review",middleware.Authorization())
	ReviewRoute.POST("/",controllers.InsertReview)
	return r
}