package main

import (
	"github.com/danielleit241/internal/app"
	"github.com/danielleit241/internal/config"
	"github.com/danielleit241/internal/db"
)

func main() {
	app.LoadEnv()

	cfg := config.NewConfig()

	if err := db.InitDB(cfg); err != nil {
		panic(err)
	}

	application := app.NewApplication(cfg)

	if err := application.Run(); err != nil {
		panic(err)
	}
}
