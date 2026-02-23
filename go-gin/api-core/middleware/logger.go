package middleware

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func Logger() gin.HandlerFunc {
	logPath := filepath.Join("logs", "http.log")

	if err := os.MkdirAll(filepath.Dir(logPath), os.ModePerm); err != nil {
		panic("Failed to create log directory: " + err.Error())
	}

	fileLog, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic("Failed to open log file: " + err.Error())
	}

	logger := zerolog.New(fileLog).With().Timestamp().Logger()

	return func(c *gin.Context) {

		start := time.Now()

		c.Next()

		duration := time.Since(start)

		statusCode := c.Writer.Status()

		logEvent := logger.Info()
		if statusCode >= 500 {
			logEvent = logger.Error()
		} else if statusCode >= 400 {
			logEvent = logger.Warn()
		}

		logEvent.Str("protocol", c.Request.Proto).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Str("query", c.Request.URL.RawQuery).
			Str("client_ip", c.ClientIP()).
			Str("user_agent", c.Request.UserAgent()).
			Str("referer", c.Request.Referer()).
			Int("status", statusCode).
			Str("remote_address", c.Request.RemoteAddr).
			Str("headers", fmt.Sprint(c.Request.Header)).
			Int64("duration_ms", duration.Milliseconds()).
			Msg("HTTP request")
	}
}
