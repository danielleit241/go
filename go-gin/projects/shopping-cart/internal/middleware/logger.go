package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
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

type LogEntry struct {
	Protocol      string              `json:"protocol"`
	Method        string              `json:"method"`
	Path          string              `json:"path"`
	Query         string              `json:"query"`
	ClientIP      string              `json:"client_ip"`
	UserAgent     string              `json:"user_agent"`
	Referer       string              `json:"referer"`
	Status        int                 `json:"status"`
	RemoteAddress string              `json:"remote_address"`
	Headers       map[string][]string `json:"headers"`
	RequestBody   map[string]any      `json:"request_body,omitempty"`
	ResponseBody  any                 `json:"response_body,omitempty"`
	DurationMs    int64               `json:"duration_ms"`
	Timestamp     time.Time           `json:"timestamp"`
}

const (
	multipartMaxMemory   = 32 << 20
	responsePreviewLimit = 1000
)

func (w *CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func getLogger() zerolog.Logger {
	_, file, _, ok := runtime.Caller(0)
	baseDir := "."
	if ok {
		baseDir = filepath.Dir(file)
	}

	logDir := filepath.Clean(filepath.Join(baseDir, "..", "logs"))
	_ = os.MkdirAll(logDir, 0o755)
	logPath := filepath.Join(logDir, "http.log")
	return zerolog.New(&lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     5,
		Compress:   true,
		LocalTime:  true,
	}).With().Timestamp().Logger()
}

func Logger() gin.HandlerFunc {
	logger := getLogger()

	return func(c *gin.Context) {

		start := time.Now()

		requestBody := parseRequestBody(c, logger)

		customWriter := &CustomResponseWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBuffer(nil),
		}
		c.Writer = customWriter

		c.Next()

		duration := time.Since(start)

		statusCode := c.Writer.Status()

		responseBody := parseResponseBody(c.Writer.Header().Get("Content-Type"), customWriter.body.Bytes())
		logEntry := buildLogEntry(c, statusCode, duration, requestBody, responseBody)

		logEvent := getLoggerEvent(statusCode, logger)
		writeLogEvent(logEvent, logEntry)
	}
}

func buildLogEntry(c *gin.Context, statusCode int, duration time.Duration, requestBody map[string]any, responseBody any) LogEntry {
	return LogEntry{
		Protocol:      c.Request.Proto,
		Method:        c.Request.Method,
		Path:          c.Request.URL.Path,
		Query:         c.Request.URL.RawQuery,
		ClientIP:      c.ClientIP(),
		UserAgent:     c.Request.UserAgent(),
		Referer:       c.Request.Referer(),
		Status:        statusCode,
		RemoteAddress: c.Request.RemoteAddr,
		Headers:       c.Request.Header,
		RequestBody:   requestBody,
		ResponseBody:  responseBody,
		DurationMs:    duration.Milliseconds(),
		Timestamp:     time.Now(),
	}
}

func writeLogEvent(logEvent *zerolog.Event, entry LogEntry) {
	logEvent.Str("protocol", entry.Protocol).
		Str("method", entry.Method).
		Str("path", entry.Path).
		Str("query", entry.Query).
		Str("client_ip", entry.ClientIP).
		Str("user_agent", entry.UserAgent).
		Str("referer", entry.Referer).
		Int("status", entry.Status).
		Str("remote_address", entry.RemoteAddress).
		Interface("headers", entry.Headers).
		Interface("request_body", entry.RequestBody).
		Interface("response_body", entry.ResponseBody).
		Int64("duration_ms", entry.DurationMs).
		Time("timestamp", entry.Timestamp).
		Msg("HTTP request")
}

func getLoggerEvent(statusCode int, logger zerolog.Logger) *zerolog.Event {
	if statusCode >= 500 {
		return logger.Error()
	} else if statusCode >= 400 {
		return logger.Warn()
	} else {
		return logger.Info()
	}
}

func parseRequestBody(c *gin.Context, logger zerolog.Logger) map[string]any {
	contentType := c.GetHeader("Content-Type")
	requestBody := make(map[string]any)
	var formFiles []map[string]any

	if strings.HasPrefix(contentType, "multipart/form-data") {
		if err := c.Request.ParseMultipartForm(multipartMaxMemory); err != nil {
			logger.Error().Err(err).Msg("Failed to parse multipart form")
			return requestBody
		}

		if c.Request.MultipartForm != nil {
			for key, vals := range c.Request.MultipartForm.Value {
				if len(vals) == 1 {
					requestBody[key] = vals[0]
				} else {
					requestBody[key] = vals
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
		}

		if len(formFiles) > 0 {
			requestBody["files"] = formFiles
		}

		return requestBody
	}

	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to read request body")
		return requestBody
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if strings.HasPrefix(contentType, "application/json") {
		_ = json.Unmarshal(bodyBytes, &requestBody)
		return requestBody
	}

	values, _ := url.ParseQuery(string(bodyBytes))
	for key, vals := range values {
		if len(vals) == 1 {
			requestBody[key] = vals[0]
		} else {
			requestBody[key] = vals
		}
	}

	return requestBody
}

func parseResponseBody(contentType string, bodyBytes []byte) any {
	responseBodyRaw := string(bodyBytes)

	if strings.HasPrefix(contentType, "image/") || strings.HasPrefix(contentType, "application/octet-stream") {
		return fmt.Sprintf("[Binary data of type %s, size %s]", contentType, formatFileSize(int64(len(bodyBytes))))
	}

	trimmedBody := strings.TrimSpace(responseBodyRaw)
	if strings.HasPrefix(contentType, "application/json") ||
		strings.HasPrefix(contentType, "text/") ||
		strings.HasPrefix(trimmedBody, "{") ||
		strings.HasPrefix(trimmedBody, "[") {
		var parsedBody any
		if err := json.Unmarshal(bodyBytes, &parsedBody); err != nil {
			return responseBodyRaw
		}
		return parsedBody
	}

	if len(responseBodyRaw) > responsePreviewLimit {
		return fmt.Sprintf("%s... [truncated, total %d bytes]", responseBodyRaw[:responsePreviewLimit], len(responseBodyRaw))
	}

	return responseBodyRaw
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
