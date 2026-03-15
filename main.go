package main

import (
	"blogapi/config"
	"blogapi/routes"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	// .env file load karo
	godotenv.Load()

	// Database connect karo
	config.ConnectDB()

	// Routes setup karo
	r := routes.SetupRoutes()

	fmt.Println("✅ Server chal raha hai: http://localhost:8080")
	r.Run(":8080")
} 
