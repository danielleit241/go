package main

import (
	"log"

	"github.com/danielleit241/internal/config"
	"github.com/danielleit241/internal/db"
	"github.com/danielleit241/internal/handler"
	"github.com/danielleit241/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	config := config.NewConfig()

	if err := db.InitDB(config); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	r := gin.Default()

	userRepo := repository.NewUserRepository()
	userHandler := handler.NewUserHandler(userRepo)

	r.GET("/api/v1/users", userHandler.GetUser)
	r.POST("/api/v1/users", userHandler.CreateUser)

	r.Run("127.0.0.1:8080")
}
