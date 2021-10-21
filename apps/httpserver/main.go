package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mteixidorc/trips/apps/httpserver/controllers"
)

// ControllerHandler
// Interface that all controllers must implement to expose their endpoints
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

	// Adding controllers
	s.addController(controllers.NewTripHTTPController(nil))

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
	// TODO port and other variables should be setted using environment variables
	port := 8080

	log.Printf("Starting server at port %d", port)
	s := NewServer()
	s.HandleFunc("/", s.home)
	http.Handle("/", s)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
