package repository

import (
	"encoding/json"
	"os"

	"github.com/maha0894/s-coding-challenge/pkg/entities"
)

var (
	usersDB       = make(map[int]entities.User)
	userActionsDB = make(map[int][]entities.Action)
	actionsDB     []entities.Action
)

// Repository implementation
type Repository struct {
}

// New returns a new Repository instance
func New() (*Repository, error) {
	return &Repository{}, initialise()
}

// initialise loads users and actions from json
func initialise() error {
	buf, err := os.ReadFile("db/users.json")
	if err != nil {
		return err
	}
	var users []entities.User
	err = json.Unmarshal(buf, &users)
	if err != nil {
		return err
	}

	for _, u := range users {
		usersDB[u.ID] = u
	}

	buf1, err := os.ReadFile("db/actions.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(buf1, &actionsDB)
	if err != nil {
		return err
	}

	for _, a := range actionsDB {
		userActionsDB[a.UserID] = append(userActionsDB[a.UserID], a)
	}

	return err
}
