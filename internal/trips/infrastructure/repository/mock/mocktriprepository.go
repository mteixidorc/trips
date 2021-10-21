package mock

import (
	"github.com/mteixidorc/trips/internal/shared/domain/errors"
	"github.com/mteixidorc/trips/internal/shared/domain/value"
	"github.com/mteixidorc/trips/internal/trips/domain"
)

var (
	MockTrip1ID = value.NewValueObjectUniqueID()
	MockTrip2ID = value.NewValueObjectUniqueID()
)

type MockTripRepository struct {
	data map[value.ValueObjectUniqueID]*domain.Trip
}

func NewMockTripRepository() MockTripRepository {

	return MockTripRepository{
		data: map[value.ValueObjectUniqueID]*domain.Trip{
			MockTrip1ID: domain.NewTrip(MockTrip1ID,
				value.NewValueObjectID(1),
				value.NewValueObjectID(2),
				domain.NewValueObjectTripDatesMust("Mon Tue Wed Fri"),
				domain.NewValueObjectPriceMust(40.55)),
			MockTrip2ID: domain.NewTrip(MockTrip2ID,
				value.NewValueObjectID(2),
				value.NewValueObjectID(1),
				domain.NewValueObjectTripDatesMust("Sat Sun"),
				domain.NewValueObjectPriceMust(20.55)),
		},
	}
}

func (repository MockTripRepository) Get(id value.ValueObjectUniqueID) (*domain.Trip, error) {
	trip, exists := repository.data[id]
	if !exists {
		return nil, errors.NewRecordNotExistsError(id)
	}
	return trip, nil
}

func (repository MockTripRepository) GetAll() ([]*domain.Trip, error) {
	result := make([]*domain.Trip, 0)
	for _, value := range repository.data {
		result = append(result, value)
	}

	return result, nil
}

func (repository MockTripRepository) Create(trip *domain.Trip) (string, error) {
	id := value.NewValueObjectUniqueID()
	repository.data[id] = trip
	return id.String(), nil
}
