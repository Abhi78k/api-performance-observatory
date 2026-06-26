package models

import (
	"time"
)

type Incident struct {
	ID         uint
	EndpointID uint
	StartedAt  time.Time
	ResolvedAt *time.Time
	IsResolved bool
}
