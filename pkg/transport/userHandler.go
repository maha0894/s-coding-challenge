package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/maha0894/s-coding-challenge/pkg/entities"
)

// UserHandler handles user related HTTP requests
type UserHandler struct {
	service UserService
}

// UserService ...
type UserService interface {
	FetchUserInfo(ctx context.Context, id int) (entities.User, error)
	FetchUserActionsCount(ctx context.Context, userId int) (entities.Actions, error)
	FetchReferralIndex(ctx context.Context) (map[int]int, error)
}

// NewUserHandler creates a new user HTTP handler
func NewUserHandler(s UserService) *UserHandler {
	return &UserHandler{service: s}
}

// RegisterRoutes registers the endpoints with the router
func (h *UserHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/users/{id:[0-9]+}", h.getUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{id:[0-9]+}/actions", h.getUserActionsCount).Methods(http.MethodGet)
	router.HandleFunc("/users/referral-index", h.getUsersReferralIndexes).Methods(http.MethodGet)
}

func (h *UserHandler) getUser(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing/incorrect id field"))
		return
	}
	user, err := h.service.FetchUserInfo(req.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
	}
}

func (h *UserHandler) getUserActionsCount(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing/incorrect id field"))
		return
	}
	actionsCount, err := h.service.FetchUserActionsCount(req.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(actionsCount)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
	}
}

func (h *UserHandler) getUsersReferralIndexes(w http.ResponseWriter, req *http.Request) {
	referralIndex, err := h.service.FetchReferralIndex(req.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(referralIndex)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
	}
}
