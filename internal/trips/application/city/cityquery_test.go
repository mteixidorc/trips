package city_test

import (
	"testing"

	"github.com/mteixidorc/trips/internal/trips/application/city"
	"github.com/mteixidorc/trips/internal/trips/infrastructure/repository/mock"
)

func TestQueryCity(t *testing.T) {
	tt := []struct {
		name      string
		id        int64
		wantError bool
	}{
		{
			name:      "OK Query city, exists and must work",
			id:        mock.MockCity1ID.Value().(int64),
			wantError: false,
		},
		{
			name:      "KO Query city, not exists and will fail",
			id:        999,
			wantError: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			query := city.NewCityQuery(mock.NewMockCityRepository())
			_, err := query.Get(tc.id)
			if (tc.wantError && err == nil) || (!tc.wantError && err != nil) {
				t.Fail()
			}
		})
	}
}
