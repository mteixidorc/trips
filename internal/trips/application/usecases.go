package application

import (
	"fmt"

	cityapp "github.com/mteixidorc/trips/internal/trips/application/city"
	tripapp "github.com/mteixidorc/trips/internal/trips/application/trip"
	"github.com/mteixidorc/trips/internal/trips/domain"
)

// UseCases
// Interface for usecases implemented into trips bounded context
type UseCases interface {
	CreateTrip(*tripapp.TripDTO) (string, error)
	GetTripById(id string) (*tripapp.TripDTO, error)
	GetAllTrips() ([]*tripapp.TripDTO, error)
}

// TripsService
// An implementation of UseCases interface, out domain service of tips bounded context
type tripsService struct {
	tripquery   tripapp.TripQuery
	tripupdater tripapp.TripUpdater
	cityquery   cityapp.CityQuery
}

func NewTripService(tripRepository domain.TripRepository, cityRepository domain.CityRepository) tripsService {
	return tripsService{
		tripquery:   tripapp.NewTripQuery(tripRepository),
		tripupdater: tripapp.NewTripUpdater(tripRepository),
		cityquery:   cityapp.NewCityQuery(cityRepository),
	}
}

func (service tripsService) CreateTrip(trip *tripapp.TripDTO) (string, error) {
	// Verify if origin and destination cities exist
	_, err := service.cityquery.Get(trip.OriginId)
	if err != nil {
		return "", fmt.Errorf("origin city %d not exists", trip.OriginId)
	}

	_, err = service.cityquery.Get(trip.DestinationId)
	if err != nil {
		return "", fmt.Errorf("destination city %d not exists", trip.DestinationId)
	}

	return service.tripupdater.New(trip)
}

func (service tripsService) GetTripById(id string) (*tripapp.TripDTO, error) {
	trip, err := service.tripquery.Get(id)
	if err != nil {
		return nil, err
	}
	tripResult := tripapp.NewTripDTOFromAggregate(trip)

	err = service.getTripCities(tripResult)
	if err != nil {
		return nil, err
	}

	return tripResult, nil
}

func (service tripsService) GetAllTrips() ([]*tripapp.TripDTO, error) {
	trips, err := service.tripquery.GetAll()
	if err != nil {
		return nil, err
	}

	result := make([]*tripapp.TripDTO, len(trips))
	for pos, trip := range trips {
		tripResult := tripapp.NewTripDTOFromAggregate(trip)
		service.getTripCities(tripResult)
		result[pos] = tripResult
	}

	return result, nil
}

func (service tripsService) getTripCities(trip *tripapp.TripDTO) error {
	originCity, err := service.cityquery.Get(trip.OriginId)
	if err != nil {
		return fmt.Errorf("origin city id %d not exists", trip.OriginId)
	}

	destinationCity, err := service.cityquery.Get(trip.DestinationId)
	if err != nil {
		return fmt.Errorf("destination city id %d not exists", trip.DestinationId)
	}

	trip.OriginCityName = originCity.Name
	trip.DestinationCityName = destinationCity.Name
	return nil
}
