package trip_test

import (
	"testing"

	"github.com/mteixidorc/trips/internal/trips/application/trip"
	"github.com/mteixidorc/trips/internal/trips/infrastructure/repository/mock"
)

func TestUpdaterNewTripWithWrongDatesMustFail(t *testing.T) {
	updater := trip.NewTripUpdater(mock.NewMockTripRepository())
	_, err := updater.New(&trip.TripDTO{
		OriginId:      mock.MockCity1ID.Value().(int64),
		DestinationId: mock.MockCity2ID.Value().(int64),
		Dates:         "2020-01",
		Price:         12.34,
	})

	if err == nil {
		t.Fail()
	}
}

func TestUpdaterNewTripMustWork(t *testing.T) {
	updater := trip.NewTripUpdater(mock.NewMockTripRepository())
	_, err := updater.New(&trip.TripDTO{
		OriginId:      mock.MockCity1ID.Value().(int64),
		DestinationId: mock.MockCity2ID.Value().(int64),
		Dates:         "Mon Tue",
		Price:         64.34,
	})

	if err != nil {
		t.Fail()
	}
}

func TestUpdaterNewTripWithWrongPriceMustFail(t *testing.T) {
	updater := trip.NewTripUpdater(mock.NewMockTripRepository())
	_, err := updater.New(&trip.TripDTO{
		OriginId:      mock.MockCity1ID.Value().(int64),
		DestinationId: mock.MockCity2ID.Value().(int64),
		Dates:         "Mon Tue",
		Price:         -1.34,
	})

	if err == nil {
		t.Fail()
	}
}
