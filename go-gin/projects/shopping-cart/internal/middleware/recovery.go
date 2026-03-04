package middleware

import (
	"bytes"
	"fmt"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func Recovery(recoveryLogger *zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				stack := debug.Stack()

				recoveryLogger.Error().
					Str("path", ctx.Request.URL.Path).
					Str("method", ctx.Request.Method).
					Str("client_ip", ctx.ClientIP()).
					Str("panic", fmt.Sprintf("%v", err)).
					Str("stack", ExtractFirstAppStackLine(stack)).
					Msg("Panic recovered in middleware")

				ctx.AbortWithStatusJSON(500, gin.H{"error": fmt.Sprintf("%v", err)})
			}
		}()
		ctx.Next()
	}
}

func ExtractFirstAppStackLine(stack []byte) string {
	lines := bytes.Split(stack, []byte{'\n'})
	for _, line := range lines {
		if bytes.Contains(line, []byte(".go:")) &&
			!bytes.Contains(line, []byte("/runtime/")) &&
			!bytes.Contains(line, []byte("/testing/")) &&
			!bytes.Contains(line, []byte("/debug/")) &&
			!bytes.Contains(line, []byte("/vendor/")) &&
			!bytes.Contains(line, []byte("/internal/middleware/recovery.go")) {
			cleanLine := bytes.TrimSpace(line)
			return string(cleanLine)
		}
	}
	return ""
}
