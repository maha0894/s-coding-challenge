package transport

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// ActionsHandler handles user related HTTP requests
type ActionsHandler struct {
	service Service
}

// NewActionsHandler creates a new user HTTP handler
func NewActionsHandler(s Service) *ActionsHandler {
	return &ActionsHandler{service: s}
}

// RegisterRoutes registers the endpoints with the router
func (h *ActionsHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/actions/{action}/next", h.getNextActions).Methods(http.MethodGet)
}

func (h *ActionsHandler) getNextActions(w http.ResponseWriter, req *http.Request) {
	action, ok := mux.Vars(req)["action"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing/incorrect action field"))
		return
	}
	nextActions, err := h.service.FetchNextActions(req.Context(), action)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(nextActions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
	}
}
