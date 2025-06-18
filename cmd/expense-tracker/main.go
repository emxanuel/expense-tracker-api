package main

import (
	"expense-tracker/api/api/routes"
	database_config "expense-tracker/api/internal/config"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env %s", err)
	}

	r := gin.Default()
	database_config.Connect()
	routes.RegisterRoutes(r)

	r.Run()
}
