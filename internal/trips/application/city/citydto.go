package city

import "github.com/mteixidorc/trips/internal/trips/domain"

// CityDTO
// Transfer Object to communicate with application layer
type CityDTO struct {
	Id   int64
	Name string
}

func NewCityResponseFromAggregate(city *domain.City) *CityDTO {
	return &CityDTO{
		Id:   city.GetId().Value().(int64),
		Name: city.GetName(),
	}
}
