package domain

import (
	"fmt"
)

type notAValidPriceError struct {
	price float64
}

func (err notAValidPriceError) Error() string {
	return fmt.Sprintf("%f is not a valid price", err.price)
}

var validPricesInterval = []float64{0.001, 999999}

// ValueObjectPrice
// Value Object to check price values
type ValueObjectPrice struct {
	value float64
}

func NewValueObjectPrice(price float64) (*ValueObjectPrice, error) {
	if price < validPricesInterval[0] || price > validPricesInterval[1] {
		return nil, notAValidPriceError{price}
	}

	return &ValueObjectPrice{price}, nil
}

func NewValueObjectPriceMust(price float64) *ValueObjectPrice {
	valueObject, _ := NewValueObjectPrice(price)
	return valueObject
}

func (price ValueObjectPrice) Value() interface{} {
	return price.value
}

func (price ValueObjectPrice) String() string {
	return fmt.Sprintf("%.2f", price.value)
}
