package repository

import (
	"context"
	"errors"

	"github.com/maha0894/s-coding-challenge/pkg/entities"
)

// FetchUserActionsCount returns user actions count from the database
func (*Repository) FetchUserActionsCount(ctx context.Context, userId int) (entities.Actions, error) {
	actions, ok := userActionsDB[userId]
	if !ok {
		return entities.Actions{}, errors.New("user actions not found")
	}
	return entities.Actions{Count: len(actions)}, nil
}
