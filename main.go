package main

import (
	"api_ticketing_web/config"
	"api_ticketing_web/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading env")
	}
	
	config.ConnectDatabase()
	r := routes.SetupRoute()
	r.Run(":8080")
}