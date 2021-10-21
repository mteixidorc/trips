package domain_test

import (
	"testing"

	"github.com/mteixidorc/trips/internal/trips/domain"
)

func TestValueObjectTripDatesMustFail(t *testing.T) {
	wrongDates := "hello bye"
	_, err := domain.NewValueObjectTripDates(wrongDates)
	if err == nil {
		t.Fail()
	}
}

func TestValueObjectTripDatesOk(t *testing.T) {
	wrongDates := "Mon Wed"
	_, err := domain.NewValueObjectTripDates(wrongDates)
	if err != nil {
		t.Fail()
	}
}
