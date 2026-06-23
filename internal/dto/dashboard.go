package dto

type DashboardStatusResponse struct {
	EndpointID   uint   `json:"endpoint_id"`
	EndpointName string `json:"endpoint_name"`
	Status       string `json:"status"`
}
