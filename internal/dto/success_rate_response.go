package dto

type SuccessRateResponse struct {
	TotalChecks    int     `json:"total_checks"`
	Successful     int     `json:"successful"`
	Failed         int     `json:"failed"`
	SuccessRate    float64 `json:"success_rate"`
	FailureRate    float64 `json:"failure_rate"`
}
