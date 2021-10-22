package shared

import "github.com/gorilla/mux"

// ControllerHandler
// Interface that all controllers must implement to expose its endpoints
type ControllerHandler interface {
	AddHTTPHandlers(router *mux.Router)
}
