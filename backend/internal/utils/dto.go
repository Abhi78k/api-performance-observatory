package utils

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type ValidationErrorResponse struct {
	Success bool              `json:"success"`
	Errors  map[string]string `json:"errors"`
}

type MessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
