package dto

type PerformanceStatsResponse struct {
	AverageResponseTime float64 `json:"average_response_time"`
	MinResponseTime     int64   `json:"min_response_time"`
	MaxResponseTime     int64   `json:"max_response_time"`
}
