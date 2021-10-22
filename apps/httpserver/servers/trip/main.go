package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mteixidorc/trips/apps/httpserver/controllers/trip"
	"github.com/mteixidorc/trips/apps/httpserver/servers/shared"
)

// Trip server
// HTTP server router
type TripServer struct {
	*shared.Server
}

func NewTripServer() TripServer {
	baseServer := shared.NewServer()
	return TripServer{
		Server: &baseServer,
	}
}

// Entry point of our HTTP Service for trips bounded context
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := NewTripServer()
	server.AddController(trip.NewTripHTTPController(nil))

	log.Printf("Starting Trips server at port %s", port)
	server.HandleFunc("/", server.Home)
	http.Handle("/", server)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
