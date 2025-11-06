package routes

import (
	"api_ticketing_web/controllers"
	"api_ticketing_web/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoute(r *gin.Engine) {

	api := r.Group("/api")

	categoryRoute := api.Group("/category")
	{
		categoryRoute.POST("", controllers.InsertCategory)
		categoryRoute.GET("", controllers.GetCategory)
	}

	priorityRoute := api.Group("/priority")
	{
		priorityRoute.POST("", controllers.InsertPriority)
		priorityRoute.GET("", controllers.GetPriority)
	}

	statusRoute := api.Group("/status")
	{
		statusRoute.POST("", controllers.InsertStatus)
		statusRoute.GET("", controllers.GetStatus)
	}

	userRoute := api.Group("/user")
	{
		userRoute.POST("/register", controllers.Register)
		userRoute.POST("/login", controllers.Login)
		userRoute.GET("/teknisi",controllers.GetTeknisi)
	}

	ticketRoute := api.Group("/ticket",middleware.Authorization())
	{
		ticketRoute.GET("", controllers.GetTicket)
		ticketRoute.GET("/:id",controllers.GetDetailTicket)
		ticketRoute.POST("", controllers.InsertTicket)
		ticketRoute.PATCH("/:id", controllers.UpdateTicket)
	}

	ticketLogRoute := api.Group("/ticket-log")
	{
		ticketLogRoute.GET("", controllers.GetTicketLog)
	}

	reviewRoute := api.Group("/review",middleware.Authorization())
	{
		reviewRoute.POST("/:id", controllers.InsertReview)
	}

	dashboardRoute := api.Group("/dashboard")
	{
		dashboardRoute.GET("",controllers.GetDashhoard)
	}

}