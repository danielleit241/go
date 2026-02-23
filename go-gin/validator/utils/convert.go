package utils

import (
	"regexp"
	"strings"
)

func CamelCaseToSnakeCase(s string) string {
	var snakeCaseRegex = regexp.MustCompile(`([a-z0-9])([A-Z])`)
	snake := snakeCaseRegex.ReplaceAllString(s, "${1}_${2}")
	return strings.ToLower(snake)
}
