package errors

import (
	"errors"
	"net/http"
	
	"github.com/gin-gonic/gin"
)

// Common error types
var (
	ErrNotFound = errors.New("resource not found")
	ErrUnauthorized = errors.New("unauthorized access")
	ErrBadRequest = errors.New("invalid request")
	ErrServerError = errors.New("internal server error")
	ErrForbidden = errors.New("forbidden access")
)

// ErrorResponse formats standard error responses
func ErrorResponse(c *gin.Context, statusCode int, err error, details string) {
	message := err.Error()
	if details != "" {
		message = message + ": " + details
	}
	
	c.JSON(statusCode, gin.H{
		"success": false,
		"error": message,
	})
}

// HandleError determines the appropriate status code and handles the error
func HandleError(c *gin.Context, err error, details string) {
	var statusCode int
	
	switch {
	case errors.Is(err, ErrNotFound):
		statusCode = http.StatusNotFound
	case errors.Is(err, ErrUnauthorized):
		statusCode = http.StatusUnauthorized
	case errors.Is(err, ErrBadRequest):
		statusCode = http.StatusBadRequest
	case errors.Is(err, ErrForbidden):
		statusCode = http.StatusForbidden
	default:
		statusCode = http.StatusInternalServerError
		if err == nil {
			err = ErrServerError
		}
	}
	
	ErrorResponse(c, statusCode, err, details)
}