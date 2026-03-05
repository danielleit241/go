package app

import (
	v1handler "github.com/danielleit241/internal/handler/v1"
	"github.com/danielleit241/internal/repository"
	routers "github.com/danielleit241/internal/routes"
	v1routers "github.com/danielleit241/internal/routes/v1"
	v1service "github.com/danielleit241/internal/service/v1"
)

type UserModule struct {
	routes routers.Route
}

func NewUserModule(ctx *ModuleContext) *UserModule {
	userRepo := repository.NewUserRepository(ctx.DB)
	userService := v1service.NewUserService(userRepo)
	userHandler := v1handler.NewUserHandler(userService)
	userRoutes := v1routers.NewUserRoutes(userHandler)

	return &UserModule{
		routes: userRoutes,
	}
}

func (m *UserModule) Routes() routers.Route {
	return m.routes
}
