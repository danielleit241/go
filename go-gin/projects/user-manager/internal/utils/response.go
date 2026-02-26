package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseError(c *gin.Context, err error) {
	if appErr, ok := err.(*AppError); ok {
		switch appErr.Code {
		case ErrCodeBadRequest:
			c.JSON(http.StatusBadRequest, appErr)
			return
		case ErrCodeNotFound:
			c.JSON(http.StatusNotFound, appErr)
			return
		case ErrCodeConflict:
			c.JSON(http.StatusConflict, appErr)
			return
		default:
			c.JSON(http.StatusInternalServerError, appErr)
			return
		}
	}

	c.JSON(http.StatusInternalServerError, NewError("internal server error", ErrCodeInternalServerError))
}

func ResponseSuccess(c *gin.Context, data any, message string) {
	if message == "" {
		message = "success"
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"status":  "success",
		"data":    data,
	})
}
