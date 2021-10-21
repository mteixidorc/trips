package mock

import (
	"github.com/mteixidorc/trips/internal/trips/application"
	"github.com/mteixidorc/trips/internal/trips/infrastructure/repository/mock"
)

func BuildMockUseCases() application.UseCases {
	cityRepo := mock.NewMockCityRepository()
	tripRepo := mock.NewMockTripRepository()

	return application.NewTripService(tripRepo, cityRepo)
}
