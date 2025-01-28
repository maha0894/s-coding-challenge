package entities

import "time"

type Action struct {
	ID         int       `json:"id"`
	Type       string    `json:"type"`
	UserID     int       `json:"userId"`     // The ID of the User who performed this action
	TargetUser int       `json:"targetUser"` // Supplied when "REFER_USER" action type
	CreatedAt  time.Time `json:"createdAt"`
}
