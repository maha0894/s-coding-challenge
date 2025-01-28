package repository

import (
	"context"
	"encoding/json"
	"errors"
	"os"

	"github.com/maha0894/s-coding-challenge/pkg/entities"
)

var (
	usersDB   = make(map[int]entities.User)
	actionsDB = make(map[int][]entities.Action)
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
	var actions []entities.Action
	err = json.Unmarshal(buf1, &actions)
	if err != nil {
		return err
	}

	for _, a := range actions {
		actionsDB[a.UserID] = append(actionsDB[a.UserID], a)
	}

	return err
}

func (*Repository) FetchUserInfo(ctx context.Context, id int) (entities.User, error) {
	user, ok := usersDB[id]
	if !ok {
		return entities.User{}, errors.New("user not found")
	}
	return user, nil
}
