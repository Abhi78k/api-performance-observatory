package stats

import "github.com/Abhi78k/api-performance-observatory/backend/internal/models"

type EndpointStats struct {
	TotalChecks    int     `json:"total_checks"`
	SuccessRate    float64 `json:"success_rate"`
	AverageLatency float64 `json:"average_latency"`
}

func CalculateStats(checks []models.HealthCheck) EndpointStats {
	total := len(checks)

	// If there are no checks, return empty stats
	if total == 0 {
		return EndpointStats{}
	}

	successful := 0
	var totalLatency int64

	// Loop through all health checks
	for _, check := range checks {

		// Count successful checks (HTTP 200)
		if check.StatusCode == 200 {
			successful++
		}

		// Add response time to total latency
		totalLatency += check.ResponseTime
	}

	// Calculate success rate
	successRate := (float64(successful) / float64(total)) * 100

	// Calculate average latency
	averageLatency := float64(totalLatency) / float64(total)

	// Return calculated statistics
	return EndpointStats{
		TotalChecks:    total,
		SuccessRate:    successRate,
		AverageLatency: averageLatency,
	}
}
