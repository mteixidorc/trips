package domain

import "github.com/mteixidorc/trips/internal/shared/domain/value"

// CityRepository
// Interface used by application layer to in/out city records
type CityRepository interface {
	Get(id value.ValueObjectID) (*City, error)
}
