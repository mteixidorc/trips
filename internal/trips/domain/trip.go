package domain

import "github.com/mteixidorc/trips/internal/shared/domain/value"

// Trip
// Domain object that represents a trip
type Trip struct {
	id            value.ValueObjectUniqueID
	originId      value.ValueObjectID
	destinationId value.ValueObjectID
	dates         ValueObjectTripDates
	price         float64
}

func NewTrip(id value.ValueObjectUniqueID, originId value.ValueObjectID, destinationId value.ValueObjectID, dates *ValueObjectTripDates, price float64) *Trip {
	return &Trip{
		id:            id,
		originId:      originId,
		destinationId: destinationId,
		dates:         *dates,
		price:         price,
	}
}

func (trip Trip) Id() value.ValueObjectUniqueID {
	return trip.id
}

func (trip Trip) OriginId() value.ValueObjectID {
	return trip.originId
}

func (trip Trip) DestinationId() value.ValueObjectID {
	return trip.destinationId
}

func (trip Trip) Dates() ValueObjectTripDates {
	return trip.dates
}

func (trip Trip) Price() float64 {
	return trip.price
}
