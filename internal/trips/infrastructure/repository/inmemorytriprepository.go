package repository

import (
	"github.com/mteixidorc/trips/internal/shared/domain/errors"
	"github.com/mteixidorc/trips/internal/shared/domain/value"
	"github.com/mteixidorc/trips/internal/trips/domain"
)

// InMemoryTripRepository
// Implements TripRepository interface and saves data in memory
type InMemoryTripRepository struct {
	trips map[string]*domain.Trip
}

func NewInMemoryTripRepository() InMemoryTripRepository {
	return InMemoryTripRepository{
		trips: map[string]*domain.Trip{
			"6de052bc-94f9-4bfc-8fa5-0962ff9b4b99": domain.NewTrip(
				value.NewValueObjectUniqueIDFromStringMust("6de052bc-94f9-4bfc-8fa5-0962ff9b4b99"),
				value.NewValueObjectID(1),
				value.NewValueObjectID(2),
				domain.NewValueObjectTripDatesMust("Mon Tue Wed Fri"),
				domain.NewValueObjectPriceMust(50.55)),
			"c7b33bfe-3584-4a02-96f1-d1a831f83252": domain.NewTrip(
				value.NewValueObjectUniqueIDFromStringMust("c7b33bfe-3584-4a02-96f1-d1a831f83252"),
				value.NewValueObjectID(2),
				value.NewValueObjectID(1),
				domain.NewValueObjectTripDatesMust("Sat Sun"),
				domain.NewValueObjectPriceMust(150.55)),
			"42bf5c67-8f42-4c55-8e0c-7c5020cfad58": domain.NewTrip(
				value.NewValueObjectUniqueIDFromStringMust("42bf5c67-8f42-4c55-8e0c-7c5020cfad58"),
				value.NewValueObjectID(3),
				value.NewValueObjectID(6),
				domain.NewValueObjectTripDatesMust("Mon Tue Wed Thu Fri"),
				domain.NewValueObjectPriceMust(5.01)),
		},
	}
}

func (repository InMemoryTripRepository) Get(id value.ValueObjectUniqueID) (*domain.Trip, error) {
	trip, exists := repository.trips[id.String()]
	if !exists {
		return nil, errors.NewRecordNotExistsError(id)
	}

	return trip, nil
}

func (repository InMemoryTripRepository) GetAll() ([]*domain.Trip, error) {
	result := make([]*domain.Trip, 0)
	for _, trip := range repository.trips {
		result = append(result, trip)
	}
	return result, nil
}

func (repository InMemoryTripRepository) Create(trip *domain.Trip) (string, error) {
	repository.trips[trip.Id().String()] = trip
	return trip.Id().String(), nil
}
