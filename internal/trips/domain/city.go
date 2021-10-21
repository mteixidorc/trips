package domain

import "github.com/mteixidorc/trips/internal/shared/domain/value"

// City
// Domain object that represents a city
type City struct {
	id   value.ValueObjectID
	name string
}

func NewCity(id value.ValueObjectID, name string) *City {
	return &City{
		id:   id,
		name: name,
	}
}

func (city City) GetId() value.ValueObjectID {
	return city.id
}

func (city City) GetName() string {
	return city.name
}
