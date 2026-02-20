package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleValidationError(err error) gin.H {

	if validErr, ok := err.(validator.ValidationErrors); ok {

		errors := make(map[string]string)
		for _, fieldErr := range validErr {
			switch fieldErr.Tag() {
			case "required":
				errors[fieldErr.Field()] = "is required"
			case "gt":
				errors[fieldErr.Field()] = "must be greater than " + fieldErr.Param()
			case "uuid":
				errors[fieldErr.Field()] = "must be a valid UUID"
			default:
				errors[fieldErr.Field()] = "is invalid"
			}
		}
		return gin.H{
			"error": errors,
		}
	}

	return gin.H{
		"error": "Validation error: " + err.Error(),
	}
}
