package services

import (
	"github.com/Abhi78k/api-performance-observatory/internal/models"
	"github.com/Abhi78k/api-performance-observatory/internal/repositories"
)

type EndpointService struct {
	endpointRepo *repositories.EndpointRepository
}

func NewEndpointService(
	endpointRepo *repositories.EndpointRepository,
) *EndpointService {

	return &EndpointService{
		endpointRepo: endpointRepo,
	}
}

func (s *EndpointService) CreateEndpoint(
	name string,
	url string,
	expectedStatus int,
	userID uint,
) (*models.Endpoint, error) {

	endpoint := &models.Endpoint{
		Name:           name,
		URL:            url,
		ExpectedStatus: expectedStatus,
		UserID:         userID,
	}

	err := s.endpointRepo.Create(endpoint)

	if err != nil {
		return nil, err
	}

	return endpoint, nil
}

func (s *EndpointService) GetEndpoints(
	userID uint,
) ([]models.Endpoint, error) {

	return s.endpointRepo.GetAllByUserID(userID)
}

func (s *EndpointService) GetEndpoint(
	id string,
	userID uint,
) (*models.Endpoint, error) {

	return s.endpointRepo.GetByID(id, userID)
}

func (s *EndpointService) UpdateEndpoint(
	id string,
	userID uint,
	name string,
	url string,
	expectedStatus int,
) (*models.Endpoint, error) {

	endpoint, err := s.endpointRepo.GetByID(id, userID)

	if err != nil {
		return nil, err
	}

	endpoint.Name = name
	endpoint.URL = url
	endpoint.ExpectedStatus = expectedStatus

	err = s.endpointRepo.Update(endpoint)

	if err != nil {
		return nil, err
	}

	return endpoint, nil
}

func (s *EndpointService) DeleteEndpoint(
	id string,
	userID uint,
) error {

	endpoint, err := s.endpointRepo.GetByID(id, userID)

	if err != nil {
		return err
	}

	return s.endpointRepo.Delete(endpoint)
}
