package trip_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mteixidorc/trips/apps/httpserver/controllers/trip"
	appmock "github.com/mteixidorc/trips/internal/trips/application/mock"
	"github.com/mteixidorc/trips/internal/trips/infrastructure/repository/mock"
)

func TestTripControllerGetByID(t *testing.T) {
	tt := []struct {
		name       string
		method     string
		id         string
		want       string
		statusCode int
	}{
		{
			name:       "OK get trip by ID, exists",
			method:     http.MethodGet,
			id:         mock.MockTrip1ID.String(),
			want:       `test-city-1`,
			statusCode: http.StatusOK,
		},
		{
			name:       "FAIL get trip, not exists",
			method:     http.MethodPost,
			id:         uuid.NewString(),
			want:       `not exists`,
			statusCode: http.StatusBadRequest,
		},
	}

	useCases := appmock.BuildMockUseCases()
	controller := trip.NewTripHTTPController(useCases)

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/trip/{id}", nil)
			req = mux.SetURLVars(req, map[string]string{
				"id": tc.id,
			})

			w := httptest.NewRecorder()
			controller.GetTripByIdHandler(w, req)
			res := w.Result()

			defer res.Body.Close()
			data, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}

			if res.StatusCode != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, res.StatusCode)
			}

			if !strings.Contains(string(data), tc.want) {
				t.Errorf("expected %s got %v", tc.want, string(data))
			}
		})
	}
}

func TestTripControllerGetALL(t *testing.T) {
	useCases := appmock.BuildMockUseCases()
	controller := trip.NewTripHTTPController(useCases)

	req := httptest.NewRequest(http.MethodGet, "/trip", nil)

	w := httptest.NewRecorder()
	controller.GetAllTripsHandler(w, req)
	res := w.Result()

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	if !(strings.Contains(string(data), "test-city-1") && strings.Contains(string(data), "test-city-2")) {
		t.Errorf("expected %s got %v", "test-city-1 and test-city-2", string(data))
	}
}

func TestTripControllerPOST(t *testing.T) {
	tt := []struct {
		name       string
		method     string
		trip       *trip.TripRequest
		body       string
		want       string
		statusCode int
	}{
		{
			name:       "OK add a new trip",
			method:     http.MethodPost,
			body:       `{"originId":1,"destinationId":2,"dates":"Mon Tue","price":10.12}`,
			want:       `trip record saved`,
			statusCode: http.StatusOK,
		},
		{
			name:       "FAIL add a new trip dates are incorrect",
			method:     http.MethodPost,
			body:       `{"originId":1,"destinationId":2,"dates":"Moc Tuc","price":10.12}`,
			want:       `not a valid trip date`,
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "FAIL add a new trip, destinationId not exists",
			method:     http.MethodPost,
			body:       `{"originId":1,"destinationId":332,"dates":"Moc Tuc","price":10.12}`,
			want:       `destination city 332 not exists`,
			statusCode: http.StatusBadRequest,
		},
	}

	useCases := appmock.BuildMockUseCases()
	controller := trip.NewTripHTTPController(useCases)

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, "/trip", strings.NewReader(tc.body))

			w := httptest.NewRecorder()
			controller.PostTripHandler(w, req)
			res := w.Result()

			defer res.Body.Close()
			data, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}

			if res.StatusCode != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, res.StatusCode)
			}

			if !strings.Contains(string(data), tc.want) {
				t.Errorf("expected %s got %v", tc.want, string(data))
			}
		})
	}
}
