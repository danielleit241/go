package routers

import (
	"github.com/danielleit241/internal/config"
	"github.com/danielleit241/internal/middleware"
	"github.com/gin-gonic/gin"
)

type Route interface {
	Register(router *gin.RouterGroup)
}

func RegisterRoutes(r *gin.Engine, routes ...Route) {
	r.Use(
		middleware.Logger(),
		middleware.ApiKeyMiddleware(),
		middleware.AuthMiddleware(),
		middleware.RateLimit(),
	)

	apiGroup := r.Group(config.NewConfig().ApiPrefix)
	for _, route := range routes {
		route.Register(apiGroup)
	}
}
