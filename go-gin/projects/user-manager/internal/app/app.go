package app

import (
	"github.com/danielleit241/internal/config"
	routers "github.com/danielleit241/internal/routes"
	"github.com/danielleit241/internal/validation"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Module interface {
	Routes() routers.Route
}

type Application struct {
	config  *config.Config
	router  *gin.Engine
	modules []Module
}

func NewApplication(config *config.Config) *Application {
	r := gin.Default()

	validation.Initialize()

	loadEnv()

	modules := []Module{
		NewUserModule(),
	}

	routers.RegisterRoutes(r, getModuleRoutes(modules)...)

	return &Application{
		config:  config,
		router:  r,
		modules: modules,
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

func loadEnv() {
	if err := godotenv.Load("../../.env"); err != nil {
		panic("Failed to load environment variables: " + err.Error())
	}
}
