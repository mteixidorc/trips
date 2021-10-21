package trip

import (
	"github.com/mteixidorc/trips/internal/shared/domain/value"
	"github.com/mteixidorc/trips/internal/trips/domain"
)


// TripUpdater
// Encapsulates all operations that modifies info (CQS)
type TripUpdater struct {
	repository domain.TripRepository
}

func NewTripUpdater(repository domain.TripRepository) TripUpdater {
	return TripUpdater{
		repository: repository,
	}
}

func (updater TripUpdater) New(trip *TripDTO) error {
	// Verify dates
	dates, err := domain.NewValueObjectTripDates(trip.Dates)

	if err != nil {
		return err
	}

	return updater.repository.Create(
		domain.NewTrip(
			value.NewValueObjectUniqueID(),
			value.NewValueObjectID(trip.OriginId),
			value.NewValueObjectID(trip.DestinationId),
			dates,
			trip.Price),
	)
}
