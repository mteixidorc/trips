package application_test

import (
	"testing"

	"github.com/mteixidorc/trips/internal/shared/domain/value"
	"github.com/mteixidorc/trips/internal/trips/application/mock"
	"github.com/mteixidorc/trips/internal/trips/application/trip"
	appmock "github.com/mteixidorc/trips/internal/trips/infrastructure/repository/mock"
)

func TestCreateTripOriginCityNotExistsMustFail(t *testing.T) {
	useCases := mock.BuildMockUseCases()

	_, err := useCases.CreateTrip(&trip.TripDTO{
		OriginId:      999,
		DestinationId: appmock.MockCity1ID.Value().(int64),
		Dates:         "Thu",
		Price:         345.34,
	})

	if err == nil {
		t.Fail()
	}
}

func TestCreateTripMustWork(t *testing.T) {
	useCases := mock.BuildMockUseCases()

	_, err := useCases.CreateTrip(&trip.TripDTO{
		OriginId:      appmock.MockCity2ID.Value().(int64),
		DestinationId: appmock.MockCity1ID.Value().(int64),
		Dates:         "Thu",
		Price:         345.34,
	})

	if err != nil {
		t.Fail()
	}
}

func TestGetNonExistIDTripMustFail(t *testing.T) {
	useCases := mock.BuildMockUseCases()

	_, err := useCases.GetTripById(value.NewValueObjectUniqueID().String())

	if err == nil {
		t.Fail()
	}
}

func TestGetAllTripsMustWork(t *testing.T) {
	useCases := mock.BuildMockUseCases()

	trips, err := useCases.GetAllTrips()
	if err != nil {
		t.FailNow()
	}

	if len(trips) == 0 {
		t.Fail()
	}
}
