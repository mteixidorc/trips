package domain_test

import (
	"testing"

	"github.com/mteixidorc/trips/internal/trips/domain"
)

func TestValueObjectPriceMustFail(t *testing.T) {
	wrongPrice := -0.01
	_, err := domain.NewValueObjectPrice(wrongPrice)
	if err == nil {
		t.Fail()
	}
}

func TestValueObjectPriceOk(t *testing.T) {
	price := 0.01
	_, err := domain.NewValueObjectPrice(price)
	if err != nil {
		t.Fail()
	}
}
