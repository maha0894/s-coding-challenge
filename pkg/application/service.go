package application

import (
	"context"

	"github.com/maha0894/s-coding-challenge/pkg/entities"
)

// Service internal implementation of the Service
type Service struct {
	repository Repository
}

// Repository is a data repository
type Repository interface {
	FetchUserInfo(ctx context.Context, id int) (entities.User, error)
}

// NewService creates new service
func NewService(r Repository) *Service {
	return &Service{repository: r}
}

// FetchUserInfo returns user info from the database.
func (s *Service) FetchUserInfo(ctx context.Context, id int) (entities.User, error) {
	return s.repository.FetchUserInfo(ctx, id)
}
