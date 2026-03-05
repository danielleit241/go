package routers

import (
	"github.com/danielleit241/internal/config"
	"github.com/danielleit241/internal/middleware"
	"github.com/danielleit241/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Route interface {
	Register(router *gin.RouterGroup)
}

func RegisterRoutes(r *gin.Engine, conf *config.Config, routes ...Route) {
	rateLimitLogger := logger.NewWithPath("internal/logs/ratelimit.log", "warn", conf.Environment)
	httpLogger := logger.NewWithPath("internal/logs/http.log", "info", conf.Environment)
	recoveryLogger := logger.NewWithPath("internal/logs/recovery.log", "error", conf.Environment)

	r.Use(
		middleware.RateLimit(rateLimitLogger),
		middleware.Logger(httpLogger),
		middleware.Recovery(recoveryLogger),
		middleware.ApiKeyMiddleware(),
		middleware.AuthMiddleware(),
	)

	apiGroup := r.Group(conf.ApiPrefix)
	for _, route := range routes {
		route.Register(apiGroup)
	}
}
