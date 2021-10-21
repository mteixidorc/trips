package trip_test

import (
	"testing"

	"github.com/mteixidorc/trips/internal/trips/application/trip"
	"github.com/mteixidorc/trips/internal/trips/infrastructure/repository/mock"
)

func TestQueryTripNotExistsMustFail(t *testing.T) {
	query := trip.NewTripQuery(mock.NewMockTripRepository())
	_, err := query.Get("this-id-not-exists")
	if err == nil {
		t.Fail()
	}
}

func TestQueryTripGetByIDExistsAndMustWork(t *testing.T) {
	query := trip.NewTripQuery(mock.NewMockTripRepository())
	_, err := query.Get(mock.MockTrip1ID.String())
	if err != nil {
		t.Fail()
	}
}

func TestQueryTripGetALLMustWork(t *testing.T) {
	query := trip.NewTripQuery(mock.NewMockTripRepository())
	_, err := query.GetAll()
	if err != nil {
		t.Fail()
	}
}
