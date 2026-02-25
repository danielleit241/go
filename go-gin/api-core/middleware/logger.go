package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"
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
		contentType := c.GetHeader("Content-Type")
		requestBody := make(map[string]any)
		var formFiles []map[string]any

		if strings.HasPrefix(contentType, "multipart/form-data") {
			if err := c.Request.ParseMultipartForm(32 << 20); err == nil && c.Request.MultipartForm != nil {
				for key, vals := range c.Request.MultipartForm.Value {
					if len(vals) == 1 {
						requestBody[key] = vals[0]
					} else {
						requestBody[key] = vals
					}
				}
			}
			for field, files := range c.Request.MultipartForm.File {
				for _, fileHeader := range files {
					formFiles = append(formFiles, map[string]any{
						"field":        field,
						"filename":     fileHeader.Filename,
						"size":         formatFileSize(fileHeader.Size),
						"content_type": fileHeader.Header.Get("Content-Type"),
					})
				}
			}

			if len(formFiles) > 0 {
				requestBody["files"] = formFiles
			}

		} else {

			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err != nil {
				logger.Error().Err(err).Msg("Failed to read request body")
			}

			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			if strings.HasPrefix(contentType, "application/json") {
				_ = json.Unmarshal(bodyBytes, &requestBody)
			} else {
				values, _ := url.ParseQuery(string(bodyBytes))
				for key, vals := range values {
					if len(vals) == 1 {
						requestBody[key] = vals[0]
					} else {
						requestBody[key] = vals
					}
				}
			}
		}

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
			Interface("request_body", requestBody).
			Int64("duration_ms", duration.Milliseconds()).
			Msg("HTTP request")
	}
}

func formatFileSize(size int64) string {
	const (
		KB = 1 << (10 * 1)
		MB = 1 << (10 * 2)
		GB = 1 << (10 * 3)
	)
	switch {
	case size >= GB:
		return fmt.Sprintf("%.2f GB", float64(size)/GB)
	case size >= MB:
		return fmt.Sprintf("%.2f MB", float64(size)/MB)
	case size >= KB:
		return fmt.Sprintf("%.2f KB", float64(size)/KB)
	default:
		return fmt.Sprintf("%d B", size)
	}
}
