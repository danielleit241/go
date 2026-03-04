package main

import (
	"github.com/danielleit241/internal/app"
	"github.com/danielleit241/internal/config"
)

func main() {
	cfg := config.NewConfig()

	application := app.NewApplication(cfg)

	if err := application.Run(); err != nil {
		panic(err)
	}
}
