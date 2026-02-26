package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logger() gin.HandlerFunc {
	logPath := filepath.Join("logs", "http.log")

	logger := zerolog.New(&lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    1, // megabytes
		MaxBackups: 5,
		MaxAge:     5,    //days
		Compress:   true, // disabled by default
		LocalTime:  true,
	}).With().Timestamp().Logger()

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

		customWriter := &CustomResponseWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBuffer(nil),
		}
		c.Writer = customWriter

		c.Next()

		duration := time.Since(start)

		statusCode := c.Writer.Status()

		responseContentType := c.Writer.Header().Get("Content-Type")
		responseBodyRaw := customWriter.body.String()
		var responseBody any

		if strings.HasPrefix(responseContentType, "image/") || strings.HasPrefix(responseContentType, "application/octet-stream") {
			responseBody = fmt.Sprintf("[Binary data of type %s, size %s]", responseContentType, formatFileSize(int64(customWriter.body.Len())))
		} else if strings.HasPrefix(responseContentType, "application/json") ||
			strings.HasPrefix(responseContentType, "text/") ||
			strings.HasPrefix(strings.TrimSpace(responseBodyRaw), "{") ||
			strings.HasPrefix(strings.TrimSpace(responseBodyRaw), "[") {
			if err := json.Unmarshal(customWriter.body.Bytes(), &responseBody); err != nil {
				responseBody = responseBodyRaw
			}
		} else {
			responseBody = responseBodyRaw
			if len(responseBodyRaw) > 1000 {
				responseBody = fmt.Sprintf("%s... [truncated, total %d bytes]", responseBodyRaw[:1000], len(responseBodyRaw))
			}
		}

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
			Interface("response_body", responseBody).
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
