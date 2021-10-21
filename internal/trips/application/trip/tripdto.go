package trip

import (
	"github.com/mteixidorc/trips/internal/trips/domain"
)

// TripDTO
// Transfer Object to communicate with application layer
type TripDTO struct {
	Id                  string
	OriginId            int64
	OriginCityName      string
	DestinationId       int64
	DestinationCityName string
	Dates               string
	Price               float64
}

func NewTripDTOFromAggregate(trip *domain.Trip) *TripDTO {
	return &TripDTO{
		Id:            trip.Id().String(),
		OriginId:      trip.OriginId().Value().(int64),
		DestinationId: trip.DestinationId().Value().(int64),
		Dates:         trip.Dates().String(),
		Price:         trip.Price(),
	}
}
