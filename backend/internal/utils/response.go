package utils

import (
	"math"
	"net/http"
	"strconv"

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

func GetPaginationParams(c *gin.Context) (page int, limit int) {
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	page = 1
	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	limit = 10
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	if limit > 100 {
		limit = 100
	}

	return page, limit
}

func PaginatedOK(c *gin.Context, data any, page, limit int, totalItems int64) {
	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))
	if totalPages < 0 {
		totalPages = 0
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
		"pagination": gin.H{
			"page":        page,
			"limit":       limit,
			"totalItems":  totalItems,
			"totalPages":  totalPages,
			"hasNext":     page < totalPages,
			"hasPrevious": page > 1,
		},
	})
}
