package main

import (
	"database/sql"
	"log"

	"github.com/danielleit241/internal/config"
	"github.com/danielleit241/internal/db"
	"github.com/danielleit241/internal/handler"
	"github.com/danielleit241/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found, using environment variables")
	}

	config := config.NewConfig()

	DB, err := db.Init(config)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer DB.Close()

	r := gin.Default()

	userRepo := repository.NewUserRepository(DB)
	userHandler := handler.NewUserHandler(userRepo)

	r.GET("/api/v1/users/:id", userHandler.GetUser)
	r.POST("/api/v1/users", userHandler.CreateUser)

	r.Run("127.0.0.1:8080")
}
