package repository

import (
	"context"
	"errors"

	"github.com/maha0894/s-coding-challenge/pkg/entities"
)

// FetchUserInfo returns user info from the database
func (*Repository) FetchUserInfo(ctx context.Context, id int) (entities.User, error) {
	user, ok := usersDB[id]
	if !ok {
		return entities.User{}, errors.New("user not found")
	}
	return user, nil
}
