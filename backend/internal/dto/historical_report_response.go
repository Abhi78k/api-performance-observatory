package dto

type HistoricalReportResponse struct {
	Period string `json:"period"`

	AverageResponseTime float64 `json:"average_response_time"`

	SuccessRate float64 `json:"success_rate"`

	TotalChecks int `json:"total_checks"`
}
