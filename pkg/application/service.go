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
	FetchUserActionsCount(ctx context.Context, userId int) (entities.Actions, error)
	FetchReferralIndex(ctx context.Context) (map[int]int, error)
}

// NewService creates new service
func NewService(r Repository) *Service {
	return &Service{repository: r}
}

// FetchUserInfo returns user info from the database
func (s *Service) FetchUserInfo(ctx context.Context, id int) (entities.User, error) {
	return s.repository.FetchUserInfo(ctx, id)
}

// FetchUserActionsCount returns user actions count from the database
func (s *Service) FetchUserActionsCount(ctx context.Context, userId int) (entities.Actions, error) {
	return s.repository.FetchUserActionsCount(ctx, userId)
}

// FetchReferralIndex returns all users ReferralIndexes
func (s *Service) FetchReferralIndex(ctx context.Context) (map[int]int, error) {
	return s.repository.FetchReferralIndex(ctx)
}
