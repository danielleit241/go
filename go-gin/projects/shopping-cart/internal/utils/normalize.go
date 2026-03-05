package utils

import "strings"

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
