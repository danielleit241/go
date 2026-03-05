package utils

import (
	"regexp"
	"time"
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

const DateTimeFormat = "2006-01-02"

func CamelCaseToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return snake
}

func ConvertToInt32Ptr(i *int) *int32 {
	if i == nil {
		return nil
	}
	i32 := int32(*i)
	return &i32
}

func ConvertInt32PtrToInt(i *int32) *int {
	if i == nil {
		return nil
	}
	result := int(*i)
	return &result
}

func FormatTimePtr(t *time.Time, layout string) *string {
	if t == nil {
		return nil
	}
	formatted := t.Format(layout)
	return &formatted
}
