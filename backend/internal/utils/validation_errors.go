package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationErrors(err error) map[string]string {
	result := make(map[string]string)

	for _, e := range err.(validator.ValidationErrors) {

		field := strings.ToLower(e.Field())

		switch e.Tag() {

		case "required":
			result[field] = "is required"

		case "email":
			result[field] = "must be a valid email"

		case "url":
			result[field] = "must be a valid URL"

		case "min":
			result[field] = "is too short"

		case "max":
			result[field] = "is too long"

		case "gte":
			result[field] = "must be greater than or equal to " + e.Param()

		case "lte":
			result[field] = "must be less than or equal to " + e.Param()

		default:
			result[field] = "is invalid"
		}
	}

	return result
}
