package service

import "github.com/fernandoocampo/pack/util"

// IHealthService defines health behavior for microservices.
type IHealthService interface {
	// Health census the resources that the service use.
	Health() *util.HealthStatus
}
