package domain

import "github.com/mteixidorc/trips/internal/shared/domain/value"

// TripRepository
// Interface used by application layer to in/out trip records 
type TripRepository interface {
	GetAll() ([]*Trip, error)
	Get(id value.ValueObjectUniqueID) (*Trip, error)
	Create(trip *Trip) error
}
