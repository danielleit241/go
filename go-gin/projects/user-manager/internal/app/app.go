package app

import (
	"github.com/danielleit241/internal/config"
	routers "github.com/danielleit241/internal/routes"
	"github.com/gin-gonic/gin"
)

type Module interface {
	Routes() routers.Route
}

type Application struct {
	config *config.Config
	router *gin.Engine
}

func NewApplication(config *config.Config) *Application {
	r := gin.Default()

	modules := []Module{
		NewUserModule(),
	}

	routers.RegisterRoutes(r, getModuleRoutes(modules)...)

	return &Application{
		config: config,
		router: r,
	}
}

func (app *Application) Run() error {
	return app.router.Run(app.config.ServerPort)
}

func getModuleRoutes(modules []Module) []routers.Route {
	var routes []routers.Route
	for _, module := range modules {
		routes = append(routes, module.Routes())
	}
	return routes
}
