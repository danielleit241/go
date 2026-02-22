package utils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type ErrorMessageFunc func(field, tag, param string) string

var validationErrorMessages = map[string]ErrorMessageFunc{
	"required": func(field, tag, param string) string {
		return fmt.Sprintf("%s is required", field)
	},
	"omitempty": func(field, tag, param string) string {
		return fmt.Sprintf("%s is optional", field)
	},
	"min": func(field, tag, param string) string {
		return fmt.Sprintf("%s must be at least %s characters long", field, param)
	},
	"max": func(field, tag, param string) string {
		return fmt.Sprintf("%s must be at most %s characters long", field, param)
	},
	"gt": func(field, tag, param string) string {
		return fmt.Sprintf("%s must be greater than %s", field, param)
	},
	"gte": func(field, tag, param string) string {
		return fmt.Sprintf("%s must be greater than or equal to %s", field, param)
	},
	"oneof": func(field, tag, param string) string {
		return fmt.Sprintf("%s must be one of the following values: %s", field, param)
	},
	"uuid": func(field, tag, param string) string {
		return fmt.Sprintf("%s must be a valid UUID", field)
	},
	"slug": func(field, tag, param string) string {
		return fmt.Sprintf("%s must be a valid slug (lowercase letters, numbers, and hyphens) without spaces", field)
	},
	"category": func(field, tag, param string) string {
		return fmt.Sprintf("%s must be one of the following categories: electronics, books, clothing, home", field)
	},
	"search": func(field, tag, param string) string {
		return fmt.Sprintf("%s must be a valid search query (alphanumeric characters and spaces only)", field)
	},
	"date": func(field, tag, param string) string {
		return fmt.Sprintf("%s must be a valid date in the format %s", field, param)
	},
	"image": func(field, tag, param string) string {
		return fmt.Sprintf("%s must be a valid image URL (must end with .jpg, .jpeg, .png, or .gif)", field)
	},
}

type CustomValidator struct {
	Tag string
	Fn  validator.Func
}

var customValidators = []CustomValidator{
	{
		Tag: "slug",
		Fn:  validateSlug,
	},
	{
		Tag: "category",
		Fn:  validateCategory,
	},
	{
		Tag: "search",
		Fn:  validateSearch,
	},
	{
		Tag: "image",
		Fn:  validateImage,
	},
}

func validateSlug(fl validator.FieldLevel) bool {
	var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)
	return slugRegex.MatchString(fl.Field().String())
}

func validateCategory(fl validator.FieldLevel) bool {
	var validCategory = map[string]bool{
		"electronics": true,
		"books":       true,
		"clothing":    true,
		"home":        true,
	}
	return validCategory[strings.TrimSpace(strings.ToLower(fl.Field().String()))]
}

func validateSearch(fl validator.FieldLevel) bool {
	var searchRegex = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)
	return searchRegex.MatchString(fl.Field().String())
}

func validateImage(fl validator.FieldLevel) bool {
	validImage := []string{".jpg", ".jpeg", ".png", ".gif"}
	url := fl.Field().String()
	var urlRegex = regexp.MustCompile(`^https?://[^\s]+$`)
	return urlRegex.MatchString(url) && func() bool {
		for _, ext := range validImage {
			if strings.HasSuffix(url, ext) {
				return true
			}
		}
		return false
	}()
}

func HandleValidationError(err error) gin.H {
	validErr, ok := err.(validator.ValidationErrors)
	if !ok {
		return gin.H{"error": "Invalid error type"}
	}

	errorMessages := make(map[string]string)
	for _, e := range validErr {
		fieldName := e.Field()
		tag := e.Tag()
		param := e.Param()
		if msgFunc, exists := validationErrorMessages[tag]; exists {
			if tag == "oneof" {
				param = strings.Join(strings.Split(param, " "), ", ")
			}
			errorMessages[fieldName] = msgFunc(fieldName, tag, param)
		} else {
			errorMessages[fieldName] = fmt.Sprintf("%s is not valid", fieldName)
		}
	}
	return gin.H{"errors": errorMessages}
}

func RegisterValidators() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return fmt.Errorf("failed to get validator engine")
	}

	for _, cv := range customValidators {
		if err := v.RegisterValidation(cv.Tag, cv.Fn); err != nil {
			return fmt.Errorf("failed to register validator for tag '%s': %w", cv.Tag, err)
		}
	}

	return nil
}
