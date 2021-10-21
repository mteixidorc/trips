package trip

import "github.com/mteixidorc/trips/internal/trips/application/trip"

// TripResponse
// DTO used by controllers, represents a trip object response
type TripResponse struct {
	Origin      string  `json:"origin"`
	Destination string  `json:"destination"`
	Dates       string  `json:"dates"`
	Price       float64 `json:"price"`
}

func TripsToTripResponses(trips ...*trip.TripDTO) []TripResponse {
	responses := make([]TripResponse, len(trips))
	for pos, trip := range trips {
		responses[pos] = TripResponse{
			Origin:      trip.OriginCityName,
			Destination: trip.DestinationCityName,
			Dates:       trip.Dates,
			Price:       trip.Price,
		}
	}

	return responses
}

// TripRequest
// DTO used by controllers, represents a trip object request
type TripRequest struct {
	OriginId      int64   `json:"originId"`
	DestinationId int64   `json:"destinationId"`
	Dates         string  `json:"dates"`
	Price         float64 `json:"price"`
}
