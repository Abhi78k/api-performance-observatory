package services

import (
	"log"
	"sync"
	"time"

	"github.com/Abhi78k/api-performance-observatory/internal/models"
	"github.com/Abhi78k/api-performance-observatory/internal/repositories"
)

type SchedulerService struct {
	EndpointRepo       *repositories.EndpointRepository
	HealthCheckService *HealthCheckService
}

func NewSchedulerService(
	endpointRepo *repositories.EndpointRepository,
	healthCheckService *HealthCheckService,
) *SchedulerService {
	return &SchedulerService{
		EndpointRepo:       endpointRepo,
		HealthCheckService: healthCheckService,
	}
}

func (s *SchedulerService) Start() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	log.Println("Scheduler started...")

	for {
		// NOTE:
		// This requires EndpointRepository to have:
		//
		// func (r *EndpointRepository) GetAllEndpoints() ([]models.Endpoint, error)
		//
		// The scheduler needs all endpoints across all users.
		endpoints, err := s.EndpointRepo.GetAllEndpoints()

		if err != nil {
			log.Println("failed to load endpoints:", err)
			<-ticker.C
			continue
		}

		var wg sync.WaitGroup

		for _, endpoint := range endpoints {
			wg.Add(1)

			go func(ep models.Endpoint) {
				defer wg.Done()

				log.Printf("Checking endpoint: %s", ep.Name)

				if err := s.HealthCheckService.CheckEndpoint(ep); err != nil {
					log.Printf(
						"health check failed for endpoint %d: %v",
						ep.ID,
						err,
					)
				}
			}(endpoint)
		}

		wg.Wait()

		log.Println("Monitoring cycle completed")

		<-ticker.C
	}
}
