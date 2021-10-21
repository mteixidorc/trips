package city_test

import (
	"testing"

	"github.com/mteixidorc/trips/internal/trips/application/city"
	"github.com/mteixidorc/trips/internal/trips/infrastructure/repository/mock"
)

func TestQueryCityNotExistsMustFail(t *testing.T) {
	query := city.NewCityQuery(mock.NewMockCityRepository())
	_, err := query.Get(999)
	if err == nil {
		t.Fail()
	}
}

func TestQueryCityExistsAndMustWork(t *testing.T) {
	query := city.NewCityQuery(mock.NewMockCityRepository())
	_, err := query.Get(mock.MockCity1ID.Value().(int64))
	if err != nil {
		t.Fail()
	}
}
