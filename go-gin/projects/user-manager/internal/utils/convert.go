package utils

import (
	"regexp"
	"strings"
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func CamelCaseToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return snake
}

func NormalizeString(text string) string {
	return strings.ToLower(strings.TrimSpace(text))
}

func NormalizePaginationParams(page, limit *int) (nP int, nL int) {
	normalizedPage, normalizedLimit := 1, 10
	if page != nil && *page > 0 {
		normalizedPage = *page
	}
	if limit != nil && *limit > 0 && *limit <= 100 {
		normalizedLimit = *limit
	}
	return normalizedPage, normalizedLimit
}

func ContainsIgnoreCase(str, substr string) bool {
	str = NormalizeString(str)
	substr = NormalizeString(substr)
	return strings.Contains(str, substr)
}
