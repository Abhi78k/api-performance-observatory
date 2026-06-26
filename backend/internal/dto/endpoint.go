package dto

type CreateEndpointRequest struct {
	Name string `json:"name" validate:"required,min=2,max=100"`

	URL string `json:"url" validate:"required,url"`

	ExpectedStatus int `json:"expected_status" validate:"required,gte=100,lte=599"`
}

type UpdateEndpointRequest struct {
	Name string `json:"name" validate:"required,min=2,max=100"`

	URL string `json:"url" validate:"required,url"`

	ExpectedStatus int `json:"expected_status" validate:"required,gte=100,lte=599"`
}
