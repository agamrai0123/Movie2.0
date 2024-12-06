package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(code int, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

func HandleError(c *gin.Context, err error) {
	var appErr *AppError
	if ok := errors.As(err, &appErr); ok {
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
	}
}
