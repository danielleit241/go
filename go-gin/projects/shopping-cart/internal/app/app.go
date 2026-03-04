package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

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

	modules := []Module{
		NewUserModule(),
	}

	routers.RegisterRoutes(r, config, getModuleRoutes(modules)...)

	return &Application{
		config:  config,
		router:  r,
		modules: modules,
	}
}

func (app *Application) Run() error {
	serv := &http.Server{
		Addr:    app.config.ServerPort,
		Handler: app.router,
	}

	quitCh := make(chan os.Signal, 1)

	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	// syscall.SIGINT: Ctrl+C
	// syscall.SIGTERM: Kill
	// syscall.SIGHUP: Reload

	go func() {
		log.Printf("Server is running on %s", app.config.ServerPort)
		if err := serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-quitCh
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) // 15 seconds timeout for graceful shutdown
	defer cancel()
	if err := serv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exitted gracefully")
	return nil
}

func getModuleRoutes(modules []Module) []routers.Route {
	var routes []routers.Route
	for _, module := range modules {
		routes = append(routes, module.Routes())
	}
	return routes
}

func LoadEnv() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Printf("failed to get current working directory, skip loading .env: %v", err)
		return
	}

	dir := currentDir
	for {
		envPath := filepath.Join(dir, ".env")
		if _, statErr := os.Stat(envPath); statErr == nil {
			if loadErr := godotenv.Load(envPath); loadErr != nil {
				log.Printf("failed to load .env file at %s: %v", envPath, loadErr)
			}
			return
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	log.Println(".env file not found, using existing system environment variables")
}
