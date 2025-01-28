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

// FetchReferralIndex calculates and returns ReferralIndex
func (*Repository) FetchReferralIndex(ctx context.Context) (map[int]int, error) {
	res := make(map[int]int)
	referrers := make(map[int][]int)
	for _, a := range actionsDB {
		if a.Type == "REFER_USER" {
			res[a.UserID]++
			referrers[a.TargetUser] = append(referrers[a.TargetUser], a.UserID)
			referrers[a.TargetUser] = append(referrers[a.TargetUser], referrers[a.UserID]...)
			for _, referrer := range referrers[a.UserID] {
				res[referrer]++
			}
			res[a.TargetUser] = 0
		}
	}
	return res, nil
}
