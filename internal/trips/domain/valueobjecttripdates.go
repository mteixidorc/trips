package domain

import (
	"fmt"
	"strings"
)

type notAValidTripDateError struct {
	date string
}

func (err notAValidTripDateError) Error() string {
	return fmt.Sprintf("%s is not a valid trip date", err.date)
}

var validTripDates = map[string]int{
	"Mon": 1,
	"Tue": 2,
	"Wed": 3,
	"Thu": 4,
	"Fri": 5,
	"Sat": 6,
	"Sun": 7,
}

// ValueObjectTripDates
// Value Object to check dates string values
type ValueObjectTripDates struct {
	value string
}

func NewValueObjectTripDates(dates string) (*ValueObjectTripDates, error) {
	if dates == "" {
		return nil, notAValidTripDateError{dates}
	}
	datesValues := strings.Split(dates, " ")

	for _, date := range datesValues {
		if _, ok := validTripDates[date]; !ok {
			return nil, notAValidTripDateError{date}
		}
	}

	return &ValueObjectTripDates{dates}, nil
}

func NewValueObjectTripDatesMust(dates string) *ValueObjectTripDates {
	valueObject, _ := NewValueObjectTripDates(dates)
	return valueObject
}

func (dates ValueObjectTripDates) Value() interface{} {
	return dates.value
}

func (dates ValueObjectTripDates) String() string {
	return dates.value
}
