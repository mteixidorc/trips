package application_test

import (
	"testing"

	"github.com/mteixidorc/trips/internal/shared/domain/errors"
	"github.com/mteixidorc/trips/internal/shared/domain/value"
	"github.com/mteixidorc/trips/internal/trips/application"
	"github.com/mteixidorc/trips/internal/trips/application/trip"
	"github.com/mteixidorc/trips/internal/trips/domain"
)

type MockCityRepository struct {
	data map[value.ValueObjectID]*domain.City
}

func NewMockCityRepository() MockCityRepository {
	city1ID := value.NewValueObjectID(1)
	city2ID := value.NewValueObjectID(2)

	return MockCityRepository{
		data: map[value.ValueObjectID]*domain.City{
			city1ID: domain.NewCity(city1ID, "test-city-1"),
			city2ID: domain.NewCity(city2ID, "test-city-2"),
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

type MockTripRepository struct {
	data map[value.ValueObjectUniqueID]*domain.Trip
}

func NewMockTripRepository() MockTripRepository {
	trip1ID := value.NewValueObjectUniqueID()
	trip2ID := value.NewValueObjectUniqueID()

	return MockTripRepository{
		data: map[value.ValueObjectUniqueID]*domain.Trip{
			trip1ID: domain.NewTrip(trip1ID,
				value.NewValueObjectID(1),
				value.NewValueObjectID(2),
				domain.NewValueObjectTripDatesMust("Mon Tue Wed Fri"),
				40.55),
			trip2ID: domain.NewTrip(trip2ID,
				value.NewValueObjectID(2),
				value.NewValueObjectID(1),
				domain.NewValueObjectTripDatesMust("Sat Sun"),
				40.55),
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

func TestCreateTripOriginCityNotExistsMustFail(t *testing.T) {
	useCases := buildUseCases()

	_, err := useCases.CreateTrip(&trip.TripDTO{
		OriginId:      34,
		DestinationId: 1,
		Dates:         "Thu",
		Price:         345.34,
	})

	if err == nil {
		t.Fail()
	}
}

func TestCreateTripMustWork(t *testing.T) {
	useCases := buildUseCases()

	_, err := useCases.CreateTrip(&trip.TripDTO{
		OriginId:      2,
		DestinationId: 1,
		Dates:         "Thu",
		Price:         345.34,
	})

	if err != nil {
		t.Fail()
	}
}

func TestGetNonExistIDTripMustFail(t *testing.T) {
	useCases := buildUseCases()

	_, err := useCases.GetTripById(value.NewValueObjectUniqueID().String())

	if err == nil {
		t.Fail()
	}
}

func TestGetTripsMustWork(t *testing.T) {
	useCases := buildUseCases()

	trips, err := useCases.GetAllTrips()
	if err != nil {
		t.FailNow()
	}

	if len(trips) == 0 {
		t.Fail()
	}
}

// DRY
func buildUseCases() application.UseCases {
	cityRepo := NewMockCityRepository()
	tripRepo := NewMockTripRepository()

	return application.NewTripService(tripRepo, cityRepo)
}
