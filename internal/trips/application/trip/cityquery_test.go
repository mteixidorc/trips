package trip_test

import (
	"testing"

	"github.com/mteixidorc/trips/internal/trips/application/trip"
	"github.com/mteixidorc/trips/internal/trips/infrastructure/repository/mock"
)

func TestQueryTripNotExistsMustFail(t *testing.T) {
	tt := []struct {
		name      string
		id        string
		wantError bool
	}{
		{
			name:      "OK Query trip, exists and must work",
			id:        mock.MockTrip1ID.String(),
			wantError: false,
		},
		{
			name:      "KO Query trip, not exists and will fail",
			id:        "7253a0a1-1ee6-4c0c-bfde-a9594b2a237b",
			wantError: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			query := trip.NewTripQuery(mock.NewMockTripRepository())
			_, err := query.Get(tc.id)
			if (tc.wantError && err == nil) || (!tc.wantError && err != nil) {
				t.Fail()
			}
		})
	}
}

func TestQueryTripGetALLMustWork(t *testing.T) {
	query := trip.NewTripQuery(mock.NewMockTripRepository())
	_, err := query.GetAll()
	if err != nil {
		t.Fail()
	}
}
