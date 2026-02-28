package validation

import (
	"fmt"
	"strings"

	"github.com/danielleit241/internal/utils"
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
	"search": func(field, tag, param string) string {
		return fmt.Sprintf("%s must be a valid search query (alphanumeric characters and spaces only)", field)
	},
	"date": func(field, tag, param string) string {
		return fmt.Sprintf("%s must be a valid date in the format %s", field, param)
	},
	"blocked_email_domain": func(field, tag, param string) string {
		return fmt.Sprintf("%s is not allowed to be from a blocked email domain", field)
	},
	"strong_password": func(field, tag, param string) string {
		return fmt.Sprintf("%s must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, one number, and one special character", field)
	},
}

func Initialize() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		fmt.Println("Failed to initialize validator")
	}

	RegisterCustomValidations(v)

	return nil
}

func HandleValidationError(err error) gin.H {
	validErr, ok := err.(validator.ValidationErrors)
	if !ok {
		return gin.H{"error": "Invalid error type"}
	}

	errorMessages := make(map[string]string, len(validErr))
	for _, e := range validErr {
		namespaceParts := strings.SplitN(e.Namespace(), ".", 2)
		rawPath := e.Namespace()
		if len(namespaceParts) == 2 {
			rawPath = namespaceParts[1]
		}

		parts := strings.Split(rawPath, ".")
		for i, part := range parts {
			indexPos := strings.Index(part, "[")
			if indexPos == -1 {
				parts[i] = utils.CamelCaseToSnakeCase(part)
				continue
			}

			base := part[:indexPos]
			index := part[indexPos:]
			parts[i] = fmt.Sprintf("%s%s", utils.CamelCaseToSnakeCase(base), index)
		}

		fieldName := strings.Join(parts, ".")
		tag := e.Tag()
		param := formatValidationParam(tag, e.Param())

		if msgFunc, exists := validationErrorMessages[tag]; exists {
			errorMessages[fieldName] = msgFunc(fieldName, tag, param)
		} else {
			errorMessages[fieldName] = fmt.Sprintf("%s is not valid", fieldName)
		}
	}

	return gin.H{"errors": errorMessages}
}

func formatValidationParam(tag, param string) string {
	if tag == "oneof" {
		return strings.Join(strings.Fields(param), ", ")
	}

	return param
}
