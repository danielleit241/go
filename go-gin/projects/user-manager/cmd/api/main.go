package main

import (
	"github.com/danielleit241/internal/config"
	"github.com/danielleit241/internal/handler"
	"github.com/danielleit241/internal/repository"
	"github.com/danielleit241/internal/routers"
	"github.com/danielleit241/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewConfig()

	userRepo := repository.NewInMemoryUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	userRoutes := routers.NewUserRoutes(userHandler)

	r := gin.Default()

	routers.RegisterRoutes(r, userRoutes)

	if err := r.Run(cfg.ServerPort); err != nil {
		panic(err)
	}
}
