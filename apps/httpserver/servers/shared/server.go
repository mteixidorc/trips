package shared

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mteixidorc/trips/apps/httpserver/controllers/shared"
)

// server
// HTTP server router
type Server struct {
	*mux.Router
	controllers []shared.ControllerHandler
}

func NewServer() Server {
	s := Server{
		Router: mux.NewRouter(),
	}

	return s
}

func (s *Server) AddController(controller shared.ControllerHandler) {
	s.controllers = append(s.controllers, controller)
	controller.AddHTTPHandlers(s.Router)
}

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "I'm alive"}`))
}
