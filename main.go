package main

import (
	"api_ticketing_web/config"
	"api_ticketing_web/model"
	"api_ticketing_web/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r:=gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS","PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// r.Use(cors.Default())

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading env")
	}
	
	config.ConnectDatabase()
	config.DB.AutoMigrate(&model.User{},&model.Status{},&model.Priority{},&model.Category{},&model.Ticket{},&model.TicketLog{},&model.Review{})

	routes.SetupRoute(r) 
	
	r.Run("localhost:8081")
}