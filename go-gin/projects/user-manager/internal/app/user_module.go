package app

import (
	"github.com/danielleit241/internal/handler"
	"github.com/danielleit241/internal/repository"
	routers "github.com/danielleit241/internal/routes"
	"github.com/danielleit241/internal/service"
)

type UserModule struct {
	routes routers.Route
}

func NewUserModule() *UserModule {
	userRepo := repository.NewInMemoryUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	userRoutes := routers.NewUserRoutes(userHandler)

	return &UserModule{
		routes: userRoutes,
	}
}

func (m *UserModule) Routes() routers.Route {
	return m.routes
}
