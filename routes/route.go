package routes

import (
	"api_ticketing_web/controllers"
	"api_ticketing_web/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoute(r *gin.Engine) {

	api := r.Group("/api")

	categoryRoute := api.Group("/category",middleware.Authorization())
	{
		categoryRoute.POST("", controllers.InsertCategory)
		categoryRoute.GET("", controllers.GetCategory)
	}

	priorityRoute := api.Group("/priority",middleware.Authorization())
	{
		priorityRoute.POST("", controllers.InsertPriority)
		priorityRoute.GET("", controllers.GetPriority)
	}

	statusRoute := api.Group("/status",middleware.Authorization())
	{
		statusRoute.POST("", controllers.InsertStatus)
		statusRoute.GET("", controllers.GetStatus)
	}

	userRoute := api.Group("/user")
	{
		userRoute.POST("/register", controllers.Register)
		userRoute.POST("/login", controllers.Login)
		userRoute.GET("/technician",controllers.GetTeknisi)
		userRoute.GET("/general",controllers.GetGeneral)
		userRoute.DELETE("/:id",controllers.DeleteUser)
	}

	ticketRoute := api.Group("/ticket",middleware.Authorization())
	{
		ticketRoute.GET("", controllers.GetTicket)
		ticketRoute.GET("/:id",controllers.GetDetailTicket)
		ticketRoute.POST("", controllers.InsertTicket)
		ticketRoute.PATCH("/:id", controllers.UpdateTicket)
	}

	ticketLogRoute := api.Group("/ticket-log",middleware.Authorization())
	{
		ticketLogRoute.GET("", controllers.GetTicketLog)
	}

	reviewRoute := api.Group("/review",middleware.Authorization())
	{
		reviewRoute.POST("/:id", controllers.InsertReview)
		reviewRoute.GET("",controllers.GetReview)
	}

	dashboardRoute := api.Group("/dashboard",middleware.Authorization())
	{
		dashboardRoute.GET("",controllers.GetDashhoard)
		dashboardRoute.GET("/performance",controllers.GetPerformance)
	}

}