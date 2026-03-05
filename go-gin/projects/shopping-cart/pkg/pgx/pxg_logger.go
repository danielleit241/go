package pgx

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/tracelog"
	"github.com/rs/zerolog"
)

type PgxZerologTracer struct {
	Logger         *zerolog.Logger
	SlowQueryLimit time.Duration
}

type QueryInfo struct {
	QueryName   string
	Opperation  string
	CleanSql    string
	OriginalSql string
}

func (t *PgxZerologTracer) Log(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]any) {
	sql, _ := data["sql"].(string)
	duration, _ := data["duration"].(time.Duration)
	args, _ := data["args"].([]any)

	queryInfo := parseSql(sql, args)

	baseLogger := t.Logger.With().
		Str("query_name", queryInfo.QueryName).
		Str("operation", queryInfo.Opperation).
		Str("original_sql", queryInfo.OriginalSql).
		Str("clean_sql", queryInfo.CleanSql).
		Dur("duration", duration).
		Interface("args", args)

	logger := baseLogger.Logger()

	if msg == "Query" && duration > t.SlowQueryLimit {
		logger.Warn().Str("event", "slow query").Msgf("Slow query detected: %s", msg)
		return
	}

	if msg == "Query" {
		logger.Info().Str("event", "query").Msg(msg)
		return
	}
}

var (
	sqlcNameRegex = regexp.MustCompile(`-- name:\s*(\w+)\s*:(\w+)`)
	spaceRegex    = regexp.MustCompile(`\s+`)
	commentRegex  = regexp.MustCompile(`--.*?$`)
)

func parseSql(sql string, args []any) QueryInfo {
	info := QueryInfo{
		OriginalSql: sql,
	}
	matches := sqlcNameRegex.FindStringSubmatch(sql)
	if len(matches) == 3 {
		info.QueryName = matches[1]
		info.Opperation = matches[2]
	}

	cleanSql := sqlcNameRegex.ReplaceAllString(sql, "")
	cleanSql = strings.TrimSpace(cleanSql)
	cleanSql = spaceRegex.ReplaceAllString(cleanSql, " ")
	cleanSql = commentRegex.ReplaceAllString(cleanSql, "")

	for i, arg := range args {
		placeholder := fmt.Sprintf("$%d", i+1)
		cleanSql = strings.ReplaceAll(cleanSql, placeholder, formatArg(arg))
	}

	info.CleanSql = cleanSql

	return info
}

func formatArg(arg any) string {
	val := reflect.ValueOf(arg)
	if arg == nil || (val.Kind() == reflect.Ptr && val.IsNil()) {
		return "NULL"
	}

	if val.Kind() == reflect.Ptr {
		arg = val.Elem().Interface()
	}

	switch v := arg.(type) {
	case string:
		return fmt.Sprintf("'%s'", v)
	case time.Time:
		return fmt.Sprintf("'%s'", v.Format(time.RFC3339))
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", v)
	case float32, float64:
		return fmt.Sprintf("%f", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}
