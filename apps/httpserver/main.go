package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mteixidorc/trips/apps/httpserver/controllers/trip"
)

// ControllerHandler
// Interface that all controllers must implement to expose its endpoints
type ControllerHandler interface {
	AddHTTPHandlers(router *mux.Router)
}

// server
// HTTP server router
type server struct {
	*mux.Router
	controllers []ControllerHandler
}

func NewServer() server {
	s := server{
		Router: mux.NewRouter(),
	}

	// Add here all controllers
	s.addController(trip.NewTripHTTPController(nil))

	return s
}

func (s *server) addController(controller ControllerHandler) {
	s.controllers = append(s.controllers, controller)
	controller.AddHTTPHandlers(s.Router)
}

func (s *server) home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "I'm alive"}`))
}

// Entry point of our HTTP Service for trips bounded context
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server at port %s", port)
	s := NewServer()
	s.HandleFunc("/", s.home)
	http.Handle("/", s)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
