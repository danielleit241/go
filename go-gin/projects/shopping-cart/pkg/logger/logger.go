package logger

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerConfig struct {
	Level       string
	FileName    string
	MaxSize     int
	MaxBackups  int
	MaxAge      int
	Compress    bool
	Environment string
}

func NewLogger(config LoggerConfig) *zerolog.Logger {
	if dir := filepath.Dir(config.FileName); dir != "." {
		_ = os.MkdirAll(dir, 0o755)
	}

	zerolog.TimeFieldFormat = time.RFC3339

	lvl, err := zerolog.ParseLevel(config.Level)
	if err != nil {
		lvl = zerolog.InfoLevel
	}

	var writer io.Writer

	if config.Environment == "production" {
		writer = &lumberjack.Logger{
			Filename:   config.FileName,
			MaxSize:    config.MaxSize,
			MaxBackups: config.MaxBackups,
			MaxAge:     config.MaxAge,
			Compress:   config.Compress,
		}
	} else {
		writer = &PrettyJSONWriter{Writer: os.Stdout}
	}

	logger := zerolog.New(writer).Level(lvl).With().Timestamp().Logger()
	return &logger
}

type PrettyJSONWriter struct {
	Writer io.Writer
}

func (w *PrettyJSONWriter) Write(p []byte) (n int, err error) {
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, p, "", "  ")
	if err != nil {
		return w.Writer.Write(p)
	}
	return w.Writer.Write(prettyJSON.Bytes())
}
