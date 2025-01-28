package transport

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ActionsHandler handles user related HTTP requests
type ActionsHandler struct {
	service UserService
}

// NewActionsHandler creates a new user HTTP handler
func NewActionsHandler(s UserService) *ActionsHandler {
	return &ActionsHandler{service: s}
}

// RegisterRoutes registers the endpoints with the router
func (h *ActionsHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/actions/{id:[0-9]+}", h.getUser).Methods(http.MethodGet)
}

func (h *ActionsHandler) getUser(w http.ResponseWriter, req *http.Request) {
}
