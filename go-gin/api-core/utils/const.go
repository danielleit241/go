package utils

import "regexp"

var maxImageSize int64 = 1 << 20 // 1MB
var validImageContentTypes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
	"image/gif":  true,
}
var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)
var searchRegex = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)
var validImage = []string{".jpg", ".jpeg", ".png", ".gif"}
var urlRegex = regexp.MustCompile(`^https?://[^\s]+$`)
var validCategory = map[string]bool{
	"electronics": true,
	"books":       true,
	"clothing":    true,
	"home":        true,
}
