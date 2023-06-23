package main

import (
	"ginredis/app/handler"
	"ginredis/db"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitPostgres()
	db.InitRedis() // Inisialisasi koneksi Redis

	container := handler.NewContainer()
	router := handler.SetupRouter(container)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server started on port %s", port)
	router.Run(":" + port)
}
