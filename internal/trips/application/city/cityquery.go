package city

import (
	"github.com/mteixidorc/trips/internal/shared/domain/value"
	"github.com/mteixidorc/trips/internal/trips/domain"
)

// CityQuery
// Queries for cities (CQS)
type CityQuery struct {
	repository domain.CityRepository
}

func NewCityQuery(repository domain.CityRepository) CityQuery {
	return CityQuery{
		repository: repository,
	}
}

func (cityQuery CityQuery) Get(id int64) (*CityDTO, error) {
	city, err := cityQuery.repository.Get(value.NewValueObjectID(id))
	if err != nil {
		return nil, err
	}

	return NewCityResponseFromAggregate(city), nil
}
