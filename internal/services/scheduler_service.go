package services

import (
	"context"
	"sync"
	"time"

	"github.com/Abhi78k/api-performance-observatory/internal/logger"
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

func ShouldCheck(endpoint models.Endpoint) bool {

	if endpoint.LastCheckedAt == nil {
		return true
	}

	nextCheck := endpoint.LastCheckedAt.Add(
		time.Duration(endpoint.CheckInterval) * time.Minute,
	)

	return time.Now().After(nextCheck)
}

func (s *SchedulerService) Start(ctx context.Context) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	logger.Info("Scheduler started.")

	for {
		// NOTE:
		// This requires EndpointRepository to have:
		//
		// func (r *EndpointRepository) GetAllEndpoints() ([]models.Endpoint, error)
		//
		// The scheduler needs all endpoints across all users.
		endpoints, err := s.EndpointRepo.GetAllEndpoints()

		if err != nil {
			logger.Error(
				"Failed to load endpoints",
				"error",
				err,
			)
			select {

			case <-ctx.Done():

				logger.Info("Scheduler stopped.")

				return

			case <-ticker.C:

			}
			continue
		}

		var wg sync.WaitGroup

		semaphore := make(chan struct{}, 10)

		for _, endpoint := range endpoints {
			if !ShouldCheck(endpoint) {
				continue
			}

			wg.Add(1)

			go func(ep models.Endpoint) {

				semaphore <- struct{}{}

				defer func() {
					<-semaphore
					wg.Done()
				}()

				logger.Info(
					"Running health check",
					"endpoint_id", ep.ID,
					"name", ep.Name,
				)

				if err := s.HealthCheckService.CheckEndpoint(ep); err != nil {
					logger.Error(
						"Health check failed",
						"endpoint_id", ep.ID,
						"name", ep.Name,
						"error", err,
					)
					return
				}

				now := time.Now()
				ep.LastCheckedAt = &now

				if err := s.EndpointRepo.Update(&ep); err != nil {
					logger.Error(
						"Failed to update last_checked_at",
						"endpoint_id", ep.ID,
						"name", ep.Name,
						"error", err,
					)
				}

			}(endpoint)
		}

		wg.Wait()

		logger.Info("Monitoring cycle completed.")

		select {

		case <-ctx.Done():

			logger.Info("Scheduler stopped.")

			return

		case <-ticker.C:

		}
	}
}
