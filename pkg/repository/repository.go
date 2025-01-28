package repository

import (
	"context"

	"github.com/maha0894/s-coding-challenge/pkg/entities"
)

// Repository implementation
type Repository struct {
}

// New returns a new Repository instance
func New() *Repository {
	return &Repository{}
}

func (*Repository) FetchUserInfo(ctx context.Context, id int) (entities.User, error) {
	return entities.User{}, nil
}
