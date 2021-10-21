package trip

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mteixidorc/trips/apps/httpserver/controllers/shared"
	"github.com/mteixidorc/trips/internal/trips/application"
	"github.com/mteixidorc/trips/internal/trips/application/trip"
	"github.com/mteixidorc/trips/internal/trips/infrastructure/repository"
)

// TripHTTPController
// HTTP controller with all endpoints related to trips use cases
type TripHTTPController struct {
	useCases application.UseCases
}

func NewTripHTTPController(useCases application.UseCases) TripHTTPController {
	// TODO - Dependency Injection
	if useCases == nil {
		tripRepository := repository.NewInMemoryTripRepository()
		cityRepository, err := repository.NewFileCityRepository()
		if err != nil {
			log.Fatal(err)
		}

		useCases = application.NewTripService(&tripRepository, cityRepository)
	}
	return TripHTTPController{
		useCases: useCases,
	}
}

func (controller TripHTTPController) AddHTTPHandlers(router *mux.Router) {
	router.HandleFunc("/trip", controller.GetAllTripsHandler).Methods(http.MethodGet)
	router.HandleFunc("/trip/{id}", controller.GetTripByIdHandler).Methods(http.MethodGet)
	router.HandleFunc("/trip", controller.PostTripHandler).Methods(http.MethodPost)
}

func (c *TripHTTPController) GetAllTripsHandler(w http.ResponseWriter, r *http.Request) {
	trips, err := c.useCases.GetAllTrips()
	if err != nil {
		http.Error(w, shared.NewHTTPErrorJSON(0, err.Error()), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(TripsToTripResponses(trips...))
}

func (s *TripHTTPController) GetTripByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	trip, err := s.useCases.GetTripById(id)
	if err != nil {
		http.Error(w, shared.NewHTTPErrorJSON(0, err.Error()), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(TripsToTripResponses(trip))
}

func (s *TripHTTPController) PostTripHandler(w http.ResponseWriter, r *http.Request) {
	tripRequestBody := TripRequest{}
	err := json.NewDecoder(r.Body).Decode(&tripRequestBody) 
	if err != nil {
		http.Error(w, fmt.Sprintf("JSON format error: %s", err.Error()), http.StatusBadRequest)
		return
	}

	id, err := s.useCases.CreateTrip(
		&trip.TripDTO{
			OriginId:      tripRequestBody.OriginId,
			DestinationId: tripRequestBody.DestinationId,
			Dates:         tripRequestBody.Dates,
			Price:         tripRequestBody.Price,
		},
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(shared.HTTPPostCreationOk{Id: id, Message: "trip record saved"})

}
