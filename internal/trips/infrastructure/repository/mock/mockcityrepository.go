package mock

import (
	"github.com/mteixidorc/trips/internal/shared/domain/errors"
	"github.com/mteixidorc/trips/internal/shared/domain/value"
	"github.com/mteixidorc/trips/internal/trips/domain"
)

var (
	MockCity1ID = value.NewValueObjectID(1)
	MockCity2ID = value.NewValueObjectID(2)
)

type MockCityRepository struct {
	data map[value.ValueObjectID]*domain.City
}

func NewMockCityRepository() MockCityRepository {

	return MockCityRepository{
		data: map[value.ValueObjectID]*domain.City{
			MockCity1ID: domain.NewCity(MockCity1ID, "test-city-1"),
			MockCity2ID: domain.NewCity(MockCity2ID, "test-city-2"),
		},
	}
}

func (repository MockCityRepository) Get(id value.ValueObjectID) (*domain.City, error) {
	city, exists := repository.data[id]
	if !exists {
		return nil, errors.NewRecordNotExistsError(id)
	}

	return city, nil
}
