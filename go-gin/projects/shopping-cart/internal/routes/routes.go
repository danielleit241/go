package routers

import (
	"github.com/danielleit241/internal/config"
	"github.com/danielleit241/internal/middleware"
	"github.com/danielleit241/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Route interface {
	Register(router *gin.RouterGroup)
}

func RegisterRoutes(r *gin.Engine, conf *config.Config, routes ...Route) {
	rateLimitLogger := newLoggerWithPath("internal/logs/ratelimit.log", "warn", conf.Environment)
	httpLogger := newLoggerWithPath("internal/logs/http.log", "info", conf.Environment)
	recoveryLogger := newLoggerWithPath("internal/logs/recovery.log", "error", conf.Environment)

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

func newLoggerWithPath(logPath string, level string, env string) *zerolog.Logger {
	config := logger.LoggerConfig{
		FileName:    logPath,
		MaxSize:     1,
		MaxBackups:  5,
		MaxAge:      5,
		Compress:    true,
		Level:       level,
		Environment: env,
	}
	return logger.NewLogger(config)
}
