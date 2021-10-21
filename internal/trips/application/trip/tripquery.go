package trip

import (
	"github.com/mteixidorc/trips/internal/shared/domain/value"
	"github.com/mteixidorc/trips/internal/trips/domain"
)

// TripQuery
// Queries for trips (CQS)
type TripQuery struct {
	repository domain.TripRepository
}

func NewTripQuery(repository domain.TripRepository) TripQuery {
	return TripQuery{
		repository: repository,
	}
}

func (query TripQuery) Get(id string) (*domain.Trip, error) {
	tripID, err := value.NewValueObjectUniqueIDFromString(id)
	if err != nil {
		return nil, err
	}
	return query.repository.Get(*tripID)
}

func (query TripQuery) GetAll() ([]*domain.Trip, error) {
	return query.repository.GetAll()
}
