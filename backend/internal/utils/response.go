package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, status int, data any) {
	c.JSON(status, SuccessResponse{
		Success: true,
		Data:    data,
	})
}

func Message(c *gin.Context, status int, message string) {
	c.JSON(status, SuccessResponse{
		Success: true,
		Message: message,
	})
}

func Error(c *gin.Context, status int, err string) {
	c.JSON(status, ErrorResponse{
		Success: false,
		Error:   err,
	})
}

func ValidationError(c *gin.Context, errors map[string]string) {
	c.JSON(http.StatusBadRequest, ValidationErrorResponse{
		Success: true,
		Errors:  errors,
	})
}

func OK(c *gin.Context, data any) {
	Success(c, http.StatusOK, data)
}

func Created(c *gin.Context, data any) {
	Success(c, http.StatusCreated, data)
}

func BadRequest(c *gin.Context, err string) {
	Error(c, http.StatusBadRequest, err)
}

func Unauthorized(c *gin.Context, err string) {
	Error(c, http.StatusUnauthorized, err)
}

func NotFound(c *gin.Context, err string) {
	Error(c, http.StatusNotFound, err)
}

func Internal(c *gin.Context, err string) {
	Error(c, http.StatusInternalServerError, err)
}

func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func Conflict(c *gin.Context, err string) {
	Error(c, http.StatusConflict, err)
}

func Forbidden(c *gin.Context, err string) {
	Error(c, http.StatusForbidden, err)
}
